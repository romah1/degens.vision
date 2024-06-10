package helius

import (
	"context"
	"net/http"
	"solana_project/libraries/programs"
	"solana_project/libraries/queues/mint_tracker_mints"

	"github.com/gagliardetto/solana-go"
	"github.com/gin-gonic/gin"
)

type NftMint struct {
	Instructions []Instruction `json:"instructions" binding:"required"`
	Signature    string        `json:"signature" binding:"required"`
}

type Instruction struct {
	ProgramId string `json:"programId" binding:"required"`
}

func HandleHeliusMints(
	ctx context.Context,
	mintsTrackerQueueProducer *mint_tracker_mints.MintTrackerMintsQueueProducer,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "mrbeastzxc495812" {
			c.JSON(http.StatusForbidden, gin.H{"error": "bad auth header"})
			return
		}
		var mints []NftMint
		if err := c.ShouldBindJSON(&mints); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		for _, mint := range mints {
			var programId *solana.PublicKey
			for _, instruction := range mint.Instructions {
				switch instruction.ProgramId {
				case programs.LMNFTProgramID.String():
					programId = &programs.LMNFTProgramID
				case programs.LMNFTBubblegumProgramID.String():
					programId = &programs.LMNFTBubblegumProgramID
				case programs.TruffleProgramId.String():
					programId = &programs.TruffleProgramId
				}
			}
			if programId == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve program id"})
				return
			}
			err := processSignature(ctx, mintsTrackerQueueProducer, *programId, mint.Signature)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish message"})
				return
			}
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

func processSignature(
	ctx context.Context,
	mintsTrackerQueueProducer *mint_tracker_mints.MintTrackerMintsQueueProducer,
	programId solana.PublicKey,
	sig string,
) error {
	err := mintsTrackerQueueProducer.Produce(ctx, &mint_tracker_mints.Message{
		ProgramId: programId.String(),
		Signature: sig,
	})
	if err != nil {
		return err
	}

	return nil
}
