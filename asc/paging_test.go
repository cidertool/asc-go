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

func TestNewRelationships(t *testing.T) {
	var rel *relationshipDeclaration
	rel = newRelationshipDeclaration(nil, "dog")
	assert.Nil(t, rel)
	id := "10"
	rel = newRelationshipDeclaration(&id, "dog")
	assert.Equal(t, &relationshipDeclaration{RelationshipData{"10", "dog"}}, rel)

	var rels pagedRelationshipDeclaration
	rels = newPagedRelationshipDeclaration(nil, "dog")
	assert.Empty(t, rels.Data)
	rels = newPagedRelationshipDeclaration([]string{}, "dog")
	assert.Empty(t, rels.Data)
	rels = newPagedRelationshipDeclaration([]string{"10", "20", "30"}, "dog")
	assert.Equal(t, pagedRelationshipDeclaration{[]RelationshipData{{"10", "dog"}, {"20", "dog"}, {"30", "dog"}}}, rels)
}
