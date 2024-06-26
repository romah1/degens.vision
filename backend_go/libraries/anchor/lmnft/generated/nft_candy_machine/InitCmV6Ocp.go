// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitCmV6Ocp is the `initCmV6Ocp` instruction.
type InitCmV6Ocp struct {
	Data             *CandyMachineDataV3
	Seed             *ag_solanago.PublicKey
	ThawDate         *int64 `bin:"optional"`
	EnforceRoyalties *bool
	Name             *string
	Collection       *ag_solanago.PublicKey `bin:"optional"`

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] wallet
	//
	// [2] = [SIGNER] authority
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitCmV6OcpInstructionBuilder creates a new `InitCmV6Ocp` instruction builder.
func NewInitCmV6OcpInstructionBuilder() *InitCmV6Ocp {
	nd := &InitCmV6Ocp{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *InitCmV6Ocp) SetData(data CandyMachineDataV3) *InitCmV6Ocp {
	inst.Data = &data
	return inst
}

// SetSeed sets the "seed" parameter.
func (inst *InitCmV6Ocp) SetSeed(seed ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.Seed = &seed
	return inst
}

// SetThawDate sets the "thawDate" parameter.
func (inst *InitCmV6Ocp) SetThawDate(thawDate int64) *InitCmV6Ocp {
	inst.ThawDate = &thawDate
	return inst
}

// SetEnforceRoyalties sets the "enforceRoyalties" parameter.
func (inst *InitCmV6Ocp) SetEnforceRoyalties(enforceRoyalties bool) *InitCmV6Ocp {
	inst.EnforceRoyalties = &enforceRoyalties
	return inst
}

// SetName sets the "name" parameter.
func (inst *InitCmV6Ocp) SetName(name string) *InitCmV6Ocp {
	inst.Name = &name
	return inst
}

// SetCollection sets the "collection" parameter.
func (inst *InitCmV6Ocp) SetCollection(collection ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.Collection = &collection
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *InitCmV6Ocp) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *InitCmV6Ocp) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWalletAccount sets the "wallet" account.
func (inst *InitCmV6Ocp) SetWalletAccount(wallet ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *InitCmV6Ocp) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitCmV6Ocp) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitCmV6Ocp) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *InitCmV6Ocp) SetPayerAccount(payer ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *InitCmV6Ocp) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitCmV6Ocp) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitCmV6Ocp {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitCmV6Ocp) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitCmV6Ocp) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitCmV6Ocp,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitCmV6Ocp) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitCmV6Ocp) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
		if inst.Seed == nil {
			return errors.New("Seed parameter is not set")
		}
		if inst.EnforceRoyalties == nil {
			return errors.New("EnforceRoyalties parameter is not set")
		}
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
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

func (inst *InitCmV6Ocp) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitCmV6Ocp")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("            Data", *inst.Data))
						paramsBranch.Child(ag_format.Param("            Seed", *inst.Seed))
						paramsBranch.Child(ag_format.Param("        ThawDate (OPT)", inst.ThawDate))
						paramsBranch.Child(ag_format.Param("EnforceRoyalties", *inst.EnforceRoyalties))
						paramsBranch.Child(ag_format.Param("            Name", *inst.Name))
						paramsBranch.Child(ag_format.Param("      Collection (OPT)", inst.Collection))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       wallet", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitCmV6Ocp) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `Seed` param:
	err = encoder.Encode(obj.Seed)
	if err != nil {
		return err
	}
	// Serialize `ThawDate` param (optional):
	{
		if obj.ThawDate == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ThawDate)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `EnforceRoyalties` param:
	err = encoder.Encode(obj.EnforceRoyalties)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *InitCmV6Ocp) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `Seed`:
	err = decoder.Decode(&obj.Seed)
	if err != nil {
		return err
	}
	// Deserialize `ThawDate` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ThawDate)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `EnforceRoyalties`:
	err = decoder.Decode(&obj.EnforceRoyalties)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewInitCmV6OcpInstruction declares a new InitCmV6Ocp instruction with the provided parameters and accounts.
func NewInitCmV6OcpInstruction(
	// Parameters:
	data CandyMachineDataV3,
	seed ag_solanago.PublicKey,
	thawDate int64,
	enforceRoyalties bool,
	name string,
	collection ag_solanago.PublicKey,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitCmV6Ocp {
	return NewInitCmV6OcpInstructionBuilder().
		SetData(data).
		SetSeed(seed).
		SetThawDate(thawDate).
		SetEnforceRoyalties(enforceRoyalties).
		SetName(name).
		SetCollection(collection).
		SetCandyMachineAccount(candyMachine).
		SetWalletAccount(wallet).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}
