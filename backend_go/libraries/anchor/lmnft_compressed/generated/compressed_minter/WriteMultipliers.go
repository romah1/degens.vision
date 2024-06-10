// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package compressed_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WriteMultipliers is the `writeMultipliers` instruction.
type WriteMultipliers struct {
	Start *uint32
	Data  *[]byte

	// [0] = [] machine
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [WRITE] multiplier
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWriteMultipliersInstructionBuilder creates a new `WriteMultipliers` instruction builder.
func NewWriteMultipliersInstructionBuilder() *WriteMultipliers {
	nd := &WriteMultipliers{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetStart sets the "start" parameter.
func (inst *WriteMultipliers) SetStart(start uint32) *WriteMultipliers {
	inst.Start = &start
	return inst
}

// SetData sets the "data" parameter.
func (inst *WriteMultipliers) SetData(data []byte) *WriteMultipliers {
	inst.Data = &data
	return inst
}

// SetMachineAccount sets the "machine" account.
func (inst *WriteMultipliers) SetMachineAccount(machine ag_solanago.PublicKey) *WriteMultipliers {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(machine)
	return inst
}

// GetMachineAccount gets the "machine" account.
func (inst *WriteMultipliers) GetMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WriteMultipliers) SetAuthorityAccount(authority ag_solanago.PublicKey) *WriteMultipliers {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WriteMultipliers) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMultiplierAccount sets the "multiplier" account.
func (inst *WriteMultipliers) SetMultiplierAccount(multiplier ag_solanago.PublicKey) *WriteMultipliers {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(multiplier).WRITE()
	return inst
}

// GetMultiplierAccount gets the "multiplier" account.
func (inst *WriteMultipliers) GetMultiplierAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *WriteMultipliers) SetPayerAccount(payer ag_solanago.PublicKey) *WriteMultipliers {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *WriteMultipliers) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *WriteMultipliers) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *WriteMultipliers {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *WriteMultipliers) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst WriteMultipliers) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WriteMultipliers,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WriteMultipliers) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WriteMultipliers) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Start == nil {
			return errors.New("Start parameter is not set")
		}
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Machine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Multiplier is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *WriteMultipliers) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WriteMultipliers")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Start", *inst.Start))
						paramsBranch.Child(ag_format.Param(" Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      machine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   multiplier", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj WriteMultipliers) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WriteMultipliers) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewWriteMultipliersInstruction declares a new WriteMultipliers instruction with the provided parameters and accounts.
func NewWriteMultipliersInstruction(
	// Parameters:
	start uint32,
	data []byte,
	// Accounts:
	machine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	multiplier ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *WriteMultipliers {
	return NewWriteMultipliersInstructionBuilder().
		SetStart(start).
		SetData(data).
		SetMachineAccount(machine).
		SetAuthorityAccount(authority).
		SetMultiplierAccount(multiplier).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}