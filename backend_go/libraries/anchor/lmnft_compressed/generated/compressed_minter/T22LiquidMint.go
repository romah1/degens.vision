// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package compressed_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// T22LiquidMint is the `t22LiquidMint` instruction.
type T22LiquidMint struct {
	Input *MachineMintInput

	// [0] = [WRITE, SIGNER] buyer
	//
	// [1] = [SIGNER] requiredSigner
	//
	// [2] = [WRITE] machine
	//
	// [3] = [] multipliers
	//
	// [4] = [WRITE] totalMints
	//
	// [5] = [WRITE] wallet
	//
	// [6] = [WRITE] wallet2
	//
	// [7] = [] systemProgram
	//
	// [8] = [WRITE] buyerWlTokenWallet
	//
	// [9] = [] wlNftMetadata
	//
	// [10] = [WRITE] mintsPerWlNft
	//
	// [11] = [WRITE] whitelistMint
	//
	// [12] = [WRITE] buyerPaymentTokenWallet
	//
	// [13] = [] instructionSysvarAccount
	//
	// [14] = [] tokenProgram
	//
	// [15] = [] token22Program
	//
	// [16] = [] tokenMetadataProgram
	//
	// [17] = [] accountCompressionProgram
	//
	// [18] = [] noopProgram
	//
	// [19] = [] recentSlothashes
	//
	// [20] = [] bubblegumProgram
	//
	// [21] = [] wlNftTree
	//
	// [22] = [WRITE] wlNftTreeAuthority
	//
	// [23] = [WRITE] liquidity
	//
	// [24] = [WRITE] deploymentFungibleTokenAccount
	//
	// [25] = [WRITE] deploymentNonFungibleTokenAccount
	//
	// [26] = [WRITE] deployment
	//
	// [27] = [WRITE] deploymentConfig
	//
	// [28] = [WRITE] hashlist
	//
	// [29] = [WRITE] hashlistMarker
	//
	// [30] = [WRITE] pooledHashlistMarket
	//
	// [31] = [WRITE] fungibleMint
	//
	// [32] = [WRITE] liquidityFungibleTokenAccount
	//
	// [33] = [WRITE, SIGNER] nonFungibleMint
	//
	// [34] = [WRITE] nonFungibleTokenAccount
	//
	// [35] = [WRITE, SIGNER] pooledNonFungibleMint
	//
	// [36] = [WRITE] pooledNonFungibleTokenAccount
	//
	// [37] = [] fairLaunch
	//
	// [38] = [] liquidityProgram
	//
	// [39] = [] associatedTokenProgram
	//
	// [40] = [] libreplexRoyaltyHook
	//
	// [41] = [WRITE] royaltyHookExtraAccountMeta
	//
	// [42] = [WRITE] pooledRoyaltyHookExtraAccountMeta
	//
	// [43] = [] globalDenyListAccount
	//
	// [44] = [] collectionRoyaltyListAccount
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewT22LiquidMintInstructionBuilder creates a new `T22LiquidMint` instruction builder.
func NewT22LiquidMintInstructionBuilder() *T22LiquidMint {
	nd := &T22LiquidMint{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 45),
	}
	return nd
}

// SetInput sets the "input" parameter.
func (inst *T22LiquidMint) SetInput(input MachineMintInput) *T22LiquidMint {
	inst.Input = &input
	return inst
}

// SetBuyerAccount sets the "buyer" account.
func (inst *T22LiquidMint) SetBuyerAccount(buyer ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(buyer).WRITE().SIGNER()
	return inst
}

// GetBuyerAccount gets the "buyer" account.
func (inst *T22LiquidMint) GetBuyerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRequiredSignerAccount sets the "requiredSigner" account.
func (inst *T22LiquidMint) SetRequiredSignerAccount(requiredSigner ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(requiredSigner).SIGNER()
	return inst
}

// GetRequiredSignerAccount gets the "requiredSigner" account.
func (inst *T22LiquidMint) GetRequiredSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMachineAccount sets the "machine" account.
func (inst *T22LiquidMint) SetMachineAccount(machine ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(machine).WRITE()
	return inst
}

