package asc

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReference(t *testing.T) {
	marshaled := `{"self":"https://api.appstoreconnect.apple.com/me?cursor=TEST"}`
	var links DocumentLinks
	err := json.Unmarshal([]byte(marshaled), &links)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", links.Self.Cursor())
	remarshaled, err := json.Marshal(links)
	assert.NoError(t, err)
	assert.Equal(t, marshaled, string(remarshaled))
}

func TestReferenceNoCursor(t *testing.T) {
	ref := new(Reference)
	assert.Empty(t, ref.Cursor())
	ref = &Reference{url.URL{}}
	assert.Empty(t, ref.Cursor())
}

func TestReferenceBadUnmarshal(t *testing.T) {
	marshaledNoString := `{"self":0}`
	marshaledNoURL := `{"self":":"}`
	var links DocumentLinks
	var err error
	err = json.Unmarshal([]byte(marshaledNoString), &links)
	assert.Error(t, err)
	err = json.Unmarshal([]byte(marshaledNoURL), &links)
	assert.Error(t, err)
}
