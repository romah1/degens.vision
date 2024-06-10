// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_core_program

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateCollectionPluginV1 is the `UpdateCollectionPluginV1` instruction.
type UpdateCollectionPluginV1 struct {
	UpdateCollectionPluginV1Args *UpdateCollectionPluginV1Args

	// [0] = [WRITE] collection
	// ··········· The address of the asset
	//
	// [1] = [WRITE, SIGNER] payer
	// ··········· The account paying for the storage fees
	//
	// [2] = [SIGNER] authority
	// ··········· The owner or delegate of the asset
	//
	// [3] = [] systemProgram
	// ··········· The system program
	//
	// [4] = [] logWrapper
	// ··········· The SPL Noop Program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateCollectionPluginV1InstructionBuilder creates a new `UpdateCollectionPluginV1` instruction builder.
func NewUpdateCollectionPluginV1InstructionBuilder() *UpdateCollectionPluginV1 {
	nd := &UpdateCollectionPluginV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetUpdateCollectionPluginV1Args sets the "updateCollectionPluginV1Args" parameter.
func (inst *UpdateCollectionPluginV1) SetUpdateCollectionPluginV1Args(updateCollectionPluginV1Args UpdateCollectionPluginV1Args) *UpdateCollectionPluginV1 {
	inst.UpdateCollectionPluginV1Args = &updateCollectionPluginV1Args
	return inst
}

// SetCollectionAccount sets the "collection" account.
// The address of the asset
func (inst *UpdateCollectionPluginV1) SetCollectionAccount(collection ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
// The address of the asset
func (inst *UpdateCollectionPluginV1) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
// The account paying for the storage fees
func (inst *UpdateCollectionPluginV1) SetPayerAccount(payer ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// The account paying for the storage fees
func (inst *UpdateCollectionPluginV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
// The owner or delegate of the asset
func (inst *UpdateCollectionPluginV1) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The owner or delegate of the asset
func (inst *UpdateCollectionPluginV1) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// The system program
func (inst *UpdateCollectionPluginV1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// The system program
func (inst *UpdateCollectionPluginV1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLogWrapperAccount sets the "logWrapper" account.
// The SPL Noop Program
func (inst *UpdateCollectionPluginV1) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
// The SPL Noop Program
func (inst *UpdateCollectionPluginV1) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateCollectionPluginV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateCollectionPluginV1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateCollectionPluginV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateCollectionPluginV1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UpdateCollectionPluginV1Args == nil {
			return errors.New("UpdateCollectionPluginV1Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LogWrapper is not set")
		}
	}
	return nil
}

func (inst *UpdateCollectionPluginV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateCollectionPluginV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UpdateCollectionPluginV1Args", *inst.UpdateCollectionPluginV1Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   collection", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   logWrapper", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateCollectionPluginV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UpdateCollectionPluginV1Args` param:
	err = encoder.Encode(obj.UpdateCollectionPluginV1Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateCollectionPluginV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UpdateCollectionPluginV1Args`:
	err = decoder.Decode(&obj.UpdateCollectionPluginV1Args)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateCollectionPluginV1Instruction declares a new UpdateCollectionPluginV1 instruction with the provided parameters and accounts.
func NewUpdateCollectionPluginV1Instruction(
	// Parameters:
	updateCollectionPluginV1Args UpdateCollectionPluginV1Args,
	// Accounts:
	collection ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey) *UpdateCollectionPluginV1 {
	return NewUpdateCollectionPluginV1InstructionBuilder().
		SetUpdateCollectionPluginV1Args(updateCollectionPluginV1Args).
		SetCollectionAccount(collection).
		SetPayerAccount(payer).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetLogWrapperAccount(logWrapper)
}