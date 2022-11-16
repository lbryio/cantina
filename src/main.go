package main

import (
	"encoding/json"
	"fmt"
	"lbryio/cantina/objects"
	"lbryio/cantina/sdk"
	"log"
	"net/http"
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

func handleChannel(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		claim := sdk.GetClaim("aoeu")
		channel = objects.channelFromClaim(claim)
	}
}

const PathChannel = "/channel/"

func main() {
	http.HandleFunc(PathChannel, handleChannel)
	http.ListenAndServe("localhost:8000", nil)
}
