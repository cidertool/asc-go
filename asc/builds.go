package asc

import (
	"context"
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/build
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
		App                      *Relationship      `json:"app,omitempty"`
		AppEncryptionDeclaration *Relationship      `json:"appEncryptionDeclaration,omitempty"`
		AppStoreVersion          *Relationship      `json:"appStoreVersion,omitempty"`
		BetaAppReviewSubmission  *Relationship      `json:"betaAppReviewSubmission,omitempty"`
		BetaBuildLocalizations   *PagedRelationship `json:"betaBuildLocalizations,omitempty"`
		BuildBetaDetail          *Relationship      `json:"buildBetaDetail,omitempty"`
		Icons                    *PagedRelationship `json:"icons,omitempty"`
		IndividualTesters        *PagedRelationship `json:"individualTesters,omitempty"`
		PreReleaseVersion        *Relationship      `json:"preReleaseVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildResponse defines model for BuildResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildresponse
type BuildResponse struct {
	Data     Build          `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// BuildsResponse defines model for BuildsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildsresponse
type BuildsResponse struct {
	Data     []Build            `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// BuildUpdateRequest defines model for BuildUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildupdaterequest
type BuildUpdateRequest struct {
	Attributes    *BuildUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                           `json:"id"`
	Relationships *BuildUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                           `json:"type"`
}

// BuildUpdateRequestAttributes are attributes for BuildUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildupdaterequest/data/attributes
type BuildUpdateRequestAttributes struct {
	Expired                 *bool `json:"expired,omitempty"`
	UsesNonExemptEncryption *bool `json:"usesNonExemptEncryption,omitempty"`
}

// BuildUpdateRequestRelationships are relationships for BuildUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildupdaterequest/data/relationships
type BuildUpdateRequestRelationships struct {
	AppEncryptionDeclaration *RelationshipDeclaration `json:"appEncryptionDeclaration,omitempty"`
}

// BuildAppEncryptionDeclarationLinkageResponse defines model for BuildAppEncryptionDeclarationLinkageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildappencryptiondeclarationlinkageresponse
type BuildAppEncryptionDeclarationLinkageResponse struct {
	Data  RelationshipData `json:"data"`
	Links DocumentLinks    `json:"links"`
}

// BuildIndividualTestersLinkagesResponse defines model for BuildIndividualTestersLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildindividualtesterslinkagesresponse
type BuildIndividualTestersLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListBuildsQuery are query options for ListBuilds
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_builds
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_an_app
type ListBuildsForAppQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// GetBuildQuery are query options for GetBuilds
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_build_information
type GetBuildQuery struct {
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_build
type GetAppForBuildQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// GetAppStoreVersionForBuildQuery are query options for GetAppStoreVersionForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_information_of_a_build
type GetAppStoreVersionForBuildQuery struct {
	FieldsAppStoreVersions []string `url:"fields[appStoreVersions],omitempty"`
}

// GetBuildForAppStoreVersionQuery are query options for GetBuildForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_an_app_store_version
type GetBuildForAppStoreVersionQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
}

// ListResourceIDsForIndividualTestersForBuildQuery are query options for ListResourceIDsForIndividualTestersForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_resource_ids_of_individual_testers_for_a_build
type ListResourceIDsForIndividualTestersForBuildQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetAppEncryptionDeclarationForBuildQuery are query options for GetAppEncryptionDeclarationForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_encryption_declaration_of_a_build
type GetAppEncryptionDeclarationForBuildQuery struct {
	FieldsAppEncryptionDeclarations []string `url:"fields[appEncryptionDeclarations],omitempty"`
}

// ListBuilds finds and lists builds for all apps in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_builds
func (s *BuildsService) ListBuilds(ctx context.Context, params *ListBuildsQuery) (*BuildsResponse, *Response, error) {
	res := new(BuildsResponse)
	resp, err := s.client.get(ctx, "builds", params, res)
	return res, resp, err
}

// ListBuildsForApp gets a list of builds associated with a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_of_an_app
func (s *BuildsService) ListBuildsForApp(ctx context.Context, id string, params *ListBuildsForAppQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBuild gets information about a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_build_information
func (s *BuildsService) GetBuild(ctx context.Context, id string, params *GetBuildQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppForBuild gets the app information for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_build
func (s *BuildsService) GetAppForBuild(ctx context.Context, id string, params *GetAppForBuildQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppStoreVersionForBuild gets the App Store version of a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_information_of_a_build
func (s *BuildsService) GetAppStoreVersionForBuild(ctx context.Context, id string, params *GetAppStoreVersionForBuildQuery) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/appStoreVersion", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetBuildForAppStoreVersion gets the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_build_information_of_an_app_store_version
func (s *BuildsService) GetBuildForAppStoreVersion(ctx context.Context, id string, params *GetBuildForAppStoreVersionQuery) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/build", id)
	res := new(BuildResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateBuild expires a build or changes its encryption exemption setting.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_build
func (s *BuildsService) UpdateBuild(ctx context.Context, id string, body *BuildUpdateRequest) (*BuildResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s", id)
	res := new(BuildResponse)
	resp, err := s.client.post(ctx, url, body, res)
	return res, resp, err
}

// UpdateAppEncryptionDeclarationForBuild assigns an app encryption declaration to a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_the_app_encryption_declaration_for_a_build
func (s *BuildsService) UpdateAppEncryptionDeclarationForBuild(ctx context.Context, id string, linkage *RelationshipData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	return s.client.patch(ctx, url, linkage, nil)
}

// CreateAccessForBetaGroupsToBuild adds or creates a beta group to a build to enable testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_access_for_beta_groups_to_a_build
func (s *BuildsService) CreateAccessForBetaGroupsToBuild(ctx context.Context, id string, linkages *[]RelationshipData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.client.post(ctx, url, linkages, nil)
}

// RemoveAccessForBetaGroupsFromBuild removes access to a specific build for all beta testers in one or more beta groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_access_for_beta_groups_to_a_build
func (s *BuildsService) RemoveAccessForBetaGroupsFromBuild(ctx context.Context, id string, linkages *[]RelationshipData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/betaGroups", id)
	return s.client.delete(ctx, url, linkages)
}

// CreateAccessForIndividualTestersToBuild enables a beta tester who is not a part of a beta group to test a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/assign_individual_testers_to_a_build
func (s *BuildsService) CreateAccessForIndividualTestersToBuild(ctx context.Context, id string, linkages *[]RelationshipData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.client.post(ctx, url, linkages, nil)
}

// RemoveAccessForIndividualTestersFromBuild removes access to test a specific build from one or more individually assigned testers.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_individual_testers_from_a_build
func (s *BuildsService) RemoveAccessForIndividualTestersFromBuild(ctx context.Context, id string, linkages *[]RelationshipData) (*Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	return s.client.delete(ctx, url, linkages)
}

// ListResourceIDsForIndividualTestersForBuild gets a list of resource IDs of individual testers associated with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_resource_ids_of_individual_testers_for_a_build
func (s *BuildsService) ListResourceIDsForIndividualTestersForBuild(ctx context.Context, id string, params *ListResourceIDsForIndividualTestersForBuildQuery) (*BuildIndividualTestersLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/individualTesters", id)
	res := new(BuildIndividualTestersLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationForBuild reads an app encryption declaration associated with a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_encryption_declaration_of_a_build
func (s *BuildsService) GetAppEncryptionDeclarationForBuild(ctx context.Context, id string, params *GetAppEncryptionDeclarationForBuildQuery) (*AppEncryptionDeclarationResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/appEncryptionDeclaration", id)
	res := new(AppEncryptionDeclarationResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppEncryptionDeclarationIDForBuild gets the beta app encryption declaration resource ID associated with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_the_app_encryption_declaration_id_for_a_build
func (s *BuildsService) GetAppEncryptionDeclarationIDForBuild(ctx context.Context, id string) (*BuildAppEncryptionDeclarationLinkageResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/relationships/appEncryptionDeclaration", id)
	res := new(BuildAppEncryptionDeclarationLinkageResponse)
	resp, err := s.client.get(ctx, url, nil, res)
	return res, resp, err
}
