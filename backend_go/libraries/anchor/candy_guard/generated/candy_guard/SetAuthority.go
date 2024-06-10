// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_guard

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetAuthority is the `setAuthority` instruction.
type SetAuthority struct {
	NewAuthority *ag_solanago.PublicKey

	// [0] = [WRITE] candyGuard
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetAuthorityInstructionBuilder creates a new `SetAuthority` instruction builder.
func NewSetAuthorityInstructionBuilder() *SetAuthority {
	nd := &SetAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewAuthority sets the "newAuthority" parameter.
func (inst *SetAuthority) SetNewAuthority(newAuthority ag_solanago.PublicKey) *SetAuthority {
	inst.NewAuthority = &newAuthority
	return inst
}

// SetCandyGuardAccount sets the "candyGuard" account.
func (inst *SetAuthority) SetCandyGuardAccount(candyGuard ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyGuard).WRITE()
	return inst
}

// GetCandyGuardAccount gets the "candyGuard" account.
func (inst *SetAuthority) GetCandyGuardAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetAuthority) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetAuthority) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst SetAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewAuthority == nil {
			return errors.New("NewAuthority parameter is not set")
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
	}
	return nil
}

func (inst *SetAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewAuthority", *inst.NewAuthority))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("candyGuard", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj SetAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAuthority` param:
	err = encoder.Encode(obj.NewAuthority)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAuthority`:
	err = decoder.Decode(&obj.NewAuthority)
	if err != nil {
		return err
	}
	return nil
}

// NewSetAuthorityInstruction declares a new SetAuthority instruction with the provided parameters and accounts.
func NewSetAuthorityInstruction(
	// Parameters:
	newAuthority ag_solanago.PublicKey,
	// Accounts:
	candyGuard ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetAuthority {
	return NewSetAuthorityInstructionBuilder().
		SetNewAuthority(newAuthority).
		SetCandyGuardAccount(candyGuard).
		SetAuthorityAccount(authority)
}