// GetMachineAccount gets the "machine" account.
func (inst *T22LiquidMint) GetMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMultipliersAccount sets the "multipliers" account.
func (inst *T22LiquidMint) SetMultipliersAccount(multipliers ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(multipliers)
	return inst
}

// GetMultipliersAccount gets the "multipliers" account.
func (inst *T22LiquidMint) GetMultipliersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTotalMintsAccount sets the "totalMints" account.
func (inst *T22LiquidMint) SetTotalMintsAccount(totalMints ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(totalMints).WRITE()
	return inst
}

// GetTotalMintsAccount gets the "totalMints" account.
func (inst *T22LiquidMint) GetTotalMintsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetWalletAccount sets the "wallet" account.
func (inst *T22LiquidMint) SetWalletAccount(wallet ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(wallet).WRITE()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *T22LiquidMint) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetWallet2Account sets the "wallet2" account.
func (inst *T22LiquidMint) SetWallet2Account(wallet2 ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(wallet2).WRITE()
	return inst
}

// GetWallet2Account gets the "wallet2" account.
func (inst *T22LiquidMint) GetWallet2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *T22LiquidMint) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *T22LiquidMint) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetBuyerWlTokenWalletAccount sets the "buyerWlTokenWallet" account.
func (inst *T22LiquidMint) SetBuyerWlTokenWalletAccount(buyerWlTokenWallet ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(buyerWlTokenWallet).WRITE()
	return inst
}

// GetBuyerWlTokenWalletAccount gets the "buyerWlTokenWallet" account.
func (inst *T22LiquidMint) GetBuyerWlTokenWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetWlNftMetadataAccount sets the "wlNftMetadata" account.
func (inst *T22LiquidMint) SetWlNftMetadataAccount(wlNftMetadata ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(wlNftMetadata)
	return inst
}

// GetWlNftMetadataAccount gets the "wlNftMetadata" account.
func (inst *T22LiquidMint) GetWlNftMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMintsPerWlNftAccount sets the "mintsPerWlNft" account.
func (inst *T22LiquidMint) SetMintsPerWlNftAccount(mintsPerWlNft ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(mintsPerWlNft).WRITE()
	return inst
}

// GetMintsPerWlNftAccount gets the "mintsPerWlNft" account.
func (inst *T22LiquidMint) GetMintsPerWlNftAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetWhitelistMintAccount sets the "whitelistMint" account.
func (inst *T22LiquidMint) SetWhitelistMintAccount(whitelistMint ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(whitelistMint).WRITE()
	return inst
}

// GetWhitelistMintAccount gets the "whitelistMint" account.
func (inst *T22LiquidMint) GetWhitelistMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetBuyerPaymentTokenWalletAccount sets the "buyerPaymentTokenWallet" account.
func (inst *T22LiquidMint) SetBuyerPaymentTokenWalletAccount(buyerPaymentTokenWallet ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(buyerPaymentTokenWallet).WRITE()
	return inst
}

// GetBuyerPaymentTokenWalletAccount gets the "buyerPaymentTokenWallet" account.
func (inst *T22LiquidMint) GetBuyerPaymentTokenWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetInstructionSysvarAccountAccount sets the "instructionSysvarAccount" account.
func (inst *T22LiquidMint) SetInstructionSysvarAccountAccount(instructionSysvarAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(instructionSysvarAccount)
	return inst
}

// GetInstructionSysvarAccountAccount gets the "instructionSysvarAccount" account.
func (inst *T22LiquidMint) GetInstructionSysvarAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *T22LiquidMint) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *T22LiquidMint) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetToken22ProgramAccount sets the "token22Program" account.
func (inst *T22LiquidMint) SetToken22ProgramAccount(token22Program ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(token22Program)
	return inst
}

// GetToken22ProgramAccount gets the "token22Program" account.
func (inst *T22LiquidMint) GetToken22ProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *T22LiquidMint) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *T22LiquidMint) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetAccountCompressionProgramAccount sets the "accountCompressionProgram" account.
func (inst *T22LiquidMint) SetAccountCompressionProgramAccount(accountCompressionProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(accountCompressionProgram)
	return inst
}

