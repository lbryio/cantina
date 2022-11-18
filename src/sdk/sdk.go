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

const NOT_FOUND = "NOT_FOUND"

type ClaimId string
type ClaimName string

type Claim struct {
	ClaimId      ClaimId   `json:"claim_id"`
	Name         ClaimName `json:"name"`
	PermanentURL string    `json:"permanent_url"`
	Value        struct {
		Title string `json:"title"`
	} `json:"value"`

	Error interface{} `json:"error"`
}

type ClaimResponse struct {
	// TODO? Blocked map[string]interface{} `json:"blocked"`

	Result map[string]Claim `json:"result"`
}

type ClaimNotFound struct {
}

func (ce *ClaimNotFound) Error() string {
	return "Claim not found"
}

func GetClaim(claimId ClaimId) (claim Claim, err error) {
	url := "lbry://channel#" + string(claimId)

	reqBodyStr, _ := json.Marshal(map[string]interface{}{
		"method": "resolve",
		"params": map[string]([]string){
			"urls": []string{url},
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

	claim, ok := respBody.Result[url]

	if ok {
		switch claim.Error.(type) {
		case map[string]interface{}:
			claimErr := claim.Error.(map[string]interface{})
			if name, _ := claimErr["name"].(string); name == NOT_FOUND {
				err = &ClaimNotFound{}
				claim = Claim{}
			} else {
				err = fmt.Errorf("Unknown SDK response error")
				claim = Claim{}
			}
		case nil:
		default:
			// Some other error format, which does exist
			err = fmt.Errorf("Unknown SDK response error")
			claim = Claim{}
		}
	} else {
		err = fmt.Errorf("Unknown SDK response error")
	}

	if err == nil && claim.Name[0] != '@' {
		// TODO - different error, probably not a 500
		err = fmt.Errorf("Invalid claim")
	}

	return
}
