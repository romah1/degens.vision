package collection_updater

import (
	"context"
	"net/http"
	clients "solana_project/libraries/clients/lmnft"
	"solana_project/services/mints-tracker/pkg/collection_updater"
	"solana_project/services/mints-tracker/pkg/models"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var collectionUpdater *collection_updater.CollectionUpdater

func init() {
	lmnftClient := clients.NewLmnftClient(&http.Client{})
	logger := zap.Must(zap.NewProduction()).Sugar()
	rpcClient := rpc.New(rpc.MainNetBeta_RPC)
	collectionUpdater = collection_updater.NewCollectionUpdater(
		logger,
		rpcClient,
		&pgxpool.Pool{},
		lmnftClient,
	)
}

func TestUpdate(t *testing.T) {
	collectionKey := solana.MustPublicKeyFromBase58("8HhShEpX4NnafoPGKo1wNoCXHRN2ot2g9C63WJ8uMwas")
	dataForUpdate, err := collectionUpdater.UpdateCollection(context.TODO(), &collection_updater.CollectionForUpdate{
		CollectionKey: collectionKey,
		CandyMachine:  solana.MustPublicKeyFromBase58("CrnjzhTzEYHwyA1ByWEBvR9ATdrGupDBGhw7tMuJSSMK").ToPointer(),
		Launchpad:     models.LaunchpadLmnft,
		Type:          "mplx",
		Name:          "Medieval Hunt Genesis Access",
	})
	assert.Equal(t, nil, err, "update failed")
	spew.Dump(dataForUpdate)
}
