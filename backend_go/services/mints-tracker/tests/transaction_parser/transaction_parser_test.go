package transaction_parser

import (
	"context"
	"encoding/json"
	"net/http"
	"solana_project/libraries/programs"
	"solana_project/libraries/queues/mint_tracker_mints"
	"solana_project/services/mints-tracker/pkg/ipfs_client"
	"solana_project/services/mints-tracker/pkg/transaction_parser"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

var transactionParser *transaction_parser.TransactionParser

func init() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	rpcClient := rpc.New(rpc.MainNetBeta_RPC)
	ipfsClient := ipfs_client.NewIpfsClient(
		[]*http.Client{&http.Client{}},
	)
	transactionParser = transaction_parser.NewTransactionParser(
		logger,
		rpcClient,
		rpc.CommitmentConfirmed,
		ipfsClient,
	)
}

func TestLmnftBubblegumMint(t *testing.T) {
	sig := "4R7YUD6jK6WQsCS77hymgDRt9ZGGJkxqHcMj2KLko4D47oHfx6kzRJ6Xf3mVg7fFXkTspM5Uo2JrdCoYDBQxtvdd"
	testParser(
		t,
		sig,
		programs.LMNFTBubblegumProgramID.String(),
		[]*transaction_parser.FullMintData{
			{
				Mint: transaction_parser.Mint{
					TxSignature:   solana.MustSignatureFromBase58(sig),
					ProgramKey:    programs.LMNFTBubblegumProgramID,
					BlockTime:     1713234266,
					CollectionKey: solana.MustPublicKeyFromBase58("2Betcy9AVFpCJzf11uigvBjPDCDwtYPGG81w15SqsiJW"),
					AssetKey:      solana.MustPublicKeyFromBase58("7exHWD6W53zQwsC4LKWVxaBDF3TAhbB5SuMKbvpWVf7o"),
					AssetOwnerKey: solana.MustPublicKeyFromBase58("Hys9bmGzQ4FnReW6CWYZRf2Lj6HGt2TQRZCbCybA2gyw"),
					IpfsData: ipfs_client.IpfsData{
						Name:        "SolRuneStone #3897",
						Symbol:      "SOLRUNESTONE",
						Image:       "https://ap-assets.pinit.io/7JnbBZWri4fCzMy2WVPV8ub7qEZvpQvEDRZLdgxZGKEd/62850252-c983-4bce-9083-994d8a62c416/0",
						Description: "The Epic RuneStone On Solana Chain.",
						Attributes:  []ipfs_client.Attribute{},
					},
					MintPriceSol: 0.1,
				},
				Collection: transaction_parser.CollectionWithMeta{
					CollectionKey: solana.MustPublicKeyFromBase58("2Betcy9AVFpCJzf11uigvBjPDCDwtYPGG81w15SqsiJW"),
					CandyMachine:  solana.MustPublicKeyFromBase58("CJaMtS5NgfmAyzVdFZCrrThQqXGu7zYU28VcBYUzAW2n").ToPointer(),
					Name:          "SolRuneStone",
					Symbol:        "SOLRUNESTO",
					Image:         "https://ap-assets.pinit.io/7JnbBZWri4fCzMy2WVPV8ub7qEZvpQvEDRZLdgxZGKEd/62850252-c983-4bce-9083-994d8a62c416/0",
					Size:          10000,
					Launchpad:     "lmnft",
					Type:          "cnft",
				},
			},
		},
	)
}

