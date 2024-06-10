package transaction_parser

import (
	"context"
	"fmt"
	"solana_project/libraries/programs"
	"strings"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	lookup "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/mr-tron/base58"
)

// func (mp *MintTrackerMintsProcessor) processLMNFT(
// 	ctx context.Context,
// 	payload *mint_tracker_mints.MessagePayloadLMNFT,
// ) (*models.MintData, error) {
// 	sig, err := solana.SignatureFromBase58(payload.Signature)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var maxSupportedTransactionVersion uint64 = 0
// 	transactionResult, err := mp.rpcClient.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
// 		Commitment:                     mp.commitment,
// 		Encoding:                       solana.EncodingBase64,
// 		MaxSupportedTransactionVersion: &maxSupportedTransactionVersion,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(transactionResult.Transaction.GetBinary()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = mp.resolveLookups(tx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, instruction := range tx.Message.Instructions {
// 		accounts, err := instruction.ResolveInstructionAccounts(&tx.Message)
// 		if err != nil {
// 			return nil, err
// 		}
// 		decoded, err := nft_candy_machine.DecodeInstruction(accounts, instruction.Data)
// 		if err != nil {
// 			continue
// 		}
// 		mintV6, ok := decoded.Impl.(*nft_candy_machine.MintV6)
// 		if ok {
// 			return &models.MintData{
// 				Signature: sig,
// 				Mint:      mintV6.GetMintAccount().PublicKey,
// 				Payer:     mintV6.GetPayerAccount().PublicKey,
// 				CandyMachine: mintV6.GetCandyMachineAccount().PublicKey,
// 				Program: tx.Message.AccountKeys[instruction.ProgramIDIndex],
// 				IpfsURL:   payload.IpfsURL,
// 				MintedAt: transactionResult.BlockTime.Time(),
// 			}, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("failed to find mint instruction")
// }

func (tp *TransactionParser) resolveLookups(txx *solana.Transaction) error {
	if !txx.Message.IsVersioned() {
		// tp.logger.Warn("tx is not versioned; only versioned transactions can contain lookups")
		return nil
	}
	tblKeys := txx.Message.GetAddressTableLookups().GetTableIDs()
	if len(tblKeys) == 0 {
		// tp.logger.Warn("no lookup tables in versioned transaction")
		return nil
	}
	numLookups := txx.Message.GetAddressTableLookups().NumLookups()
	if numLookups == 0 {
		tp.logger.Warn("no lookups in versioned transaction")
		return nil
	}
	resolutions := make(map[solana.PublicKey]solana.PublicKeySlice)
	for _, key := range tblKeys {
		// tp.logger.Info("Getting table " + key.String())

		info, err := tp.rpcClient.GetAccountInfo(
			context.Background(),
			key,
		)
		if err != nil {
			return err
		}
		// tp.logger.Info("got table " + key.String())

		tableContent, err := lookup.DecodeAddressLookupTableState(info.GetBinary())
		if err != nil {
			return err
		}

		resolutions[key] = tableContent.Addresses
	}

	err := txx.Message.SetAddressTables(resolutions)
	if err != nil {
		return err
	}

	err = txx.Message.ResolveLookups()
	if err != nil {
		return err
	}
	return nil
}

func retrieveChangeLogEventsFromInstructions(instructions []solana.CompiledInstruction) []ChangeLogEvent {
	var events []ChangeLogEvent
	for _, instruction := range instructions {
		acEvent, err := decodeAccountCompressionEvent(instruction.Data)
		if err != nil {
			continue
		}
		if acEvent.Enum != 0 {
			continue
		}
		events = append(events, acEvent.ChangeLog)
	}
	return events
}

func decodeAccountCompressionEvent(data solana.Base58) (*AccountCompressionEvent, error) {
	binary, err := base58.Decode(data.String())
	if err != nil {
		return nil, err
	}
	acEvent := AccountCompressionEvent{}
	decoder := bin.NewBorshDecoder(binary)
	err = decoder.Decode(&acEvent.Enum)
	if err != nil {
		return nil, err
	}
	switch acEvent.Enum {
	case 0:
		err = decoder.Decode(&acEvent.ChangeLog)
		if err != nil {
			return nil, err
		}
	case 1:
		err = decoder.Decode(&acEvent.ApplicationData)
		if err != nil {
			return nil, err
		}
	}

	return &acEvent, nil
}

func MetadataDeserialize(data []byte) (Metadata, error) {
	var metadata Metadata
	dec := bin.NewBorshDecoder(data)
	err := dec.Decode(&metadata)
	// if err != nil {
	// 	return Metadata{}, fmt.Errorf("failed to decode data; error=%s", err.Error())
	// }
	if err != nil {
		// https://github.com/samuelvanderwaal/metaboss/issues/121
		// https://github.com/metaplex-foundation/metaplex-program-library/pull/407
		// C.f. https://github.com/metaplex-foundation/metaplex-program-library/blob/master/token-metadata/program/src/deser.rs#L12
		var metadataPreV11 metadataPreV11
		decoder := bin.NewBorshDecoder(data)
		err := decoder.Decode(&metadataPreV11)
		if err != nil {
			return Metadata{}, fmt.Errorf("failed to deserialize data, err: %v", err)
		} else {
			metadata.Key = metadataPreV11.Key
			metadata.UpdateAuthority = metadataPreV11.UpdateAuthority
			metadata.Mint = metadataPreV11.Mint
			metadata.Data = metadataPreV11.Data
			// metadata.PrimarySaleHappened = metadataPreV11.PrimarySaleHappened
			// metadata.IsMutable = metadataPreV11.IsMutable
			// metadata.EditionNonce = metadataPreV11.EditionNonce
		}
	}
	// trim null byte
	metadata.Data.Name = strings.TrimRight(metadata.Data.Name, "\x00")
	// if metadata.Data.Name == "" {
	// 	return Metadata{}, fmt.Errorf("got empty metadata name")
	// }
	metadata.Data.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
	metadata.Data.Uri = strings.TrimRight(metadata.Data.Uri, "\x00")
	return metadata, nil
}

func trimNullBytes(s string) string {
	return strings.Trim(s, "\x00")
}

func RetrieveMetadataForAccount(
	ctx context.Context,
	rpcClient *rpc.Client,
	address solana.PublicKey,
) (*Metadata, error) {
	metadataAccount, err := getMetadataAccount(address)
	if err != nil {
		return nil, err
	}
	out, err := rpcClient.GetAccountInfo(ctx, metadataAccount)
	if err != nil {
		return nil, err
	}
	metadata, err := MetadataDeserialize(out.Bytes())
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func getMetadataAccount(
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
