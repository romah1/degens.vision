// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_core_program

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BurnV1 is the `BurnV1` instruction.
type BurnV1 struct {
	BurnV1Args *BurnV1Args

	// [0] = [WRITE] asset
	// ··········· The address of the asset
	//
	// [1] = [WRITE] collection
	// ··········· The collection to which the asset belongs
	//
	// [2] = [WRITE, SIGNER] payer
	// ··········· The account paying for the storage fees
	//
	// [3] = [SIGNER] authority
	// ··········· The owner or delegate of the asset
	//
	// [4] = [] systemProgram
	// ··········· The system program
	//
	// [5] = [] logWrapper
	// ··········· The SPL Noop Program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBurnV1InstructionBuilder creates a new `BurnV1` instruction builder.
func NewBurnV1InstructionBuilder() *BurnV1 {
	nd := &BurnV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetBurnV1Args sets the "burnV1Args" parameter.
func (inst *BurnV1) SetBurnV1Args(burnV1Args BurnV1Args) *BurnV1 {
	inst.BurnV1Args = &burnV1Args
	return inst
}

// SetAssetAccount sets the "asset" account.
// The address of the asset
func (inst *BurnV1) SetAssetAccount(asset ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(asset).WRITE()
	return inst
}

// GetAssetAccount gets the "asset" account.
// The address of the asset
func (inst *BurnV1) GetAssetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAccount sets the "collection" account.
// The collection to which the asset belongs
func (inst *BurnV1) SetCollectionAccount(collection ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
// The collection to which the asset belongs
func (inst *BurnV1) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// The account paying for the storage fees
func (inst *BurnV1) SetPayerAccount(payer ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// The account paying for the storage fees
func (inst *BurnV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
// The owner or delegate of the asset
func (inst *BurnV1) SetAuthorityAccount(authority ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The owner or delegate of the asset
func (inst *BurnV1) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// The system program
func (inst *BurnV1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// The system program
func (inst *BurnV1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetLogWrapperAccount sets the "logWrapper" account.
// The SPL Noop Program
func (inst *BurnV1) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *BurnV1 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
// The SPL Noop Program
func (inst *BurnV1) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst BurnV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BurnV1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BurnV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BurnV1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.BurnV1Args == nil {
			return errors.New("BurnV1Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Asset is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.LogWrapper is not set")
		}
	}
	return nil
}

func (inst *BurnV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BurnV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("BurnV1Args", *inst.BurnV1Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        asset", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   collection", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   logWrapper", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj BurnV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BurnV1Args` param:
	err = encoder.Encode(obj.BurnV1Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BurnV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BurnV1Args`:
	err = decoder.Decode(&obj.BurnV1Args)
	if err != nil {
		return err
	}
	return nil
}

// NewBurnV1Instruction declares a new BurnV1 instruction with the provided parameters and accounts.
func NewBurnV1Instruction(
	// Parameters:
	burnV1Args BurnV1Args,
	// Accounts:
	asset ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey) *BurnV1 {
	return NewBurnV1InstructionBuilder().
		SetBurnV1Args(burnV1Args).
		SetAssetAccount(asset).
		SetCollectionAccount(collection).
		SetPayerAccount(payer).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetLogWrapperAccount(logWrapper)
}
