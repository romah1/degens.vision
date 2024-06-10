// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package bubblegum

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelRedeem is the `cancelRedeem` instruction.
type CancelRedeem struct {
	Root *[32]uint8

	// [0] = [] treeAuthority
	//
	// [1] = [WRITE, SIGNER] leafOwner
	//
	// [2] = [WRITE] merkleTree
	//
	// [3] = [WRITE] voucher
	//
	// [4] = [] logWrapper
	//
	// [5] = [] compressionProgram
	//
	// [6] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelRedeemInstructionBuilder creates a new `CancelRedeem` instruction builder.
func NewCancelRedeemInstructionBuilder() *CancelRedeem {
	nd := &CancelRedeem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetRoot sets the "root" parameter.
func (inst *CancelRedeem) SetRoot(root [32]uint8) *CancelRedeem {
	inst.Root = &root
	return inst
}

// SetTreeAuthorityAccount sets the "treeAuthority" account.
func (inst *CancelRedeem) SetTreeAuthorityAccount(treeAuthority ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(treeAuthority)
	return inst
}

// GetTreeAuthorityAccount gets the "treeAuthority" account.
func (inst *CancelRedeem) GetTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLeafOwnerAccount sets the "leafOwner" account.
func (inst *CancelRedeem) SetLeafOwnerAccount(leafOwner ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(leafOwner).WRITE().SIGNER()
	return inst
}

// GetLeafOwnerAccount gets the "leafOwner" account.
func (inst *CancelRedeem) GetLeafOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMerkleTreeAccount sets the "merkleTree" account.
func (inst *CancelRedeem) SetMerkleTreeAccount(merkleTree ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(merkleTree).WRITE()
	return inst
}

// GetMerkleTreeAccount gets the "merkleTree" account.
func (inst *CancelRedeem) GetMerkleTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetVoucherAccount sets the "voucher" account.
func (inst *CancelRedeem) SetVoucherAccount(voucher ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(voucher).WRITE()
	return inst
}

// GetVoucherAccount gets the "voucher" account.
func (inst *CancelRedeem) GetVoucherAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLogWrapperAccount sets the "logWrapper" account.
func (inst *CancelRedeem) SetLogWrapperAccount(logWrapper ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(logWrapper)
	return inst
}

// GetLogWrapperAccount gets the "logWrapper" account.
func (inst *CancelRedeem) GetLogWrapperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCompressionProgramAccount sets the "compressionProgram" account.
func (inst *CancelRedeem) SetCompressionProgramAccount(compressionProgram ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(compressionProgram)
	return inst
}

// GetCompressionProgramAccount gets the "compressionProgram" account.
func (inst *CancelRedeem) GetCompressionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CancelRedeem) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CancelRedeem {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CancelRedeem) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CancelRedeem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelRedeem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelRedeem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelRedeem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Root == nil {
			return errors.New("Root parameter is not set")
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
			return errors.New("accounts.MerkleTree is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Voucher is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LogWrapper is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CompressionProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CancelRedeem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelRedeem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Root", *inst.Root))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     treeAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         leafOwner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        merkleTree", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           voucher", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        logWrapper", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("compressionProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CancelRedeem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelRedeem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelRedeemInstruction declares a new CancelRedeem instruction with the provided parameters and accounts.
func NewCancelRedeemInstruction(
	// Parameters:
	root [32]uint8,
	// Accounts:
	treeAuthority ag_solanago.PublicKey,
	leafOwner ag_solanago.PublicKey,
	merkleTree ag_solanago.PublicKey,
	voucher ag_solanago.PublicKey,
	logWrapper ag_solanago.PublicKey,
	compressionProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CancelRedeem {
	return NewCancelRedeemInstructionBuilder().
		SetRoot(root).
		SetTreeAuthorityAccount(treeAuthority).
		SetLeafOwnerAccount(leafOwner).
		SetMerkleTreeAccount(merkleTree).
		SetVoucherAccount(voucher).
		SetLogWrapperAccount(logWrapper).
		SetCompressionProgramAccount(compressionProgram).
		SetSystemProgramAccount(systemProgram)
}
