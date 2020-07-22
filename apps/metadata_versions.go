package apps

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/internal"
)

// AppStoreVersionState defines model for AppStoreVersionState.
type AppStoreVersionState string

// List of AppStoreVersionState
const (
	DeveloperRejected          AppStoreVersionState = "DEVELOPER_REJECTED"
	DeveloperRemovedFromSale   AppStoreVersionState = "DEVELOPER_REMOVED_FROM_SALE"
	InvalidBinary              AppStoreVersionState = "INVALID_BINARY"
	InReview                   AppStoreVersionState = "IN_REVIEW"
	MetadataRejected           AppStoreVersionState = "METADATA_REJECTED"
	PendingAppleRelease        AppStoreVersionState = "PENDING_APPLE_RELEASE"
	PendingContract            AppStoreVersionState = "PENDING_CONTRACT"
	PendingDeveloperRelease    AppStoreVersionState = "PENDING_DEVELOPER_RELEASE"
	PreorderReadyForSale       AppStoreVersionState = "PREORDER_READY_FOR_SALE"
	PrepareForSubmission       AppStoreVersionState = "PREPARE_FOR_SUBMISSION"
	ProcessingForAppStore      AppStoreVersionState = "PROCESSING_FOR_APP_STORE"
	ReadyForSale               AppStoreVersionState = "READY_FOR_SALE"
	Rejected                   AppStoreVersionState = "REJECTED"
	RemovedFromSale            AppStoreVersionState = "REMOVED_FROM_SALE"
	ReplacedWithNewVersion     AppStoreVersionState = "REPLACED_WITH_NEW_VERSION"
	WaitingForExportCompliance AppStoreVersionState = "WAITING_FOR_EXPORT_COMPLIANCE"
	WaitingForReview           AppStoreVersionState = "WAITING_FOR_REVIEW"
)