// GetAccountCompressionProgramAccount gets the "accountCompressionProgram" account.
func (inst *T22LiquidMint) GetAccountCompressionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetNoopProgramAccount sets the "noopProgram" account.
func (inst *T22LiquidMint) SetNoopProgramAccount(noopProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(noopProgram)
	return inst
}

// GetNoopProgramAccount gets the "noopProgram" account.
func (inst *T22LiquidMint) GetNoopProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetRecentSlothashesAccount sets the "recentSlothashes" account.
func (inst *T22LiquidMint) SetRecentSlothashesAccount(recentSlothashes ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(recentSlothashes)
	return inst
}

// GetRecentSlothashesAccount gets the "recentSlothashes" account.
func (inst *T22LiquidMint) GetRecentSlothashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetBubblegumProgramAccount sets the "bubblegumProgram" account.
func (inst *T22LiquidMint) SetBubblegumProgramAccount(bubblegumProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(bubblegumProgram)
	return inst
}

// GetBubblegumProgramAccount gets the "bubblegumProgram" account.
func (inst *T22LiquidMint) GetBubblegumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetWlNftTreeAccount sets the "wlNftTree" account.
func (inst *T22LiquidMint) SetWlNftTreeAccount(wlNftTree ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(wlNftTree)
	return inst
}

// GetWlNftTreeAccount gets the "wlNftTree" account.
func (inst *T22LiquidMint) GetWlNftTreeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

// SetWlNftTreeAuthorityAccount sets the "wlNftTreeAuthority" account.
func (inst *T22LiquidMint) SetWlNftTreeAuthorityAccount(wlNftTreeAuthority ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[22] = ag_solanago.Meta(wlNftTreeAuthority).WRITE()
	return inst
}

// GetWlNftTreeAuthorityAccount gets the "wlNftTreeAuthority" account.
func (inst *T22LiquidMint) GetWlNftTreeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(22)
}

// SetLiquidityAccount sets the "liquidity" account.
func (inst *T22LiquidMint) SetLiquidityAccount(liquidity ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[23] = ag_solanago.Meta(liquidity).WRITE()
	return inst
}

// GetLiquidityAccount gets the "liquidity" account.
func (inst *T22LiquidMint) GetLiquidityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(23)
}

// SetDeploymentFungibleTokenAccountAccount sets the "deploymentFungibleTokenAccount" account.
func (inst *T22LiquidMint) SetDeploymentFungibleTokenAccountAccount(deploymentFungibleTokenAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[24] = ag_solanago.Meta(deploymentFungibleTokenAccount).WRITE()
	return inst
}

// GetDeploymentFungibleTokenAccountAccount gets the "deploymentFungibleTokenAccount" account.
func (inst *T22LiquidMint) GetDeploymentFungibleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(24)
}

// SetDeploymentNonFungibleTokenAccountAccount sets the "deploymentNonFungibleTokenAccount" account.
func (inst *T22LiquidMint) SetDeploymentNonFungibleTokenAccountAccount(deploymentNonFungibleTokenAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[25] = ag_solanago.Meta(deploymentNonFungibleTokenAccount).WRITE()
	return inst
}

// GetDeploymentNonFungibleTokenAccountAccount gets the "deploymentNonFungibleTokenAccount" account.
func (inst *T22LiquidMint) GetDeploymentNonFungibleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(25)
}

// SetDeploymentAccount sets the "deployment" account.
func (inst *T22LiquidMint) SetDeploymentAccount(deployment ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[26] = ag_solanago.Meta(deployment).WRITE()
	return inst
}

// GetDeploymentAccount gets the "deployment" account.
func (inst *T22LiquidMint) GetDeploymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(26)
}

// SetDeploymentConfigAccount sets the "deploymentConfig" account.
func (inst *T22LiquidMint) SetDeploymentConfigAccount(deploymentConfig ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[27] = ag_solanago.Meta(deploymentConfig).WRITE()
	return inst
}

