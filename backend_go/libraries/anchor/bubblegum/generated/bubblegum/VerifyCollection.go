// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// VerifyCollection is the `verifyCollection` instruction.
type VerifyCollection struct {
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
	// [5] = [] treeDelegate
	//
	// [6] = [SIGNER] collectionAuthority
	//
	// [7] = [] collectionAuthorityRecordPda
	//
	// [8] = [] collectionMint
	//
	// [9] = [WRITE] collectionMetadata
	//
	// [10] = [] editionAccount
	//
	// [11] = [] bubblegumSigner
	//
	// [12] = [] logWrapper
	//
	// [13] = [] compressionProgram
	//
	// [14] = [] tokenMetadataProgram
	//
	// [15] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVerifyCollectionInstructionBuilder creates a new `VerifyCollection` instruction builder.
func NewVerifyCollectionInstructionBuilder() *VerifyCollection {
	nd := &VerifyCollection{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetRoot sets the "root" parameter.
func (inst *VerifyCollection) SetRoot(root [32]uint8) *VerifyCollection {
	inst.Root = &root
	return inst
}

// SetDataHash sets the "dataHash" parameter.
func (inst *VerifyCollection) SetDataHash(dataHash [32]uint8) *VerifyCollection {
	inst.DataHash = &dataHash
	return inst
}

// SetCreatorHash sets the "creatorHash" parameter.
func (inst *VerifyCollection) SetCreatorHash(creatorHash [32]uint8) *VerifyCollection {
	inst.CreatorHash = &creatorHash
	return inst
}

// SetNonce sets the "nonce" parameter.
func (inst *VerifyCollection) SetNonce(nonce uint64) *VerifyCollection {
	inst.Nonce = &nonce
	return inst
}

// SetIndex sets the "index" parameter.
func (inst *VerifyCollection) SetIndex(index uint32) *VerifyCollection {
	inst.Index = &index
	return inst
}

// SetMessage sets the "message" parameter.
func (inst *VerifyCollection) SetMessage(message MetadataArgs) *VerifyCollection {
	inst.Message = &message
	return inst
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *VerifyCollection) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority)
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *VerifyCollection) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLeafOwnerAccount sets the "leafOwner" account.
func (inst *VerifyCollection) SetLeafOwnerAccount(leafOwner ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(leafOwner)
	return inst
}

// GetLeafOwnerAccount gets the "leafOwner" account.
func (inst *VerifyCollection) GetLeafOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLeafDelegateAccount sets the "leafDelegate" account.
func (inst *VerifyCollection) SetLeafDelegateAccount(leafDelegate ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(leafDelegate)
	return inst
}

// GetLeafDelegateAccount gets the "leafDelegate" account.
func (inst *VerifyCollection) GetLeafDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMerkleTreeAccount sets the "merkleTree" account.
func (inst *VerifyCollection) SetMerkleTreeAccount(merkleTree ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(merkleTree).WRITE()
	return inst
}

// GetMerkleTreeAccount gets the "merkleTree" account.
func (inst *VerifyCollection) GetMerkleTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *VerifyCollection) SetPayerAccount(payer ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *VerifyCollection) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTreeDelegateAccount sets the "treeDelegate" account.
func (inst *VerifyCollection) SetTreeDelegateAccount(treeDelegate ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(treeDelegate)
	return inst
}

// GetTreeDelegateAccount gets the "treeDelegate" account.
func (inst *VerifyCollection) GetTreeDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCollectionAuthorityAccount sets the "collectionAuthority" account.
func (inst *VerifyCollection) SetCollectionAuthorityAccount(collectionAuthority ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(collectionAuthority).SIGNER()
	return inst
}

// GetCollectionAuthorityAccount gets the "collectionAuthority" account.
func (inst *VerifyCollection) GetCollectionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionAuthorityRecordPdaAccount sets the "collectionAuthorityRecordPda" account.
func (inst *VerifyCollection) SetCollectionAuthorityRecordPdaAccount(collectionAuthorityRecordPda ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionAuthorityRecordPda)
	return inst
}

// GetCollectionAuthorityRecordPdaAccount gets the "collectionAuthorityRecordPda" account.
func (inst *VerifyCollection) GetCollectionAuthorityRecordPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *VerifyCollection) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *VerifyCollection) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
func (inst *VerifyCollection) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(collectionMetadata).WRITE()
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
func (inst *VerifyCollection) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetEditionAccountAccount sets the "editionAccount" account.
func (inst *VerifyCollection) SetEditionAccountAccount(editionAccount ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(editionAccount)
	return inst
}

// GetEditionAccountAccount gets the "editionAccount" account.
func (inst *VerifyCollection) GetEditionAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetBubblegumSignerAccount sets the "bubblegumSigner" account.
func (inst *VerifyCollection) SetBubblegumSignerAccount(bubblegumSigner ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(bubblegumSigner)
	return inst
}

// GetBubblegumSignerAccount gets the "bubblegumSigner" account.
func (inst *VerifyCollection) GetBubblegumSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetLogWrapperAccount sets the "logWrapper" account.
func (inst *VerifyCollection) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
func (inst *VerifyCollection) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetCompressionProgramAccount sets the "compressionProgram" account.
func (inst *VerifyCollection) SetCompressionProgramAccount(compressionProgram ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(compressionProgram)
	return inst
}

// GetCompressionProgramAccount gets the "compressionProgram" account.
func (inst *VerifyCollection) GetCompressionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *VerifyCollection) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *VerifyCollection) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *VerifyCollection) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *VerifyCollection {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *VerifyCollection) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst VerifyCollection) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_VerifyCollection,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst VerifyCollection) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *VerifyCollection) Validate() error {
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
			return errors.New("accounts.TreeDelegate is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CollectionAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionAuthorityRecordPda is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.EditionAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.BubblegumSigner is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.LogWrapper is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.CompressionProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *VerifyCollection) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("VerifyCollection")).
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
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   leafOwner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                leafDelegate", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  merkleTree", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                       payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                treeDelegate", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         collectionAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("collectionAuthorityRecordPda", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              collectionMint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          collectionMetadata", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                     edition", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("             bubblegumSigner", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                  logWrapper", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("          compressionProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("        tokenMetadataProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("               systemProgram", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj VerifyCollection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *VerifyCollection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewVerifyCollectionInstruction declares a new VerifyCollection instruction with the provided parameters and accounts.
func NewVerifyCollectionInstruction(
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
	treeDelegate ag_solanago.PublicKey,
	collectionAuthority ag_solanago.PublicKey,
	collectionAuthorityRecordPda ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	editionAccount ag_solanago.PublicKey,
	bubblegumSigner ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey,
	compressionProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *VerifyCollection {
	return NewVerifyCollectionInstructionBuilder().
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
		SetTreeDelegateAccount(treeDelegate).
		SetCollectionAuthorityAccount(collectionAuthority).
		SetCollectionAuthorityRecordPdaAccount(collectionAuthorityRecordPda).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetEditionAccountAccount(editionAccount).
		SetBubblegumSignerAccount(bubblegumSigner).
		SetLogWrapperAccount(logWrapper).
		SetCompressionProgramAccount(compressionProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetSystemProgramAccount(systemProgram)
}
