package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/divyaanshjha/validator_rewards/model"
)

func ConsensusPerformance(validatorIndexOrPubKey string) (*model.ApiResponseConsensus, error) {
    url := fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s/performance", validatorIndexOrPubKey)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var apiResponseConsensus model.ApiResponseConsensus
    err = json.Unmarshal(body, &apiResponseConsensus)
    if err != nil {
        return nil, err
    }
    //fmt.Printf("%+v\n", apiResponseConsensus)
    return &apiResponseConsensus, nil
}


func ExecutionPerformance(validatorIndexOrPubKey string) (*model.ApiResponseExecution, error) {
    url := fmt.Sprintf("https://beaconcha.in/api/v1/validator/%s/execution/performance", validatorIndexOrPubKey)
	//https://beaconcha.in/api/v1/validator/0x80000001677f23a227dfed6f61b132d114be83b8ad0aa5f3c5d1d77e6ee0bf5f73b0af750cc34e8f2dae73c21dc36f4a/incomedetailhistory

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var apiResponseExecution model.ApiResponseExecution
    err = json.Unmarshal(body, &apiResponseExecution)
    if err != nil {
        return nil, err
    }

    return &apiResponseExecution, nil
}