// GetDeploymentConfigAccount gets the "deploymentConfig" account.
func (inst *T22LiquidMint) GetDeploymentConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(27)
}

// SetHashlistAccount sets the "hashlist" account.
func (inst *T22LiquidMint) SetHashlistAccount(hashlist ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[28] = ag_solanago.Meta(hashlist).WRITE()
	return inst
}

// GetHashlistAccount gets the "hashlist" account.
func (inst *T22LiquidMint) GetHashlistAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(28)
}

// SetHashlistMarkerAccount sets the "hashlistMarker" account.
func (inst *T22LiquidMint) SetHashlistMarkerAccount(hashlistMarker ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[29] = ag_solanago.Meta(hashlistMarker).WRITE()
	return inst
}

// GetHashlistMarkerAccount gets the "hashlistMarker" account.
func (inst *T22LiquidMint) GetHashlistMarkerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(29)
}

// SetPooledHashlistMarketAccount sets the "pooledHashlistMarket" account.
func (inst *T22LiquidMint) SetPooledHashlistMarketAccount(pooledHashlistMarket ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[30] = ag_solanago.Meta(pooledHashlistMarket).WRITE()
	return inst
}

// GetPooledHashlistMarketAccount gets the "pooledHashlistMarket" account.
func (inst *T22LiquidMint) GetPooledHashlistMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(30)
}

// SetFungibleMintAccount sets the "fungibleMint" account.
func (inst *T22LiquidMint) SetFungibleMintAccount(fungibleMint ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[31] = ag_solanago.Meta(fungibleMint).WRITE()
	return inst
}

// GetFungibleMintAccount gets the "fungibleMint" account.
func (inst *T22LiquidMint) GetFungibleMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(31)
}

// SetLiquidityFungibleTokenAccountAccount sets the "liquidityFungibleTokenAccount" account.
func (inst *T22LiquidMint) SetLiquidityFungibleTokenAccountAccount(liquidityFungibleTokenAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[32] = ag_solanago.Meta(liquidityFungibleTokenAccount).WRITE()
	return inst
}

// GetLiquidityFungibleTokenAccountAccount gets the "liquidityFungibleTokenAccount" account.
func (inst *T22LiquidMint) GetLiquidityFungibleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(32)
}

// SetNonFungibleMintAccount sets the "nonFungibleMint" account.
func (inst *T22LiquidMint) SetNonFungibleMintAccount(nonFungibleMint ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[33] = ag_solanago.Meta(nonFungibleMint).WRITE().SIGNER()
	return inst
}

// GetNonFungibleMintAccount gets the "nonFungibleMint" account.
func (inst *T22LiquidMint) GetNonFungibleMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(33)
}

// SetNonFungibleTokenAccountAccount sets the "nonFungibleTokenAccount" account.
func (inst *T22LiquidMint) SetNonFungibleTokenAccountAccount(nonFungibleTokenAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[34] = ag_solanago.Meta(nonFungibleTokenAccount).WRITE()
	return inst
}

// GetNonFungibleTokenAccountAccount gets the "nonFungibleTokenAccount" account.
func (inst *T22LiquidMint) GetNonFungibleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(34)
}

// SetPooledNonFungibleMintAccount sets the "pooledNonFungibleMint" account.
func (inst *T22LiquidMint) SetPooledNonFungibleMintAccount(pooledNonFungibleMint ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[35] = ag_solanago.Meta(pooledNonFungibleMint).WRITE().SIGNER()
	return inst
}

// GetPooledNonFungibleMintAccount gets the "pooledNonFungibleMint" account.
func (inst *T22LiquidMint) GetPooledNonFungibleMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(35)
}

// SetPooledNonFungibleTokenAccountAccount sets the "pooledNonFungibleTokenAccount" account.
func (inst *T22LiquidMint) SetPooledNonFungibleTokenAccountAccount(pooledNonFungibleTokenAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[36] = ag_solanago.Meta(pooledNonFungibleTokenAccount).WRITE()
	return inst
}