// AppStoreVersionUpdateRequest defines model for AppStoreVersionUpdateRequest.
type AppStoreVersionUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Copyright           *string    `json:"copyright,omitempty"`
			Downloadable        *bool      `json:"downloadable,omitempty"`
			EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
			ReleaseType         *string    `json:"releaseType,omitempty"`
			UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
			VersionString       *string    `json:"versionString,omitempty"`
		} `json:"attributes,omitempty"`
		ID            string `json:"id"`
		Relationships *struct {
			Build *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"build,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
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
	ID    string                 `json:"id"`
	Links internal.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
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
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		AgeRatingDeclaration *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"ageRatingDeclaration,omitempty"`
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppStoreReviewDetail *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreReviewDetail,omitempty"`
		AppStoreVersionLocalizations *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"appStoreVersionLocalizations,omitempty"`
		AppStoreVersionPhasedRelease *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionPhasedRelease,omitempty"`
		AppStoreVersionSubmission *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersionSubmission,omitempty"`
		Build *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"build,omitempty"`
		IDFADeclaration *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"idfaDeclaration,omitempty"`
		RoutingAppCoverage *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"routingAppCoverage,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionResponse defines model for AppStoreVersionResponse.
type AppStoreVersionResponse struct {
	Data     AppStoreVersion        `json:"data"`
	Included *[]interface{}         `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
}

// AppStoreVersionsResponse defines model for AppStoreVersionsResponse.
type AppStoreVersionsResponse struct {
	Data     []AppStoreVersion           `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// AppStoreVersionCreateRequest defines model for AppStoreVersionCreateRequest.
type AppStoreVersionCreateRequest struct {
	Data struct {
		Attributes struct {
			Copyright           *string    `json:"copyright,omitempty"`
			EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
			Platform            Platform   `json:"platform"`
			ReleaseType         *string    `json:"releaseType,omitempty"`
			UsesIDFA            *bool      `json:"usesIdfa,omitempty"`
			VersionString       string     `json:"versionString"`
		} `json:"attributes"`
		Relationships struct {
			App struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"app"`
			Build *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"build,omitempty"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionBuildLinkageRequest defines model for AppStoreVersionBuildLinkageRequest.
type AppStoreVersionBuildLinkageRequest struct {
	Data internal.RelationshipsData `json:"data"`
}

// AppStoreVersionBuildLinkageResponse defines model for AppStoreVersionBuildLinkageResponse.
type AppStoreVersionBuildLinkageResponse struct {
	Data  internal.RelationshipsData `json:"data"`
	Links internal.DocumentLinks     `json:"links"`
}

type ListAppStoreVersionsQuery struct {
	FieldsApps                          *[]string `url:"fields[apps],omitempty"`
	FieldsAppStoreVersionSubmissions    *[]string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        *[]string `url:"fields[builds],omitempty"`
	FieldsAppStoreVersions              *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreReviewDetails         *[]string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAgeRatingDeclarations         *[]string `url:"fields[ageRatingDeclarations],omitempty"`
	FieldsAppStoreVersionPhasedReleases *[]string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           *[]string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              *[]string `url:"fields[idfaDeclarations],omitempty"`
	Limit                               *int      `url:"limit,omitempty"`
	Include                             *[]string `url:"include,omitempty"`
	FilterID                            *[]string `url:"filter[id],omitempty"`
	FilterVersionString                 *[]string `url:"filter[versionString],omitempty"`
	FilterPlatform                      *[]string `url:"filter[platform],omitempty"`
	FilterAppStoreState                 *[]string `url:"filter[appStoreState],omitempty"`
	Cursor                              *string   `url:"cursor,omitempty"`
}

type GetAppStoreVersionQuery struct {
	FieldsAppStoreVersions              *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions    *[]string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        *[]string `url:"fields[builds],omitempty"`
	FieldsAppStoreReviewDetails         *[]string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAgeRatingDeclarations         *[]string `url:"fields[ageRatingDeclarations],omitempty"`
	FieldsAppStoreVersionPhasedReleases *[]string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           *[]string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              *[]string `url:"fields[idfaDeclarations],omitempty"`
	FieldsAppStoreVersionLocalizations  *[]string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                             *[]string `url:"include,omitempty"`
	LimitAppStoreVersionLocalizations   *int      `url:"limit[appStoreVersionLocalizations],omitempty"`
}

type GetAgeRatingDeclarationForAppStoreVersionQuery struct {
	FieldsAgeRatingDeclarations *[]string `url:"fields[ageRatingDeclarations],omitempty"`
}

type GetRoutingAppCoverageForVersionQuery struct {
	FieldsRoutingAppCoverages *[]string `url:"fields[routingAppCoverages],omitempty"`
}

// ListAppStoreVersionsForApp gets a list of all App Store versions of an app across all platforms.
func (s *Service) ListAppStoreVersionsForApp(id string, params *ListAppStoreVersionsQuery) (*AppStoreVersionsResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/appStoreVersions", id)
	res := new(AppStoreVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppStoreVersion gets information for a specific app store version.
func (s *Service) GetAppStoreVersion(id string, params *GetAppStoreVersionQuery) (*AppStoreVersionResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppStoreVersion adds a new App Store version or platform to an app.
func (s *Service) CreateAppStoreVersion(body *AppStoreVersionCreateRequest) (*AppStoreVersionResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions")
	res := new(AppStoreVersionResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateAppStoreVersion updates the app store version for a specific app.
func (s *Service) UpdateAppStoreVersion(id string, body *AppStoreVersionUpdateRequest) (*AppStoreVersionResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersion deletes an app store version that is associated with an app.
func (s *Service) DeleteAppStoreVersion(id string) (*internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	return s.Delete(url, nil)
}

// GetBuildIDForAppStoreVersion gets the ID of the build that is attached to a specific App Store version.
func (s *Service) GetBuildIDForAppStoreVersion(id string) (*AppStoreVersionBuildLinkageResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.GetWithQuery(url, nil, res)
	return res, resp, err
}

// UpdateBuildForAppStoreVersion changes the build that is attached to a specific App Store version.
func (s *Service) UpdateBuildForAppStoreVersion(id string, body *AppStoreVersionBuildLinkageRequest) (*AppStoreVersionBuildLinkageResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// GetAgeRatingDeclarationForAppStoreVersion gets the age-related information declared for your app.
func (s *Service) GetAgeRatingDeclarationForAppStoreVersion(id string, params *GetAgeRatingDeclarationForAppStoreVersionQuery) (*AgeRatingDeclarationResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/ageRatingDeclaration", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetRoutingAppCoverageForAppStoreVersion gets the routing app coverage file that is associated with a specific App Store version
func (s *Service) GetRoutingAppCoverageForAppStoreVersion(id string, params *GetRoutingAppCoverageForVersionQuery) (*RoutingAppCoverageResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/routingAppCoverage", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
