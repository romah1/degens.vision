// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetName is the `setName` instruction.
type SetName struct {
	Name       *string
	ReturnAuth *bool

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [] candyMachine
	//
	// [2] = [WRITE] metadata
	//
	// [3] = [] mint
	//
	// [4] = [] masterEdition
	//
	// [5] = [] authorizationRules
	//
	// [6] = [] authorizationRulesProgram
	//
	// [7] = [] instructionSysvarAccount
	//
	// [8] = [] tokenMetadataProgram
	//
	// [9] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetNameInstructionBuilder creates a new `SetName` instruction builder.
func NewSetNameInstructionBuilder() *SetName {
	nd := &SetName{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetName sets the "name" parameter.
func (inst *SetName) SetName(name string) *SetName {
	inst.Name = &name
	return inst
}

// SetReturnAuth sets the "returnAuth" parameter.
func (inst *SetName) SetReturnAuth(returnAuth bool) *SetName {
	inst.ReturnAuth = &returnAuth
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *SetName) SetPayerAccount(payer ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *SetName) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *SetName) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine)
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *SetName) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *SetName) SetMetadataAccount(metadata ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *SetName) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
func (inst *SetName) SetMintAccount(mint ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *SetName) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *SetName) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(masterEdition)
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *SetName) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAuthorizationRulesAccount sets the "authorizationRules" account.
func (inst *SetName) SetAuthorizationRulesAccount(authorizationRules ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(authorizationRules)
	return inst
}

// GetAuthorizationRulesAccount gets the "authorizationRules" account.
func (inst *SetName) GetAuthorizationRulesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAuthorizationRulesProgramAccount sets the "authorizationRulesProgram" account.
func (inst *SetName) SetAuthorizationRulesProgramAccount(authorizationRulesProgram ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(authorizationRulesProgram)
	return inst
}

// GetAuthorizationRulesProgramAccount gets the "authorizationRulesProgram" account.
func (inst *SetName) GetAuthorizationRulesProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetInstructionSysvarAccountAccount sets the "instructionSysvarAccount" account.
func (inst *SetName) SetInstructionSysvarAccountAccount(instructionSysvarAccount ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(instructionSysvarAccount)
	return inst
}

// GetInstructionSysvarAccountAccount gets the "instructionSysvarAccount" account.
func (inst *SetName) GetInstructionSysvarAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *SetName) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *SetName) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetName) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetName {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetName) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst SetName) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetName,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetName) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetName) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
		if inst.ReturnAuth == nil {
			return errors.New("ReturnAuth parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuthorizationRules is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AuthorizationRulesProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.InstructionSysvarAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *SetName) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetName")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      Name", *inst.Name))
						paramsBranch.Child(ag_format.Param("ReturnAuth", *inst.ReturnAuth))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                     mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            masterEdition", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       authorizationRules", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("authorizationRulesProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("        instructionSysvar", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     tokenMetadataProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("            systemProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj SetName) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `ReturnAuth` param:
	err = encoder.Encode(obj.ReturnAuth)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetName) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `ReturnAuth`:
	err = decoder.Decode(&obj.ReturnAuth)
	if err != nil {
		return err
	}
	return nil
}

// NewSetNameInstruction declares a new SetName instruction with the provided parameters and accounts.
func NewSetNameInstruction(
	// Parameters:
	name string,
	returnAuth bool,
	// Accounts:
	payer ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	authorizationRules ag_solanago.PublicKey,
	authorizationRulesProgram ag_solanago.PublicKey,
	instructionSysvarAccount ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetName {
	return NewSetNameInstructionBuilder().
		SetName(name).
		SetReturnAuth(returnAuth).
		SetPayerAccount(payer).
		SetCandyMachineAccount(candyMachine).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMasterEditionAccount(masterEdition).
		SetAuthorizationRulesAccount(authorizationRules).
		SetAuthorizationRulesProgramAccount(authorizationRulesProgram).
		SetInstructionSysvarAccountAccount(instructionSysvarAccount).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetSystemProgramAccount(systemProgram)
}
