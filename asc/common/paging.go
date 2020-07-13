package common

import (
	"encoding/json"
	"net/url"
)

type Reference struct {
	url.URL
}

func (r *Reference) Cursor() string {
	var cursor string
	if r == nil {
		return cursor
	}
	values := r.Query()
	cursor = values.Get("cursor")
	return cursor
}

func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

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
