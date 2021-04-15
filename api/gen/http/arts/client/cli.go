// Code generated by goa v3.3.1, DO NOT EDIT.
//
// arts HTTP client CLI support package
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package client

import (
	"encoding/json"
	"fmt"

	arts "github.com/pastelnetwork/walletnode/api/gen/arts"
	goa "goa.design/goa/v3/pkg"
)

// BuildRegisterPayload builds the payload for the arts register endpoint from
// CLI flags.
func BuildRegisterPayload(artsRegisterBody string) (*arts.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(artsRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"address\": \"12349231421309dsfdf\",\n      \"artist_name\": \"Leonardo da Vinci\",\n      \"artwork_name\": \"Mona Lisa\",\n      \"fee\": 100,\n      \"issued_copies\": 5,\n      \"pastel_id\": \"123456789\"\n   }'")
		}
	}
	v := &arts.RegisterPayload{
		ArtistName:   body.ArtistName,
		ArtworkName:  body.ArtworkName,
		IssuedCopies: body.IssuedCopies,
		Fee:          body.Fee,
		PastelID:     body.PastelID,
		Address:      body.Address,
	}
	{
		var zero int
		if v.IssuedCopies == zero {
			v.IssuedCopies = 1
		}
	}

	return v, nil
}

// BuildUploadImagePayload builds the payload for the arts uploadImage endpoint
// from CLI flags.
func BuildUploadImagePayload(artsUploadImageBody string) (*arts.ImageUploadPayload, error) {
	var err error
	var body UploadImageRequestBody
	{
		err = json.Unmarshal([]byte(artsUploadImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"file\": \"TWF4aW1lIGV0Lg==\"\n   }'")
		}
		if body.File == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &arts.ImageUploadPayload{
		File: body.File,
	}

	return v, nil
}
