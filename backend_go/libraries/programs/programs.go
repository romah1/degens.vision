package programs

import "github.com/gagliardetto/solana-go"

var (
	LMNFTProgramID            = solana.MustPublicKeyFromBase58("ArAA6CZC123yMJLUe4uisBEgvfuw2WEvex9iFmFCYiXv")
	LMNFTBubblegumProgramID   = solana.MustPublicKeyFromBase58("F9SixdqdmEBP5kprp2gZPZNeMmfHJRCTMFjN22dx3akf")
	BubblegumProgramId        = solana.MustPublicKeyFromBase58("BGUMAp9Gq7iTEuizy4pqaxsTyUCBK68MDfK752saRPUY")
	TruffleProgramId          = solana.MustPublicKeyFromBase58("CSGrdwbJ5z58tLGKjjcmiNMj8bG1Zazthk3cXMrbSZoX")
	MetaplexMetadataProgramId = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
	MetaplexCoreProgramId     = solana.MustPublicKeyFromBase58("CoREENxT6tW1HoK8ypY1SxRMZTcVPm7R94rH4PZNhX7d")
)

func ProgramToString(programId string) *string {
	switch programId {
	case LMNFTProgramID.String():
		s := "lmnft"
		return &s
	case LMNFTBubblegumProgramID.String():
		s := "lmnft"
		return &s
	case TruffleProgramId.String():
		s := "truffle"
		return &s
	default:
		return nil
	}
}