func TestLmnftCoreMint(t *testing.T) {
	sig := "5g8VWDTRX7GvxekANQq8dHyw4cfyhz6qY6H92VYP1HV5Zk27Ssm6utyVnEexBoYufEvNCEFE8R4HJbCMCV6PvRW"
	testParser(
		t,
		sig,
		programs.LMNFTBubblegumProgramID.String(),
		[]*transaction_parser.FullMintData{
			{
				Mint: transaction_parser.Mint{
					TxSignature:   solana.MustSignatureFromBase58(sig),
					ProgramKey:    programs.LMNFTBubblegumProgramID,
					BlockTime:     1713208768,
					CollectionKey: solana.MustPublicKeyFromBase58("DJLJqxwvVVgN9DcR5sjPZ8DmzJUi3pTYcfSMwqjBWh3n"),
					AssetKey:      solana.MustPublicKeyFromBase58("8quyRuPggV4H834Q9irAVRBtwuQBfcSeSthsScTFg7gn"),
					AssetOwnerKey: solana.MustPublicKeyFromBase58("xYpqNxyYrydQdacjGmfCNdCU6BsrPhBwhSjbgfmxLQ7"),
					IpfsData: ipfs_client.IpfsData{
						Name:        "PIXEL SAMOYED #0",
						Symbol:      "PXLSAMO",
						Image:       "https://we-assets.pinit.io/xYpqNxyYrydQdacjGmfCNdCU6BsrPhBwhSjbgfmxLQ7/9c945966-c584-49ed-b843-f0bec8584a0a/0",
						Description: "Your favorite samoyeds, but now pixelated.\n\nHave you ever wanted to have a good boy? This is your chance, full free! 3333 supply",
						Attributes: []ipfs_client.Attribute{
							{
								TraitType: "Accessory",
								Value:     json.RawMessage("\"Horns Green\""),
							},
							{
								TraitType: "Background",
								Value:     json.RawMessage("\"Black\""),
							},
							{
								TraitType: "Face",
								Value:     json.RawMessage("\"Round Yellow\""),
							},
							{
								TraitType: "Left Eye",
								Value:     json.RawMessage("\"X Blue\""),
							},
							{
								TraitType: "Mouth",
								Value:     json.RawMessage("\"Grin Green\""),
							},
							{
								TraitType: "Right Eye",
								Value:     json.RawMessage("\"X Green\""),
							},
							{
								TraitType: "Rarity Rank",
								Value:     json.RawMessage("9"),
							},
						},
					},
					MintPriceSol: 0,
				},
				Collection: transaction_parser.CollectionWithMeta{
					CollectionKey: solana.MustPublicKeyFromBase58("DJLJqxwvVVgN9DcR5sjPZ8DmzJUi3pTYcfSMwqjBWh3n"),
					CandyMachine:  solana.MustPublicKeyFromBase58("4BJUPi8K6ZZsA78mdJRnRaCRXfLevHC9B5nTMGE4DL2V").ToPointer(),
					Name:          "test",
					Symbol:        "test",
					Image:         "https://we-assets.pinit.io/xYpqNxyYrydQdacjGmfCNdCU6BsrPhBwhSjbgfmxLQ7/9c945966-c584-49ed-b843-f0bec8584a0a/0",
					Size:          1,
					Launchpad:     "lmnft",
					Type:          "core",
				},
			},
		},
	)
}

func TestTruffleMint(t *testing.T) {
	sig := "2kJaUp54X1vLm6r5iEH53PFNmp8haax2xWPm3zBoyD1PuUfhX6TYanb62CWP5DmovLrBbWcHaftdsiFgihH99hUB"
	testParser(
		t,
		sig,
		programs.TruffleProgramId.String(),
		[]*transaction_parser.FullMintData{
			{
				Mint: transaction_parser.Mint{
					TxSignature:   solana.MustSignatureFromBase58(sig),
					ProgramKey:    programs.TruffleProgramId,
					BlockTime:     1709238210,
					CollectionKey: solana.MustPublicKeyFromBase58("hwepzH43NBgYkUs8cHkzcB8K7PxTvUgg88faWrEezpp"),
					AssetKey:      solana.MustPublicKeyFromBase58("Edpc4y9yQKR51s3SRK1dnDo9kQPKW34asWKnAhJKkM6S"),
					AssetOwnerKey: solana.MustPublicKeyFromBase58("5EvQJvjke8wpz9tcq57gcSh135gwQpavVSE1yjS8jEwY"),
					IpfsData: ipfs_client.IpfsData{
						Name:        "epoch #2082",
						Symbol:      "",
						Image:       "https://nftstorage.link/ipfs/bafybeihzar5ejflxq32dhxowrlkxg5epbkzkj6qw3evusdkf6vwnbsss5i/2082.png",
						Description: "A love letter to those that dared to move against the tide. An homage to the few that stood the test of time. The Thinkers. The Creators. The Founders. The Builders. The Artists. The Collectors. The Traders. The Contributors. The Validators. Today they stand together, side by side, on the brink of a rebirth…a Renaissance if you will. A shift from dark to light. Solana is our home, and our time is now. The belief of epoch is simple. Keep pushing forward",
						Attributes: []ipfs_client.Attribute{
							{
								TraitType: "Faction",
								Value:     json.RawMessage("\"Cognitive\""),
							},
							{
								TraitType: "Background",
								Value:     json.RawMessage("\"Complex\""),
							},
							{
								TraitType: "Head",
								Value:     json.RawMessage("\"Observe\""),
							},
							{
								TraitType: "Body",
								Value:     json.RawMessage("\"Simple\""),
							},
						},
					},
					MintPriceSol: 0.36001,
				},
				Collection: transaction_parser.CollectionWithMeta{
					CollectionKey: solana.MustPublicKeyFromBase58("hwepzH43NBgYkUs8cHkzcB8K7PxTvUgg88faWrEezpp"),
					Name:          "epoch.001",
					Symbol:        "",
					Image:         "https://nftstorage.link/ipfs/bafybeihzar5ejflxq32dhxowrlkxg5epbkzkj6qw3evusdkf6vwnbsss5i/2082.png",
					Size:          2550,
					Launchpad:     "truffle",
					Type:          "mplx",
				},
			},
		},
	)
}

