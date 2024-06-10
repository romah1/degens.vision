package transaction_parser

import (
	"context"
	"solana_project/libraries/anchor/lmnft/generated/nft_candy_machine"
	"solana_project/libraries/programs"
	"solana_project/services/mints-tracker/pkg/models"
	"strings"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func (tp *TransactionParser) parseLmnftInstruction(
	ctx context.Context,
	sig solana.Signature,
	transactionResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
	instruction *solana.CompiledInstruction,
	launchpad string,
) (*FullMintData, error) {
	accounts, err := instruction.ResolveInstructionAccounts(&tx.Message)
	if err != nil {
		return nil, err
	}
	decoded, err := nft_candy_machine.DecodeInstruction(accounts, instruction.Data)
	if err != nil {
		return nil, err
	}
	mintV6, ok := decoded.Impl.(*nft_candy_machine.MintV6)
	if ok {
		result, err := tp.parseLmnftInstructionMintV6(ctx, sig, transactionResult, mintV6, launchpad)
		return result, err
	}

	mintV3, ok := decoded.Impl.(*nft_candy_machine.MintV2)
	if ok {
		result, err := tp.parseLmnftInstructionMintV2(ctx, sig, mintV3)
		return result, err
	}

	return nil, nil
}

func (tp *TransactionParser) parseLmnftInstructionMintV6(
	ctx context.Context,
	sig solana.Signature,
	transactionResult *rpc.GetTransactionResult,
	mintV6 *nft_candy_machine.MintV6,
	launchpad string,
) (*FullMintData, error) {
	nftMetadataAccountInfo, err := tp.rpcClient.GetAccountInfo(
		ctx,
		mintV6.GetMetadataAccount().PublicKey,
	)
	if err != nil {
		return nil, err
	}
	nftMeta, err := MetadataDeserialize(nftMetadataAccountInfo.Bytes())
	if err != nil {
		return nil, err
	}

	candyMachineData, err := tp.retrieveLmnftCandyMachineData(ctx, mintV6.GetCandyMachineAccount().PublicKey)
	if err != nil {
		return nil, err
	}

	curTime := time.Now().Unix()
	var curPrice float64
	for _, salePhase := range candyMachineData.Data.SaleFazes {
		if salePhase.Start < curTime {
			curPrice = float64(salePhase.Price) / float64(solana.LAMPORTS_PER_SOL)
		}
	}

	ipfsData, err := tp.ipfsClient.Get(nftMeta.Data.Uri)
	if err != nil {
		return nil, err
	}

	collectionMetadataAccount := mintV6.GetCollectionMetadataAccount()
	var collectionName string
	var collectionSymbol string
	if collectionMetadataAccount == nil {
		collectionName = strings.TrimSpace(strings.Split(ipfsData.Name, "#")[0])
		collectionSymbol = ipfsData.Symbol
	} else {
		collectionMetadataAccountInfo, err := tp.rpcClient.GetAccountInfo(
			ctx,
			mintV6.GetCollectionMetadataAccount().PublicKey,
		)
		if err != nil {
			return nil, err
		}
		collectionMeta, err := MetadataDeserialize(collectionMetadataAccountInfo.Bytes())
		if err != nil {
			return nil, err
		}
		collectionName = collectionMeta.Data.Name
		collectionSymbol = collectionMeta.Data.Symbol
	}

	blockTime := solana.UnixTimeSeconds(time.Now().Unix())
	if transactionResult.BlockTime != nil {
		blockTime = *transactionResult.BlockTime
	}

	return &FullMintData{
		Mint: Mint{
			TxSignature:   sig,
			ProgramKey:    programs.LMNFTProgramID,
			BlockTime:     blockTime,
			CollectionKey: mintV6.GetCollectionMintAccount().PublicKey,
			AssetKey:      mintV6.GetMintAccount().PublicKey,
			AssetOwnerKey: mintV6.GetPayerAccount().PublicKey,
			MintPriceSol:  curPrice,
			IpfsData:      *ipfsData,
		},
		Collection: CollectionWithMeta{
			CollectionKey: mintV6.GetCollectionMintAccount().PublicKey,
			CandyMachine:  mintV6.GetCandyMachineAccount().PublicKey.ToPointer(),
			Name:          collectionName,
			Symbol:        collectionSymbol,
			Image:         ipfsData.Image,
			Size:          candyMachineData.Data.ItemsAvailable,
			Launchpad:     launchpad,
			Type:          models.CollectionTypeMPLX,
		},
	}, nil
}

func (tp *TransactionParser) parseLmnftInstructionMintV2(
	ctx context.Context,
	sig solana.Signature,
	mintV2 *nft_candy_machine.MintV2,
) (*FullMintData, error) {
	// TODO: implement if needed
	return nil, nil
}

func (tp *TransactionParser) retrieveLmnftCandyMachineData(
	ctx context.Context,
	candyMachineAddress solana.PublicKey,
) (*nft_candy_machine.CandyMachineV6, error) {
	candyMachineAccountInfo, err := tp.rpcClient.GetAccountInfo(
		ctx,
		candyMachineAddress,
	)
	if err != nil {
		return nil, err
	}
	var candyMachineData nft_candy_machine.CandyMachineV6
	candyMachineData.UnmarshalWithDecoder(bin.NewBorshDecoder(candyMachineAccountInfo.Bytes()))
	if err != nil {
		return nil, err
	}
	return &candyMachineData, nil
}
