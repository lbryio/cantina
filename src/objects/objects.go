package objects

import (
	"lbryio/cantina/sdk"

	vocab "github.com/go-ap/activitypub"
)

type Channel vocab.Actor

func channelFromClaim(claim sdk.Claim) Channel {
	return Channel{
		ID:   claim.PermanentURL,
		Type: vocab.Person,
	}
}
