package main

import (
    "github.com/divyaanshjha/validator_rewards/api"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // API endpoint for fetching validator rewards
    r.GET("/validator-rewards/:validator", api.GetValidatorRewards)

    // Start the server
    r.Run(":8080") // Server runs on localhost:8080
}
