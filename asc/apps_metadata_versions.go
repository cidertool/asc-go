package asc

import (
	"fmt"
	"net/http"
	"time"
)

// AppStoreVersionState defines model for AppStoreVersionState.
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
type AppStoreVersionUpdateRequest struct {
	Attributes    *AppStoreVersionUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Relationships *AppStoreVersionUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// AppStoreVersionUpdateRequestAttributes are attributes for AppStoreVersionUpdateRequest
type AppStoreVersionUpdateRequestAttributes struct {
	Copyright           *string    `json:"copyright,omitempty"`
	Downloadable        *bool      `json:"downloadable,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	ReleaseType         *string    `json:"releaseType,omitempty"`
	UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
	VersionString       *string    `json:"versionString,omitempty"`
}

// AppStoreVersionUpdateRequestRelationships are relationships for AppStoreVersionUpdateRequest
type AppStoreVersionUpdateRequestRelationships struct {
	Build *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"build,omitempty"`
}

// AgeRatingDeclaration defines model for AgeRatingDeclaration.
type AgeRatingDeclaration struct {
	Attributes *struct {
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
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// AppStoreVersion defines model for AppStoreVersion.
type AppStoreVersion struct {
	Attributes *struct {
		AppStoreState       *AppStoreVersionState `json:"appStoreState,omitempty"`
		Copyright           *string               `json:"copyright,omitempty"`
		CreatedDate         *time.Time            `json:"createdDate,omitempty"`
		Downloadable        *bool                 `json:"downloadable,omitempty"`
		EarliestReleaseDate *time.Time            `json:"earliestReleaseDate,omitempty"`
		Platform            *Platform             `json:"platform,omitempty"`
		ReleaseType         *string               `json:"releaseType,omitempty"`
		UsesIDFA            *bool                 `json:"usesIdfa,omitempty"`
		VersionString       *string               `json:"versionString,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AgeRatingDeclaration *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"ageRatingDeclaration,omitempty"`
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppStoreReviewDetail *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreReviewDetail,omitempty"`
		AppStoreVersionLocalizations *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appStoreVersionLocalizations,omitempty"`
		AppStoreVersionPhasedRelease *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionPhasedRelease,omitempty"`
		AppStoreVersionSubmission *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionSubmission,omitempty"`
		Build *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
		IDFADeclaration *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"idfaDeclaration,omitempty"`
		RoutingAppCoverage *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"routingAppCoverage,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionResponse defines model for AppStoreVersionResponse.
type AppStoreVersionResponse struct {
	Data     AppStoreVersion `json:"data"`
	Included *[]interface{}  `json:"included,omitempty"`
	Links    DocumentLinks   `json:"links"`
}

// AppStoreVersionsResponse defines model for AppStoreVersionsResponse.
type AppStoreVersionsResponse struct {
	Data     []AppStoreVersion  `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppStoreVersionCreateRequest defines model for AppStoreVersionCreateRequest.
type AppStoreVersionCreateRequest struct {
	Attributes    AppStoreVersionCreateRequestAttributes    `json:"attributes"`
	Relationships AppStoreVersionCreateRequestRelationships `json:"relationships"`
	Type          string                                    `json:"type"`
}

// AppStoreVersionCreateRequestAttributes are attributes for AppStoreVersionCreateRequest
type AppStoreVersionCreateRequestAttributes struct {
	Copyright           *string    `json:"copyright,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	Platform            Platform   `json:"platform"`
	ReleaseType         *string    `json:"releaseType,omitempty"`
	UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
	VersionString       string     `json:"versionString"`
}

// AppStoreVersionCreateRequestRelationships are relationships for AppStoreVersionCreateRequest
type AppStoreVersionCreateRequestRelationships struct {
	App struct {
		Data RelationshipsData `json:"data"`
	} `json:"app"`
	Build *struct {
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"build,omitempty"`
}

// AppStoreVersionBuildLinkageResponse defines model for AppStoreVersionBuildLinkageResponse.
type AppStoreVersionBuildLinkageResponse struct {
	Data  RelationshipsData `json:"data"`
	Links DocumentLinks     `json:"links"`
}

// ListAppStoreVersionsQuery are query options for ListAppStoreVersions
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
type GetAgeRatingDeclarationForAppStoreVersionQuery struct {
	FieldsAgeRatingDeclarations []string `url:"fields[ageRatingDeclarations],omitempty"`
}

// GetRoutingAppCoverageForVersionQuery are query options for GetRoutingAppCoverageForVersion
type GetRoutingAppCoverageForVersionQuery struct {
	FieldsRoutingAppCoverages []string `url:"fields[routingAppCoverages],omitempty"`
}

// ListAppStoreVersionsForApp gets a list of all App Store versions of an app across all platforms.
func (s *AppsService) ListAppStoreVersionsForApp(id string, params *ListAppStoreVersionsQuery) (*AppStoreVersionsResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/appStoreVersions", id)
	res := new(AppStoreVersionsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppStoreVersion gets information for a specific app store version.
func (s *AppsService) GetAppStoreVersion(id string, params *GetAppStoreVersionQuery) (*AppStoreVersionResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateAppStoreVersion adds a new App Store version or platform to an app.
func (s *AppsService) CreateAppStoreVersion(body *AppStoreVersionCreateRequest) (*AppStoreVersionResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions")
	res := new(AppStoreVersionResponse)
	resp, err := s.client.post(url, body, res)
	return res, resp, err
}

// UpdateAppStoreVersion updates the app store version for a specific app.
func (s *AppsService) UpdateAppStoreVersion(id string, body *AppStoreVersionUpdateRequest) (*AppStoreVersionResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersion deletes an app store version that is associated with an app.
func (s *AppsService) DeleteAppStoreVersion(id string) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	return s.client.delete(url, nil)
}

// GetBuildIDForAppStoreVersion gets the ID of the build that is attached to a specific App Store version.
func (s *AppsService) GetBuildIDForAppStoreVersion(id string) (*AppStoreVersionBuildLinkageResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.get(url, nil, res)
	return res, resp, err
}

// UpdateBuildForAppStoreVersion changes the build that is attached to a specific App Store version.
func (s *AppsService) UpdateBuildForAppStoreVersion(id string, linkages *[]RelationshipsData) (*AppStoreVersionBuildLinkageResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.patch(url, linkages, res)
	return res, resp, err
}

// GetAgeRatingDeclarationForAppStoreVersion gets the age-related information declared for your app.
func (s *AppsService) GetAgeRatingDeclarationForAppStoreVersion(id string, params *GetAgeRatingDeclarationForAppStoreVersionQuery) (*AgeRatingDeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/ageRatingDeclaration", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetRoutingAppCoverageForAppStoreVersion gets the routing app coverage file that is associated with a specific App Store version
func (s *AppsService) GetRoutingAppCoverageForAppStoreVersion(id string, params *GetRoutingAppCoverageForVersionQuery) (*RoutingAppCoverageResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/routingAppCoverage", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
