package main

import (
	"encoding/json"
	"fmt"
	"lbryio/cantina/objects"
	"lbryio/cantina/sdk"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// TODO - Does ActivityPub specify a specific error format?
type ErrorResponse struct {
	Error string `json:"error"`
}

func readableErrorJson(w http.ResponseWriter, code int, extra string) {
	errorStr := http.StatusText(code)
	if extra != "" {
		errorStr = errorStr + ": " + extra
	}
	errorJson, err := json.Marshal(ErrorResponse{Error: errorStr})
	if err != nil {
		// In case something really stupid happens
		http.Error(w, `{"error": "error when JSON-encoding error message"}`, code)
	}
	http.Error(w, string(errorJson), code)
	return
}

// Don't report any details to the user. Log it instead.
func internalServiceErrorJson(w http.ResponseWriter, serverErr error, errContext string) {
	errorStr := http.StatusText(http.StatusInternalServerError)
	errorJson, err := json.Marshal(ErrorResponse{Error: errorStr})
	if err != nil {
		// In case something really stupid happens
		http.Error(w, `{"error": "error when JSON-encoding error message"}`, http.StatusInternalServerError)
		log.Printf("error when JSON-encoding error message")
		return
	}
	http.Error(w, string(errorJson), http.StatusInternalServerError)
	log.Printf("%s: %+v\n", errContext, serverErr)

	return
}

var IsHex = regexp.MustCompile(`^[a-z0-9]+$`).MatchString

func handleChannel(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		segments := strings.Split(req.URL.Path, PathChannel)
		if len(segments) != 2 || segments[0] != "" {
			// Not sure exactly how this could happen but it violates expectations.
			readableErrorJson(w, http.StatusNotFound, "")
			return
		}
		if !IsHex(segments[1]) {
			readableErrorJson(w, http.StatusNotFound, "Need hex claim ID")
			return
		}
		claimID := sdk.ClaimId(segments[1])

		claim, err := sdk.GetClaim(claimID)
		if err != nil {
			switch err.(type) {
			case *sdk.ClaimNotFound:
				// It's not the expected format
				readableErrorJson(w, http.StatusNotFound, "Claim not found")
				return
			default:
				internalServiceErrorJson(w, err, "Error getting claim")
				return
			}
		}
		channel := objects.ChannelFromClaim(claim)

		var response []byte
		response, err = json.Marshal(channel)

		if err != nil {
			internalServiceErrorJson(w, err, "Error generating channel response")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(response))
	}

}

const PathChannel = "/channel/"

func main() {
	http.HandleFunc(PathChannel, handleChannel)
	http.ListenAndServe("localhost:8000", nil)
}
