package transaction_parser

import (
	"context"
	"encoding/binary"
	"fmt"
	"solana_project/libraries/anchor/bubblegum/generated/bubblegum"
	"solana_project/libraries/programs"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func (tp *TransactionParser) parseBubblegumInstruction(
	ctx context.Context,
	sig solana.Signature,
	transactionResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
	instruction *solana.CompiledInstruction,
	instructionsWithNoopLogs []solana.CompiledInstruction,
	launchpad string,
) (*FullMintData, error) {
	accounts, err := instruction.ResolveInstructionAccounts(&tx.Message)
	if err != nil {
		return nil, err
	}
	decoded, err := bubblegum.DecodeInstruction(accounts, instruction.Data)
	if err != nil {
		return nil, err
	}
	mintToCollectionV1, ok := decoded.Impl.(*bubblegum.MintToCollectionV1)
	if !ok {
		return nil, nil
	}

	changeLogEvents := retrieveChangeLogEventsFromInstructions(instructionsWithNoopLogs)
	if len(changeLogEvents) == 0 {
		return nil, fmt.Errorf("expected to get at least single changeLogEvent; got len=%d", len(changeLogEvents))
	}
	changeLogEvent := changeLogEvents[len(changeLogEvents)-1]
	assetId, err := getLeafAssetId(changeLogEvent.V1.Id, changeLogEvent.V1.Index)
	if err != nil {
		return nil, err
	}
	out, err := tp.rpcClient.GetAccountInfo(
		ctx,
		mintToCollectionV1.GetCollectionMetadataAccount().PublicKey,
	)
	if err != nil {
		return nil, err
	}
	meta, err := MetadataDeserialize(out.Bytes())
	if err != nil {
		return nil, err
	}

	blockTime := solana.UnixTimeSeconds(time.Now().Unix())
	if transactionResult.BlockTime != nil {
		blockTime = *transactionResult.BlockTime
	}

	ipfsData, err := tp.ipfsClient.Get(mintToCollectionV1.MetadataArgs.Uri)
	if err != nil {
		return nil, err
	}

	return &FullMintData{
		Mint: Mint{
			TxSignature:   sig,
			ProgramKey:    programs.BubblegumProgramId,
			BlockTime:     blockTime,
			CollectionKey: mintToCollectionV1.MetadataArgs.Collection.Key,
			AssetKey:      assetId,
			AssetOwnerKey: mintToCollectionV1.GetLeafOwnerAccount().PublicKey,
			MintPriceSol:  float64(meta.Data.SellerFeeBasisPoints) / 1e4,
			IpfsData:      *ipfsData,
		},
		Collection: CollectionWithMeta{
			CollectionKey: mintToCollectionV1.MetadataArgs.Collection.Key,
			Name:          meta.Data.Name,
			Symbol:        meta.Data.Symbol,
			Image:         ipfsData.Image,
			Size:          meta.CollectionDetails.V1.Size / 2,
			Launchpad:     launchpad,
		},
	}, nil
}

func getLeafAssetId(tree solana.PublicKey, index uint32) (solana.PublicKey, error) {
	binIndex := make([]byte, 8)
	binary.LittleEndian.PutUint32(binIndex, index)
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("asset"),
			tree.Bytes(),
			binIndex,
		},
		programs.BubblegumProgramId,
	)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return addr, err
}

// func (tp *MintTrackerMintsProcessor) processBubblegum(
// 	ctx context.Context,
// 	payload *mint_tracker_mints.MessagePayloadBubblegum,
// ) (*models.MintData, error) {
// 	payload.Signature = "5Yx6mh1ENnMtt41vptWA24gm6tZvBKmmvEq6W1nqNKFoKPLuDXSjNEC4bx6UFMEzRLEoNGhju58a8NFavZbkJDUd"
// 	sig, err := solana.SignatureFromBase58(payload.Signature)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var maxSupportedTransactionVersion uint64 = 0
// 	transactionResult, err := tp.rpcClient.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
// 		Commitment:                     mp.commitment,
// 		Encoding:                       solana.EncodingBase64,
// 		MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// spew.Dump(transactionResult)
// 	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(transactionResult.Transaction.GetBinary()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// spew.Dump(tx)
// 	err = mp.resolveLookups(tx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var allInstructions []solana.CompiledInstruction
// 	allInstructions = append(allInstructions, tx.Message.Instructions...)
// 	for _, inner := range transactionResult.Meta.InnerInstructions {
// 		allInstructions = append(allInstructions, inner.Instructions...)
// 	}
// 	for _, instruction := range allInstructions {
// 		// spew.Dump(instruction)
// 		accounts, err := instruction.ResolveInstructionAccounts(&tx.Message)
// 		if err != nil {
// 			return nil, err
// 		}
// 		// spew.Dump(tx.Message.AccountKeys[instruction.ProgramIDIndex])
// 		decoded, err := bubblegum.DecodeInstruction(accounts, instruction.Data)
// 		if err != nil {
// 			continue
// 		}
// 		mintToCollectionV1, ok := decoded.Impl.(*bubblegum.MintToCollectionV1)
// 		if ok {
// 			// spew.Dump(mintToCollectionV1.GetAccounts())
// 			return &models.MintData{
// 				Signature: sig,
// 				Mint:      mintToCollectionV1.GetCollectionMintAccount().PublicKey,
// 				Payer:     mintToCollectionV1.GetTreeDelegateAccount().PublicKey,
// 				CandyMachine: mintToCollectionV1.MetadataArgs.Collection.Key,
// 				Program: tx.Message.AccountKeys[instruction.ProgramIDIndex],
// 				IpfsURL:   mintToCollectionV1.MetadataArgs.Uri,
// 				MintedAt: transactionResult.BlockTime.Time(),
// 			}, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("failed to find mint instruction")
// }
