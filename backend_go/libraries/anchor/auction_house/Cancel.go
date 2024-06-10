// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Cancel is the `cancel` instruction.
type Cancel struct {
	BuyerPrice *uint64
	TokenSize  *uint64

	// [0] = [WRITE] wallet
	//
	// [1] = [WRITE] tokenAccount
	//
	// [2] = [] tokenMint
	//
	// [3] = [] authority
	//
	// [4] = [] auctionHouse
	//
	// [5] = [WRITE] auctionHouseFeeAccount
	//
	// [6] = [WRITE] tradeState
	//
	// [7] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelInstructionBuilder creates a new `Cancel` instruction builder.
func NewCancelInstructionBuilder() *Cancel {
	nd := &Cancel{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetBuyerPrice sets the "buyerPrice" parameter.
func (inst *Cancel) SetBuyerPrice(buyerPrice uint64) *Cancel {
	inst.BuyerPrice = &buyerPrice
	return inst
}

// SetTokenSize sets the "tokenSize" parameter.
func (inst *Cancel) SetTokenSize(tokenSize uint64) *Cancel {
	inst.TokenSize = &tokenSize
	return inst
}

// SetWalletAccount sets the "wallet" account.
func (inst *Cancel) SetWalletAccount(wallet ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(wallet).WRITE()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *Cancel) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAccountAccount sets the "tokenAccount" account.
func (inst *Cancel) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAccount).WRITE()
	return inst
}

// GetTokenAccountAccount gets the "tokenAccount" account.
func (inst *Cancel) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *Cancel) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *Cancel) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Cancel) SetAuthorityAccount(authority ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Cancel) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *Cancel) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *Cancel) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *Cancel) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *Cancel) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTradeStateAccount sets the "tradeState" account.
func (inst *Cancel) SetTradeStateAccount(tradeState ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tradeState).WRITE()
	return inst
}

// GetTradeStateAccount gets the "tradeState" account.
func (inst *Cancel) GetTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Cancel) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Cancel {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Cancel) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst Cancel) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Cancel,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Cancel) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Cancel) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.BuyerPrice == nil {
			return errors.New("BuyerPrice parameter is not set")
		}
		if inst.TokenSize == nil {
			return errors.New("TokenSize parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TradeState is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *Cancel) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Cancel")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("BuyerPrice", *inst.BuyerPrice))
						paramsBranch.Child(ag_format.Param(" TokenSize", *inst.TokenSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          token", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      tokenMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   auctionHouse", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("auctionHouseFee", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     tradeState", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj Cancel) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BuyerPrice` param:
	err = encoder.Encode(obj.BuyerPrice)
	if err != nil {
		return err
	}
	// Serialize `TokenSize` param:
	err = encoder.Encode(obj.TokenSize)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Cancel) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BuyerPrice`:
	err = decoder.Decode(&obj.BuyerPrice)
	if err != nil {
		return err
	}
	// Deserialize `TokenSize`:
	err = decoder.Decode(&obj.TokenSize)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelInstruction declares a new Cancel instruction with the provided parameters and accounts.
func NewCancelInstruction(
	// Parameters:
	buyerPrice uint64,
	tokenSize uint64,
	// Accounts:
	wallet ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	tradeState ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *Cancel {
	return NewCancelInstructionBuilder().
		SetBuyerPrice(buyerPrice).
		SetTokenSize(tokenSize).
		SetWalletAccount(wallet).
		SetTokenAccountAccount(tokenAccount).
		SetTokenMintAccount(tokenMint).
		SetAuthorityAccount(authority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetTradeStateAccount(tradeState).
		SetTokenProgramAccount(tokenProgram)
}
