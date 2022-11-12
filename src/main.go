package main

import (
	"lbryio/cantina/sdk"
	"net/http"
)

func handleChannel(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		claim := sdk.GetClaim("aoeu")
		channel = objects.channelFromClaim(claim)
	}
}

const PathChannel = "/channel/"

func main() {
	http.HandleFunc(PathChannel, handleChannel)
}
