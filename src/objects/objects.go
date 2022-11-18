package objects

import (
	"lbryio/cantina/sdk"

	"github.com/go-ap/activitypub"
	vocab "github.com/go-ap/activitypub"
)

type Channel vocab.Actor

func ChannelFromClaim(claim sdk.Claim) Channel {
	return Channel{
		ID:   activitypub.IRI(claim.PermanentURL),
		Type: vocab.PersonType,
		Name: vocab.NaturalLanguageValues{{Ref: vocab.NilLangRef, Value: vocab.Content(claim.Value.Title)}},

		PreferredUsername: vocab.NaturalLanguageValues{{Ref: vocab.NilLangRef, Value: vocab.Content(claim.Name[1:])}},
	}
}
