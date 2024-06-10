// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine_core

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Withdraw is the `withdraw` instruction.
type Withdraw struct {

	// [0] = [WRITE] candyMachine
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawInstructionBuilder creates a new `Withdraw` instruction builder.
func NewWithdrawInstructionBuilder() *Withdraw {
	nd := &Withdraw{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *Withdraw) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *Withdraw) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Withdraw) SetAuthorityAccount(authority ag_solanago.PublicKey) *Withdraw {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Withdraw) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst Withdraw) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Withdraw,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Withdraw) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Withdraw) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *Withdraw) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Withdraw")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj Withdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Withdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewWithdrawInstruction declares a new Withdraw instruction with the provided parameters and accounts.
func NewWithdrawInstruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *Withdraw {
	return NewWithdrawInstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority)
}
