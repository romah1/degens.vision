// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_core_program

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DecompressV1 is the `DecompressV1` instruction.
type DecompressV1 struct {
	DecompressV1Args *DecompressV1Args

	// [0] = [WRITE] asset
	// ··········· The address of the asset
	//
	// [1] = [] collection
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

// NewDecompressV1InstructionBuilder creates a new `DecompressV1` instruction builder.
func NewDecompressV1InstructionBuilder() *DecompressV1 {
	nd := &DecompressV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetDecompressV1Args sets the "decompressV1Args" parameter.
func (inst *DecompressV1) SetDecompressV1Args(decompressV1Args DecompressV1Args) *DecompressV1 {
	inst.DecompressV1Args = &decompressV1Args
	return inst
}

// SetAssetAccount sets the "asset" account.
// The address of the asset
func (inst *DecompressV1) SetAssetAccount(asset ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(asset).WRITE()
	return inst
}

// GetAssetAccount gets the "asset" account.
// The address of the asset
func (inst *DecompressV1) GetAssetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCollectionAccount sets the "collection" account.
// The collection to which the asset belongs
func (inst *DecompressV1) SetCollectionAccount(collection ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(collection)
	return inst
}

// GetCollectionAccount gets the "collection" account.
// The collection to which the asset belongs
func (inst *DecompressV1) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// The account paying for the storage fees
func (inst *DecompressV1) SetPayerAccount(payer ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// The account paying for the storage fees
func (inst *DecompressV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
// The owner or delegate of the asset
func (inst *DecompressV1) SetAuthorityAccount(authority ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The owner or delegate of the asset
func (inst *DecompressV1) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// The system program
func (inst *DecompressV1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// The system program
func (inst *DecompressV1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetLogWrapperAccount sets the "logWrapper" account.
// The SPL Noop Program
func (inst *DecompressV1) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *DecompressV1 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
// The SPL Noop Program
func (inst *DecompressV1) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst DecompressV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DecompressV1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DecompressV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DecompressV1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DecompressV1Args == nil {
			return errors.New("DecompressV1Args parameter is not set")
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

func (inst *DecompressV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecompressV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DecompressV1Args", *inst.DecompressV1Args))
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

func (obj DecompressV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DecompressV1Args` param:
	err = encoder.Encode(obj.DecompressV1Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DecompressV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DecompressV1Args`:
	err = decoder.Decode(&obj.DecompressV1Args)
	if err != nil {
		return err
	}
	return nil
}

// NewDecompressV1Instruction declares a new DecompressV1 instruction with the provided parameters and accounts.
func NewDecompressV1Instruction(
	// Parameters:
	decompressV1Args DecompressV1Args,
	// Accounts:
	asset ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey) *DecompressV1 {
	return NewDecompressV1InstructionBuilder().
		SetDecompressV1Args(decompressV1Args).
		SetAssetAccount(asset).
		SetCollectionAccount(collection).
		SetPayerAccount(payer).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetLogWrapperAccount(logWrapper)
}
