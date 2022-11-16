package objects

import (
	"lbryio/cantina/sdk"

	"github.com/go-ap/activitypub"
	vocab "github.com/go-ap/activitypub"
)

type Channel vocab.Actor

func ChannelFromClaim(claim sdk.Claim) Channel {
	return Channel{
		ID:   claim.PermanentURL,
		Type: vocab.Person,
	}
}
