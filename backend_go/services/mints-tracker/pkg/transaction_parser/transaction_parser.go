package transaction_parser

import (
	"context"
	"fmt"
	"solana_project/libraries/programs"
	"solana_project/libraries/queues/mint_tracker_mints"
	"solana_project/services/mints-tracker/pkg/ipfs_client"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"go.uber.org/zap"
)

type TransactionParser struct {
	logger     *zap.SugaredLogger
	rpcClient  *rpc.Client
	commitment rpc.CommitmentType
	ipfsClient *ipfs_client.IpfsClient
}

func NewTransactionParser(
	logger *zap.SugaredLogger,
	rpcClient *rpc.Client,
	commitment rpc.CommitmentType,
	ipfsClient *ipfs_client.IpfsClient) *TransactionParser {
	return &TransactionParser{
		logger:     logger,
		rpcClient:  rpcClient,
		commitment: commitment,
		ipfsClient: ipfsClient,
	}
}

func (tp *TransactionParser) RetrieveMintData(
	ctx context.Context,
	message *mint_tracker_mints.Message,
) ([]*FullMintData, error) {
	var results []*FullMintData

	launchpadOpt := programs.ProgramToString(message.ProgramId)
	if launchpadOpt == nil {
		return nil, fmt.Errorf("unknown launchpad: " + message.ProgramId)
	}
	launchpad := *launchpadOpt

	sig, err := solana.SignatureFromBase58(message.Signature)
	if err != nil {
		return nil, err
	}
	var maxSupportedTransactionVersion uint64 = 0
	transactionResult, err := tp.rpcClient.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
		Commitment:                     tp.commitment,
		Encoding:                       solana.EncodingBase64,
		MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
	})
	if err != nil {
		return nil, err
	}
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(transactionResult.Transaction.GetBinary()))
	if err != nil {
		return nil, err
	}
	err = tp.resolveLookups(tx)
	if err != nil {
		return nil, err
	}

	instructionIdxToInner := make(map[int]*rpc.InnerInstruction)
	for _, innerInstruction := range transactionResult.Meta.InnerInstructions {
		instructionIdxToInner[int(innerInstruction.Index)] = &innerInstruction
	}

	for i, instruction := range tx.Message.Instructions {
		programId := tx.Message.AccountKeys[instruction.ProgramIDIndex]
		switch programId.String() {
		case programs.LMNFTBubblegumProgramID.String():
			mintData, err := tp.parseLmnftCompressedInstruction(
				ctx,
				sig,
				transactionResult,
				tx,
				&instruction,
				launchpad,
			)
			if err != nil {
				return nil, err
			}
			if mintData != nil {
				results = append(results, mintData)
			}
		// case programs.BubblegumProgramId.String():
		// 	innerInstruction := instructionIdxToInner[i]
		// 	if innerInstruction == nil {
		// 		continue
		// 	}
		// 	instructionsWithNoopLog := innerInstruction.Instructions
		// 	mintData, err := tp.parseBubblegumInstruction(
		// 		ctx,
		// 		sig,
		// 		transactionResult,
		// 		tx,
		// 		&instruction,
		// 		instructionsWithNoopLog,
		// 		launchpad)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	if mintData != nil {
		// 		results = append(results, mintData)
		// 	}
		case programs.TruffleProgramId.String():
			innerInstruction := instructionIdxToInner[i]
			if innerInstruction == nil {
				continue
			}
			mintData, err := tp.parseTruffleInstruction(
				ctx,
				sig,
				transactionResult,
				tx,
				&instruction,
				innerInstruction.Instructions,
				launchpad,
			)
			if err != nil {
				return nil, err
			}
			if mintData != nil {
				results = append(results, mintData)
			}
		case programs.LMNFTProgramID.String():
			mintData, err := tp.parseLmnftInstruction(
				ctx,
				sig,
				transactionResult,
				tx,
				&instruction,
				launchpad,
			)
			if err != nil {
				return nil, err
			}
			if mintData != nil {
				results = append(results, mintData)
			}
		}
	}

	return results, nil
}
