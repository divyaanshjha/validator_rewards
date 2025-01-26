package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetValidatorRewards(c *gin.Context) {
    validatorIndexOrPubKey := c.Param("validator")


    ConsensusPerformance, err := ConsensusPerformance(validatorIndexOrPubKey)
	wrappedErrC := fmt.Errorf("additional context: %w", err)
	fmt.Println(wrappedErrC)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching Consensus details"})
        return
    }
    //fmt.Printf("%+v\n", ConsensusPerformance)
    var CPerformance1d, CPerformance31d, CPerformance365d, CPerformance7d,  CPerformanceTotal  float64

    for _, detail := range ConsensusPerformance.Data {    
    CPerformance1d += float64(detail.Performance1d)/1e9
    CPerformance31d += float64(detail.Performance31d)/1e9
    CPerformance365d += float64(detail.Performance365d)/1e9
    CPerformance7d += float64(detail.Performance7d)/1e9
    CPerformanceTotal += float64(detail.PerformanceTotal)/1e9
    }
    fmt.Printf("%f\n", CPerformanceTotal)


    ExecutionPerformance, err := ExecutionPerformance(validatorIndexOrPubKey)
	wrappedErrE := fmt.Errorf("additional context: %w", err)
	fmt.Println(wrappedErrE)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching Execution details"})
        return
    }

    var EPerformance1d, EPerformance31d, EPerformance365d, EPerformance7d,  EPerformanceTotal  float64

    for _, detail := range ExecutionPerformance.Data {    
    EPerformance1d += float64(detail.Performance1d)/1e18
    EPerformance7d += float64(detail.Performance7d)/1e18
    EPerformance31d += float64(detail.Performance31d)/1e18
    EPerformance365d += float64(detail.Performance365d)/1e18
    EPerformanceTotal += float64(detail.PerformanceTotal)/1e18
    }

    TotalRewards := CPerformanceTotal + EPerformanceTotal
    Rewards1D := CPerformance1d + EPerformance1d
    Rewards7D := CPerformance7d + EPerformance7d
    Rewards31D := CPerformance31d + EPerformance31d
    Rewards365D := CPerformance365d + EPerformance365d


    response := map[string]interface{}{
        "validatorAddress":        validatorIndexOrPubKey,
        "Total Rewards":            fmt.Sprintf("%.4f ETH", TotalRewards),
        "1-Day Rewards":      fmt.Sprintf("%.4f ETH", Rewards1D),
        "7 Day Rewards":         fmt.Sprintf("%.4f ETH", Rewards7D),
        "31 Day Rewards":    fmt.Sprintf("%.4f ETH", Rewards31D),
        "365 Day Rewards":   fmt.Sprintf("%.4f ETH", Rewards365D),
    }

    c.JSON(http.StatusOK, response)
}