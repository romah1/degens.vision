// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package compressed_minter

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type CreateMinterInput struct {
	Ordered              bool
	Name                 string
	Url                  string
	Creators             []Creator
	SellerFeeBasisPoints uint16
	Supply               uint32
	Collection           ag_solanago.PublicKey
	SalePhases           []SaleFaze
	Symbol               string
}

func (obj CreateMinterInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Ordered` param:
	err = encoder.Encode(obj.Ordered)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Url` param:
	err = encoder.Encode(obj.Url)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	// Serialize `Collection` param:
	err = encoder.Encode(obj.Collection)
	if err != nil {
		return err
	}
	// Serialize `SalePhases` param:
	err = encoder.Encode(obj.SalePhases)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CreateMinterInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Ordered`:
	err = decoder.Decode(&obj.Ordered)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Url`:
	err = decoder.Decode(&obj.Url)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	// Deserialize `Collection`:
	err = decoder.Decode(&obj.Collection)
	if err != nil {
		return err
	}
	// Deserialize `SalePhases`:
	err = decoder.Decode(&obj.SalePhases)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	return nil
}

type EditMinterInput struct {
	SalePhases []SaleFaze
	Symbol     string
}

func (obj EditMinterInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SalePhases` param:
	err = encoder.Encode(obj.SalePhases)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	return nil
}

func (obj *EditMinterInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SalePhases`:
	err = decoder.Decode(&obj.SalePhases)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	return nil
}

type EditPhaseInput struct {
	Name     string
	MaxMints *uint32 `bin:"optional"`
}

func (obj EditPhaseInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `MaxMints` param (optional):
	{
		if obj.MaxMints == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxMints)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *EditPhaseInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `MaxMints` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxMints)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ResizeInput struct {
	TargetSize uint32
}

func (obj ResizeInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TargetSize` param:
	err = encoder.Encode(obj.TargetSize)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ResizeInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TargetSize`:
	err = decoder.Decode(&obj.TargetSize)
	if err != nil {
		return err
	}
	return nil
}

type InitInput struct {
	Seed             ag_solanago.PublicKey
	Data             MachineData
	MachineType      MachineType
	SalePhases       []MachinePhase
	FundReceivers    []Creator
	TableSlot        uint64
	EnforceRoyalties bool
	Frozen           bool
}

