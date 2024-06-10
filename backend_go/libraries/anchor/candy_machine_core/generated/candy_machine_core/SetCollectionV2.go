// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine_core

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetCollectionV2 is the `setCollectionV2` instruction.
type SetCollectionV2 struct {

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [WRITE] authorityPda
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] collectionUpdateAuthority
	//
	// [5] = [] collectionMint
	//
	// [6] = [WRITE] collectionMetadata
	//
	// [7] = [WRITE] collectionDelegateRecord
	//
	// [8] = [SIGNER] newCollectionUpdateAuthority
	//
	// [9] = [] newCollectionMint
	//
	// [10] = [WRITE] newCollectionMetadata
	//
	// [11] = [] newCollectionMasterEdition
	//
	// [12] = [WRITE] newCollectionDelegateRecord
	//
	// [13] = [] tokenMetadataProgram
	//
	// [14] = [] systemProgram
	//
	// [15] = [] sysvarInstructions
	//
	// [16] = [] authorizationRulesProgram
	//
	// [17] = [] authorizationRules
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetCollectionV2InstructionBuilder creates a new `SetCollectionV2` instruction builder.
func NewSetCollectionV2InstructionBuilder() *SetCollectionV2 {
	nd := &SetCollectionV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 18),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *SetCollectionV2) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *SetCollectionV2) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetCollectionV2) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetCollectionV2) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityPdaAccount sets the "authorityPda" account.
func (inst *SetCollectionV2) SetAuthorityPdaAccount(authorityPda ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authorityPda).WRITE()
	return inst
}

// GetAuthorityPdaAccount gets the "authorityPda" account.
func (inst *SetCollectionV2) GetAuthorityPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *SetCollectionV2) SetPayerAccount(payer ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *SetCollectionV2) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionUpdateAuthorityAccount sets the "collectionUpdateAuthority" account.
func (inst *SetCollectionV2) SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionUpdateAuthority)
	return inst
}

// GetCollectionUpdateAuthorityAccount gets the "collectionUpdateAuthority" account.
func (inst *SetCollectionV2) GetCollectionUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *SetCollectionV2) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *SetCollectionV2) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
func (inst *SetCollectionV2) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(collectionMetadata).WRITE()
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
func (inst *SetCollectionV2) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionDelegateRecordAccount sets the "collectionDelegateRecord" account.
func (inst *SetCollectionV2) SetCollectionDelegateRecordAccount(collectionDelegateRecord ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionDelegateRecord).WRITE()
	return inst
}

// GetCollectionDelegateRecordAccount gets the "collectionDelegateRecord" account.
func (inst *SetCollectionV2) GetCollectionDelegateRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetNewCollectionUpdateAuthorityAccount sets the "newCollectionUpdateAuthority" account.
func (inst *SetCollectionV2) SetNewCollectionUpdateAuthorityAccount(newCollectionUpdateAuthority ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(newCollectionUpdateAuthority).SIGNER()
	return inst
}

// GetNewCollectionUpdateAuthorityAccount gets the "newCollectionUpdateAuthority" account.
func (inst *SetCollectionV2) GetNewCollectionUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetNewCollectionMintAccount sets the "newCollectionMint" account.
func (inst *SetCollectionV2) SetNewCollectionMintAccount(newCollectionMint ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(newCollectionMint)
	return inst
}

// GetNewCollectionMintAccount gets the "newCollectionMint" account.
func (inst *SetCollectionV2) GetNewCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetNewCollectionMetadataAccount sets the "newCollectionMetadata" account.
func (inst *SetCollectionV2) SetNewCollectionMetadataAccount(newCollectionMetadata ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(newCollectionMetadata).WRITE()
	return inst
}

// GetNewCollectionMetadataAccount gets the "newCollectionMetadata" account.
func (inst *SetCollectionV2) GetNewCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetNewCollectionMasterEditionAccount sets the "newCollectionMasterEdition" account.
func (inst *SetCollectionV2) SetNewCollectionMasterEditionAccount(newCollectionMasterEdition ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(newCollectionMasterEdition)
	return inst
}

