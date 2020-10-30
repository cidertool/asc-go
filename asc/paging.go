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
)

// Reference is a wrapper type for a URL that contains a cursor parameter.
type Reference struct {
	url.URL
}

// Cursor returns the cursor parameter on the Reference's internal URL.
func (r Reference) Cursor() string {
	return r.Query().Get("cursor")
}

// MarshalJSON marshals the Reference into a JSON fragment.
func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON unmarshals the JSON fragment into a Reference.
func (r *Reference) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	if u != nil {
		r.URL = *u
	}

	return nil
}

// PagedDocumentLinks defines model for PagedDocumentLinks.
type PagedDocumentLinks struct {
	First *Reference `json:"first,omitempty"`
	Next  *Reference `json:"next,omitempty"`
	Self  Reference  `json:"self"`
}

// PagingInformation defines model for PagingInformation.
type PagingInformation struct {
	Paging struct {
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"paging"`
}

// ResourceLinks defines model for ResourceLinks.
type ResourceLinks struct {
	Self Reference `json:"self"`
}

// DocumentLinks defines model for DocumentLinks.
type DocumentLinks struct {
	Self Reference `json:"self"`
}

// Relationship contains data about a related resources as well as API references that can be followed.
type Relationship struct {
	Data  *RelationshipData  `json:"data,omitempty"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

// PagedRelationship is a relationship to multiple resources that have paging information.
type PagedRelationship struct {
	Data  []RelationshipData `json:"data,omitempty"`
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// RelationshipData contains data on the given relationship.
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipLinks contains links on the given relationship.
type RelationshipLinks struct {
	Related *Reference `json:"related,omitempty"`
	Self    *Reference `json:"self,omitempty"`
}

// relationshipDeclaration represents a declared relationship to a single resource.
type relationshipDeclaration struct {
	Data RelationshipData `json:"data"`
}

// pagedRelationshipDeclaration represents a declared relationship to multiple resources.
type pagedRelationshipDeclaration struct {
	Data []RelationshipData `json:"data"`
}

func newRelationshipDeclaration(id *string, relationshipType string) *relationshipDeclaration {
	if id == nil {
		return nil
	}

	return &relationshipDeclaration{
		Data: RelationshipData{
			ID:   *id,
			Type: relationshipType,
		},
	}
}

func newPagedRelationshipDeclaration(ids []string, relationshipType string) pagedRelationshipDeclaration {
	datas := []RelationshipData{}
	for _, id := range ids {
		datas = append(datas, RelationshipData{
			ID:   id,
			Type: relationshipType,
		})
	}

	return pagedRelationshipDeclaration{Data: datas}
}
