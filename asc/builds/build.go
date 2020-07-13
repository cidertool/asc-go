package builds

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// Build defines model for Build.
type Build struct {
	Attributes *struct {
		ExpirationDate          *time.Time       `json:"expirationDate,omitempty"`
		Expired                 *bool            `json:"expired,omitempty"`
		IconAssetToken          *apps.ImageAsset `json:"iconAssetToken,omitempty"`
		MinOsVersion            *string          `json:"minOsVersion,omitempty"`
		ProcessingState         *string          `json:"processingState,omitempty"`
		UploadedDate            *time.Time       `json:"uploadedDate,omitempty"`
		UsesNonExemptEncryption *bool            `json:"usesNonExemptEncryption,omitempty"`
		Version                 *string          `json:"version,omitempty"`
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
		AppEncryptionDeclaration *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appEncryptionDeclaration,omitempty"`
		AppStoreVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
		BetaAppReviewSubmission *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"betaAppReviewSubmission,omitempty"`
		BetaBuildLocalizations *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"betaBuildLocalizations,omitempty"`
		BuildBetaDetail *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"buildBetaDetail,omitempty"`
		Icons *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"icons,omitempty"`
		IndividualTesters *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"individualTesters,omitempty"`
		PreReleaseVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"preReleaseVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildResponse defines model for BuildResponse.
type BuildResponse struct {
	Data     Build                `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BuildsResponse defines model for BuildsResponse.
type BuildsResponse struct {
	Data     []Build                   `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// BuildUpdateRequest defines model for BuildUpdateRequest.
type BuildUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Expired                 *bool `json:"expired,omitempty"`
			UsesNonExemptEncryption *bool `json:"usesNonExemptEncryption,omitempty"`
		} `json:"attributes,omitempty"`
		ID            string `json:"id"`
		Relationships *struct {
			AppEncryptionDeclaration *struct {
				Data *struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"appEncryptionDeclaration,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

// BuildAppEncryptionDeclarationLinkageRequest defines model for BuildAppEncryptionDeclarationLinkageRequest.
type BuildAppEncryptionDeclarationLinkageRequest struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BuildAppEncryptionDeclarationLinkageResponse defines model for BuildAppEncryptionDeclarationLinkageResponse.
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// BuildIndividualTestersLinkagesRequest defines model for BuildIndividualTestersLinkagesRequest.
type BuildIndividualTestersLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BuildIndividualTestersLinkagesResponse defines model for BuildIndividualTestersLinkagesResponse.
type BuildIndividualTestersLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BuildBetaGroupsLinkagesRequest defines model for BuildBetaGroupsLinkagesRequest.
type BuildBetaGroupsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

type ListBuildsOptions struct {
	Fields *struct {
		AppEncryptionDeclarations *[]string `url:"appEncryptionDeclarations,omitempty"`
		Apps                      *[]string `url:"apps,omitempty"`
		BetaTesters               *[]string `url:"betaTesters,omitempty"`
		Builds                    *[]string `url:"builds,omitempty"`
		PreReleaseVersions        *[]string `url:"preReleaseVersions,omitempty"`
		BuildBetaDetails          *[]string `url:"buildBetaDetails,omitempty"`
		BetaAppReviewSubmissions  *[]string `url:"betaAppReviewSubmissions,omitempty"`
		BetaBuildLocalizations    *[]string `url:"betaBuildLocalizations,omitempty"`
		DiagnosticSignatures      *[]string `url:"diagnosticSignatures,omitempty"`
		AppStoreVersions          *[]string `url:"appStoreVersions,omitempty"`
		PerfPowerMetrics          *[]string `url:"perfPowerMetrics,omitempty"`
		BuildIcons                *[]string `url:"buildIcons,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App                                *[]string `url:"app,omitempty"`
		Expired                            *[]string `url:"expired,omitempty"`
		ID                                 *[]string `url:"id,omitempty"`
		PreReleaseVersion                  *[]string `url:"preReleaseVersion,omitempty"`
		ProcessingState                    *[]string `url:"processingState,omitempty"`
		Version                            *[]string `url:"version,omitempty"`
		UsesNonExemptEncryption            *[]string `url:"usesNonExemptEncryption,omitempty"`
		PreReleaseVersionVersion           *[]string `url:"preReleaseVersion.version,omitempty"`
		PreReleaseVersionPlatform          *[]string `url:"preReleaseVersion.platform,omitempty"`
		BetaGroups                         *[]string `url:"betaGroups,omitempty"`
		BetaAppReviewSubmissionReviewState *[]string `url:"betaAppReviewSubmission.betaReviewState,omitempty"`
		AppStoreVersion                    *[]string `url:"appStoreVersion,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	Sort     *[]string `url:"sort,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		IndividualTesters      *int `url:"individualTesters,omitempty"`
		BetaBuildLocalizations *int `url:"betaBuildLocalizations,omitempty"`
		Icons                  *int `url:"icons,omitempty"`
	} `url:"limit,omitempty"`
}

type ListBuildsForAppOptions struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type GetBuildsOptions struct {
	Fields *struct {
		AppEncryptionDeclarations *[]string `url:"appEncryptionDeclarations,omitempty"`
		Apps                      *[]string `url:"apps,omitempty"`
		BetaTesters               *[]string `url:"betaTesters,omitempty"`
		Builds                    *[]string `url:"builds,omitempty"`
		PreReleaseVersions        *[]string `url:"preReleaseVersions,omitempty"`
		BuildBetaDetails          *[]string `url:"buildBetaDetails,omitempty"`
		BetaAppReviewSubmissions  *[]string `url:"betaAppReviewSubmissions,omitempty"`
		BetaBuildLocalizations    *[]string `url:"betaBuildLocalizations,omitempty"`
		DiagnosticSignatures      *[]string `url:"diagnosticSignatures,omitempty"`
		AppStoreVersions          *[]string `url:"appStoreVersions,omitempty"`
		PerfPowerMetrics          *[]string `url:"perfPowerMetrics,omitempty"`
		BuildIcons                *[]string `url:"buildIcons,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		IndividualTesters      *int `url:"individualTesters,omitempty"`
		BetaBuildLocalizations *int `url:"betaBuildLocalizations,omitempty"`
		Icons                  *int `url:"icons,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAppForBuildOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type GetAppStoreVersionForBuildOptions struct {
	Fields *struct {
		AppStoreVersions *[]string `url:"appStoreVersions,omitempty"`
	} `url:"fields,omitempty"`
}

type GetBuildForAppStoreVersionOptions struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
}

type ListResourceIDsForIndividualTestersForBuildOptions struct {
	Limit *int `url:"limit,omitempty"`
}

type GetAppEncryptionDeclarationForBuildOptions struct {
	Fields *struct {
		AppEncryptionDeclarations *[]string `url:"appEncryptionDeclarations,omitempty"`
	} `url:"fields,omitempty"`
}

// ListBuilds finds and lists builds for all apps in App Store Connect.
func (s *Service) ListBuilds(params *ListBuildsOptions) (*BuildsResponse, *common.Response, error) {
	res := new(BuildsResponse)
	resp, err := s.Get("builds", params, res)
	return res, resp, err
}

// ListBuildsForApp gets a list of builds associated with a specific app.
func (s *Service) ListBuildsForApp(id string, params *ListBuildsForAppOptions) (*BuildsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetBuild gets information about a specific build.
func (s *Service) GetBuild(id string, params *GetBuildsOptions) (*BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForBuild gets the app information for a specific build.
func (s *Service) GetAppForBuild(id string, params *GetAppForBuildOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppStoreVersionForBuild gets the App Store version of a specific build.
func (s *Service) GetAppStoreVersionForBuild(id string, params *GetAppStoreVersionForBuildOptions) (*apps.AppStoreVersionResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/appStoreVersion", id)
	res := new(apps.AppStoreVersionResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetBuildForAppStoreVersion gets the build that is attached to a specific App Store version.
func (s *Service) GetBuildForAppStoreVersion(id string, params *GetBuildForAppStoreVersionOptions) (*BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// UpdateBuild expires a build or changes its encryption exemption setting.
func (s *Service) UpdateBuild(id string, body BuildUpdateRequest) (*BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateAppEncryptionDeclarationForBuild assigns an app encryption declaration to a build.
func (s *Service) UpdateAppEncryptionDeclarationForBuild(id string, body *BuildAppEncryptionDeclarationLinkageRequest) (*common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	return s.Patch(url, body, nil)
}

// CreateAccessForBetaGroupsToBuild adds or create a beta group to a build to enable testing.
func (s *Service) CreateAccessForBetaGroupsToBuild(id string, body *BuildBetaGroupsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.Post(url, body, nil)
}

// RemoveAccessForBetaGroupsFromBuild removes access to a specific build for all beta testers in one or more beta groups.
func (s *Service) RemoveAccessForBetaGroupsFromBuild(id string, body *BuildBetaGroupsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.Delete(url, body)
}

// CreateAccessForIndividualTestersToBuild enables a beta tester who is not a part of a beta group to test a build.
func (s *Service) CreateAccessForIndividualTestersToBuild(id string, body *BuildIndividualTestersLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.Post(url, body, nil)
}

// RemoveAccessForIndividualTestersFromBuild removes access to test a specific build from one or more individually assigned testers.
func (s *Service) RemoveAccessForIndividualTestersFromBuild(id string, body *BuildIndividualTestersLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.Delete(url, body)
}

// ListResourceIDsForIndividualTestersForBuild gets a list of resource IDs of individual testers associated with a build.
func (s *Service) ListResourceIDsForIndividualTestersForBuild(id string, params *ListResourceIDsForIndividualTestersForBuildOptions) (*BuildIndividualTestersLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	res := new(BuildIndividualTestersLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationForBuild reads an app encryption declaration associated with a specific build.
func (s *Service) GetAppEncryptionDeclarationForBuild(id string, params *GetAppEncryptionDeclarationForBuildOptions) (*AppEncryptionDeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/appEncryptionDeclaration", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationIDForBuild gets the beta app encryption declaration resource ID associated with a build.
func (s *Service) GetAppEncryptionDeclarationIDForBuild(id string) (*BuildAppEncryptionDeclarationLinkageResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	res := new(BuildAppEncryptionDeclarationLinkageResponse)
	resp, err := s.Get(url, nil, res)
	return res, resp, err
}
