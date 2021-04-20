/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReference(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	ref := new(Reference)
	assert.Empty(t, ref.Cursor())
	ref = &Reference{url.URL{}}
	assert.Empty(t, ref.Cursor())
}

func TestReferenceBadUnmarshal(t *testing.T) {
	t.Parallel()

	marshaledNoString := `{"self":0}`
	marshaledNoURL := `{"self":":"}`

	var links DocumentLinks
	err := json.Unmarshal([]byte(marshaledNoString), &links)
	assert.Error(t, err)
	err = json.Unmarshal([]byte(marshaledNoURL), &links)
	assert.Error(t, err)
}

func TestNewRelationships(t *testing.T) {
	t.Parallel()

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
