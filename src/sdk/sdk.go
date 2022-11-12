package sdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TODO - replace me with communication straight to hub

// Interactions with the LBRY sdk

const SDK_URL = "http://localhost:5279"

type ClaimId string
type ClaimName string

type ClaimResponse struct {
	Result struct {
		Items []Claim `json:"items"`
		// TODO "blocked" - ?
	} `json:"result"`
}

type Claim struct {
	ClaimId      ClaimId   `json:"claim_id"`
	PermanentURL string    `json:"permanent_url"`
	Name         ClaimName `json:"name"`
}

func GetClaim(claimId ClaimId) (claim Claim, err error) {
	reqBodyStr, _ := json.Marshal(map[string]string{
		"method":   "claim_search",
		"claim_id": string(claimId),
	})

	resp, err := http.Post(SDK_URL, "application/json", bytes.NewBuffer(reqBodyStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var respBody ClaimResponse
	respBodyStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBodyStr, respBody)

	claim = respBody.Result.Items[0]

	return
}
