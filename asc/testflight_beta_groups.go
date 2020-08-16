package asc

import (
	"context"
	"fmt"
	"time"
)

// BetaGroup defines model for BetaGroup.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroup
type BetaGroup struct {
	Attributes    *BetaGroupAttributes    `json:"attributes,omitempty"`
	ID            string                  `json:"id"`
	Links         ResourceLinks           `json:"links"`
	Relationships *BetaGroupRelationships `json:"relationships,omitempty"`
	Type          string                  `json:"type"`
}

// BetaGroupAttributes defines model for BetaGroup.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroup/attributes
type BetaGroupAttributes struct {
	CreatedDate            *time.Time `json:"createdDate,omitempty"`
	FeedbackEnabled        *bool      `json:"feedbackEnabled,omitempty"`
	IsInternalGroup        *bool      `json:"isInternalGroup,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	PublicLink             *string    `json:"publicLink,omitempty"`
	PublicLinkEnabled      *bool      `json:"publicLinkEnabled,omitempty"`
	PublicLinkID           *string    `json:"publicLinkId,omitempty"`
	PublicLinkLimit        *int       `json:"publicLinkLimit,omitempty"`
	PublicLinkLimitEnabled *bool      `json:"publicLinkLimitEnabled,omitempty"`
}

// BetaGroupRelationships defines model for BetaGroup.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroup/relationships
type BetaGroupRelationships struct {
	App         *Relationship      `json:"app,omitempty"`
	BetaTesters *PagedRelationship `json:"betaTesters,omitempty"`
	Builds      *PagedRelationship `json:"builds,omitempty"`
}

// BetaGroupResponse defines model for BetaGroupResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupresponse
type BetaGroupResponse struct {
	Data     BetaGroup                   `json:"data"`
	Included []BetaGroupResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks               `json:"links"`
}

// BetaGroupResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a BetaGroupResponse or BetaGroupsResponse.
type BetaGroupResponseIncluded included

// BetaGroupCreateRequest defines model for BetaGroupCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupcreaterequest/data
type betaGroupCreateRequest struct {
	Attributes    BetaGroupCreateRequestAttributes    `json:"attributes"`
	Relationships betaGroupCreateRequestRelationships `json:"relationships"`
	Type          string                              `json:"type"`
}

// BetaGroupCreateRequestAttributes are attributes for BetaGroupCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupcreaterequest/data/attributes
type BetaGroupCreateRequestAttributes struct {
	FeedbackEnabled        *bool  `json:"feedbackEnabled,omitempty"`
	Name                   string `json:"name"`
	PublicLinkEnabled      *bool  `json:"publicLinkEnabled,omitempty"`
	PublicLinkLimit        *int   `json:"publicLinkLimit,omitempty"`
	PublicLinkLimitEnabled *bool  `json:"publicLinkLimitEnabled,omitempty"`
}

// BetaGroupCreateRequestRelationships are relationships for BetaGroupCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupcreaterequest/data/relationships
type betaGroupCreateRequestRelationships struct {
	App         relationshipDeclaration       `json:"app"`
	BetaTesters *pagedRelationshipDeclaration `json:"betaTesters,omitempty"`
	Builds      *pagedRelationshipDeclaration `json:"builds,omitempty"`
}

// BetaGroupUpdateRequest defines model for BetaGroupUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupupdaterequest/data
type betaGroupUpdateRequest struct {
	Attributes *BetaGroupUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                            `json:"id"`
	Type       string                            `json:"type"`
}

// BetaGroupUpdateRequestAttributes are attributes for BetaGroupUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupupdaterequest/data/attributes
type BetaGroupUpdateRequestAttributes struct {
	FeedbackEnabled        *bool   `json:"feedbackEnabled,omitempty"`
	Name                   *string `json:"name,omitempty"`
	PublicLinkEnabled      *bool   `json:"publicLinkEnabled,omitempty"`
	PublicLinkLimit        *int    `json:"publicLinkLimit,omitempty"`
	PublicLinkLimitEnabled *bool   `json:"publicLinkLimitEnabled,omitempty"`
}

// BetaGroupBetaTestersLinkagesResponse defines model for BetaGroupBetaTestersLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupbetatesterslinkagesresponse
type BetaGroupBetaTestersLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// BetaGroupBuildsLinkagesResponse defines model for BetaGroupBuildsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupbuildslinkagesresponse
type BetaGroupBuildsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// BetaGroupsResponse defines model for BetaGroupsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betagroupsresponse
type BetaGroupsResponse struct {
	Data     []BetaGroup                 `json:"data"`
	Included []BetaGroupResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks          `json:"links"`
	Meta     *PagingInformation          `json:"meta,omitempty"`
}

// ListBetaGroupsQuery defines model for ListBetaGroups
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_groups
type ListBetaGroupsQuery struct {
	FieldsApps                   []string `url:"fields[apps],omitempty"`
	FieldsBetaGroups             []string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters            []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                 []string `url:"fields[builds],omitempty"`
	FilterApp                    []string `url:"filter[app],omitempty"`
	FilterBuilds                 []string `url:"filter[builds],omitempty"`
	FilterID                     []string `url:"filter[id],omitempty"`
	FilterIsInternalGroup        []string `url:"filter[isInternalGroup],omitempty"`
	FilterName                   []string `url:"filter[name],omitempty"`
	FilterPublicLinkEnabled      []string `url:"filter[publicLinkEnabled],omitempty"`
	FilterPublicLinkLimitEnabled []string `url:"filter[publicLinkLimitEnabled],omitempty"`
	FilterPublicLink             []string `url:"filter[publicLink],omitempty"`
	Include                      []string `url:"include,omitempty"`
	Sort                         []string `url:"sort,omitempty"`
	Limit                        int      `url:"limit,omitempty"`
	LimitBuilds                  int      `url:"limit[builds],omitempty"`
	LimitBetaTesters             int      `url:"limit[betaTesters],omitempty"`
	Cursor                       string   `url:"cursor,omitempty"`
}

// GetBetaGroupQuery defines model for GetBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_group_information
type GetBetaGroupQuery struct {
	FieldsApps        []string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  []string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      []string `url:"fields[builds],omitempty"`
	Include           []string `url:"include,omitempty"`
	LimitBuilds       int      `url:"limit[builds],omitempty"`
	LimitBetaTesters  int      `url:"limit[betaTesters],omitempty"`
}

// GetAppForBetaGroupQuery defines model for GetAppForBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_group
type GetAppForBetaGroupQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
}

// ListBetaGroupsForAppQuery defines model for ListBetaGroupsForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_groups_for_an_app
type ListBetaGroupsForAppQuery struct {
	FieldsBetaGroups []string `url:"fields[betaGroups],omitempty"`
	Limit            int      `url:"limit,omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListBuildsForBetaGroupQuery defines model for ListBuildsForBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_for_a_betagroup
type ListBuildsForBetaGroupQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// ListBuildIDsForBetaGroupQuery defines model for ListBuildIDsForBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_build_ids_in_a_beta_group
type ListBuildIDsForBetaGroupQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListBetaTestersForBetaGroupQuery defines model for ListBetaTestersForBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_testers_in_a_betagroup
type ListBetaTestersForBetaGroupQuery struct {
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// ListBetaTesterIDsForBetaGroupQuery defines model for ListBetaTesterIDsForBetaGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_beta_tester_ids_in_a_beta_group
type ListBetaTesterIDsForBetaGroupQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// CreateBetaGroup creates a beta group associated with an app, optionally enabling TestFlight public links.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_beta_group
func (s *TestflightService) CreateBetaGroup(ctx context.Context, attributes BetaGroupCreateRequestAttributes, appID string, betaTesterIDs []string, buildIDs []string) (*BetaGroupResponse, *Response, error) {
	req := betaGroupCreateRequest{
		Attributes: attributes,
		Relationships: betaGroupCreateRequestRelationships{
			App: *newRelationshipDeclaration(&appID, "apps"),
		},
		Type: "betaGroups",
	}
	if len(betaTesterIDs) > 0 {
		relationship := newPagedRelationshipDeclaration(betaTesterIDs, "betaTesters")
		req.Relationships.BetaTesters = &relationship
	}
	if len(buildIDs) > 0 {
		relationship := newPagedRelationshipDeclaration(buildIDs, "betaTesters")
		req.Relationships.Builds = &relationship
	}
	res := new(BetaGroupResponse)
	resp, err := s.client.post(ctx, "betaGroups", req, res)
	return res, resp, err
}

// UpdateBetaGroup modifies a beta group's metadata, including changing its Testflight public link status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_beta_group
func (s *TestflightService) UpdateBetaGroup(ctx context.Context, id string, attributes *BetaGroupUpdateRequestAttributes) (*BetaGroupResponse, *Response, error) {
	req := betaGroupUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "betaGroups",
	}
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.client.patch(ctx, url, req, res)
	return res, resp, err
}

// DeleteBetaGroup deletes a beta group and remove beta tester access to associated builds.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_beta_group
func (s *TestflightService) DeleteBetaGroup(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	return s.client.delete(ctx, url, nil)
}

// ListBetaGroups finds and lists beta groups for all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_groups
func (s *TestflightService) ListBetaGroups(ctx context.Context, params *ListBetaGroupsQuery) (*BetaGroupsResponse, *Response, error) {
	res := new(BetaGroupsResponse)
	resp, err := s.client.get(ctx, "betaGroups", params, res)
	return res, resp, err
}

// GetBetaGroup gets a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_group_information
func (s *TestflightService) GetBetaGroup(ctx context.Context, id string, params *GetBetaGroupQuery) (*BetaGroupResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppForBetaGroup gets the app information for a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_information_of_a_beta_group
func (s *TestflightService) GetAppForBetaGroup(ctx context.Context, id string, params *GetAppForBetaGroupQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListBetaGroupsForApp gets a list of beta groups associated with a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_groups_for_an_app
func (s *TestflightService) ListBetaGroupsForApp(ctx context.Context, id string, params *ListBetaGroupsForAppQuery) (*BetaGroupsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// AddBetaTestersToBetaGroup adds a specific beta tester to one or more beta groups for beta testing.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_beta_testers_to_a_beta_group
func (s *TestflightService) AddBetaTestersToBetaGroup(ctx context.Context, id string, betaTesterIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(betaTesterIDs, "betaTesters")
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.client.post(ctx, url, linkages, nil)
}

// RemoveBetaTestersFromBetaGroup removes a specific beta tester from a one or more beta groups, revoking their access to test builds associated with those groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_beta_testers_from_a_beta_group
func (s *TestflightService) RemoveBetaTestersFromBetaGroup(ctx context.Context, id string, betaTesterIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(betaTesterIDs, "betaTesters")
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.client.delete(ctx, url, linkages)
}

// AddBuildsToBetaGroup associates builds with a beta group to enable the group to test the builds.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_builds_to_a_beta_group
func (s *TestflightService) AddBuildsToBetaGroup(ctx context.Context, id string, buildIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(buildIDs, "builds")
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.client.post(ctx, url, linkages, nil)
}

// RemoveBuildsFromBetaGroup removes access to test one or more builds from beta testers in a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_builds_from_a_beta_group
func (s *TestflightService) RemoveBuildsFromBetaGroup(ctx context.Context, id string, buildIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(buildIDs, "builds")
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.client.delete(ctx, url, linkages)
}

// ListBuildsForBetaGroup gets a list of builds associated with a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_for_a_betagroup
func (s *TestflightService) ListBuildsForBetaGroup(ctx context.Context, id string, params *ListBuildsForBetaGroupQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListBuildIDsForBetaGroup gets a list of build resource IDs in a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_build_ids_in_a_beta_group
func (s *TestflightService) ListBuildIDsForBetaGroup(ctx context.Context, id string, params *ListBuildIDsForBetaGroupQuery) (*BetaGroupBuildsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	res := new(BetaGroupBuildsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListBetaTestersForBetaGroup gets a list of beta testers contained in a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_testers_in_a_betagroup
func (s *TestflightService) ListBetaTestersForBetaGroup(ctx context.Context, id string, params *ListBetaTestersForBetaGroupQuery) (*BetaTestersResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s/betaTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListBetaTesterIDsForBetaGroup gets a list of the beta tester resource IDs in a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_beta_tester_ids_in_a_beta_group
func (s *TestflightService) ListBetaTesterIDsForBetaGroup(ctx context.Context, id string, params *ListBetaTesterIDsForBetaGroupQuery) (*BetaGroupBetaTestersLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	res := new(BetaGroupBetaTestersLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in BetaGroupResponseIncluded.
func (i *BetaGroupResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner
	return err
}

// App returns the App stored within, if one is present.
func (i *BetaGroupResponseIncluded) App() *App {
	return extractIncludedApp(i.inner)
}

// Build returns the Build stored within, if one is present.
func (i *BetaGroupResponseIncluded) Build() *Build {
	return extractIncludedBuild(i.inner)
}

// BetaTester returns the BetaTester stored within, if one is present.
func (i *BetaGroupResponseIncluded) BetaTester() *BetaTester {
	return extractIncludedBetaTester(i.inner)
}
