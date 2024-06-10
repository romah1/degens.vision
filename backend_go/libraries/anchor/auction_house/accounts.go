// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type BidReceipt struct {
	TradeState      ag_solanago.PublicKey
	Bookkeeper      ag_solanago.PublicKey
	AuctionHouse    ag_solanago.PublicKey
	Buyer           ag_solanago.PublicKey
	Metadata        ag_solanago.PublicKey
	TokenAccount    *ag_solanago.PublicKey `bin:"optional"`
	PurchaseReceipt *ag_solanago.PublicKey `bin:"optional"`
	Price           uint64
	TokenSize       uint64
	Bump            uint8
	TradeStateBump  uint8
	CreatedAt       int64
	CanceledAt      *int64 `bin:"optional"`
}

var BidReceiptDiscriminator = [8]byte{186, 150, 141, 135, 59, 122, 39, 99}

func (obj BidReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BidReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TradeState` param:
	err = encoder.Encode(obj.TradeState)
	if err != nil {
		return err
	}
	// Serialize `Bookkeeper` param:
	err = encoder.Encode(obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Serialize `AuctionHouse` param:
	err = encoder.Encode(obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Serialize `Buyer` param:
	err = encoder.Encode(obj.Buyer)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `TokenAccount` param (optional):
	{
		if obj.TokenAccount == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenAccount)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PurchaseReceipt` param (optional):
	{
		if obj.PurchaseReceipt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PurchaseReceipt)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `TokenSize` param:
	err = encoder.Encode(obj.TokenSize)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `TradeStateBump` param:
	err = encoder.Encode(obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	// Serialize `CanceledAt` param (optional):
	{
		if obj.CanceledAt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.CanceledAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *BidReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BidReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[186 150 141 135 59 122 39 99]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TradeState`:
	err = decoder.Decode(&obj.TradeState)
	if err != nil {
		return err
	}
	// Deserialize `Bookkeeper`:
	err = decoder.Decode(&obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Deserialize `AuctionHouse`:
	err = decoder.Decode(&obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Deserialize `Buyer`:
	err = decoder.Decode(&obj.Buyer)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `TokenAccount` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenAccount)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PurchaseReceipt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PurchaseReceipt)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `TokenSize`:
	err = decoder.Decode(&obj.TokenSize)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `TradeStateBump`:
	err = decoder.Decode(&obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	// Deserialize `CanceledAt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.CanceledAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ListingReceipt struct {
	TradeState      ag_solanago.PublicKey
	Bookkeeper      ag_solanago.PublicKey
	AuctionHouse    ag_solanago.PublicKey
	Seller          ag_solanago.PublicKey
	Metadata        ag_solanago.PublicKey
	PurchaseReceipt *ag_solanago.PublicKey `bin:"optional"`
	Price           uint64
	TokenSize       uint64
	Bump            uint8
	TradeStateBump  uint8
	CreatedAt       int64
	CanceledAt      *int64 `bin:"optional"`
}

var ListingReceiptDiscriminator = [8]byte{240, 71, 225, 94, 200, 75, 84, 231}

func (obj ListingReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ListingReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TradeState` param:
	err = encoder.Encode(obj.TradeState)
	if err != nil {
		return err
	}
	// Serialize `Bookkeeper` param:
	err = encoder.Encode(obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Serialize `AuctionHouse` param:
	err = encoder.Encode(obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Serialize `Seller` param:
	err = encoder.Encode(obj.Seller)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `PurchaseReceipt` param (optional):
	{
		if obj.PurchaseReceipt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PurchaseReceipt)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `TokenSize` param:
	err = encoder.Encode(obj.TokenSize)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `TradeStateBump` param:
	err = encoder.Encode(obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	// Serialize `CanceledAt` param (optional):
	{
		if obj.CanceledAt == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.CanceledAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *ListingReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ListingReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[240 71 225 94 200 75 84 231]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TradeState`:
	err = decoder.Decode(&obj.TradeState)
	if err != nil {
		return err
	}
	// Deserialize `Bookkeeper`:
	err = decoder.Decode(&obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Deserialize `AuctionHouse`:
	err = decoder.Decode(&obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Deserialize `Seller`:
	err = decoder.Decode(&obj.Seller)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `PurchaseReceipt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PurchaseReceipt)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `TokenSize`:
	err = decoder.Decode(&obj.TokenSize)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `TradeStateBump`:
	err = decoder.Decode(&obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	// Deserialize `CanceledAt` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.CanceledAt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type PurchaseReceipt struct {
	Bookkeeper   ag_solanago.PublicKey
	Buyer        ag_solanago.PublicKey
	Seller       ag_solanago.PublicKey
	AuctionHouse ag_solanago.PublicKey
	Metadata     ag_solanago.PublicKey
	TokenSize    uint64
	Price        uint64
	Bump         uint8
	CreatedAt    int64
}

var PurchaseReceiptDiscriminator = [8]byte{79, 127, 222, 137, 154, 131, 150, 134}

func (obj PurchaseReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PurchaseReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bookkeeper` param:
	err = encoder.Encode(obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Serialize `Buyer` param:
	err = encoder.Encode(obj.Buyer)
	if err != nil {
		return err
	}
	// Serialize `Seller` param:
	err = encoder.Encode(obj.Seller)
	if err != nil {
		return err
	}
	// Serialize `AuctionHouse` param:
	err = encoder.Encode(obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `TokenSize` param:
	err = encoder.Encode(obj.TokenSize)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PurchaseReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PurchaseReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[79 127 222 137 154 131 150 134]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bookkeeper`:
	err = decoder.Decode(&obj.Bookkeeper)
	if err != nil {
		return err
	}
	// Deserialize `Buyer`:
	err = decoder.Decode(&obj.Buyer)
	if err != nil {
		return err
	}
	// Deserialize `Seller`:
	err = decoder.Decode(&obj.Seller)
	if err != nil {
		return err
	}
	// Deserialize `AuctionHouse`:
	err = decoder.Decode(&obj.AuctionHouse)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `TokenSize`:
	err = decoder.Decode(&obj.TokenSize)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

type AuctionHouse struct {
	AuctionHouseFeeAccount        ag_solanago.PublicKey
	AuctionHouseTreasury          ag_solanago.PublicKey
	TreasuryWithdrawalDestination ag_solanago.PublicKey
	FeeWithdrawalDestination      ag_solanago.PublicKey
	TreasuryMint                  ag_solanago.PublicKey
	Authority                     ag_solanago.PublicKey
	Creator                       ag_solanago.PublicKey
	Bump                          uint8
	TreasuryBump                  uint8
	FeePayerBump                  uint8
	SellerFeeBasisPoints          uint16
	RequiresSignOff               bool
	CanChangeSalePrice            bool
}

var AuctionHouseDiscriminator = [8]byte{40, 108, 215, 107, 213, 85, 245, 48}

func (obj AuctionHouse) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AuctionHouseDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `AuctionHouseFeeAccount` param:
	err = encoder.Encode(obj.AuctionHouseFeeAccount)
	if err != nil {
		return err
	}
	// Serialize `AuctionHouseTreasury` param:
	err = encoder.Encode(obj.AuctionHouseTreasury)
	if err != nil {
		return err
	}
	// Serialize `TreasuryWithdrawalDestination` param:
	err = encoder.Encode(obj.TreasuryWithdrawalDestination)
	if err != nil {
		return err
	}
	// Serialize `FeeWithdrawalDestination` param:
	err = encoder.Encode(obj.FeeWithdrawalDestination)
	if err != nil {
		return err
	}
	// Serialize `TreasuryMint` param:
	err = encoder.Encode(obj.TreasuryMint)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Creator` param:
	err = encoder.Encode(obj.Creator)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `TreasuryBump` param:
	err = encoder.Encode(obj.TreasuryBump)
	if err != nil {
		return err
	}
	// Serialize `FeePayerBump` param:
	err = encoder.Encode(obj.FeePayerBump)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `RequiresSignOff` param:
	err = encoder.Encode(obj.RequiresSignOff)
	if err != nil {
		return err
	}
	// Serialize `CanChangeSalePrice` param:
	err = encoder.Encode(obj.CanChangeSalePrice)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AuctionHouse) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AuctionHouseDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[40 108 215 107 213 85 245 48]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `AuctionHouseFeeAccount`:
	err = decoder.Decode(&obj.AuctionHouseFeeAccount)
	if err != nil {
		return err
	}
	// Deserialize `AuctionHouseTreasury`:
	err = decoder.Decode(&obj.AuctionHouseTreasury)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryWithdrawalDestination`:
	err = decoder.Decode(&obj.TreasuryWithdrawalDestination)
	if err != nil {
		return err
	}
	// Deserialize `FeeWithdrawalDestination`:
	err = decoder.Decode(&obj.FeeWithdrawalDestination)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryMint`:
	err = decoder.Decode(&obj.TreasuryMint)
	if err != nil {
		return err
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Creator`:
	err = decoder.Decode(&obj.Creator)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryBump`:
	err = decoder.Decode(&obj.TreasuryBump)
	if err != nil {
		return err
	}
	// Deserialize `FeePayerBump`:
	err = decoder.Decode(&obj.FeePayerBump)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `RequiresSignOff`:
	err = decoder.Decode(&obj.RequiresSignOff)
	if err != nil {
		return err
	}
	// Deserialize `CanChangeSalePrice`:
	err = decoder.Decode(&obj.CanChangeSalePrice)
	if err != nil {
		return err
	}
	return nil
}