func (obj InitInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Seed` param:
	err = encoder.Encode(obj.Seed)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `MachineType` param:
	{
		tmp := machineTypeContainer{}
		switch realvalue := obj.MachineType.(type) {
		case *MachineTypeCompressed:
			tmp.Enum = 0
			tmp.Compressed = *realvalue
		case *MachineTypeFairLaunch:
			tmp.Enum = 1
			tmp.FairLaunch = *realvalue
		case *MachineTypeLiquid:
			tmp.Enum = 2
			tmp.Liquid = *realvalue
		case *MachineTypeCore:
			tmp.Enum = 3
			tmp.Core = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `SalePhases` param:
	err = encoder.Encode(obj.SalePhases)
	if err != nil {
		return err
	}
	// Serialize `FundReceivers` param:
	err = encoder.Encode(obj.FundReceivers)
	if err != nil {
		return err
	}
	// Serialize `TableSlot` param:
	err = encoder.Encode(obj.TableSlot)
	if err != nil {
		return err
	}
	// Serialize `EnforceRoyalties` param:
	err = encoder.Encode(obj.EnforceRoyalties)
	if err != nil {
		return err
	}
	// Serialize `Frozen` param:
	err = encoder.Encode(obj.Frozen)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Seed`:
	err = decoder.Decode(&obj.Seed)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `MachineType`:
	{
		tmp := new(machineTypeContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.MachineType = &tmp.Compressed
		case 1:
			obj.MachineType = &tmp.FairLaunch
		case 2:
			obj.MachineType = &tmp.Liquid
		case 3:
			obj.MachineType = &tmp.Core
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `SalePhases`:
	err = decoder.Decode(&obj.SalePhases)
	if err != nil {
		return err
	}
	// Deserialize `FundReceivers`:
	err = decoder.Decode(&obj.FundReceivers)
	if err != nil {
		return err
	}
	// Deserialize `TableSlot`:
	err = decoder.Decode(&obj.TableSlot)
	if err != nil {
		return err
	}
	// Deserialize `EnforceRoyalties`:
	err = decoder.Decode(&obj.EnforceRoyalties)
	if err != nil {
		return err
	}
	// Deserialize `Frozen`:
	err = decoder.Decode(&obj.Frozen)
	if err != nil {
		return err
	}
	return nil
}

type UpdateInput struct {
	Data          *MachineData    `bin:"optional"`
	MachineType   *MachineType    `bin:"optional"`
	SalePhases    *[]MachinePhase `bin:"optional"`
	FundReceivers *[]Creator      `bin:"optional"`
}

func (obj UpdateInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param (optional):
	{
		if obj.Data == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MachineType` param (optional):
	{
		if obj.MachineType == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MachineType)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SalePhases` param (optional):
	{
		if obj.SalePhases == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SalePhases)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `FundReceivers` param (optional):
	{
		if obj.FundReceivers == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.FundReceivers)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UpdateInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MachineType` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MachineType)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SalePhases` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SalePhases)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `FundReceivers` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.FundReceivers)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MachineMintInput struct {
	Proof         [][32]uint8
	ExpectedPrice uint64
	Amount        uint32
	CnftWlEntry   *CNFTWhitelistEntry `bin:"optional"`
	TargetPhase   *uint8              `bin:"optional"`
}

func (obj MachineMintInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Proof` param:
	err = encoder.Encode(obj.Proof)
	if err != nil {
		return err
	}
	// Serialize `ExpectedPrice` param:
	err = encoder.Encode(obj.ExpectedPrice)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `CnftWlEntry` param (optional):
	{
		if obj.CnftWlEntry == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.CnftWlEntry)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `TargetPhase` param (optional):
	{
		if obj.TargetPhase == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TargetPhase)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *MachineMintInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Proof`:
	err = decoder.Decode(&obj.Proof)
	if err != nil {
		return err
	}
	// Deserialize `ExpectedPrice`:
	err = decoder.Decode(&obj.ExpectedPrice)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `CnftWlEntry` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.CnftWlEntry)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `TargetPhase` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TargetPhase)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MintInput struct {
	Proof         [][32]uint8
	ExpectedPrice uint64
	Amount        uint32
}

func (obj MintInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Proof` param:
	err = encoder.Encode(obj.Proof)
	if err != nil {
		return err
	}
	// Serialize `ExpectedPrice` param:
	err = encoder.Encode(obj.ExpectedPrice)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Proof`:
	err = decoder.Decode(&obj.Proof)
	if err != nil {
		return err
	}
	// Deserialize `ExpectedPrice`:
	err = decoder.Decode(&obj.ExpectedPrice)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type CNFTWhitelistEntry struct {
	Root        [32]uint8
	DataHash    [32]uint8
	CreatorHash [32]uint8
	AssetId     ag_solanago.PublicKey
	Nonce       uint64
	Index       uint32
	Creators    []BGumCreator
}

func (obj CNFTWhitelistEntry) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `DataHash` param:
	err = encoder.Encode(obj.DataHash)
	if err != nil {
		return err
	}
	// Serialize `CreatorHash` param:
	err = encoder.Encode(obj.CreatorHash)
	if err != nil {
		return err
	}
	// Serialize `AssetId` param:
	err = encoder.Encode(obj.AssetId)
	if err != nil {
		return err
	}
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CNFTWhitelistEntry) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `DataHash`:
	err = decoder.Decode(&obj.DataHash)
	if err != nil {
		return err
	}
	// Deserialize `CreatorHash`:
	err = decoder.Decode(&obj.CreatorHash)
	if err != nil {
		return err
	}
	// Deserialize `AssetId`:
	err = decoder.Decode(&obj.AssetId)
	if err != nil {
		return err
	}
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	return nil
}

type MachineData struct {
	Authority            ag_solanago.PublicKey
	Ordered              bool
	Name                 string
	Url                  string
	Symbol               string
	Creators             []Creator
	SellerFeeBasisPoints uint16
	Supply               uint32
}

func (obj MachineData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Ordered` param:
	err = encoder.Encode(obj.Ordered)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Url` param:
	err = encoder.Encode(obj.Url)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachineData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Ordered`:
	err = decoder.Decode(&obj.Ordered)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Url`:
	err = decoder.Decode(&obj.Url)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	return nil
}

type MachinePhase struct {
	Start           int64
	Price           uint64
	Currency        *ag_solanago.PublicKey `bin:"optional"`
	WhitelistMode   WhitelistMode
	End             *int64  `bin:"optional"`
	GlobalMintLimit *uint32 `bin:"optional"`
	MintsInPhase    uint32
	RequiredSigner  *ag_solanago.PublicKey `bin:"optional"`
	Name            string
	Padding         [10]uint8
}

