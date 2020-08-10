package asc

import (
	"context"
	"fmt"
	"time"
)

// AppEncryptionDeclarationState defines model for AppEncryptionDeclarationState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationstate
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclaration
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
		App *Relationship `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppEncryptionDeclarationResponse defines model for AppEncryptionDeclarationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationresponse
type AppEncryptionDeclarationResponse struct {
	Data     AppEncryptionDeclaration `json:"data"`
	Included *[]App                   `json:"included,omitempty"`
	Links    DocumentLinks            `json:"links"`
}

// AppEncryptionDeclarationsResponse defines model for AppEncryptionDeclarationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appencryptiondeclarationsresponse
type AppEncryptionDeclarationsResponse struct {
	Data     []AppEncryptionDeclaration `json:"data"`
	Included *[]App                     `json:"included,omitempty"`
	Links    PagedDocumentLinks         `json:"links"`
	Meta     *PagingInformation         `json:"meta,omitempty"`
}

// ListAppEncryptionDeclarationsQuery are query options for ListAppEncryptionDeclarations
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_encryption_declarations
type ListAppEncryptionDeclarationsQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FilterApp                       []string `url:"filter[app],omitempty"`
	FilterBuilds                    []string `url:"filter[builds],omitempty"`
	FilterPlatforms                 []string `url:"filter[platforms],omitempty"`
	Include                         []string `url:"include,omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	Cursor                          *string  `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationQuery are query options for GetAppEncryptionDeclaration
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_encryption_declaration_information
type GetAppEncryptionDeclarationQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	Include                         []string `url:"include,omitempty"`
}

// GetAppForEncryptionDeclarationQuery are query options for GetAppForEncryptionDeclaration
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_an_app_encryption_declaration
type GetAppForEncryptionDeclarationQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListAppEncryptionDeclarations finds and lists all available app encryption declarations.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_app_encryption_declarations
func (s *BuildsService) ListAppEncryptionDeclarations(ctx context.Context, params *ListAppEncryptionDeclarationsQuery) (*AppEncryptionDeclarationsResponse, *Response, error) {
	res := new(AppEncryptionDeclarationsResponse)
	resp, err := s.client.get(ctx, "appEncryptionDeclarations", params, res)
	return res, resp, err
}

// GetAppEncryptionDeclaration gets information about a specific app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_encryption_declaration_information
func (s *BuildsService) GetAppEncryptionDeclaration(ctx context.Context, id string, params *GetAppEncryptionDeclarationQuery) (*AppEncryptionDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppForAppEncryptionDeclaration gets the app information from a specific app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_an_app_encryption_declaration
func (s *BuildsService) GetAppForAppEncryptionDeclaration(ctx context.Context, id string, params *GetAppForEncryptionDeclarationQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("appEncryptionDeclarations/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// AssignBuildsToAppEncryptionDeclaration assigns one or more builds to an app encryption declaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_builds_to_an_app_encryption_declaration
func (s *BuildsService) AssignBuildsToAppEncryptionDeclaration(ctx context.Context, id string, linkages *[]RelationshipData) (*Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.client.post(ctx, url, linkages, nil)
}
