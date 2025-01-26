package model



type ConsensusPerformance struct {
    Balance        int64    `json:"balance"`
    Performance1d        int64    `json:"performance1d"`
    Performance31d        int64    `json:"performance31d"`
    Performance365d        int64    `json:"performance365d"`
    Performance7d        int64    `json:"performance7d"`
    PerformanceToday        int64    `json:"performancetoday"`
    PerformanceTotal        int64    `json:"performancetotal"`
    Rank7d        int64    `json:"rank7d"`
    ValidatorIndex        int64    `json:"validatorindex"`
}

type ExecutionPerformance struct{
    Performance1d        int64    `json:"performance1d"`
    Performance7d        int64    `json:"performance7d"`
    Performance31d        int64    `json:"performance31d"`
    Performance365d        int64    `json:"performance365d"`
    PerformanceTotal        int64    `json:"performancetotal"`
    ValidatorIndex        int64    `json:"validatorindex"`

}



type ApiResponseConsensus struct {
    Status string         `json:"status"`
    Data   []ConsensusPerformance `json:"data"`

}

type ApiResponseExecution struct {
    Data   []ExecutionPerformance `json:"data"`
    Status string         `json:"status"`
}