// GetPooledNonFungibleTokenAccountAccount gets the "pooledNonFungibleTokenAccount" account.
func (inst *T22LiquidMint) GetPooledNonFungibleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(36)
}

// SetFairLaunchAccount sets the "fairLaunch" account.
func (inst *T22LiquidMint) SetFairLaunchAccount(fairLaunch ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[37] = ag_solanago.Meta(fairLaunch)
	return inst
}

// GetFairLaunchAccount gets the "fairLaunch" account.
func (inst *T22LiquidMint) GetFairLaunchAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(37)
}

// SetLiquidityProgramAccount sets the "liquidityProgram" account.
func (inst *T22LiquidMint) SetLiquidityProgramAccount(liquidityProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[38] = ag_solanago.Meta(liquidityProgram)
	return inst
}

// GetLiquidityProgramAccount gets the "liquidityProgram" account.
func (inst *T22LiquidMint) GetLiquidityProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(38)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *T22LiquidMint) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[39] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *T22LiquidMint) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(39)
}

// SetLibreplexRoyaltyHookAccount sets the "libreplexRoyaltyHook" account.
func (inst *T22LiquidMint) SetLibreplexRoyaltyHookAccount(libreplexRoyaltyHook ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[40] = ag_solanago.Meta(libreplexRoyaltyHook)
	return inst
}

// GetLibreplexRoyaltyHookAccount gets the "libreplexRoyaltyHook" account.
func (inst *T22LiquidMint) GetLibreplexRoyaltyHookAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(40)
}

// SetRoyaltyHookExtraAccountMetaAccount sets the "royaltyHookExtraAccountMeta" account.
func (inst *T22LiquidMint) SetRoyaltyHookExtraAccountMetaAccount(royaltyHookExtraAccountMeta ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[41] = ag_solanago.Meta(royaltyHookExtraAccountMeta).WRITE()
	return inst
}

// GetRoyaltyHookExtraAccountMetaAccount gets the "royaltyHookExtraAccountMeta" account.
func (inst *T22LiquidMint) GetRoyaltyHookExtraAccountMetaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(41)
}

// SetPooledRoyaltyHookExtraAccountMetaAccount sets the "pooledRoyaltyHookExtraAccountMeta" account.
func (inst *T22LiquidMint) SetPooledRoyaltyHookExtraAccountMetaAccount(pooledRoyaltyHookExtraAccountMeta ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[42] = ag_solanago.Meta(pooledRoyaltyHookExtraAccountMeta).WRITE()
	return inst
}

// GetPooledRoyaltyHookExtraAccountMetaAccount gets the "pooledRoyaltyHookExtraAccountMeta" account.
func (inst *T22LiquidMint) GetPooledRoyaltyHookExtraAccountMetaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(42)
}

// SetGlobalDenyListAccountAccount sets the "globalDenyListAccount" account.
func (inst *T22LiquidMint) SetGlobalDenyListAccountAccount(globalDenyListAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[43] = ag_solanago.Meta(globalDenyListAccount)
	return inst
}

// GetGlobalDenyListAccountAccount gets the "globalDenyListAccount" account.
func (inst *T22LiquidMint) GetGlobalDenyListAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(43)
}

// SetCollectionRoyaltyListAccountAccount sets the "collectionRoyaltyListAccount" account.
func (inst *T22LiquidMint) SetCollectionRoyaltyListAccountAccount(collectionRoyaltyListAccount ag_solanago.PublicKey) *T22LiquidMint {
	inst.AccountMetaSlice[44] = ag_solanago.Meta(collectionRoyaltyListAccount)
	return inst
}

// GetCollectionRoyaltyListAccountAccount gets the "collectionRoyaltyListAccount" account.
func (inst *T22LiquidMint) GetCollectionRoyaltyListAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(44)
}

