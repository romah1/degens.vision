// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintV1 is the `mintV1` instruction.
type MintV1 struct {
	Message *MetadataArgs

	// [0] = [WRITE] treeAuthority
	//
	// [1] = [] leafOwner
	//
	// [2] = [] leafDelegate
	//
	// [3] = [WRITE] merkleTree
	//
	// [4] = [SIGNER] payer
	//
	// [5] = [SIGNER] treeDelegate
	//
	// [6] = [] logWrapper
	//
	// [7] = [] compressionProgram
	//
	// [8] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintV1InstructionBuilder creates a new `MintV1` instruction builder.
func NewMintV1InstructionBuilder() *MintV1 {
	nd := &MintV1{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetMessage sets the "message" parameter.
func (inst *MintV1) SetMessage(message MetadataArgs) *MintV1 {
	inst.Message = &message
	return inst
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *MintV1) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority).WRITE()
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *MintV1) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLeafOwnerAccount sets the "leafOwner" account.
func (inst *MintV1) SetLeafOwnerAccount(leafOwner ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(leafOwner)
	return inst
}

// GetLeafOwnerAccount gets the "leafOwner" account.
func (inst *MintV1) GetLeafOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLeafDelegateAccount sets the "leafDelegate" account.
func (inst *MintV1) SetLeafDelegateAccount(leafDelegate ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(leafDelegate)
	return inst
}

// GetLeafDelegateAccount gets the "leafDelegate" account.
func (inst *MintV1) GetLeafDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMerkleTreeAccount sets the "merkleTree" account.
func (inst *MintV1) SetMerkleTreeAccount(merkleTree ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(merkleTree).WRITE()
	return inst
}

// GetMerkleTreeAccount gets the "merkleTree" account.
func (inst *MintV1) GetMerkleTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintV1) SetPayerAccount(payer ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintV1) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTreeDelegateAccount sets the "treeDelegate" account.
func (inst *MintV1) SetTreeDelegateAccount(treeDelegate ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(treeDelegate).SIGNER()
	return inst
}

// GetTreeDelegateAccount gets the "treeDelegate" account.
func (inst *MintV1) GetTreeDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetLogWrapperAccount sets the "logWrapper" account.
func (inst *MintV1) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
func (inst *MintV1) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCompressionProgramAccount sets the "compressionProgram" account.
func (inst *MintV1) SetCompressionProgramAccount(compressionProgram ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(compressionProgram)
	return inst
}

// GetCompressionProgramAccount gets the "compressionProgram" account.
func (inst *MintV1) GetCompressionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintV1) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintV1 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintV1) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst MintV1) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintV1,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintV1) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintV1) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TreeAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LeafOwner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.LeafDelegate is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MerkleTree is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TreeDelegate is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.LogWrapper is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CompressionProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *MintV1) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintV1")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         leafOwner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      leafDelegate", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        merkleTree", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      treeDelegate", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        logWrapper", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("compressionProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj MintV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewMintV1Instruction declares a new MintV1 instruction with the provided parameters and accounts.
func NewMintV1Instruction(
	// Parameters:
	message MetadataArgs,
	// Accounts:
	treeAuthority ag_solanago.PublicKey,
	leafOwner ag_solanago.PublicKey,
	leafDelegate ag_solanago.PublicKey,
	merkleTree ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	treeDelegate ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey,
	compressionProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *MintV1 {
	return NewMintV1InstructionBuilder().
		SetMessage(message).
		SetTreeAuthorityAccount(treeAuthority).
		SetLeafOwnerAccount(leafOwner).
		SetLeafDelegateAccount(leafDelegate).
		SetMerkleTreeAccount(merkleTree).
		SetPayerAccount(payer).
		SetTreeDelegateAccount(treeDelegate).
		SetLogWrapperAccount(logWrapper).
		SetCompressionProgramAccount(compressionProgram).
		SetSystemProgramAccount(systemProgram)
}
