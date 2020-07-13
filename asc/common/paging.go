package common

// PagedDocumentLinks defines model for PagedDocumentLinks.
type PagedDocumentLinks struct {
	First *string `json:"first,omitempty"`
	Next  *string `json:"next,omitempty"`
	Self  string  `json:"self"`
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
	Self string `json:"self"`
}

// DocumentLinks defines model for DocumentLinks.
type DocumentLinks struct {
	Self string `json:"self"`
}
