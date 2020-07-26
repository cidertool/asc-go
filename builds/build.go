package builds

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal/services"
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
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppEncryptionDeclaration *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appEncryptionDeclaration,omitempty"`
		AppStoreVersion *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
		BetaAppReviewSubmission *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"betaAppReviewSubmission,omitempty"`
		BetaBuildLocalizations *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"betaBuildLocalizations,omitempty"`
		BuildBetaDetail *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"buildBetaDetail,omitempty"`
		Icons *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"icons,omitempty"`
		IndividualTesters *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"individualTesters,omitempty"`
		PreReleaseVersion *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"preReleaseVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildResponse defines model for BuildResponse.
type BuildResponse struct {
	Data     Build                  `json:"data"`
	Included *[]interface{}         `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// BuildsResponse defines model for BuildsResponse.
type BuildsResponse struct {
	Data     []Build                     `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// BuildUpdateRequest defines model for BuildUpdateRequest.
type BuildUpdateRequest struct {
	Attributes    *BuildUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                           `json:"id"`
	Relationships *BuildUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                           `json:"type"`
}

// BuildUpdateRequestAttributes are attributes for BuildUpdateRequest
type BuildUpdateRequestAttributes struct {
	Expired                 *bool `json:"expired,omitempty"`
	UsesNonExemptEncryption *bool `json:"usesNonExemptEncryption,omitempty"`
}

// BuildUpdateRequestRelationships are relationships for BuildUpdateRequest
type BuildUpdateRequestRelationships struct {
	AppEncryptionDeclaration *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"appEncryptionDeclaration,omitempty"`
}

// BuildAppEncryptionDeclarationLinkageResponse defines model for BuildAppEncryptionDeclarationLinkageResponse.
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data  services.RelationshipsData `json:"data"`
	Links services.DocumentLinks     `json:"links"`
}

// BuildIndividualTestersLinkagesResponse defines model for BuildIndividualTestersLinkagesResponse.
type BuildIndividualTestersLinkagesResponse struct {
	Data  []services.RelationshipsData `json:"data"`
	Links services.PagedDocumentLinks  `json:"links"`
	Meta  *services.PagingInformation  `json:"meta,omitempty"`
}

// ListBuildsQuery are query options for ListBuilds
type ListBuildsQuery struct {
	FieldsAppEncryptionDeclarations          *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                               *[]string `url:"fields[apps],omitempty"`
	FieldsBetaTesters                        *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                             *[]string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions                 *[]string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBuildBetaDetails                   *[]string `url:"fields[buildBetaDetails],omitempty"`
	FieldsBetaAppReviewSubmissions           *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FieldsBetaBuildLocalizations             *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	FieldsDiagnosticSignatures               *[]string `url:"fields[diagnosticSignatures],omitempty"`
	FieldsAppStoreVersions                   *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsPerfPowerMetrics                   *[]string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsBuildIcons                         *[]string `url:"fields[buildIcons],omitempty"`
	FilterApp                                *[]string `url:"filter[app],omitempty"`
	FilterExpired                            *[]string `url:"filter[expired],omitempty"`
	FilterID                                 *[]string `url:"filter[id],omitempty"`
	FilterPreReleaseVersion                  *[]string `url:"filter[preReleaseVersion],omitempty"`
	FilterProcessingState                    *[]string `url:"filter[processingState],omitempty"`
	FilterVersion                            *[]string `url:"filter[version],omitempty"`
	FilterUsesNonExemptEncryption            *[]string `url:"filter[usesNonExemptEncryption],omitempty"`
	FilterPreReleaseVersionVersion           *[]string `url:"filter[preReleaseVersion.version],omitempty"`
	FilterPreReleaseVersionPlatform          *[]string `url:"filter[preReleaseVersion.platform],omitempty"`
	FilterBetaGroups                         *[]string `url:"filter[betaGroups],omitempty"`
	FilterBetaAppReviewSubmissionReviewState *[]string `url:"filter[betaAppReviewSubmission.betaReviewState],omitempty"`
	FilterAppStoreVersion                    *[]string `url:"filter[appStoreVersion],omitempty"`
	Include                                  *[]string `url:"include,omitempty"`
	Sort                                     *[]string `url:"sort,omitempty"`
	Limit                                    *int      `url:"limit,omitempty"`
	LimitIndividualTesters                   *int      `url:"limit[individualTesters],omitempty"`
	LimitBetaBuildLocalizations              *int      `url:"limit[betaBuildLocalizations],omitempty"`
	LimitIcons                               *int      `url:"limit[icons],omitempty"`
	Cursor                                   *string   `url:"cursor,omitempty"`
}

// ListBuildsForAppQuery are query options for ListBuildsForApp
type ListBuildsForAppQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
	Limit        *int      `url:"limit,omitempty"`
	Cursor       *string   `url:"cursor,omitempty"`
}

