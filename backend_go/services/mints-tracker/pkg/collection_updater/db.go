package collection_updater

import (
	"context"
	"solana_project/services/mints-tracker/pkg/models"
	"solana_project/services/mints-tracker/pkg/sql"

	"github.com/gagliardetto/solana-go"
	"github.com/jackc/pgx/v5"
)

type CollectionForUpdate struct {
	CollectionKey solana.PublicKey
	CandyMachine  *solana.PublicKey
	Launchpad     string
	Type          string
	Name          string
}

func (cu *CollectionUpdater) selectCollectionForUpdate(ctx context.Context) (*CollectionForUpdate, error) {
	row := cu.pg.QueryRow(ctx, sql.SelectCollectionForUpdate)
	var collectionKey string
	var candyMachineKey *string
	var launchpad string
	var collectionType string
	var name string
	err := row.Scan(
		&collectionKey,
		&candyMachineKey,
		&launchpad,
		&collectionType,
		&name,
	)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}
	collection, err := solana.PublicKeyFromBase58(collectionKey)
	if err != nil {
		return nil, err
	}
	var candyMachine *solana.PublicKey
	if candyMachineKey != nil {
		addr, err := solana.PublicKeyFromBase58(*candyMachineKey)
		if err != nil {
			return nil, err
		}
		candyMachine = addr.ToPointer()
	}
	return &CollectionForUpdate{
		CollectionKey: collection,
		CandyMachine:  candyMachine,
		Launchpad:     launchpad,
		Type:          collectionType,
		Name:          name,
	}, nil
}

type UpdateCollectionData struct {
	CollectionKey string
	TotalMinted   uint64
	Size          uint64
	Links         []UpdateCollectionDataLink
}

type UpdateCollectionDataLink struct {
	Type models.CollectionLinkType
	Uri  string
}

func (cu *CollectionUpdater) updateCollection(ctx context.Context, data *UpdateCollectionData) error {
	tx, err := cu.pg.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		sql.UpdateCollection,
		data.CollectionKey,
		data.TotalMinted,
		data.Size,
	)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if len(data.Links) != 0 {
		batch := pgx.Batch{}
		for _, link := range data.Links {
			batch.Queue(sql.UpsertCollectionLink, data.CollectionKey, link.Type, link.Uri)
		}
		br := tx.SendBatch(ctx, &batch)
		_, err = br.Exec()
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
		err = br.Close()
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