func (inst T22LiquidMint) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_T22LiquidMint,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst T22LiquidMint) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *T22LiquidMint) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Input == nil {
			return errors.New("Input parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Buyer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RequiredSigner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Machine is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Multipliers is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TotalMints is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Wallet2 is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BuyerWlTokenWallet is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.WlNftMetadata is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MintsPerWlNft is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.WhitelistMint is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.BuyerPaymentTokenWallet is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.InstructionSysvarAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Token22Program is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.AccountCompressionProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.NoopProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.RecentSlothashes is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.BubblegumProgram is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.WlNftTree is not set")
		}
		if inst.AccountMetaSlice[22] == nil {
			return errors.New("accounts.WlNftTreeAuthority is not set")
		}
		if inst.AccountMetaSlice[23] == nil {
			return errors.New("accounts.Liquidity is not set")
		}
		if inst.AccountMetaSlice[24] == nil {
			return errors.New("accounts.DeploymentFungibleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[25] == nil {
			return errors.New("accounts.DeploymentNonFungibleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[26] == nil {
			return errors.New("accounts.Deployment is not set")
		}
		if inst.AccountMetaSlice[27] == nil {
			return errors.New("accounts.DeploymentConfig is not set")
		}
		if inst.AccountMetaSlice[28] == nil {
			return errors.New("accounts.Hashlist is not set")
		}
		if inst.AccountMetaSlice[29] == nil {
			return errors.New("accounts.HashlistMarker is not set")
		}
		if inst.AccountMetaSlice[30] == nil {
			return errors.New("accounts.PooledHashlistMarket is not set")
		}
		if inst.AccountMetaSlice[31] == nil {
			return errors.New("accounts.FungibleMint is not set")
		}
		if inst.AccountMetaSlice[32] == nil {
			return errors.New("accounts.LiquidityFungibleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[33] == nil {
			return errors.New("accounts.NonFungibleMint is not set")
		}
		if inst.AccountMetaSlice[34] == nil {
			return errors.New("accounts.NonFungibleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[35] == nil {
			return errors.New("accounts.PooledNonFungibleMint is not set")
		}
		if inst.AccountMetaSlice[36] == nil {
			return errors.New("accounts.PooledNonFungibleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[37] == nil {
			return errors.New("accounts.FairLaunch is not set")
		}
		if inst.AccountMetaSlice[38] == nil {
			return errors.New("accounts.LiquidityProgram is not set")
		}
		if inst.AccountMetaSlice[39] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[40] == nil {
			return errors.New("accounts.LibreplexRoyaltyHook is not set")
		}
		if inst.AccountMetaSlice[41] == nil {
			return errors.New("accounts.RoyaltyHookExtraAccountMeta is not set")
		}
		if inst.AccountMetaSlice[42] == nil {
			return errors.New("accounts.PooledRoyaltyHookExtraAccountMeta is not set")
		}
		if inst.AccountMetaSlice[43] == nil {
			return errors.New("accounts.GlobalDenyListAccount is not set")
		}
		if inst.AccountMetaSlice[44] == nil {
			return errors.New("accounts.CollectionRoyaltyListAccount is not set")
		}
	}
	return nil
}

func (inst *T22LiquidMint) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("T22LiquidMint")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Input", *inst.Input))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=45]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                            buyer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   requiredSigner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                          machine", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                      multipliers", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                       totalMints", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                           wallet", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                          wallet2", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                    systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("               buyerWlTokenWallet", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                    wlNftMetadata", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                    mintsPerWlNft", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                    whitelistMint", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("          buyerPaymentTokenWallet", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                instructionSysvar", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                     tokenProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                   token22Program", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("             tokenMetadataProgram", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("        accountCompressionProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("                      noopProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("                 recentSlothashes", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("                 bubblegumProgram", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("                        wlNftTree", inst.AccountMetaSlice.Get(21)))
						accountsBranch.Child(ag_format.Meta("               wlNftTreeAuthority", inst.AccountMetaSlice.Get(22)))
						accountsBranch.Child(ag_format.Meta("                        liquidity", inst.AccountMetaSlice.Get(23)))
						accountsBranch.Child(ag_format.Meta("          deploymentFungibleToken", inst.AccountMetaSlice.Get(24)))
						accountsBranch.Child(ag_format.Meta("       deploymentNonFungibleToken", inst.AccountMetaSlice.Get(25)))
						accountsBranch.Child(ag_format.Meta("                       deployment", inst.AccountMetaSlice.Get(26)))
						accountsBranch.Child(ag_format.Meta("                 deploymentConfig", inst.AccountMetaSlice.Get(27)))
						accountsBranch.Child(ag_format.Meta("                         hashlist", inst.AccountMetaSlice.Get(28)))
						accountsBranch.Child(ag_format.Meta("                   hashlistMarker", inst.AccountMetaSlice.Get(29)))
						accountsBranch.Child(ag_format.Meta("             pooledHashlistMarket", inst.AccountMetaSlice.Get(30)))
						accountsBranch.Child(ag_format.Meta("                     fungibleMint", inst.AccountMetaSlice.Get(31)))
						accountsBranch.Child(ag_format.Meta("           liquidityFungibleToken", inst.AccountMetaSlice.Get(32)))
						accountsBranch.Child(ag_format.Meta("                  nonFungibleMint", inst.AccountMetaSlice.Get(33)))
						accountsBranch.Child(ag_format.Meta("                 nonFungibleToken", inst.AccountMetaSlice.Get(34)))
						accountsBranch.Child(ag_format.Meta("            pooledNonFungibleMint", inst.AccountMetaSlice.Get(35)))
						accountsBranch.Child(ag_format.Meta("           pooledNonFungibleToken", inst.AccountMetaSlice.Get(36)))
						accountsBranch.Child(ag_format.Meta("                       fairLaunch", inst.AccountMetaSlice.Get(37)))
						accountsBranch.Child(ag_format.Meta("                 liquidityProgram", inst.AccountMetaSlice.Get(38)))
						accountsBranch.Child(ag_format.Meta("           associatedTokenProgram", inst.AccountMetaSlice.Get(39)))
						accountsBranch.Child(ag_format.Meta("             libreplexRoyaltyHook", inst.AccountMetaSlice.Get(40)))
						accountsBranch.Child(ag_format.Meta("      royaltyHookExtraAccountMeta", inst.AccountMetaSlice.Get(41)))
						accountsBranch.Child(ag_format.Meta("pooledRoyaltyHookExtraAccountMeta", inst.AccountMetaSlice.Get(42)))
						accountsBranch.Child(ag_format.Meta("                   globalDenyList", inst.AccountMetaSlice.Get(43)))
						accountsBranch.Child(ag_format.Meta("            collectionRoyaltyList", inst.AccountMetaSlice.Get(44)))
					})
				})
		})
}