// GetBuildsQuery are query options for GetBuilds
type GetBuildsQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FieldsBetaTesters               *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                    *[]string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions        *[]string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBuildBetaDetails          *[]string `url:"fields[buildBetaDetails],omitempty"`
	FieldsBetaAppReviewSubmissions  *[]string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FieldsBetaBuildLocalizations    *[]string `url:"fields[betaBuildLocalizations],omitempty"`
	FieldsDiagnosticSignatures      *[]string `url:"fields[diagnosticSignatures],omitempty"`
	FieldsAppStoreVersions          *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsPerfPowerMetrics          *[]string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsBuildIcons                *[]string `url:"fields[buildIcons],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	LimitIndividualTesters          *int      `url:"limit[individualTesters],omitempty"`
	LimitBetaBuildLocalizations     *int      `url:"limit[betaBuildLocalizations],omitempty"`
	LimitIcons                      *int      `url:"limit[icons],omitempty"`
}

// GetAppForBuildQuery are query options for GetAppForBuild
type GetAppForBuildQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// GetAppStoreVersionForBuildQuery are query options for GetAppStoreVersionForBuild
type GetAppStoreVersionForBuildQuery struct {
	FieldsAppStoreVersions *[]string `url:"fields[appStoreVersions],omitempty"`
}

// GetBuildForAppStoreVersionQuery are query options for GetBuildForAppStoreVersion
type GetBuildForAppStoreVersionQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
}

// ListResourceIDsForIndividualTestersForBuildQuery are query options for ListResourceIDsForIndividualTestersForBuild
type ListResourceIDsForIndividualTestersForBuildQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationForBuildQuery are query options for GetAppEncryptionDeclarationForBuild
type GetAppEncryptionDeclarationForBuildQuery struct {
	FieldsAppEncryptionDeclarations *[]string `url:"fields[appEncryptionDeclarations],omitempty"`
}

// ListBuilds finds and lists builds for all apps in App Store Connect.
func (s *Service) ListBuilds(params *ListBuildsQuery) (*BuildsResponse, *http.Response, error) {
	res := new(BuildsResponse)
	resp, err := s.GetWithQuery("builds", params, res)
	return res, resp, err
}

// ListBuildsForApp gets a list of builds associated with a specific app.
func (s *Service) ListBuildsForApp(id string, params *ListBuildsForAppQuery) (*BuildsResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuild gets information about a specific build.
func (s *Service) GetBuild(id string, params *GetBuildsQuery) (*BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBuild gets the app information for a specific build.
func (s *Service) GetAppForBuild(id string, params *GetAppForBuildQuery) (*apps.AppResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppStoreVersionForBuild gets the App Store version of a specific build.
func (s *Service) GetAppStoreVersionForBuild(id string, params *GetAppStoreVersionForBuildQuery) (*apps.AppStoreVersionResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/appStoreVersion", id)
	res := new(apps.AppStoreVersionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForAppStoreVersion gets the build that is attached to a specific App Store version.
func (s *Service) GetBuildForAppStoreVersion(id string, params *GetBuildForAppStoreVersionQuery) (*BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBuild expires a build or changes its encryption exemption setting.
func (s *Service) UpdateBuild(id string, body BuildUpdateRequest) (*BuildResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.Post(url, body, res)
	return res, resp, err
}

// UpdateAppEncryptionDeclarationForBuild assigns an app encryption declaration to a build.
func (s *Service) UpdateAppEncryptionDeclarationForBuild(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	return s.Patch(url, linkages, nil)
}

// CreateAccessForBetaGroupsToBuild adds or create a beta group to a build to enable testing.
func (s *Service) CreateAccessForBetaGroupsToBuild(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.Post(url, linkages, nil)
}

// RemoveAccessForBetaGroupsFromBuild removes access to a specific build for all beta testers in one or more beta groups.
func (s *Service) RemoveAccessForBetaGroupsFromBuild(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.Delete(url, linkages)
}

// CreateAccessForIndividualTestersToBuild enables a beta tester who is not a part of a beta group to test a build.
func (s *Service) CreateAccessForIndividualTestersToBuild(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.Post(url, linkages, nil)
}

// RemoveAccessForIndividualTestersFromBuild removes access to test a specific build from one or more individually assigned testers.
func (s *Service) RemoveAccessForIndividualTestersFromBuild(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.Delete(url, linkages)
}

// ListResourceIDsForIndividualTestersForBuild gets a list of resource IDs of individual testers associated with a build.
func (s *Service) ListResourceIDsForIndividualTestersForBuild(id string, params *ListResourceIDsForIndividualTestersForBuildQuery) (*BuildIndividualTestersLinkagesResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	res := new(BuildIndividualTestersLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationForBuild reads an app encryption declaration associated with a specific build.
func (s *Service) GetAppEncryptionDeclarationForBuild(id string, params *GetAppEncryptionDeclarationForBuildQuery) (*AppEncryptionDeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/appEncryptionDeclaration", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationIDForBuild gets the beta app encryption declaration resource ID associated with a build.
func (s *Service) GetAppEncryptionDeclarationIDForBuild(id string) (*BuildAppEncryptionDeclarationLinkageResponse, *http.Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	res := new(BuildAppEncryptionDeclarationLinkageResponse)
	resp, err := s.GetWithQuery(url, nil, res)
	return res, resp, err
}
