package asc

import (
	"fmt"
	"time"
)

// BuildsService handles communication with build-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/builds
// https://developer.apple.com/documentation/appstoreconnectapi/build_icons
// https://developer.apple.com/documentation/appstoreconnectapi/app_encryption_declarations
type BuildsService service

// Build defines model for Build.
type Build struct {
	Attributes *struct {
		ExpirationDate          *time.Time  `json:"expirationDate,omitempty"`
		Expired                 *bool       `json:"expired,omitempty"`
		IconAssetToken          *ImageAsset `json:"iconAssetToken,omitempty"`
		MinOsVersion            *string     `json:"minOsVersion,omitempty"`
		ProcessingState         *string     `json:"processingState,omitempty"`
		UploadedDate            *time.Time  `json:"uploadedDate,omitempty"`
		UsesNonExemptEncryption *bool       `json:"usesNonExemptEncryption,omitempty"`
		Version                 *string     `json:"version,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppEncryptionDeclaration *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appEncryptionDeclaration,omitempty"`
		AppStoreVersion *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
		BetaAppReviewSubmission *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"betaAppReviewSubmission,omitempty"`
		BetaBuildLocalizations *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"betaBuildLocalizations,omitempty"`
		BuildBetaDetail *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"buildBetaDetail,omitempty"`
		Icons *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"icons,omitempty"`
		IndividualTesters *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"individualTesters,omitempty"`
		PreReleaseVersion *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"preReleaseVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildResponse defines model for BuildResponse.
type BuildResponse struct {
	Data     Build          `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// BuildsResponse defines model for BuildsResponse.
type BuildsResponse struct {
	Data     []Build            `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
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
		Data *RelationshipsData `json:"data,omitempty"`
	} `json:"appEncryptionDeclaration,omitempty"`
}

// BuildAppEncryptionDeclarationLinkageResponse defines model for BuildAppEncryptionDeclarationLinkageResponse.
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data  RelationshipsData `json:"data"`
	Links DocumentLinks     `json:"links"`
}

// BuildIndividualTestersLinkagesResponse defines model for BuildIndividualTestersLinkagesResponse.
type BuildIndividualTestersLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// ListBuildsQuery are query options for ListBuilds
type ListBuildsQuery struct {
	FieldsAppEncryptionDeclarations          []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                               []string `url:"fields[apps],omitempty"`
	FieldsBetaTesters                        []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                             []string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions                 []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBuildBetaDetails                   []string `url:"fields[buildBetaDetails],omitempty"`
	FieldsBetaAppReviewSubmissions           []string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FieldsBetaBuildLocalizations             []string `url:"fields[betaBuildLocalizations],omitempty"`
	FieldsDiagnosticSignatures               []string `url:"fields[diagnosticSignatures],omitempty"`
	FieldsAppStoreVersions                   []string `url:"fields[appStoreVersions],omitempty"`
	FieldsPerfPowerMetrics                   []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsBuildIcons                         []string `url:"fields[buildIcons],omitempty"`
	FilterApp                                []string `url:"filter[app],omitempty"`
	FilterExpired                            []string `url:"filter[expired],omitempty"`
	FilterID                                 []string `url:"filter[id],omitempty"`
	FilterPreReleaseVersion                  []string `url:"filter[preReleaseVersion],omitempty"`
	FilterProcessingState                    []string `url:"filter[processingState],omitempty"`
	FilterVersion                            []string `url:"filter[version],omitempty"`
	FilterUsesNonExemptEncryption            []string `url:"filter[usesNonExemptEncryption],omitempty"`
	FilterPreReleaseVersionVersion           []string `url:"filter[preReleaseVersion.version],omitempty"`
	FilterPreReleaseVersionPlatform          []string `url:"filter[preReleaseVersion.platform],omitempty"`
	FilterBetaGroups                         []string `url:"filter[betaGroups],omitempty"`
	FilterBetaAppReviewSubmissionReviewState []string `url:"filter[betaAppReviewSubmission.betaReviewState],omitempty"`
	FilterAppStoreVersion                    []string `url:"filter[appStoreVersion],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	Sort                                     []string `url:"sort,omitempty"`
	Limit                                    int      `url:"limit,omitempty"`
	LimitIndividualTesters                   int      `url:"limit[individualTesters],omitempty"`
	LimitBetaBuildLocalizations              int      `url:"limit[betaBuildLocalizations],omitempty"`
	LimitIcons                               int      `url:"limit[icons],omitempty"`
	Cursor                                   string   `url:"cursor,omitempty"`
}

// ListBuildsForAppQuery are query options for ListBuildsForApp
type ListBuildsForAppQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// GetBuildsQuery are query options for GetBuilds
type GetBuildsQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FieldsBetaTesters               []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                    []string `url:"fields[builds],omitempty"`
	FieldsPreReleaseVersions        []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBuildBetaDetails          []string `url:"fields[buildBetaDetails],omitempty"`
	FieldsBetaAppReviewSubmissions  []string `url:"fields[betaAppReviewSubmissions],omitempty"`
	FieldsBetaBuildLocalizations    []string `url:"fields[betaBuildLocalizations],omitempty"`
	FieldsDiagnosticSignatures      []string `url:"fields[diagnosticSignatures],omitempty"`
	FieldsAppStoreVersions          []string `url:"fields[appStoreVersions],omitempty"`
	FieldsPerfPowerMetrics          []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsBuildIcons                []string `url:"fields[buildIcons],omitempty"`
	Include                         []string `url:"include,omitempty"`
	LimitIndividualTesters          int      `url:"limit[individualTesters],omitempty"`
	LimitBetaBuildLocalizations     int      `url:"limit[betaBuildLocalizations],omitempty"`
	LimitIcons                      int      `url:"limit[icons],omitempty"`
}

// GetAppForBuildQuery are query options for GetAppForBuild
type GetAppForBuildQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// GetAppStoreVersionForBuildQuery are query options for GetAppStoreVersionForBuild
type GetAppStoreVersionForBuildQuery struct {
	FieldsAppStoreVersions []string `url:"fields[appStoreVersions],omitempty"`
}

// GetBuildForAppStoreVersionQuery are query options for GetBuildForAppStoreVersion
type GetBuildForAppStoreVersionQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
}

