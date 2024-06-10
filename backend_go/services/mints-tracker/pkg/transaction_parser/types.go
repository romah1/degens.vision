package transaction_parser

import (
	"solana_project/services/mints-tracker/pkg/ipfs_client"
	"solana_project/services/mints-tracker/pkg/models"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type FullMintData struct {
	Mint       Mint
	Collection CollectionWithMeta
}

type Mint struct {
	TxSignature   solana.Signature
	ProgramKey    solana.PublicKey
	BlockTime     solana.UnixTimeSeconds
	CollectionKey solana.PublicKey
	MintPriceSol  float64

	AssetKey      solana.PublicKey
	AssetOwnerKey solana.PublicKey
	IpfsData      ipfs_client.IpfsData
}

type CollectionWithMeta struct {
	CollectionKey solana.PublicKey
	CandyMachine  *solana.PublicKey
	Name          string
	Symbol        string
	Image         string
	Size          uint64
	Launchpad     string
	Type          models.CollectionType
}

type PathNode struct {
	Node  [32]uint8
	Index uint32
}

type ChangeLogEventV1 struct {
	Id    solana.PublicKey
	Path  []PathNode
	Seq   uint64
	Index uint32
}

type ChangeLogEvent struct {
	Enum bin.BorshEnum `borsh_enum:"true"`
	V1   ChangeLogEventV1
}

type ApplicationDataEventV1 struct {
	ApplicationData []byte
}

type AccountCompressionEvent struct {
	Enum            bin.BorshEnum `borsh_enum:"true"`
	ChangeLog       ChangeLogEvent
	ApplicationData ApplicationDataEventV1
}

type Creator struct {
	Address  solana.PublicKey
	Verified bool
	Share    uint8
}

type Data struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator `bin:"optional"`
}

type Collection struct {
	Verified bool
	Key      solana.PublicKey
}

type Uses struct {
	UseMethod bin.BorshEnum
	Remaining uint64
	Total     uint64
}

type CollectionDetailsV1 struct {
	Size uint64
}

type CollectionDetails struct {
	Enum bin.BorshEnum `borsh_enum:"true"`
	V1   CollectionDetailsV1
}

type ProgrammableConfig struct {
	Enum bin.BorshEnum `borsh_enum:"true"`
	V1   ProgrammableConfigV1
}

type ProgrammableConfigV1 struct {
	RuleSet *solana.PublicKey `bin:"optional"`
}

type Metadata struct {
	Key                 bin.BorshEnum
	UpdateAuthority     solana.PublicKey
	Mint                solana.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8              `bin:"optional"`
	TokenStandard       *bin.BorshEnum      `bin:"optional"`
	Collection          *Collection         `bin:"optional"`
	Uses                *Uses               `bin:"optional"`
	CollectionDetails   *CollectionDetails  `bin:"optional"`
	ProgrammableConfig  *ProgrammableConfig `bin:"optional"`
}

type metadataPreV11 struct {
	Key                 bin.BorshEnum
	UpdateAuthority     solana.PublicKey
	Mint                solana.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
}
