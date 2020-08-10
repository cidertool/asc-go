package asc

import (
	"context"
	"fmt"
	"time"
)

// AppStoreVersionState defines model for AppStoreVersionState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionstate
type AppStoreVersionState string

// List of AppStoreVersionState
const (
	AppStoreVersionStateDeveloperRejected          AppStoreVersionState = "DEVELOPER_REJECTED"
	AppStoreVersionStateDeveloperRemovedFromSale   AppStoreVersionState = "DEVELOPER_REMOVED_FROM_SALE"
	AppStoreVersionStateInvalidBinary              AppStoreVersionState = "INVALID_BINARY"
	AppStoreVersionStateInReview                   AppStoreVersionState = "IN_REVIEW"
	AppStoreVersionStateMetadataRejected           AppStoreVersionState = "METADATA_REJECTED"
	AppStoreVersionStatePendingAppleRelease        AppStoreVersionState = "PENDING_APPLE_RELEASE"
	AppStoreVersionStatePendingContract            AppStoreVersionState = "PENDING_CONTRACT"
	AppStoreVersionStatePendingDeveloperRelease    AppStoreVersionState = "PENDING_DEVELOPER_RELEASE"
	AppStoreVersionStatePreorderReadyForSale       AppStoreVersionState = "PREORDER_READY_FOR_SALE"
	AppStoreVersionStatePrepareForSubmission       AppStoreVersionState = "PREPARE_FOR_SUBMISSION"
	AppStoreVersionStateProcessingForAppStore      AppStoreVersionState = "PROCESSING_FOR_APP_STORE"
	AppStoreVersionStateReadyForSale               AppStoreVersionState = "READY_FOR_SALE"
	AppStoreVersionStateRejected                   AppStoreVersionState = "REJECTED"
	AppStoreVersionStateRemovedFromSale            AppStoreVersionState = "REMOVED_FROM_SALE"
	AppStoreVersionStateReplacedWithNewVersion     AppStoreVersionState = "REPLACED_WITH_NEW_VERSION"
	AppStoreVersionStateWaitingForExportCompliance AppStoreVersionState = "WAITING_FOR_EXPORT_COMPLIANCE"
	AppStoreVersionStateWaitingForReview           AppStoreVersionState = "WAITING_FOR_REVIEW"
)

