package transaction_parser

import (
	"context"
	"solana_project/libraries/anchor/candy_guard/generated/candy_guard"
	"solana_project/libraries/programs"
	"solana_project/services/mints-tracker/pkg/models"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
)

func (tp *TransactionParser) parseTruffleInstruction(
	ctx context.Context,
	sig solana.Signature,
	transactionResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
	instruction *solana.CompiledInstruction,
	innerInstructions []solana.CompiledInstruction,
	launchpad string,
) (*FullMintData, error) {
	accounts, err := instruction.ResolveInstructionAccounts(&tx.Message)
	if err != nil {
		return nil, err
	}
	decoded, err := candy_guard.DecodeInstruction(accounts, instruction.Data)
	if err != nil {
		return nil, err
	}
	mintV2, ok := decoded.Impl.(*candy_guard.MintV2)
	if !ok {
		return nil, nil
	}

	blockTime := solana.UnixTimeSeconds(time.Now().Unix())
	if transactionResult.BlockTime != nil {
		blockTime = *transactionResult.BlockTime
	}

	nftMetadataAccountInfo, err := tp.rpcClient.GetAccountInfo(
		ctx,
		mintV2.GetNftMetadataAccount().PublicKey,
	)
	if err != nil {
		return nil, err
	}
	nftMeta, err := MetadataDeserialize(nftMetadataAccountInfo.Bytes())
	if err != nil {
		return nil, err
	}

	collectionMetadataAccountInfo, err := tp.rpcClient.GetAccountInfo(
		ctx,
		mintV2.GetCollectionMetadataAccount().PublicKey,
	)
	if err != nil {
		return nil, err
	}
	collectionMeta, err := MetadataDeserialize(collectionMetadataAccountInfo.Bytes())
	if err != nil {
		return nil, err
	}

	var mintCost float64
	for _, instruction := range innerInstructions {
		dec, err := system.DecodeInstruction(accounts, instruction.Data)
		if err != nil {
			continue
		}
		transfer, ok := dec.Impl.(*system.Transfer)
		if !ok {
			continue
		}
		if transfer.Lamports == nil {
			continue
		}
		mintCost = float64(*transfer.Lamports) / float64(solana.LAMPORTS_PER_SOL)
		break
	}

	// candyMachineDataRaw, err := tp.rpcClient.GetAccountInfo(
	// 	ctx,
	// 	mintV2.GetCandyMachineAccount().PublicKey,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// var candyMachineData candy_machine_core.CandyMachine
	// err = candyMachineData.UnmarshalWithDecoder(bin.NewBorshDecoder(candyMachineDataRaw.Bytes()))
	// if err != nil {
	// 	return nil, err
	// }

	ipfsData, err := tp.ipfsClient.Get(nftMeta.Data.Uri)
	if err != nil {
		return nil, err
	}

	return &FullMintData{
		Mint: Mint{
			TxSignature:   sig,
			ProgramKey:    programs.TruffleProgramId,
			BlockTime:     blockTime,
			CollectionKey: mintV2.GetCollectionMintAccount().PublicKey,
			AssetKey:      mintV2.GetNftMintAccount().PublicKey,
			AssetOwnerKey: mintV2.GetMinterAccount().PublicKey,
			MintPriceSol:  mintCost,
			IpfsData:      *ipfsData,
		},
		Collection: CollectionWithMeta{
			CollectionKey: mintV2.GetCollectionMintAccount().PublicKey,
			Name:          collectionMeta.Data.Name,
			Symbol:        collectionMeta.Data.Symbol,
			Image:         ipfsData.Image,
			Size:          collectionMeta.CollectionDetails.V1.Size,
			Launchpad:     launchpad,
			Type:          models.CollectionTypeMPLX,
		},
	}, nil
}