func TestLMNFTMintV6(t *testing.T) {
	sig := "3xPADYwfYgnq77eHavRfvrKvHxe22LKRUfVoKr4cmgwbA3rNwxF37XykGihJVi7xLv83WsHb9kbiXJPfifyCHNDt"
	testParser(
		t,
		sig,
		programs.LMNFTProgramID.String(),
		[]*transaction_parser.FullMintData{
			{
				Mint: transaction_parser.Mint{
					TxSignature:   solana.MustSignatureFromBase58(sig),
					ProgramKey:    programs.LMNFTProgramID,
					BlockTime:     1712343231,
					CollectionKey: solana.MustPublicKeyFromBase58("3tGnT5X1AcviFtPazrKMmbF44cFCu88QHciF9bCoh8CF"),
					AssetKey:      solana.MustPublicKeyFromBase58("FPfxyXG7bifk6P9ESjWYxi9yhFjMNrhUZKrB8h429KSj"),
					AssetOwnerKey: solana.MustPublicKeyFromBase58("ENVk4qv4zB9ucQ131CfLucjBwepGQU8huVHiSonoKUss"),
					IpfsData: ipfs_client.IpfsData{
						Name:        "Demon #310",
						Symbol:      "DT",
						Image:       "https://we-assets.pinit.io/68LGoEjrenD51ukdWBjbHynrivNtwW4kV5CBisYfu29y/a705d6c2-95e4-4705-ad7c-9d8b29795950/1390",
						Description: "Demon Town is an NFT collection living, building and growing in Solana.",
						Attributes: []ipfs_client.Attribute{
							{
								TraitType: "Background",
								Value:     json.RawMessage("\"White Hypnosis\""),
							},
							{
								TraitType: "Body",
								Value:     json.RawMessage("\"Snow\""),
							},
							{
								TraitType: "Clothes",
								Value:     json.RawMessage("\"Overalls\""),
							},
							{
								TraitType: "Mouth",
								Value:     json.RawMessage("\"Thoughtful\""),
							},
							{
								TraitType: "Accessories",
								Value:     json.RawMessage("\"Fire Earring\""),
							},
							{
								TraitType: "Eyes",
								Value:     json.RawMessage("\"Unhappy\""),
							},
							{
								TraitType: "Headwear",
								Value:     json.RawMessage("\"Summer Cap\""),
							},
							{
								TraitType: "Effects",
								Value:     json.RawMessage("\"None\""),
							},
							{
								TraitType: "Rarity Rank",
								Value:     json.RawMessage("620"),
							},
						},
					},
					MintPriceSol: 0.03,
				},
				Collection: transaction_parser.CollectionWithMeta{
					CollectionKey: solana.MustPublicKeyFromBase58("3tGnT5X1AcviFtPazrKMmbF44cFCu88QHciF9bCoh8CF"),
					CandyMachine:  solana.MustPublicKeyFromBase58("3GeSkpAd6Uz9owXZmwsKxU5Rihu3MxKCDPqa62pkm1Do").ToPointer(),
					Name:          "Demon Town",
					Symbol:        "DT",
					Image:         "https://we-assets.pinit.io/68LGoEjrenD51ukdWBjbHynrivNtwW4kV5CBisYfu29y/a705d6c2-95e4-4705-ad7c-9d8b29795950/1390",
					Size:          2000,
					Launchpad:     "lmnft",
					Type:          "mplx",
				},
			},
		},
	)
}

func testParser(
	t *testing.T,
	sig string,
	programId string,
	expMintData []*transaction_parser.FullMintData,
) {
	result, err := transactionParser.RetrieveMintData(
		context.Background(),
		&mint_tracker_mints.Message{
			ProgramId: programId,
			Signature: sig,
		},
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(expMintData), len(result))
	assert.Equal(t, expMintData, result)
}
