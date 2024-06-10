package mint_tracker_mints_processor

import (
	"context"
	"encoding/json"
	"solana_project/libraries/queues/mint_tracker_mints"
	"solana_project/libraries/topics/live_feed"
	"solana_project/services/mints-tracker/pkg/transaction_parser"
	"time"

	"github.com/IBM/sarama"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
)

type MintTrackerMintsProcessor struct {
	logger            *zap.SugaredLogger
	commitment        rpc.CommitmentType
	pg                *pgxpool.Pool
	transactionParser *transaction_parser.TransactionParser
	kafkaProducer     sarama.SyncProducer
}

func NewMintTrackerMintsProcessor(
	logger *zap.SugaredLogger,
	commitment rpc.CommitmentType,
	pg *pgxpool.Pool,
	transactionParser *transaction_parser.TransactionParser,
	kafkaProducer sarama.SyncProducer,
) *MintTrackerMintsProcessor {
	return &MintTrackerMintsProcessor{
		logger:            logger,
		commitment:        commitment,
		pg:                pg,
		transactionParser: transactionParser,
		kafkaProducer:     kafkaProducer,
	}
}

func (mp *MintTrackerMintsProcessor) Process(
	ctx context.Context,
	message *mint_tracker_mints.Message,
) error {
	startedAt := time.Now()
	mintData, err := mp.transactionParser.RetrieveMintData(ctx, message)
	if err != nil {
		return err
	}
	for _, mint := range mintData {
		err = mp.UpsertMint(ctx, mint)
		if err != nil {
			return err
		}
		err = mp.broadcastLiveFeed(mint)
		if err != nil {
			return err
		}
	}
	finishedAt := time.Now()
	mp.logger.Infof("Finished in %f seconds", finishedAt.Sub(startedAt).Seconds())

	return nil
}

func (mp *MintTrackerMintsProcessor) broadcastLiveFeed(mint *transaction_parser.FullMintData) error {
	topicMessage := live_feed.Message{
		Mint: live_feed.MessageMint{
			TxSignature:      mint.Mint.TxSignature.String(),
			ProgramKey:       mint.Mint.ProgramKey.String(),
			BlockTime:        mint.Mint.BlockTime.Time(),
			CollectionKey:    mint.Mint.CollectionKey.String(),
			AssetKey:         mint.Mint.AssetKey.String(),
			AssetOwnerKey:    mint.Mint.AssetOwnerKey.String(),
			AssetName:        mint.Mint.IpfsData.Name,
			AssetSymbol:      mint.Mint.IpfsData.Symbol,
			AssetImage:       mint.Mint.IpfsData.Image,
			MintPrice:        mint.Mint.MintPriceSol,
			AssetDescription: mint.Mint.IpfsData.Description,
			AssetAttributes:  mint.Mint.IpfsData.Attributes,
		},
		Collection: live_feed.MessageCollection{
			CollectionKey: mint.Collection.CollectionKey.String(),
			Name:          mint.Collection.Name,
			Symbol:        mint.Collection.Symbol,
			Image:         mint.Collection.Image,
			Size:          mint.Collection.Size,
			Launchpad:     mint.Collection.Launchpad,
		},
	}
	encodedTopicMessage, err := json.Marshal(topicMessage)
	if err != nil {
		return err
	}
	message := &sarama.ProducerMessage{
		Topic: "live_feed",
		Key:   sarama.StringEncoder(mint.Mint.AssetKey.String()),
		Value: sarama.ByteEncoder(encodedTopicMessage),
	}
	_, _, err = mp.kafkaProducer.SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}