// GetNewCollectionMasterEditionAccount gets the "newCollectionMasterEdition" account.
func (inst *SetCollectionV2) GetNewCollectionMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetNewCollectionDelegateRecordAccount sets the "newCollectionDelegateRecord" account.
func (inst *SetCollectionV2) SetNewCollectionDelegateRecordAccount(newCollectionDelegateRecord ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(newCollectionDelegateRecord).WRITE()
	return inst
}

// GetNewCollectionDelegateRecordAccount gets the "newCollectionDelegateRecord" account.
func (inst *SetCollectionV2) GetNewCollectionDelegateRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *SetCollectionV2) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *SetCollectionV2) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetCollectionV2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetCollectionV2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
func (inst *SetCollectionV2) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
func (inst *SetCollectionV2) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetAuthorizationRulesProgramAccount sets the "authorizationRulesProgram" account.
func (inst *SetCollectionV2) SetAuthorizationRulesProgramAccount(authorizationRulesProgram ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(authorizationRulesProgram)
	return inst
}

// GetAuthorizationRulesProgramAccount gets the "authorizationRulesProgram" account.
func (inst *SetCollectionV2) GetAuthorizationRulesProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetAuthorizationRulesAccount sets the "authorizationRules" account.
func (inst *SetCollectionV2) SetAuthorizationRulesAccount(authorizationRules ag_solanago.PublicKey) *SetCollectionV2 {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(authorizationRules)
	return inst
}

// GetAuthorizationRulesAccount gets the "authorizationRules" account.
func (inst *SetCollectionV2) GetAuthorizationRulesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

func (inst SetCollectionV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetCollectionV2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetCollectionV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetCollectionV2) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuthorityPda is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionDelegateRecord is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.NewCollectionUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.NewCollectionMint is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.NewCollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.NewCollectionMasterEdition is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.NewCollectionDelegateRecord is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.AuthorizationRulesProgram is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.AuthorizationRules is not set")
		}
	}
	return nil
}

func (inst *SetCollectionV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetCollectionV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=18]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                authorityPda", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                       payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   collectionUpdateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              collectionMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          collectionMetadata", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("    collectionDelegateRecord", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("newCollectionUpdateAuthority", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           newCollectionMint", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       newCollectionMetadata", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("  newCollectionMasterEdition", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta(" newCollectionDelegateRecord", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("        tokenMetadataProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("               systemProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("          sysvarInstructions", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("   authorizationRulesProgram", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("          authorizationRules", inst.AccountMetaSlice.Get(17)))
					})
				})
		})
}

func (obj SetCollectionV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetCollectionV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetCollectionV2Instruction declares a new SetCollectionV2 instruction with the provided parameters and accounts.
func NewSetCollectionV2Instruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	authorityPda ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	collectionUpdateAuthority ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	collectionDelegateRecord ag_solanago.PublicKey,
	newCollectionUpdateAuthority ag_solanago.PublicKey,
	newCollectionMint ag_solanago.PublicKey,
	newCollectionMetadata ag_solanago.PublicKey,
	newCollectionMasterEdition ag_solanago.PublicKey,
	newCollectionDelegateRecord ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	authorizationRulesProgram ag_solanago.PublicKey,
	authorizationRules ag_solanago.PublicKey) *SetCollectionV2 {
	return NewSetCollectionV2InstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority).
		SetAuthorityPdaAccount(authorityPda).
		SetPayerAccount(payer).
		SetCollectionUpdateAuthorityAccount(collectionUpdateAuthority).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetCollectionDelegateRecordAccount(collectionDelegateRecord).
		SetNewCollectionUpdateAuthorityAccount(newCollectionUpdateAuthority).
		SetNewCollectionMintAccount(newCollectionMint).
		SetNewCollectionMetadataAccount(newCollectionMetadata).
		SetNewCollectionMasterEditionAccount(newCollectionMasterEdition).
		SetNewCollectionDelegateRecordAccount(newCollectionDelegateRecord).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetAuthorizationRulesProgramAccount(authorizationRulesProgram).
		SetAuthorizationRulesAccount(authorizationRules)
}