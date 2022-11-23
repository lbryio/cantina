package objects

import (
	"lbryio/cantina/sdk"
	"time"

	"github.com/go-ap/activitypub"
	vocab "github.com/go-ap/activitypub"
)

type Channel vocab.Actor

func ChannelFromClaim(claim sdk.Claim) Channel {
	return Channel{
		ID: activitypub.IRI(claim.PermanentURL),

		// Ideally Type should allow for Group or Organization but I don't think that distinction is in the claim
		Type: vocab.PersonType,

		Name:      vocab.NaturalLanguageValues{{Ref: vocab.NilLangRef, Value: vocab.Content(claim.Value.Title)}},
		Published: time.Unix(int64(claim.Meta.CreationTimestamp), 0),
		Updated:   time.Unix(int64(claim.Timestamp), 0),

		PreferredUsername: vocab.NaturalLanguageValues{{Ref: vocab.NilLangRef, Value: vocab.Content(claim.Name[1:])}},
	}
}
