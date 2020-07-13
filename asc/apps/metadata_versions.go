package apps

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/common"
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
			UsesIdfa            *bool      `json:"usesIdfa,omitempty"`
			VersionString       *string    `json:"versionString,omitempty"`
		} `json:"attributes,omitempty"`
		ID            string `json:"id"`
		Relationships *struct {
			Build *struct {
				Data *struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
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
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
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
		UsesIdfa            *bool                 `json:"usesIdfa,omitempty"`
		VersionString       *string               `json:"versionString,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AgeRatingDeclaration *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"ageRatingDeclaration,omitempty"`
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
		AppStoreReviewDetail *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreReviewDetail,omitempty"`
		AppStoreVersionLocalizations *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"appStoreVersionLocalizations,omitempty"`
		AppStoreVersionPhasedRelease *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersionPhasedRelease,omitempty"`
		AppStoreVersionSubmission *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersionSubmission,omitempty"`
		Build *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"build,omitempty"`
		IdfaDeclaration *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"idfaDeclaration,omitempty"`
		RoutingAppCoverage *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"routingAppCoverage,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionResponse defines model for AppStoreVersionResponse.
type AppStoreVersionResponse struct {
	Data     AppStoreVersion      `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// AppStoreVersionsResponse defines model for AppStoreVersionsResponse.
type AppStoreVersionsResponse struct {
	Data     []AppStoreVersion         `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// AppStoreVersionCreateRequest defines model for AppStoreVersionCreateRequest.
type AppStoreVersionCreateRequest struct {
	Data struct {
		Attributes struct {
			Copyright           *string    `json:"copyright,omitempty"`
			EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
			Platform            Platform   `json:"platform"`
			ReleaseType         *string    `json:"releaseType,omitempty"`
			UsesIdfa            *bool      `json:"usesIdfa,omitempty"`
			VersionString       string     `json:"versionString"`
		} `json:"attributes"`
		Relationships struct {
			App struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"app"`
			Build *struct {
				Data *struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"build,omitempty"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionBuildLinkageRequest defines model for AppStoreVersionBuildLinkageRequest.
type AppStoreVersionBuildLinkageRequest struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionBuildLinkageResponse defines model for AppStoreVersionBuildLinkageResponse.
type AppStoreVersionBuildLinkageResponse struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

type ListAppStoreVersionsQuery struct {
	Fields *struct {
		Apps                          *[]string `url:"apps,omitempty"`
		AppStoreVersionSubmissions    *[]string `url:"appStoreVersionSubmissions,omitempty"`
		Builds                        *[]string `url:"builds,omitempty"`
		AppStoreVersions              *[]string `url:"appStoreVersions,omitempty"`
		AppStoreReviewDetails         *[]string `url:"appStoreReviewDetails,omitempty"`
		AgeRatingDeclarations         *[]string `url:"ageRatingDeclarations,omitempty"`
		AppStoreVersionPhasedReleases *[]string `url:"appStoreVersionPhasedReleases,omitempty"`
		RoutingAppCoverages           *[]string `url:"routingAppCoverages,omitempty"`
		IdfaDeclarations              *[]string `url:"idfaDeclarations,omitempty"`
	} `url:"fields,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Filter  *struct {
		ID            *[]string `url:"id,omitempty"`
		VersionString *[]string `url:"versionString,omitempty"`
		Platform      *[]string `url:"platform,omitempty"`
		AppStoreState *[]string `url:"appStoreState,omitempty"`
	} `url:"filter,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type GetAppStoreVersionQuery struct {
	Fields *struct {
		AppStoreVersions              *[]string `url:"appStoreVersions,omitempty"`
		AppStoreVersionSubmissions    *[]string `url:"appStoreVersionSubmissions,omitempty"`
		Builds                        *[]string `url:"builds,omitempty"`
		AppStoreReviewDetails         *[]string `url:"appStoreReviewDetails,omitempty"`
		AgeRatingDeclarations         *[]string `url:"ageRatingDeclarations,omitempty"`
		AppStoreVersionPhasedReleases *[]string `url:"appStoreVersionPhasedReleases,omitempty"`
		RoutingAppCoverages           *[]string `url:"routingAppCoverages,omitempty"`
		IdfaDeclarations              *[]string `url:"idfaDeclarations,omitempty"`
		AppStoreVersionLocalizations  *[]string `url:"appStoreVersionLocalizations,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		AppStoreVersionLocalizations *int `url:"appStoreVersionLocalizations,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAgeRatingDeclarationForAppStoreVersionQuery struct {
	Fields *struct {
		AgeRatingDeclarations *[]string `url:"ageRatingDeclarations,omitempty"`
	} `url:"fields,omitempty"`
}

type GetRoutingAppCoverageForVersionQuery struct {
	Fields *struct {
		RoutingAppCoverages *[]string `url:"routingAppCoverages,omitempty"`
	} `url:"fields,omitempty"`
}

// ListAppStoreVersionsForApp gets a list of all App Store versions of an app across all platforms.
func (s *Service) ListAppStoreVersionsForApp(id string, params *ListAppStoreVersionsQuery) (*AppStoreVersionsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/appStoreVersions", id)
	res := new(AppStoreVersionsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppStoreVersion gets information for a specific app store version.
func (s *Service) GetAppStoreVersion(id string, params *GetAppStoreVersionQuery) (*AppStoreVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppStoreVersion adds a new App Store version or platform to an app.
func (s *Service) CreateAppStoreVersion(body *AppStoreVersionCreateRequest) (*AppStoreVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions")
	res := new(AppStoreVersionResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateAppStoreVersion updates the app store version for a specific app.
func (s *Service) UpdateAppStoreVersion(id string, body *AppStoreVersionUpdateRequest) (*AppStoreVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppStoreVersion deletes an app store version that is associated with an app.
func (s *Service) DeleteAppStoreVersion(id string) (*common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	return s.Delete(url, nil)
}

// GetBuildIDForAppStoreVersion gets the ID of the build that is attached to a specific App Store version.
func (s *Service) GetBuildIDForAppStoreVersion(id string) (*AppStoreVersionBuildLinkageResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.GetWithQuery(url, nil, res)
	return res, resp, err
}

// UpdateBuildForAppStoreVersion changes the build that is attached to a specific App Store version.
func (s *Service) UpdateBuildForAppStoreVersion(id string, body *AppStoreVersionBuildLinkageRequest) (*AppStoreVersionBuildLinkageResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// GetAgeRatingDeclarationForAppStoreVersion gets the age-related information declared for your app.
func (s *Service) GetAgeRatingDeclarationForAppStoreVersion(id string, params *GetAgeRatingDeclarationForAppStoreVersionQuery) (*AgeRatingDeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/ageRatingDeclaration", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetRoutingAppCoverageForAppStoreVersion gets the routing app coverage file that is associated with a specific App Store version
func (s *Service) GetRoutingAppCoverageForAppStoreVersion(id string, params *GetRoutingAppCoverageForVersionQuery) (*RoutingAppCoverageResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/routingAppCoverage", id)
	res := new(RoutingAppCoverageResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
