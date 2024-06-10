package collection_updater

import (
	"context"
	"fmt"
	"solana_project/libraries/anchor/lmnft/generated/nft_candy_machine"
	"solana_project/libraries/anchor/lmnft_compressed/generated/compressed_minter"
	clients "solana_project/libraries/clients/lmnft"
	"solana_project/libraries/programs"
	"solana_project/services/mints-tracker/pkg/models"
	"solana_project/services/mints-tracker/pkg/transaction_parser"
	"time"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CollectionUpdater struct {
	logger      *zap.SugaredLogger
	rpcClient   *rpc.Client
	pg          *pgxpool.Pool
	lmnftClient *clients.LmnftClient
}

func NewCollectionUpdater(
	logger *zap.SugaredLogger,
	rpcClient *rpc.Client,
	pg *pgxpool.Pool,
	lmnftClient *clients.LmnftClient,
) *CollectionUpdater {
	return &CollectionUpdater{
		logger:      logger,
		rpcClient:   rpcClient,
		pg:          pg,
		lmnftClient: lmnftClient,
	}
}

func (cu *CollectionUpdater) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			collectionForUpdate, err := cu.selectCollectionForUpdate(ctx)
			if err != nil {
				cu.logger.Errorf("failed to select collection for update; error=%s", err.Error())
				continue
			}
			if collectionForUpdate == nil {
				continue
			}
			updateCollectionData, err := cu.UpdateCollection(ctx, collectionForUpdate)
			if err != nil {
				cu.logger.Errorf(
					"failed to update collection; collection_key=%s; error=%s",
					collectionForUpdate.CollectionKey.String(),
					err.Error(),
				)
				continue
			}
			err = cu.updateCollection(ctx, updateCollectionData)
			if err != nil {
				cu.logger.Errorf("failed to save to database; error=%s", err.Error())
				continue
			}
			cu.logger.Infof(
				"updated collection; update_collection_data=%s",
				spew.Sdump(updateCollectionData),
			)
		case <-ctx.Done():
			return
		}
	}
}

func (cu *CollectionUpdater) UpdateCollection(
	ctx context.Context,
	collecollectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	switch collecollectionForUpdate.Launchpad {
	case models.LaunchpadLmnft:
		return cu.updateLmnftCollection(ctx, collecollectionForUpdate)
	case models.LaunchpadTruffle:
		return cu.updateTruffleCollection(ctx, collecollectionForUpdate)
	default:
		return nil, fmt.Errorf("unknown launchpad: %s", collecollectionForUpdate.Launchpad)
	}
}

func (cu *CollectionUpdater) updateLmnftCollection(
	ctx context.Context,
	collectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	switch collectionForUpdate.Type {
	case models.CollectionTypeMPLX:
		return cu.updateLmnftMPLXCollection(ctx, collectionForUpdate)
	case models.CollectionTypeCNFT:
		return cu.updateLmnftCNFTCollection(ctx, collectionForUpdate)
	case models.CollectionTypeCore:
		return cu.updateLmnftCoreCollection(ctx, collectionForUpdate)
	}

	return nil, fmt.Errorf("unexpected collection type; type=%s", collectionForUpdate.Type)
}

func (cu *CollectionUpdater) updateLmnftMPLXCollection(
	ctx context.Context,
	collectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	if collectionForUpdate.CandyMachine == nil {
		return nil, fmt.Errorf("candy machine must be set for mplx lmnft mints")
	}
	candyMachineAccountInfo, err := cu.rpcClient.GetAccountInfo(
		ctx,
		*collectionForUpdate.CandyMachine,
	)
	if err != nil {
		return nil, err
	}
	var candyMachineData nft_candy_machine.CandyMachineV6
	candyMachineData.UnmarshalWithDecoder(bin.NewBorshDecoder(candyMachineAccountInfo.Bytes()))
	if err != nil {
		return nil, err
	}

	lmnftData, err := cu.getDataFromLmnft(collectionForUpdate.Name, candyMachineData.Authority)
	var links []UpdateCollectionDataLink
	if err != nil {
		cu.logger.Errorf("failed to get data from lmnft; err=%s", err.Error())
	} else {
		links = lmnftData.Links
	}

	return &UpdateCollectionData{
		CollectionKey: collectionForUpdate.CollectionKey.String(),
		TotalMinted:   candyMachineData.ItemsRedeemed,
		Size:          candyMachineData.Data.ItemsAvailable,
		Links:         links,
	}, nil
}

func (cu *CollectionUpdater) updateLmnftCNFTCollection(
	ctx context.Context,
	collectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	if collectionForUpdate.CandyMachine == nil {
		return nil, fmt.Errorf("candy machine must be set for cnft lmnft mints")
	}
	minterAccountInfoRaw, err := cu.rpcClient.GetAccountInfo(
		ctx,
		*collectionForUpdate.CandyMachine,
	)
	if err != nil {
		return nil, err
	}
	var minterData compressed_minter.Minter
	minterData.UnmarshalWithDecoder(bin.NewBorshDecoder(minterAccountInfoRaw.Bytes()))
	if err != nil {
		return nil, err
	}

	lmnftData, err := cu.getDataFromLmnft(collectionForUpdate.Name, minterData.Authority)
	var links []UpdateCollectionDataLink
	if err != nil {
		cu.logger.Errorf("failed to get data from lmnft; err=%s", err.Error())
	} else {
		links = lmnftData.Links
	}

	return &UpdateCollectionData{
		CollectionKey: collectionForUpdate.CollectionKey.String(),
		TotalMinted:   uint64(minterData.Sold),
		Size:          uint64(minterData.Supply),
		Links:         links,
	}, nil
}

