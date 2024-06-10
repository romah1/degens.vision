// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UnverifyCreator is the `unverifyCreator` instruction.
type UnverifyCreator struct {
	Root        *[32]uint8
	DataHash    *[32]uint8
	CreatorHash *[32]uint8
	Nonce       *uint64
	Index       *uint32
	Message     *MetadataArgs

	// [0] = [] treeAuthority
	//
	// [1] = [] leafOwner
	//
	// [2] = [] leafDelegate
	//
	// [3] = [WRITE] merkleTree
	//
	// [4] = [SIGNER] payer
	//
	// [5] = [SIGNER] creator
	//
	// [6] = [] logWrapper
	//
	// [7] = [] compressionProgram
	//
	// [8] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUnverifyCreatorInstructionBuilder creates a new `UnverifyCreator` instruction builder.
func NewUnverifyCreatorInstructionBuilder() *UnverifyCreator {
	nd := &UnverifyCreator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetRoot sets the "root" parameter.
func (inst *UnverifyCreator) SetRoot(root [32]uint8) *UnverifyCreator {
	inst.Root = &root
	return inst
}

// SetDataHash sets the "dataHash" parameter.
func (inst *UnverifyCreator) SetDataHash(dataHash [32]uint8) *UnverifyCreator {
	inst.DataHash = &dataHash
	return inst
}

// SetCreatorHash sets the "creatorHash" parameter.
func (inst *UnverifyCreator) SetCreatorHash(creatorHash [32]uint8) *UnverifyCreator {
	inst.CreatorHash = &creatorHash
	return inst
}

// SetNonce sets the "nonce" parameter.
func (inst *UnverifyCreator) SetNonce(nonce uint64) *UnverifyCreator {
	inst.Nonce = &nonce
	return inst
}

// SetIndex sets the "index" parameter.
func (inst *UnverifyCreator) SetIndex(index uint32) *UnverifyCreator {
	inst.Index = &index
	return inst
}

// SetMessage sets the "message" parameter.
func (inst *UnverifyCreator) SetMessage(message MetadataArgs) *UnverifyCreator {
	inst.Message = &message
	return inst
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *UnverifyCreator) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority)
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *UnverifyCreator) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLeafOwnerAccount sets the "leafOwner" account.
func (inst *UnverifyCreator) SetLeafOwnerAccount(leafOwner ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(leafOwner)
	return inst
}

// GetLeafOwnerAccount gets the "leafOwner" account.
func (inst *UnverifyCreator) GetLeafOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLeafDelegateAccount sets the "leafDelegate" account.
func (inst *UnverifyCreator) SetLeafDelegateAccount(leafDelegate ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(leafDelegate)
	return inst
}

// GetLeafDelegateAccount gets the "leafDelegate" account.
func (inst *UnverifyCreator) GetLeafDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMerkleTreeAccount sets the "merkleTree" account.
func (inst *UnverifyCreator) SetMerkleTreeAccount(merkleTree ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(merkleTree).WRITE()
	return inst
}

// GetMerkleTreeAccount gets the "merkleTree" account.
func (inst *UnverifyCreator) GetMerkleTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *UnverifyCreator) SetPayerAccount(payer ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *UnverifyCreator) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCreatorAccount sets the "creator" account.
func (inst *UnverifyCreator) SetCreatorAccount(creator ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(creator).SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
func (inst *UnverifyCreator) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetLogWrapperAccount sets the "logWrapper" account.
func (inst *UnverifyCreator) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
func (inst *UnverifyCreator) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCompressionProgramAccount sets the "compressionProgram" account.
func (inst *UnverifyCreator) SetCompressionProgramAccount(compressionProgram ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(compressionProgram)
	return inst
}

// GetCompressionProgramAccount gets the "compressionProgram" account.
func (inst *UnverifyCreator) GetCompressionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UnverifyCreator) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UnverifyCreator {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UnverifyCreator) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst UnverifyCreator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UnverifyCreator,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UnverifyCreator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UnverifyCreator) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Root == nil {
			return errors.New("Root parameter is not set")
		}
		if inst.DataHash == nil {
			return errors.New("DataHash parameter is not set")
		}
		if inst.CreatorHash == nil {
			return errors.New("CreatorHash parameter is not set")
		}
		if inst.Nonce == nil {
			return errors.New("Nonce parameter is not set")
		}
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
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
			return errors.New("accounts.Creator is not set")
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

func (inst *UnverifyCreator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UnverifyCreator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("       Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("   DataHash", *inst.DataHash))
						paramsBranch.Child(ag_format.Param("CreatorHash", *inst.CreatorHash))
						paramsBranch.Child(ag_format.Param("      Nonce", *inst.Nonce))
						paramsBranch.Child(ag_format.Param("      Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("    Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         leafOwner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      leafDelegate", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        merkleTree", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           creator", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        logWrapper", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("compressionProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj UnverifyCreator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `DataHash` param:
	err = encoder.Encode(obj.DataHash)
	if err != nil {
		return err
	}
	// Serialize `CreatorHash` param:
	err = encoder.Encode(obj.CreatorHash)
	if err != nil {
		return err
	}
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UnverifyCreator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `DataHash`:
	err = decoder.Decode(&obj.DataHash)
	if err != nil {
		return err
	}
	// Deserialize `CreatorHash`:
	err = decoder.Decode(&obj.CreatorHash)
	if err != nil {
		return err
	}
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewUnverifyCreatorInstruction declares a new UnverifyCreator instruction with the provided parameters and accounts.
func NewUnverifyCreatorInstruction(
	// Parameters:
	root [32]uint8,
	dataHash [32]uint8,
	creatorHash [32]uint8,
	nonce uint64,
	index uint32,
	message MetadataArgs,
	// Accounts:
	treeAuthority ag_solanago.PublicKey,
	leafOwner ag_solanago.PublicKey,
	leafDelegate ag_solanago.PublicKey,
	merkleTree ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	creator ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey,
	compressionProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UnverifyCreator {
	return NewUnverifyCreatorInstructionBuilder().
		SetRoot(root).
		SetDataHash(dataHash).
		SetCreatorHash(creatorHash).
		SetNonce(nonce).
		SetIndex(index).
		SetMessage(message).
		SetTreeAuthorityAccount(treeAuthority).
		SetLeafOwnerAccount(leafOwner).
		SetLeafDelegateAccount(leafDelegate).
		SetMerkleTreeAccount(merkleTree).
		SetPayerAccount(payer).
		SetCreatorAccount(creator).
		SetLogWrapperAccount(logWrapper).
		SetCompressionProgramAccount(compressionProgram).
		SetSystemProgramAccount(systemProgram)
}
