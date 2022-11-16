package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO - replace me with communication straight to hub

// Interactions with the LBRY sdk

const SDK_URL = "http://localhost:5279"

type ClaimId string
type ClaimName string

type Claim struct {
	ClaimId      ClaimId   `json:"claim_id"`
	PermanentURL string    `json:"permanent_url"`
	Name         ClaimName `json:"name"`
}

type ClaimResponse struct {
	Result struct {
		Items []Claim `json:"items"`
		// TODO? Blocked map[string]interface{} `json:"blocked"`
	} `json:"result"`
}

type ClaimNotFound struct {
}

func (ce *ClaimNotFound) Error() string {
	return "Claim not found"
}

func GetClaim(claimId ClaimId) (claim Claim, err error) {
	reqBodyStr, _ := json.Marshal(map[string]interface{}{
		"method": "claim_search",
		"params": map[string]string{
			"claim_id": string(claimId),
		},
	})

	resp, err := http.Post(SDK_URL, "application/json", bytes.NewBuffer(reqBodyStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBodyStr, _ := ioutil.ReadAll(resp.Body)

	var respBody ClaimResponse
	err = json.Unmarshal(respBodyStr, &respBody)
	if err != nil {
		return
	}

	if len(respBody.Result.Items) == 0 {
		err = &ClaimNotFound{}
	} else if len(respBody.Result.Items) > 1 {
		err = fmt.Errorf("Multiple found for some reason. Skipping.")
	} else {
		claim = respBody.Result.Items[0]
	}

	return
}