func (obj T22LiquidMint) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Input` param:
	err = encoder.Encode(obj.Input)
	if err != nil {
		return err
	}
	return nil
}
func (obj *T22LiquidMint) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Input`:
	err = decoder.Decode(&obj.Input)
	if err != nil {
		return err
	}
	return nil
}

// NewT22LiquidMintInstruction declares a new T22LiquidMint instruction with the provided parameters and accounts.
func NewT22LiquidMintInstruction(
	// Parameters:
	input MachineMintInput,
	// Accounts:
	buyer ag_solanago.PublicKey,
	requiredSigner ag_solanago.PublicKey,
	machine ag_solanago.PublicKey,
	multipliers ag_solanago.PublicKey,
	totalMints ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	wallet2 ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	buyerWlTokenWallet ag_solanago.PublicKey,
	wlNftMetadata ag_solanago.PublicKey,
	mintsPerWlNft ag_solanago.PublicKey,
	whitelistMint ag_solanago.PublicKey,
	buyerPaymentTokenWallet ag_solanago.PublicKey,
	instructionSysvarAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	token22Program ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	accountCompressionProgram ag_solanago.PublicKey,
	noopProgram ag_solanago.PublicKey,
	recentSlothashes ag_solanago.PublicKey,
	bubblegumProgram ag_solanago.PublicKey,
	wlNftTree ag_solanago.PublicKey,
	wlNftTreeAuthority ag_solanago.PublicKey,
	liquidity ag_solanago.PublicKey,
	deploymentFungibleTokenAccount ag_solanago.PublicKey,
	deploymentNonFungibleTokenAccount ag_solanago.PublicKey,
	deployment ag_solanago.PublicKey,
	deploymentConfig ag_solanago.PublicKey,
	hashlist ag_solanago.PublicKey,
	hashlistMarker ag_solanago.PublicKey,
	pooledHashlistMarket ag_solanago.PublicKey,
	fungibleMint ag_solanago.PublicKey,
	liquidityFungibleTokenAccount ag_solanago.PublicKey,
	nonFungibleMint ag_solanago.PublicKey,
	nonFungibleTokenAccount ag_solanago.PublicKey,
	pooledNonFungibleMint ag_solanago.PublicKey,
	pooledNonFungibleTokenAccount ag_solanago.PublicKey,
	fairLaunch ag_solanago.PublicKey,
	liquidityProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	libreplexRoyaltyHook ag_solanago.PublicKey,
	royaltyHookExtraAccountMeta ag_solanago.PublicKey,
	pooledRoyaltyHookExtraAccountMeta ag_solanago.PublicKey,
	globalDenyListAccount ag_solanago.PublicKey,
	collectionRoyaltyListAccount ag_solanago.PublicKey) *T22LiquidMint {
	return NewT22LiquidMintInstructionBuilder().
		SetInput(input).
		SetBuyerAccount(buyer).
		SetRequiredSignerAccount(requiredSigner).
		SetMachineAccount(machine).
		SetMultipliersAccount(multipliers).
		SetTotalMintsAccount(totalMints).
		SetWalletAccount(wallet).
		SetWallet2Account(wallet2).
		SetSystemProgramAccount(systemProgram).
		SetBuyerWlTokenWalletAccount(buyerWlTokenWallet).
		SetWlNftMetadataAccount(wlNftMetadata).
		SetMintsPerWlNftAccount(mintsPerWlNft).
		SetWhitelistMintAccount(whitelistMint).
		SetBuyerPaymentTokenWalletAccount(buyerPaymentTokenWallet).
		SetInstructionSysvarAccountAccount(instructionSysvarAccount).
		SetTokenProgramAccount(tokenProgram).
		SetToken22ProgramAccount(token22Program).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetAccountCompressionProgramAccount(accountCompressionProgram).
		SetNoopProgramAccount(noopProgram).
		SetRecentSlothashesAccount(recentSlothashes).
		SetBubblegumProgramAccount(bubblegumProgram).
		SetWlNftTreeAccount(wlNftTree).
		SetWlNftTreeAuthorityAccount(wlNftTreeAuthority).
		SetLiquidityAccount(liquidity).
		SetDeploymentFungibleTokenAccountAccount(deploymentFungibleTokenAccount).
		SetDeploymentNonFungibleTokenAccountAccount(deploymentNonFungibleTokenAccount).
		SetDeploymentAccount(deployment).
		SetDeploymentConfigAccount(deploymentConfig).
		SetHashlistAccount(hashlist).
		SetHashlistMarkerAccount(hashlistMarker).
		SetPooledHashlistMarketAccount(pooledHashlistMarket).
		SetFungibleMintAccount(fungibleMint).
		SetLiquidityFungibleTokenAccountAccount(liquidityFungibleTokenAccount).
		SetNonFungibleMintAccount(nonFungibleMint).
		SetNonFungibleTokenAccountAccount(nonFungibleTokenAccount).
		SetPooledNonFungibleMintAccount(pooledNonFungibleMint).
		SetPooledNonFungibleTokenAccountAccount(pooledNonFungibleTokenAccount).
		SetFairLaunchAccount(fairLaunch).
		SetLiquidityProgramAccount(liquidityProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetLibreplexRoyaltyHookAccount(libreplexRoyaltyHook).
		SetRoyaltyHookExtraAccountMetaAccount(royaltyHookExtraAccountMeta).
		SetPooledRoyaltyHookExtraAccountMetaAccount(pooledRoyaltyHookExtraAccountMeta).
		SetGlobalDenyListAccountAccount(globalDenyListAccount).
		SetCollectionRoyaltyListAccountAccount(collectionRoyaltyListAccount)
}
