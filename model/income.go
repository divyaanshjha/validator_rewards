package model

type Income struct {
    AttestationHeadReward          float64 `json:"attestation_head_reward"`
    AttestationSourcePenalty       float64 `json:"attestation_source_penalty"`
    AttestationSourceReward        float64 `json:"attestation_source_reward"`
    AttestationTargetPenalty       float64 `json:"attestation_target_penalty"`
    AttestationTargetReward        float64 `json:"attestation_target_reward"`
    FinalityDelayPenalty           float64 `json:"finality_delay_penalty"`
    ProposalsMissed                int     `json:"proposals_missed"`
    ProposerAttestationInclusionReward float64 `json:"proposer_attestation_inclusion_reward"`
    ProposerSlashingInclusionReward    float64 `json:"proposer_slashing_inclusion_reward"`
    ProposerSyncInclusionReward        float64 `json:"proposer_sync_inclusion_reward"`
    SlashingPenalty                   float64 `json:"slashing_penalty"`
    SlashingReward                    float64 `json:"slashing_reward"`
    SyncCommitteePenalty              float64 `json:"sync_committee_penalty"`
    SyncCommitteeReward               float64 `json:"sync_committee_reward"`
    TxFeeRewardWei                    string  `json:"tx_fee_reward_wei"`
}

type IncomeDetail struct {
    Epoch        int    `json:"epoch"`
    Income       Income `json:"income"`
    ValidatorIndex int  `json:"validatorindex"`
    Week         int    `json:"week"`
    WeekEnd      string `json:"week_end"`
    WeekStart    string `json:"week_start"`
}

type ApiResponse struct {
    Data   []IncomeDetail `json:"data"`
    Status string         `json:"status"`
}
