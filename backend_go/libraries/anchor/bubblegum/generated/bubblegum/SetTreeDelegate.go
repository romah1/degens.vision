// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetTreeDelegate is the `setTreeDelegate` instruction.
type SetTreeDelegate struct {

	// [0] = [WRITE] treeAuthority
	//
	// [1] = [SIGNER] treeCreator
	//
	// [2] = [] newTreeDelegate
	//
	// [3] = [] merkleTree
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetTreeDelegateInstructionBuilder creates a new `SetTreeDelegate` instruction builder.
func NewSetTreeDelegateInstructionBuilder() *SetTreeDelegate {
	nd := &SetTreeDelegate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *SetTreeDelegate) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *SetTreeDelegate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority).WRITE()
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *SetTreeDelegate) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTreeCreatorAccount sets the "treeCreator" account.
func (inst *SetTreeDelegate) SetTreeCreatorAccount(treeCreator ag_solanago.PublicKey) *SetTreeDelegate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(treeCreator).SIGNER()
	return inst
}

// GetTreeCreatorAccount gets the "treeCreator" account.
func (inst *SetTreeDelegate) GetTreeCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetNewTreeDelegateAccount sets the "newTreeDelegate" account.
func (inst *SetTreeDelegate) SetNewTreeDelegateAccount(newTreeDelegate ag_solanago.PublicKey) *SetTreeDelegate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(newTreeDelegate)
	return inst
}

// GetNewTreeDelegateAccount gets the "newTreeDelegate" account.
func (inst *SetTreeDelegate) GetNewTreeDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMerkleTreeAccount sets the "merkleTree" account.
func (inst *SetTreeDelegate) SetMerkleTreeAccount(merkleTree ag_solanago.PublicKey) *SetTreeDelegate {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(merkleTree)
	return inst
}

// GetMerkleTreeAccount gets the "merkleTree" account.
func (inst *SetTreeDelegate) GetMerkleTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetTreeDelegate) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetTreeDelegate {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetTreeDelegate) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst SetTreeDelegate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTreeDelegate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTreeDelegate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTreeDelegate) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TreeAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TreeCreator is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NewTreeDelegate is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MerkleTree is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *SetTreeDelegate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTreeDelegate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    treeCreator", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("newTreeDelegate", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     merkleTree", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj SetTreeDelegate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetTreeDelegate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetTreeDelegateInstruction declares a new SetTreeDelegate instruction with the provided parameters and accounts.
func NewSetTreeDelegateInstruction(
	// Accounts:
	treeAuthority ag_solanago.PublicKey,
	treeCreator ag_solanago.PublicKey,
	newTreeDelegate ag_solanago.PublicKey,
	merkleTree ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetTreeDelegate {
	return NewSetTreeDelegateInstructionBuilder().
		SetTreeAuthorityAccount(treeAuthority).
		SetTreeCreatorAccount(treeCreator).
		SetNewTreeDelegateAccount(newTreeDelegate).
		SetMerkleTreeAccount(merkleTree).
		SetSystemProgramAccount(systemProgram)
}