// ListResourceIDsForIndividualTestersForBuildQuery are query options for ListResourceIDsForIndividualTestersForBuild
type ListResourceIDsForIndividualTestersForBuildQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationForBuildQuery are query options for GetAppEncryptionDeclarationForBuild
type GetAppEncryptionDeclarationForBuildQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
}

// ListBuilds finds and lists builds for all apps in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_builds
func (s *BuildsService) ListBuilds(params *ListBuildsQuery) (*BuildsResponse, *Response, error) {
	res := new(BuildsResponse)
	resp, err := s.client.get("builds", params, res)
	return res, resp, err
}

// ListBuildsForApp gets a list of builds associated with a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_an_app
func (s *BuildsService) ListBuildsForApp(id string, params *ListBuildsForAppQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBuild gets information about a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_build_information
func (s *BuildsService) GetBuild(id string, params *GetBuildsQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppForBuild gets the app information for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_build
func (s *BuildsService) GetAppForBuild(id string, params *GetAppForBuildQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppStoreVersionForBuild gets the App Store version of a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_information_of_a_build
func (s *BuildsService) GetAppStoreVersionForBuild(id string, params *GetAppStoreVersionForBuildQuery) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/appStoreVersion", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetBuildForAppStoreVersion gets the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_an_app_store_version
func (s *BuildsService) GetBuildForAppStoreVersion(id string, params *GetBuildForAppStoreVersionQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// UpdateBuild expires a build or changes its encryption exemption setting.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_build
func (s *BuildsService) UpdateBuild(id string, body BuildUpdateRequest) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.client.post(url, body, res)
	return res, resp, err
}

// UpdateAppEncryptionDeclarationForBuild assigns an app encryption declaration to a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_the_app_encryption_declaration_for_a_build
func (s *BuildsService) UpdateAppEncryptionDeclarationForBuild(id string, linkage *RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	return s.client.patch(url, linkage, nil)
}

// CreateAccessForBetaGroupsToBuild adds or creates a beta group to a build to enable testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_access_for_beta_groups_to_a_build
func (s *BuildsService) CreateAccessForBetaGroupsToBuild(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.client.post(url, linkages, nil)
}

// RemoveAccessForBetaGroupsFromBuild removes access to a specific build for all beta testers in one or more beta groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_access_for_beta_groups_to_a_build
func (s *BuildsService) RemoveAccessForBetaGroupsFromBuild(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.client.delete(url, linkages)
}

// CreateAccessForIndividualTestersToBuild enables a beta tester who is not a part of a beta group to test a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_individual_testers_to_a_build
func (s *BuildsService) CreateAccessForIndividualTestersToBuild(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.client.post(url, linkages, nil)
}

// RemoveAccessForIndividualTestersFromBuild removes access to test a specific build from one or more individually assigned testers.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_individual_testers_from_a_build
func (s *BuildsService) RemoveAccessForIndividualTestersFromBuild(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.client.delete(url, linkages)
}

// ListResourceIDsForIndividualTestersForBuild gets a list of resource IDs of individual testers associated with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_resource_ids_of_individual_testers_for_a_build
func (s *BuildsService) ListResourceIDsForIndividualTestersForBuild(id string, params *ListResourceIDsForIndividualTestersForBuildQuery) (*BuildIndividualTestersLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	res := new(BuildIndividualTestersLinkagesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationForBuild reads an app encryption declaration associated with a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_encryption_declaration_of_a_build
func (s *BuildsService) GetAppEncryptionDeclarationForBuild(id string, params *GetAppEncryptionDeclarationForBuildQuery) (*AppEncryptionDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/appEncryptionDeclaration", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationIDForBuild gets the beta app encryption declaration resource ID associated with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_the_app_encryption_declaration_id_for_a_build
func (s *BuildsService) GetAppEncryptionDeclarationIDForBuild(id string) (*BuildAppEncryptionDeclarationLinkageResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	res := new(BuildAppEncryptionDeclarationLinkageResponse)
	resp, err := s.client.get(url, nil, res)
	return res, resp, err
}
