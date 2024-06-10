package mint_tracker_mints_processor

import (
	"context"
	"solana_project/services/mints-tracker/pkg/sql"
	"solana_project/services/mints-tracker/pkg/transaction_parser"
)

func (mp *MintTrackerMintsProcessor) UpsertMint(
	ctx context.Context,
	data *transaction_parser.FullMintData,
) error {
	tx, err := mp.pg.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		ctx,
		sql.UpsertMintQuery,
		data.Mint.TxSignature.String(),
		data.Mint.ProgramKey.String(),
		data.Mint.BlockTime.Time(),
		data.Mint.CollectionKey.String(),
		data.Mint.AssetKey.String(),
		data.Mint.AssetOwnerKey.String(),
		data.Mint.IpfsData.Name,
		data.Mint.IpfsData.Symbol,
		data.Mint.IpfsData.Image,
		data.Mint.IpfsData.Description,
		data.Mint.IpfsData.Attributes,
		data.Mint.MintPriceSol,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	var candyMachine *string
	if data.Collection.CandyMachine != nil {
		s := data.Collection.CandyMachine.String()
		candyMachine = &s
	}
	_, err = tx.Exec(
		ctx,
		sql.UpsertCollectionQuery,
		data.Collection.CollectionKey.String(),
		candyMachine,
		data.Collection.Name,
		data.Collection.Symbol,
		data.Collection.Image,
		data.Collection.Size,
		data.Collection.Type,
		data.Mint.MintPriceSol,
		0,
		data.Collection.Launchpad,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	mp.logger.Infof("Upserted mint and collection; signature=%s", data.Mint.TxSignature.String())
	return nil
}
