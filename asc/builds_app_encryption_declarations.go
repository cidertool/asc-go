package asc

import (
	"fmt"
	"net/http"
	"time"
)

// AppEncryptionDeclarationState defines model for AppEncryptionDeclarationState.
type AppEncryptionDeclarationState string

// List of AppEncryptionDeclarationState
const (
	AppEncryptionDeclarationStateApproved AppEncryptionDeclarationState = "APPROVED"
	AppEncryptionDeclarationStateExpired  AppEncryptionDeclarationState = "EXPIRED"
	AppEncryptionDeclarationStateInvalid  AppEncryptionDeclarationState = "INVALID"
	AppEncryptionDeclarationStateInReview AppEncryptionDeclarationState = "IN_REVIEW"
	AppEncryptionDeclarationStateRejected AppEncryptionDeclarationState = "REJECTED"
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
		Platform                        *Platform                      `json:"platform,omitempty"`
		UploadedDate                    *time.Time                     `json:"uploadedDate,omitempty"`
		UsesEncryption                  *bool                          `json:"usesEncryption,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included *[]App                   `json:"included,omitempty"`
	Links    DocumentLinks            `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration `json:"data"`
	Included *[]App                     `json:"included,omitempty"`
	Links    PagedDocumentLinks         `json:"links"`
	Meta     *PagingInformation         `json:"meta,omitempty"`
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
func (s *BuildsService) ListAppEncryptionDeclarations(params *ListAppEncryptionDeclarationsQuery) (*AppEncryptionDeclarationsResponse, *http.Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.client.get("appEncryptionDeclarations", params, res)
	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
func (s *BuildsService) GetAppEncryptionDeclaration(id string, params *GetAppEncryptionDeclarationQuery) (*AppEncryptionDeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
func (s *BuildsService) GetAppForAppEncryptionDeclaration(id string, params *GetAppForEncryptionDeclarationQuery) (*AppResponse, *http.Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
func (s *BuildsService) AssignBuildsToAppEncryptionDeclaration(id string, linkages *[]RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.client.post(url, linkages, nil)
}
