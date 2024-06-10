// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_core_program

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateCollectionV1 is the `CreateCollectionV1` instruction.
type CreateCollectionV1 struct {
	CreateCollectionV1Args *CreateCollectionV1Args

	// [0] = [WRITE, SIGNER] collection
	// ··········· The address of the new asset
	//
	// [1] = [] updateAuthority
	// ··········· The authority of the new asset
	//
	// [2] = [WRITE, SIGNER] payer
	// ··········· The account paying for the storage fees
	//
	// [3] = [] systemProgram
	// ··········· The system program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateCollectionV1InstructionBuilder creates a new `CreateCollectionV1` instruction builder.
func NewCreateCollectionV1InstructionBuilder() *CreateCollectionV1 {
	nd := &CreateCollectionV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetCreateCollectionV1Args sets the "createCollectionV1Args" parameter.
func (inst *CreateCollectionV1) SetCreateCollectionV1Args(createCollectionV1Args CreateCollectionV1Args) *CreateCollectionV1 {
	inst.CreateCollectionV1Args = &createCollectionV1Args
	return inst
}

// SetCollectionAccount sets the "collection" account.
// The address of the new asset
func (inst *CreateCollectionV1) SetCollectionAccount(collection ag_solanago.PublicKey) *CreateCollectionV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(collection).WRITE().SIGNER()
	return inst
}

// GetCollectionAccount gets the "collection" account.
// The address of the new asset
func (inst *CreateCollectionV1) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// The authority of the new asset
func (inst *CreateCollectionV1) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateCollectionV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(updateAuthority)
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// The authority of the new asset
func (inst *CreateCollectionV1) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// The account paying for the storage fees
func (inst *CreateCollectionV1) SetPayerAccount(payer ag_solanago.PublicKey) *CreateCollectionV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// The account paying for the storage fees
func (inst *CreateCollectionV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// The system program
func (inst *CreateCollectionV1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateCollectionV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// The system program
func (inst *CreateCollectionV1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CreateCollectionV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateCollectionV1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateCollectionV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateCollectionV1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreateCollectionV1Args == nil {
			return errors.New("CreateCollectionV1Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateCollectionV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateCollectionV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreateCollectionV1Args", *inst.CreateCollectionV1Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     collection", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj CreateCollectionV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreateCollectionV1Args` param:
	err = encoder.Encode(obj.CreateCollectionV1Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateCollectionV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreateCollectionV1Args`:
	err = decoder.Decode(&obj.CreateCollectionV1Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateCollectionV1Instruction declares a new CreateCollectionV1 instruction with the provided parameters and accounts.
func NewCreateCollectionV1Instruction(
	// Parameters:
	createCollectionV1Args CreateCollectionV1Args,
	// Accounts:
	collection ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateCollectionV1 {
	return NewCreateCollectionV1InstructionBuilder().
		SetCreateCollectionV1Args(createCollectionV1Args).
		SetCollectionAccount(collection).
		SetUpdateAuthorityAccount(updateAuthority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}
