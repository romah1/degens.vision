// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_guard

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Route is the `route` instruction.
type Route struct {
	Args  *RouteArgs
	Label *string `bin:"optional"`

	// [0] = [] candyGuard
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [WRITE, SIGNER] payer
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRouteInstructionBuilder creates a new `Route` instruction builder.
func NewRouteInstructionBuilder() *Route {
	nd := &Route{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *Route) SetArgs(args RouteArgs) *Route {
	inst.Args = &args
	return inst
}

// SetLabel sets the "label" parameter.
func (inst *Route) SetLabel(label string) *Route {
	inst.Label = &label
	return inst
}

// SetCandyGuardAccount sets the "candyGuard" account.
func (inst *Route) SetCandyGuardAccount(candyGuard ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyGuard)
	return inst
}

// GetCandyGuardAccount gets the "candyGuard" account.
func (inst *Route) GetCandyGuardAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *Route) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *Route) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *Route) SetPayerAccount(payer ag_solanago.PublicKey) *Route {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *Route) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst Route) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Route,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Route) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Route) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyGuard is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
	}
	return nil
}

func (inst *Route) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Route")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" Args", *inst.Args))
						paramsBranch.Child(ag_format.Param("Label (OPT)", inst.Label))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  candyGuard", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       payer", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj Route) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	// Serialize `Label` param (optional):
	{
		if obj.Label == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Label)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *Route) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	// Deserialize `Label` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Label)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewRouteInstruction declares a new Route instruction with the provided parameters and accounts.
func NewRouteInstruction(
	// Parameters:
	args RouteArgs,
	label string,
	// Accounts:
	candyGuard ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	payer ag_solanago.PublicKey) *Route {
	return NewRouteInstructionBuilder().
		SetArgs(args).
		SetLabel(label).
		SetCandyGuardAccount(candyGuard).
		SetCandyMachineAccount(candyMachine).
		SetPayerAccount(payer)
}
