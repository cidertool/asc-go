package builds

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppEncryptionDeclarationState defines model for AppEncryptionDeclarationState.
type AppEncryptionDeclarationState string

// List of AppEncryptionDeclarationState
const (
	Approved AppEncryptionDeclarationState = "APPROVED"
	Expired  AppEncryptionDeclarationState = "EXPIRED"
	Invalid  AppEncryptionDeclarationState = "INVALID"
	InReview AppEncryptionDeclarationState = "IN_REVIEW"
	Rejected AppEncryptionDeclarationState = "REJECTED"
)

// AppEncryptionDeclaration defines model for AppEncryptionDeclaration.
type AppEncryptionDeclaration struct {
	Attributes *struct {
		AppEncryptionDeclarationState   *AppEncryptionDeclarationState `json:"appEncryptionDeclarationState,omitempty"`
		AvailableOnFrenchStore          *bool                          `json:"availableOnFrenchStore,omitempty"`
		CodeValue                       *string                        `json:"codeValue,omitempty"`
		ContainsProprietaryCryptography *bool                          `json:"containsProprietaryCryptography,omitempty"`
		ContainsThirdPartyCryptography  *bool                          `json:"containsThirdPartyCryptography,omitempty"`
		DocumentName                    *string                        `json:"documentName,omitempty"`
		DocumentType                    *string                        `json:"documentType,omitempty"`
		DocumentURL                     *string                        `json:"documentUrl,omitempty"`
		Exempt                          *bool                          `json:"exempt,omitempty"`
		Platform                        *apps.Platform                 `json:"platform,omitempty"`
		UploadedDate                    *time.Time                     `json:"uploadedDate,omitempty"`
		UsesEncryption                  *bool                          `json:"usesEncryption,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppEncryptionDeclarationBuildsLinkagesRequest defines model for AppEncryptionDeclarationBuildsLinkagesRequest.
type AppEncryptionDeclarationBuildsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included *[]apps.App              `json:"included,omitempty"`
	Links    common.DocumentLinks     `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration `json:"data"`
	Included *[]apps.App                `json:"included,omitempty"`
	Links    common.PagedDocumentLinks  `json:"links"`
	Meta     *common.PagingInformation  `json:"meta,omitempty"`
}

type ListAppEncryptionDeclarationsOptions struct {
	Fields *struct {
		AppEncryptionDeclarations *[]string `url:"appEncryptionDeclarations,omitempty"`
		Apps                      *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App       *[]string `url:"app,omitempty"`
		Builds    *[]string `url:"builds,omitempty"`
		Platforms *[]string `url:"platforms,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
}

type GetAppEncryptionDeclarationOptions struct {
	Fields *struct {
		AppEncryptionDeclarations *[]string `url:"appEncryptionDeclarations,omitempty"`
		Apps                      *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetAppForEncryptionDeclarationOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

// ListAppEncryptionDeclarations finds and lists all available app encryption declarations.
func (s *Service) ListAppEncryptionDeclarations(params *ListAppEncryptionDeclarationsOptions) (*AppEncryptionDeclarationsResponse, *common.Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.Get("appEncryptionDeclarations", params, res)
	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
func (s *Service) GetAppEncryptionDeclaration(id string, params *GetAppEncryptionDeclarationOptions) (*AppEncryptionDeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
func (s *Service) GetAppForAppEncryptionDeclaration(id string, params *GetAppForEncryptionDeclarationOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
func (s *Service) AssignBuildsToAppEncryptionDeclaration(id string, body *AppEncryptionDeclarationBuildsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.Post(url, body, nil)
}
