package transaction_parser

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"solana_project/libraries/anchor/bubblegum/generated/bubblegum"
	"solana_project/libraries/anchor/lmnft_compressed/generated/compressed_minter"
	"solana_project/libraries/programs"
	"solana_project/services/mints-tracker/pkg/models"
	"strings"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"mvdan.cc/xurls/v2"
)

var (
	rxUrl *regexp.Regexp
)

func init() {
	rxUrl = xurls.Strict()
}

func (tp *TransactionParser) parseLmnftCompressedInstruction(
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
	decoded, err := compressed_minter.DecodeInstruction(accounts, instruction.Data)
	if err != nil {
		return nil, err
	}
	mintV2, ok := decoded.Impl.(*compressed_minter.MintV2)
	if ok {
		result, err := tp.parseLmnftCompressedInstructionMintV2(ctx, sig, accounts, transactionResult, tx, mintV2, launchpad)
		return result, err
	}

	mintCore, ok := decoded.Impl.(*compressed_minter.MintCore)
	if ok {
		result, err := tp.parseLmnftCoreInstruction(ctx, sig, transactionResult, mintCore, launchpad)
		return result, err
	}

	return nil, fmt.Errorf("unknown mint instruction")
}

func (tp *TransactionParser) parseLmnftCompressedInstructionMintV2(
	ctx context.Context,
	sig solana.Signature,
	accounts []*solana.AccountMeta,
	transactionResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
	mintV2 *compressed_minter.MintV2,
	launchpad string,
) (*FullMintData, error) {
	minterAccount := mintV2.GetMinterAccount()
	minterAccountDataRaw, err := tp.rpcClient.GetAccountInfo(ctx, minterAccount.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get minter account data; err=%s", err.Error())
	}
	var minterAccountData compressed_minter.Minter
	err = minterAccountData.UnmarshalWithDecoder(bin.NewBorshDecoder(minterAccountDataRaw.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("failed to decode minter account data; err=%s", err.Error())
	}
	assetData, err := tp.retrieveAssetDataFromBubblegumInstruction(accounts, transactionResult, tx)
	if err != nil {
		return nil, err
	}

	curTime := time.Now().Unix()
	var curPrice float64
	for _, salePhase := range minterAccountData.SalePhases {
		if salePhase.Start < curTime {
			curPrice = float64(salePhase.Price) / float64(solana.LAMPORTS_PER_SOL)
		}
	}

	blockTime := solana.UnixTimeSeconds(time.Now().Unix())
	if transactionResult.BlockTime != nil {
		blockTime = *transactionResult.BlockTime
	}

	ipfsData, err := tp.ipfsClient.Get(assetData.AssetIpfsUri)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from ipfs; err=%s", err.Error())
	}

	collectionName := strings.TrimSpace(strings.Split(minterAccountData.Name, "#")[0])

	return &FullMintData{
		Mint: Mint{
			TxSignature:   sig,
			ProgramKey:    programs.LMNFTBubblegumProgramID,
			BlockTime:     blockTime,
			CollectionKey: mintV2.GetCollectionMintAccount().PublicKey,
			AssetKey:      assetData.AssetId,
			AssetOwnerKey: mintV2.GetBuyerAccount().PublicKey,
			MintPriceSol:  curPrice,
			IpfsData:      *ipfsData,
		},
		Collection: CollectionWithMeta{
			CollectionKey: mintV2.GetCollectionMintAccount().PublicKey,
			CandyMachine:  minterAccount.PublicKey.ToPointer(),
			Name:          collectionName,
			Symbol:        minterAccountData.Symbol,
			Image:         ipfsData.Image,
			Size:          uint64(minterAccountData.Supply),
			Launchpad:     launchpad,
			Type:          models.CollectionTypeCNFT,
		},
	}, nil
}

