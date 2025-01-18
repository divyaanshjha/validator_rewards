package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetValidatorRewards(c *gin.Context) {
    validatorIndexOrPubKey := c.Param("validator")


    incomeDetails, err := FetchIncomeDetails(validatorIndexOrPubKey)
	wrappedErr := fmt.Errorf("additional context: %w", err)
	fmt.Println(wrappedErr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching income details"})
        return
    }

    var totalRewards, attestationRewards, proposalRewards, syncCommitteeRewards, transactionFeeRewards, penalties float64

    for _, detail := range incomeDetails.Data {
        income := detail.Income
        attestationRewards += income.AttestationHeadReward + income.AttestationSourceReward + income.AttestationTargetReward
        proposalRewards += income.ProposerAttestationInclusionReward + income.ProposerSlashingInclusionReward + income.ProposerSyncInclusionReward
        syncCommitteeRewards += income.SyncCommitteeReward
        txFeeReward, _ := strconv.ParseFloat(income.TxFeeRewardWei, 64)
        transactionFeeRewards += txFeeReward / 1e18
        penalties -= (income.AttestationSourcePenalty + income.AttestationTargetPenalty + income.FinalityDelayPenalty + income.SlashingPenalty + income.SyncCommitteePenalty)
    }

    totalRewards = attestationRewards + proposalRewards + syncCommitteeRewards + transactionFeeRewards + penalties

    response := map[string]interface{}{
        "validatorAddress":        validatorIndexOrPubKey,
        "totalRewards":            fmt.Sprintf("%.4f ETH", totalRewards),
        "attestationRewards":      fmt.Sprintf("%.4f ETH", attestationRewards),
        "proposalRewards":         fmt.Sprintf("%.4f ETH", proposalRewards),
        "syncCommitteeRewards":    fmt.Sprintf("%.4f ETH", syncCommitteeRewards),
        "transactionFeeRewards":   fmt.Sprintf("%.4f ETH", transactionFeeRewards),
        "penalties":               fmt.Sprintf("%.4f ETH", penalties),
    }

    c.JSON(http.StatusOK, response)
}
