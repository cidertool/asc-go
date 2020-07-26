package builds

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal/services"
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
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included *[]apps.App              `json:"included,omitempty"`
	Links    services.DocumentLinks   `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration  `json:"data"`
	Included *[]apps.App                 `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// ListAppEncryptionDeclarationsQuery are query options for ListAppEncryptionDeclarations
type ListAppEncryptionDeclarationsQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FilterApp                       *[]string `url:"filter[app],omitempty"`
	FilterBuilds                    *[]string `url:"filter[builds],omitempty"`
	FilterPlatforms                 *[]string `url:"filter[platforms],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	Limit                           *int      `url:"limit,omitempty"`
	Cursor                          *string   `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationQuery are query options for GetAppEncryptionDeclaration
type GetAppEncryptionDeclarationQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
}

// GetAppForEncryptionDeclarationQuery are query options for GetAppForEncryptionDeclaration
type GetAppForEncryptionDeclarationQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// ListAppEncryptionDeclarations finds and lists all available app encryption declarations.
func (s *Service) ListAppEncryptionDeclarations(params *ListAppEncryptionDeclarationsQuery) (*AppEncryptionDeclarationsResponse, *services.Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.GetWithQuery("appEncryptionDeclarations", params, res)
	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
func (s *Service) GetAppEncryptionDeclaration(id string, params *GetAppEncryptionDeclarationQuery) (*AppEncryptionDeclarationResponse, *services.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
func (s *Service) GetAppForAppEncryptionDeclaration(id string, params *GetAppForEncryptionDeclarationQuery) (*apps.AppResponse, *services.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
func (s *Service) AssignBuildsToAppEncryptionDeclaration(id string, linkages *[]services.RelationshipsData) (*services.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.Post(url, linkages, nil)
}