func (tp *TransactionParser) parseLmnftCoreInstruction(
	ctx context.Context,
	sig solana.Signature,
	transactionResult *rpc.GetTransactionResult,
	mintCore *compressed_minter.MintCore,
	launchpad string,
) (*FullMintData, error) {
	blockTime := solana.UnixTimeSeconds(time.Now().Unix())
	if transactionResult.BlockTime != nil {
		blockTime = *transactionResult.BlockTime
	}
	machineAccount := mintCore.GetMachineAccount().PublicKey
	machineAccountInfoRaw, err := tp.rpcClient.GetAccountInfo(ctx, machineAccount)
	if err != nil {
		return nil, fmt.Errorf("failed to get machine account info; err=%s", err.Error())
	}
	var machineData compressed_minter.MintMachine
	err = machineData.UnmarshalWithDecoder(bin.NewBorshDecoder(machineAccountInfoRaw.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("failed to decode machine data; err=%s", err.Error())
	}

	curTime := time.Now().Unix()
	var curPrice float64
	for _, salePhase := range machineData.SalePhases {
		if salePhase.Start < curTime {
			curPrice = float64(salePhase.Price) / float64(solana.LAMPORTS_PER_SOL)
		}
	}

	assetIpfsUrl, err := tp.retrieveAssetIpfsUrlFromLogs(transactionResult.Meta.LogMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ipfs url from logs; err=%s", err.Error())
	}
	assetIpfsData, err := tp.ipfsClient.Get(assetIpfsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset ipfs data; err=%s", err.Error())
	}

	collectionIpfsUrl, err := tp.retrieveCollectionIpfsUrlFromLogs(transactionResult.Meta.LogMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ipfs url from logs; err=%s", err.Error())
	}
	collectionIpfsData, err := tp.ipfsClient.Get(collectionIpfsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset ipfs data; err=%s", err.Error())
	}

	return &FullMintData{
		Mint: Mint{
			TxSignature:   sig,
			ProgramKey:    programs.LMNFTBubblegumProgramID,
			BlockTime:     blockTime,
			CollectionKey: mintCore.GetCollectionAccount().PublicKey,
			AssetKey:      mintCore.GetAssetAccount().PublicKey,
			AssetOwnerKey: mintCore.GetBuyerAccount().PublicKey,
			MintPriceSol:  curPrice,
			IpfsData:      *assetIpfsData,
		},
		Collection: CollectionWithMeta{
			CollectionKey: mintCore.GetCollectionAccount().PublicKey,
			CandyMachine:  mintCore.GetMachineAccount().PublicKey.ToPointer(),
			Name:          collectionIpfsData.Name,
			Symbol:        collectionIpfsData.Symbol,
			Image:         collectionIpfsData.Image,
			Size:          uint64(machineData.Data.Supply),
			Launchpad:     launchpad,
			Type:          models.CollectionTypeCore,
		},
	}, nil
}

type assetData struct {
	AssetId      solana.PublicKey
	AssetIpfsUri string
}

func (tp *TransactionParser) retrieveAssetDataFromBubblegumInstruction(
	accounts []*solana.AccountMeta,
	transactionResult *rpc.GetTransactionResult,
	tx *solana.Transaction,
) (*assetData, error) {
	for _, innerInstruction := range transactionResult.Meta.InnerInstructions {
		for _, instructionFromInner := range innerInstruction.Instructions {
			programId := tx.Message.AccountKeys[instructionFromInner.ProgramIDIndex]
			if programId.Equals(programs.BubblegumProgramId) {
				decodedInst, err := bubblegum.DecodeInstruction(accounts, instructionFromInner.Data)
				if err != nil {
					return nil, fmt.Errorf("failed to decode bubblegum instruction; err=%s", err.Error())
				}
				mintToCollectionV1, ok := decodedInst.Impl.(*bubblegum.MintToCollectionV1)
				if !ok {
					return nil, fmt.Errorf("instruction is not mintToCollectionV1")
				}
				changeLogEvents := retrieveChangeLogEventsFromInstructions(innerInstruction.Instructions)
				if len(changeLogEvents) == 0 {
					return nil, fmt.Errorf("expected to get at least single changeLogEvent; got len=%d", len(changeLogEvents))
				}
				changeLogEvent := changeLogEvents[len(changeLogEvents)-1]
				assetId, err := getLeafAssetId(changeLogEvent.V1.Id, changeLogEvent.V1.Index)
				if err != nil {
					return nil, err
				}
				return &assetData{
					AssetId:      assetId,
					AssetIpfsUri: mintToCollectionV1.MetadataArgs.Uri,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("AssetId not found")
}

func (tp *TransactionParser) retrieveCollectionIpfsUrlFromLogs(
	logs []string,
) (string, error) {
	for _, log := range logs {
		if !strings.Contains(log, "Collection: CollectionV1") {
			continue
		}
		u := rxUrl.FindString(log)
		if u == "" {
			return "", fmt.Errorf("failed to retrieve url from log")
		}
		return u, nil
	}
	return "", fmt.Errorf("ipfs url not found")
}

func (tp *TransactionParser) retrieveAssetIpfsUrlFromLogs(
	logs []string,
) (string, error) {
	for _, log := range logs {
		if !strings.Contains(log, "Program data:") {
			continue
		}
		uriSplit := strings.Split(log, "Program data:")
		uriSplitLast := uriSplit[len(uriSplit)-1]
		data := strings.TrimSpace(strings.Split(uriSplitLast, ",")[0])
		b, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return "", err
		}
		sB := string(b)
		res := rxUrl.FindString(sB)
		if res == "" {
			return "", fmt.Errorf("failed to retrieve url from log")
		}
		return res, nil
	}
	return "", fmt.Errorf("ipfs url not found")
}