func (obj MachinePhase) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Currency` param (optional):
	{
		if obj.Currency == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `WhitelistMode` param:
	{
		tmp := whitelistModeContainer{}
		switch realvalue := obj.WhitelistMode.(type) {
		case *WhitelistModeWalletBased:
			tmp.Enum = 0
			tmp.WalletBased = *realvalue
		case *WhitelistModeTokenBased:
			tmp.Enum = 1
			tmp.TokenBased = *realvalue
		case *WhitelistModePublic:
			tmp.Enum = 2
			tmp.Public = *realvalue
		case *WhitelistModeNFTBased:
			tmp.Enum = 3
			tmp.NFTBased = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `End` param (optional):
	{
		if obj.End == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.End)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `GlobalMintLimit` param (optional):
	{
		if obj.GlobalMintLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.GlobalMintLimit)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `MintsInPhase` param:
	err = encoder.Encode(obj.MintsInPhase)
	if err != nil {
		return err
	}
	// Serialize `RequiredSigner` param (optional):
	{
		if obj.RequiredSigner == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.RequiredSigner)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachinePhase) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Currency` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `WhitelistMode`:
	{
		tmp := new(whitelistModeContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.WhitelistMode = &tmp.WalletBased
		case 1:
			obj.WhitelistMode = &tmp.TokenBased
		case 2:
			obj.WhitelistMode = &tmp.Public
		case 3:
			obj.WhitelistMode = &tmp.NFTBased
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `End` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.End)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `GlobalMintLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.GlobalMintLimit)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `MintsInPhase`:
	err = decoder.Decode(&obj.MintsInPhase)
	if err != nil {
		return err
	}
	// Deserialize `RequiredSigner` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.RequiredSigner)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type Creator struct {
	Address ag_solanago.PublicKey
	Share   uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type BGumCreator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj BGumCreator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BGumCreator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type SaleFaze struct {
	Start         int64
	Price         uint64
	Currency      *ag_solanago.PublicKey `bin:"optional"`
	WhitelistMode WhitelistMode
	Name          string
}

func (obj SaleFaze) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Currency` param (optional):
	{
		if obj.Currency == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `WhitelistMode` param:
	{
		tmp := whitelistModeContainer{}
		switch realvalue := obj.WhitelistMode.(type) {
		case *WhitelistModeWalletBased:
			tmp.Enum = 0
			tmp.WalletBased = *realvalue
		case *WhitelistModeTokenBased:
			tmp.Enum = 1
			tmp.TokenBased = *realvalue
		case *WhitelistModePublic:
			tmp.Enum = 2
			tmp.Public = *realvalue
		case *WhitelistModeNFTBased:
			tmp.Enum = 3
			tmp.NFTBased = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SaleFaze) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Currency` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Currency)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `WhitelistMode`:
	{
		tmp := new(whitelistModeContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.WhitelistMode = &tmp.WalletBased
		case 1:
			obj.WhitelistMode = &tmp.TokenBased
		case 2:
			obj.WhitelistMode = &tmp.Public
		case 3:
			obj.WhitelistMode = &tmp.NFTBased
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

type PublicMode struct {
	MintsPerUser *uint32 `bin:"optional"`
}

func (obj PublicMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *PublicMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type WhitelistTokenMode struct {
	TokenMint ag_solanago.PublicKey
	Burn      bool
}

func (obj WhitelistTokenMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenMint` param:
	err = encoder.Encode(obj.TokenMint)
	if err != nil {
		return err
	}
	// Serialize `Burn` param:
	err = encoder.Encode(obj.Burn)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistTokenMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenMint`:
	err = decoder.Decode(&obj.TokenMint)
	if err != nil {
		return err
	}
	// Deserialize `Burn`:
	err = decoder.Decode(&obj.Burn)
	if err != nil {
		return err
	}
	return nil
}

type NFTHolderMode struct {
	VerifiedCreator ag_solanago.PublicKey
	MintsPerNft     uint32
}

func (obj NFTHolderMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `VerifiedCreator` param:
	err = encoder.Encode(obj.VerifiedCreator)
	if err != nil {
		return err
	}
	// Serialize `MintsPerNft` param:
	err = encoder.Encode(obj.MintsPerNft)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NFTHolderMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `VerifiedCreator`:
	err = decoder.Decode(&obj.VerifiedCreator)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerNft`:
	err = decoder.Decode(&obj.MintsPerNft)
	if err != nil {
		return err
	}
	return nil
}

type WhitelistWalletListMode struct {
	MintsPerUser uint32
	MerkleRoot   [32]uint8
}

func (obj WhitelistWalletListMode) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintsPerUser` param:
	err = encoder.Encode(obj.MintsPerUser)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistWalletListMode) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintsPerUser`:
	err = decoder.Decode(&obj.MintsPerUser)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

type RevealState interface {
	isRevealState()
}

type revealStateContainer struct {
	Enum       ag_binary.BorshEnum `borsh_enum:"true"`
	NotAllowed RevealStateNotAllowed
	Allowed    RevealStateAllowed
}

type RevealStateNotAllowed uint8

func (obj RevealStateNotAllowed) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *RevealStateNotAllowed) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

func (_ *RevealStateNotAllowed) isRevealState() {}

type RevealStateAllowed struct {
	Uri string
}

func (obj RevealStateAllowed) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RevealStateAllowed) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	return nil
}

