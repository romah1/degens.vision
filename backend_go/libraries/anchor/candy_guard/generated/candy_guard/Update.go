// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_guard

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update is the `update` instruction.
type Update struct {
	Data *[]byte

	// [0] = [WRITE] candyGuard
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [SIGNER] payer
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateInstructionBuilder creates a new `Update` instruction builder.
func NewUpdateInstructionBuilder() *Update {
	nd := &Update{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *Update) SetData(data []byte) *Update {
	inst.Data = &data
	return inst
}

// SetCandyGuardAccount sets the "candyGuard" account.
func (inst *Update) SetCandyGuardAccount(candyGuard ag_solanago.PublicKey) *Update {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyGuard).WRITE()
	return inst
}

// GetCandyGuardAccount gets the "candyGuard" account.
func (inst *Update) GetCandyGuardAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Update) SetAuthorityAccount(authority ag_solanago.PublicKey) *Update {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Update) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *Update) SetPayerAccount(payer ag_solanago.PublicKey) *Update {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *Update) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Update) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Update {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Update) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst Update) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Update,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Update) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Update) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyGuard is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
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

func (inst *Update) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Update")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   candyGuard", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj Update) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Update) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateInstruction declares a new Update instruction with the provided parameters and accounts.
func NewUpdateInstruction(
	// Parameters:
	data []byte,
	// Accounts:
	candyGuard ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Update {
	return NewUpdateInstructionBuilder().
		SetData(data).
		SetCandyGuardAccount(candyGuard).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}