func (cu *CollectionUpdater) updateLmnftCoreCollection(
	ctx context.Context,
	collectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	if collectionForUpdate.CandyMachine == nil {
		return nil, fmt.Errorf("candy machine must be set for cnft lmnft mints")
	}
	mintMachineDataRaw, err := cu.rpcClient.GetAccountInfo(
		ctx,
		*collectionForUpdate.CandyMachine,
	)
	if err != nil {
		return nil, err
	}
	var mintMachine compressed_minter.MintMachine
	mintMachine.UnmarshalWithDecoder(bin.NewBorshDecoder(mintMachineDataRaw.Bytes()))
	if err != nil {
		return nil, err
	}

	lmnftData, err := cu.getDataFromLmnft(collectionForUpdate.Name, mintMachine.Data.Authority)
	var links []UpdateCollectionDataLink
	if err != nil {
		cu.logger.Errorf("failed to get data from lmnft; err=%s", err.Error())
	} else {
		links = lmnftData.Links
	}

	return &UpdateCollectionData{
		CollectionKey: collectionForUpdate.CollectionKey.String(),
		TotalMinted:   uint64(mintMachine.Sold),
		Size:          uint64(mintMachine.Data.Supply),
		Links:         links,
	}, nil
}

func (cu *CollectionUpdater) updateTruffleCollection(
	ctx context.Context,
	collectionForUpdate *CollectionForUpdate,
) (*UpdateCollectionData, error) {
	metadata, err := transaction_parser.RetrieveMetadataForAccount(
		ctx, cu.rpcClient, collectionForUpdate.CollectionKey,
	)
	if err != nil {
		return nil, err
	}
	if metadata.CollectionDetails == nil {
		return nil, fmt.Errorf("metadata.CollectionDetails is nil")
	}
	return &UpdateCollectionData{
		CollectionKey: collectionForUpdate.CollectionKey.String(),
		Size:          metadata.CollectionDetails.V1.Size,
		TotalMinted:   metadata.Uses.Total,
	}, nil
}

func (cu *CollectionUpdater) getMetadataAccount(
	collectionKey solana.PublicKey,
) (solana.PublicKey, error) {
	metadataAccount, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			programs.MetaplexMetadataProgramId.Bytes(),
			collectionKey.Bytes(),
		},
		programs.MetaplexMetadataProgramId,
	)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return metadataAccount, nil
}

type dataFromLmnft struct {
	Links []UpdateCollectionDataLink
}

func (cu *CollectionUpdater) getDataFromLmnft(
	collectionName string,
	collectionOwner solana.PublicKey,
) (*dataFromLmnft, error) {
	multiSearchResult, err := cu.lmnftClient.MultiSearch(&clients.MultiSearchRequest{
		Searches: []clients.MultiSearchRequestSearch{
			{
				QueryBy:    "collectionName",
				Collection: "collections",
				SortBy:     "deployed:desc,fractionMinted:desc",
				Q:          collectionName,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(multiSearchResult.Results) != 1 {
		return nil, fmt.Errorf("got not 1 result from lmnft multi search; len=%d", len(multiSearchResult.Results))
	}

	result := multiSearchResult.Results[0]
	for _, hit := range result.Hits {
		if hit.Document.Owner == collectionOwner.String() {
			collectionData, err := cu.lmnftClient.DataCollection(&clients.DataCollectionRequest{
				Owner: hit.Document.Owner,
				Id:    hit.Document.Id,
			})
			if err != nil {
				return nil, err
			}

			var links []UpdateCollectionDataLink
			if collectionData.PageProps.Collection.Twitter != nil {
				s := "https://twitter.com/" + *collectionData.PageProps.Collection.Twitter
				links = append(links, UpdateCollectionDataLink{
					Type: models.CollectionLinkTypeTwitter,
					Uri:  s,
				})
			}

			if collectionData.PageProps.Collection.Discord != nil {
				s := "https://discord.gg/" + *collectionData.PageProps.Collection.Discord
				links = append(links, UpdateCollectionDataLink{
					Type: models.CollectionLinkTypeDiscord,
					Uri:  s,
				})
			}

			mintLink := fmt.Sprintf(
				"https://www.launchmynft.io/collections/%s/%s",
				hit.Document.Owner,
				hit.Document.Id,
			)
			links = append(links, UpdateCollectionDataLink{
				Type: models.CollectionLinkTypeMint,
				Uri:  mintLink,
			})

			return &dataFromLmnft{
				Links: links,
			}, nil
		}
	}

	return nil, fmt.Errorf("failed to find collection on lmnft; name=%s", collectionName)
}