func (_ *RevealStateAllowed) isRevealState() {}

type MachineType interface {
	isMachineType()
}

type machineTypeContainer struct {
	Enum       ag_binary.BorshEnum `borsh_enum:"true"`
	Compressed MachineTypeCompressed
	FairLaunch MachineTypeFairLaunch
	Liquid     MachineTypeLiquid
	Core       MachineTypeCore
}

type MachineTypeCompressed struct {
	Collection ag_solanago.PublicKey
}

func (obj MachineTypeCompressed) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Collection` param:
	err = encoder.Encode(obj.Collection)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachineTypeCompressed) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Collection`:
	err = decoder.Decode(&obj.Collection)
	if err != nil {
		return err
	}
	return nil
}

func (_ *MachineTypeCompressed) isMachineType() {}

type MachineTypeFairLaunch struct {
	Deployment ag_solanago.PublicKey
}

func (obj MachineTypeFairLaunch) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Deployment` param:
	err = encoder.Encode(obj.Deployment)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachineTypeFairLaunch) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Deployment`:
	err = decoder.Decode(&obj.Deployment)
	if err != nil {
		return err
	}
	return nil
}

func (_ *MachineTypeFairLaunch) isMachineType() {}

type MachineTypeLiquid struct {
	Liquidity ag_solanago.PublicKey
	Url       string
}

func (obj MachineTypeLiquid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `Url` param:
	err = encoder.Encode(obj.Url)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachineTypeLiquid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `Url`:
	err = decoder.Decode(&obj.Url)
	if err != nil {
		return err
	}
	return nil
}

func (_ *MachineTypeLiquid) isMachineType() {}

type MachineTypeCore struct {
	Collection  ag_solanago.PublicKey
	RevealLater *RevealState `bin:"optional"`
}

func (obj MachineTypeCore) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Collection` param:
	err = encoder.Encode(obj.Collection)
	if err != nil {
		return err
	}
	// Serialize `RevealLater` param (optional):
	{
		if obj.RevealLater == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.RevealLater)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *MachineTypeCore) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Collection`:
	err = decoder.Decode(&obj.Collection)
	if err != nil {
		return err
	}
	// Deserialize `RevealLater` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.RevealLater)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (_ *MachineTypeCore) isMachineType() {}

type WhitelistMode interface {
	isWhitelistMode()
}

type whitelistModeContainer struct {
	Enum        ag_binary.BorshEnum `borsh_enum:"true"`
	WalletBased WhitelistModeWalletBased
	TokenBased  WhitelistModeTokenBased
	Public      WhitelistModePublic
	NFTBased    WhitelistModeNFTBased
}

type WhitelistModeWalletBased struct {
	Info WhitelistWalletListMode
}

func (obj WhitelistModeWalletBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeWalletBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeWalletBased) isWhitelistMode() {}

type WhitelistModeTokenBased struct {
	Info WhitelistTokenMode
}

func (obj WhitelistModeTokenBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeTokenBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeTokenBased) isWhitelistMode() {}

type WhitelistModePublic struct {
	Info PublicMode
}

func (obj WhitelistModePublic) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModePublic) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModePublic) isWhitelistMode() {}

type WhitelistModeNFTBased struct {
	Info NFTHolderMode
}

func (obj WhitelistModeNFTBased) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Info` param:
	err = encoder.Encode(obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhitelistModeNFTBased) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Info`:
	err = decoder.Decode(&obj.Info)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WhitelistModeNFTBased) isWhitelistMode() {}