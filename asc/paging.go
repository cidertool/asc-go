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
func (r *Reference) Cursor() string {
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

// RelationshipsData contains data on the given relationship
type RelationshipsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipsLinks contains links on the given relationship
type RelationshipsLinks struct {
	Related *Reference `json:"related,omitempty"`
	Self    *Reference `json:"self,omitempty"`
}
