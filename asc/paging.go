package asc

import (
	"encoding/json"
	"net/url"
)

// Reference is a wrapper type for a URL that contains a cursor parameter
type Reference struct {
	url.URL
}

// Cursor returns the cursor parameter on the Reference's internal URL.
func (r Reference) Cursor() string {
	return r.Query().Get("cursor")
}

// MarshalJSON marshals the Reference into a JSON fragment
func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON unmarshals the JSON fragment into a Reference
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

// Relationship contains data about a related resources as well as API references that can be followed
type Relationship struct {
	Data  *RelationshipData  `json:"data,omitempty"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

// PagedRelationship is a relationship to multiple resources that have paging information
type PagedRelationship struct {
	Data  *[]RelationshipData `json:"data,omitempty"`
	Links *RelationshipLinks  `json:"links,omitempty"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// RelationshipDeclaration represents a declared relationship to a single resource
type RelationshipDeclaration struct {
	Data *RelationshipData `json:"data"`
}

// PagedRelationshipDeclaration represents a declared relationship to multiple resources
type PagedRelationshipDeclaration struct {
	Data *[]RelationshipData `json:"data"`
}

// RelationshipData contains data on the given relationship
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipLinks contains links on the given relationship
type RelationshipLinks struct {
	Related *Reference `json:"related,omitempty"`
	Self    *Reference `json:"self,omitempty"`
}

func (d *RelationshipDeclaration) applyType(t string) {
	if d == nil ||
		d.Data == nil ||
		d.Data.Type != "" {
		return
	}
	d.Data.Type = t
}

func (d *PagedRelationshipDeclaration) applyType(t string) {
	if d == nil || d.Data == nil {
		return
	}
	data := *d.Data
	for i := range data {
		data[i].applyType(t)
	}
}

func (d *RelationshipData) applyType(t string) {
	if d == nil || d.Type != "" {
		return
	}
	d.Type = t
}