// AppStoreVersionUpdateRequest defines model for AppStoreVersionUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest
type AppStoreVersionUpdateRequest struct {
	Attributes    *AppStoreVersionUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Relationships *AppStoreVersionUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// AppStoreVersionUpdateRequestAttributes are attributes for AppStoreVersionUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest/data/attributes
type AppStoreVersionUpdateRequestAttributes struct {
	Copyright           *string    `json:"copyright,omitempty"`
	Downloadable        *bool      `json:"downloadable,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	ReleaseType         *string    `json:"releaseType,omitempty"`
	UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
	VersionString       *string    `json:"versionString,omitempty"`
}

// AppStoreVersionUpdateRequestRelationships are relationships for AppStoreVersionUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest/data/relationships
type AppStoreVersionUpdateRequestRelationships struct {
	Build *RelationshipDeclaration `json:"build,omitempty"`
}

// AgeRatingDeclaration defines model for AgeRatingDeclaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclaration
type AgeRatingDeclaration struct {
	Attributes *AgeRatingDeclarationAttributes `json:"attributes,omitempty"`
	ID         string                          `json:"id"`
	Links      ResourceLinks                   `json:"links"`
	Type       string                          `json:"type"`
}

// AgeRatingDeclarationAttributes defines model for AgeRatingDeclaration.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclaration/attributes
type AgeRatingDeclarationAttributes struct {
	AlcoholTobaccoOrDrugUseOrReferences         *string      `json:"alcoholTobaccoOrDrugUseOrReferences,omitempty"`
	GamblingAndContests                         *bool        `json:"gamblingAndContests,omitempty"`
	GamblingSimulated                           *string      `json:"gamblingSimulated,omitempty"`
	HorrorOrFearThemes                          *string      `json:"horrorOrFearThemes,omitempty"`
	KidsAgeBand                                 *KidsAgeBand `json:"kidsAgeBand,omitempty"`
	MatureOrSuggestiveThemes                    *string      `json:"matureOrSuggestiveThemes,omitempty"`
	MedicalOrTreatmentInformation               *string      `json:"medicalOrTreatmentInformation,omitempty"`
	ProfanityOrCrudeHumor                       *string      `json:"profanityOrCrudeHumor,omitempty"`
	SexualContentGraphicAndNudity               *string      `json:"sexualContentGraphicAndNudity,omitempty"`
	SexualContentOrNudity                       *string      `json:"sexualContentOrNudity,omitempty"`
	UnrestrictedWebAccess                       *bool        `json:"unrestrictedWebAccess,omitempty"`
	ViolenceCartoonOrFantasy                    *string      `json:"violenceCartoonOrFantasy,omitempty"`
	ViolenceRealistic                           *string      `json:"violenceRealistic,omitempty"`
	ViolenceRealisticProlongedGraphicOrSadistic *string      `json:"violenceRealisticProlongedGraphicOrSadistic,omitempty"`
}

// AppStoreVersion defines model for AppStoreVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion
type AppStoreVersion struct {
	Attributes    *AppStoreVersionAttributes    `json:"attributes,omitempty"`
	ID            string                        `json:"id"`
	Links         ResourceLinks                 `json:"links"`
	Relationships *AppStoreVersionRelationships `json:"relationships,omitempty"`
	Type          string                        `json:"type"`
}

// AppStoreVersionAttributes defines model for AppStoreVersion.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion/attributes
type AppStoreVersionAttributes struct {
	AppStoreState       *AppStoreVersionState `json:"appStoreState,omitempty"`
	Copyright           *string               `json:"copyright,omitempty"`
	CreatedDate         *time.Time            `json:"createdDate,omitempty"`
	Downloadable        *bool                 `json:"downloadable,omitempty"`
	EarliestReleaseDate *time.Time            `json:"earliestReleaseDate,omitempty"`
	Platform            *Platform             `json:"platform,omitempty"`
	ReleaseType         *string               `json:"releaseType,omitempty"`
	UsesIDFA            *bool                 `json:"usesIdfa,omitempty"`
	VersionString       *string               `json:"versionString,omitempty"`
}

// AppStoreVersionRelationships defines model for AppStoreVersion.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion/relationships
type AppStoreVersionRelationships struct {
	AgeRatingDeclaration         *Relationship      `json:"ageRatingDeclaration,omitempty"`
	App                          *Relationship      `json:"app,omitempty"`
	AppStoreReviewDetail         *Relationship      `json:"appStoreReviewDetail,omitempty"`
	AppStoreVersionLocalizations *PagedRelationship `json:"appStoreVersionLocalizations,omitempty"`
	AppStoreVersionPhasedRelease *Relationship      `json:"appStoreVersionPhasedRelease,omitempty"`
	AppStoreVersionSubmission    *Relationship      `json:"appStoreVersionSubmission,omitempty"`
	Build                        *Relationship      `json:"build,omitempty"`
	IDFADeclaration              *Relationship      `json:"idfaDeclaration,omitempty"`
	RoutingAppCoverage           *Relationship      `json:"routingAppCoverage,omitempty"`
}

// AppStoreVersionResponse defines model for AppStoreVersionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionresponse
type AppStoreVersionResponse struct {
	Data     AppStoreVersion `json:"data"`
	Included *[]interface{}  `json:"included,omitempty"`
	Links    DocumentLinks   `json:"links"`
}

// AppStoreVersionsResponse defines model for AppStoreVersionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsresponse
type AppStoreVersionsResponse struct {
	Data     []AppStoreVersion  `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppStoreVersionCreateRequest defines model for AppStoreVersionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest
type AppStoreVersionCreateRequest struct {
	Attributes    AppStoreVersionCreateRequestAttributes    `json:"attributes"`
	Relationships AppStoreVersionCreateRequestRelationships `json:"relationships"`
	Type          string                                    `json:"type"`
}

// AppStoreVersionCreateRequestAttributes are attributes for AppStoreVersionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest/data/attributes
type AppStoreVersionCreateRequestAttributes struct {
	Copyright           *string    `json:"copyright,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	Platform            Platform   `json:"platform"`
	ReleaseType         *string    `json:"releaseType,omitempty"`
	UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
	VersionString       string     `json:"versionString"`
}

// AppStoreVersionCreateRequestRelationships are relationships for AppStoreVersionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest/data/relationships
type AppStoreVersionCreateRequestRelationships struct {
	App   RelationshipDeclaration  `json:"app"`
	Build *RelationshipDeclaration `json:"build,omitempty"`
}

// AppStoreVersionBuildLinkageRequest is a list of relationships to Build objects
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionbuildlinkagerequest
type AppStoreVersionBuildLinkageRequest RelationshipData

// AppStoreVersionBuildLinkageResponse defines model for AppStoreVersionBuildLinkageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionbuildlinkageresponse
type AppStoreVersionBuildLinkageResponse struct {
	Data  RelationshipData `json:"data"`
	Links DocumentLinks    `json:"links"`
}

// ListAppStoreVersionsQuery are query options for ListAppStoreVersions
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_versions_for_an_app
type ListAppStoreVersionsQuery struct {
	FieldsApps                          []string `url:"fields[apps],omitempty"`
	FieldsAppStoreVersionSubmissions    []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreReviewDetails         []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAgeRatingDeclarations         []string `url:"fields[ageRatingDeclarations],omitempty"`
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           []string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              []string `url:"fields[idfaDeclarations],omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	Include                             []string `url:"include,omitempty"`
	FilterID                            []string `url:"filter[id],omitempty"`
	FilterVersionString                 []string `url:"filter[versionString],omitempty"`
	FilterPlatform                      []string `url:"filter[platform],omitempty"`
	FilterAppStoreState                 []string `url:"filter[appStoreState],omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetAppStoreVersionQuery are query options for GetAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_information
type GetAppStoreVersionQuery struct {
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions    []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsAppStoreReviewDetails         []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAgeRatingDeclarations         []string `url:"fields[ageRatingDeclarations],omitempty"`
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           []string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              []string `url:"fields[idfaDeclarations],omitempty"`
	FieldsAppStoreVersionLocalizations  []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                             []string `url:"include,omitempty"`
	LimitAppStoreVersionLocalizations   int      `url:"limit[appStoreVersionLocalizations],omitempty"`
}

// GetAgeRatingDeclarationForAppStoreVersionQuery are query options for GetAgeRatingDeclarationForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_age_rating_declaration_information_of_an_app_store_version
type GetAgeRatingDeclarationForAppStoreVersionQuery struct {
	FieldsAgeRatingDeclarations []string `url:"fields[ageRatingDeclarations],omitempty"`
}

func (r *AppStoreVersionCreateRequest) applyTypes() {
	if r == nil {
		return
	}
	r.Type = "appStoreVersions"
	r.Relationships.App.applyType("apps")
	r.Relationships.Build.applyType("builds")
}

func (r *AppStoreVersionUpdateRequest) applyTypes() {
	if r == nil {
		return
	}
	r.Type = "appStoreVersions"
	if r.Relationships == nil {
		return
	}
	r.Relationships.Build.applyType("builds")
}

func (r *AppStoreVersionBuildLinkageRequest) applyTypes() {
	if r == nil {
		return
	}
	r.Type = "builds"
}

// ListAppStoreVersionsForApp gets a list of all App Store versions of an app across all platforms.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_versions_for_an_app
func (s *AppsService) ListAppStoreVersionsForApp(ctx context.Context, id string, params *ListAppStoreVersionsQuery) (*AppStoreVersionsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/appStoreVersions", id)
	res := new(AppStoreVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppStoreVersion gets information for a specific app store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_information
func (s *AppsService) GetAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionQuery) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateAppStoreVersion adds a new App Store version or platform to an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version
func (s *AppsService) CreateAppStoreVersion(ctx context.Context, body *AppStoreVersionCreateRequest) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions")
	res := new(AppStoreVersionResponse)
	resp, err := s.client.post(ctx, url, body, res)
	return res, resp, err
}

// UpdateAppStoreVersion updates the app store version for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_version
func (s *AppsService) UpdateAppStoreVersion(ctx context.Context, id string, body *AppStoreVersionUpdateRequest) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersion deletes an app store version that is associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version
func (s *AppsService) DeleteAppStoreVersion(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	return s.client.delete(ctx, url, nil)
}

// GetBuildIDForAppStoreVersion gets the ID of the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_the_build_id_for_an_app_store_version
func (s *AppsService) GetBuildIDForAppStoreVersion(ctx context.Context, id string) (*AppStoreVersionBuildLinkageResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}

// UpdateBuildForAppStoreVersion changes the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_the_build_for_an_app_store_version
func (s *AppsService) UpdateBuildForAppStoreVersion(ctx context.Context, id string, linkage *AppStoreVersionBuildLinkageRequest) (*AppStoreVersionBuildLinkageResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.patch(ctx, url, linkage, res)
	return res, resp, err
}

// GetAgeRatingDeclarationForAppStoreVersion gets the age-related information declared for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_age_rating_declaration_information_of_an_app_store_version
func (s *AppsService) GetAgeRatingDeclarationForAppStoreVersion(ctx context.Context, id string, params *GetAgeRatingDeclarationForAppStoreVersionQuery) (*AgeRatingDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/ageRatingDeclaration", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
