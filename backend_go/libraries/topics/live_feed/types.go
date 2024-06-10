package live_feed

import (
	"solana_project/services/mints-tracker/pkg/ipfs_client"
	"time"
)

const (
	TopicName = "live_feed"
)

type Message struct {
	Mint       MessageMint       `json:"mint"`
	Collection MessageCollection `json:"collection"`
}

type MessageMint struct {
	TxSignature   string    `json:"tx_signature"`
	ProgramKey    string    `json:"program_key"`
	BlockTime     time.Time `json:"block_time"`
	CollectionKey string    `json:"collection_key"`
	MintPrice     float64   `json:"mint_price"`
	AssetKey      string    `json:"asset_key"`
	AssetOwnerKey string    `json:"asset_owner_key"`

	AssetName        string                 `json:"asset_name"`
	AssetSymbol      string                 `json:"asset_symbol"`
	AssetImage       string                 `json:"asset_image"`
	AssetDescription string                 `json:"asset_description"`
	AssetAttributes  []MessageMintAttribute `json:"asset_attributes"`
}

type MessageMintAttribute = ipfs_client.Attribute

type MessageCollection struct {
	CollectionKey string `json:"collection_key"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Image         string `json:"image"`
	Size          uint64 `json:"size"`
	Launchpad     string `json:"launchpad"`
}
