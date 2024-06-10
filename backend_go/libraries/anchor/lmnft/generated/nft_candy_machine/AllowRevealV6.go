// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AllowRevealV6 is the `allowRevealV6` instruction.
type AllowRevealV6 struct {
	NewUri *string

	// [0] = [WRITE, SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] systemProgram
	//
	// [3] = [WRITE] extensions
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAllowRevealV6InstructionBuilder creates a new `AllowRevealV6` instruction builder.
func NewAllowRevealV6InstructionBuilder() *AllowRevealV6 {
	nd := &AllowRevealV6{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetNewUri sets the "newUri" parameter.
func (inst *AllowRevealV6) SetNewUri(newUri string) *AllowRevealV6 {
	inst.NewUri = &newUri
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AllowRevealV6) SetAuthorityAccount(authority ag_solanago.PublicKey) *AllowRevealV6 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AllowRevealV6) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *AllowRevealV6) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *AllowRevealV6 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *AllowRevealV6) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AllowRevealV6) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AllowRevealV6 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AllowRevealV6) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetExtensionsAccount sets the "extensions" account.
func (inst *AllowRevealV6) SetExtensionsAccount(extensions ag_solanago.PublicKey) *AllowRevealV6 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(extensions).WRITE()
	return inst
}

// GetExtensionsAccount gets the "extensions" account.
func (inst *AllowRevealV6) GetExtensionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst AllowRevealV6) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AllowRevealV6,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AllowRevealV6) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AllowRevealV6) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewUri == nil {
			return errors.New("NewUri parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Extensions is not set")
		}
	}
	return nil
}

func (inst *AllowRevealV6) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AllowRevealV6")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewUri", *inst.NewUri))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   extensions", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj AllowRevealV6) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewUri` param:
	err = encoder.Encode(obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AllowRevealV6) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewUri`:
	err = decoder.Decode(&obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}

// NewAllowRevealV6Instruction declares a new AllowRevealV6 instruction with the provided parameters and accounts.
func NewAllowRevealV6Instruction(
	// Parameters:
	newUri string,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	extensions ag_solanago.PublicKey) *AllowRevealV6 {
	return NewAllowRevealV6InstructionBuilder().
		SetNewUri(newUri).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine).
		SetSystemProgramAccount(systemProgram).
		SetExtensionsAccount(extensions)
}
