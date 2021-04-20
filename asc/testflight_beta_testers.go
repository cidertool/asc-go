/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// BetaInviteType defines model for BetaInviteType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betainvitetype
type BetaInviteType string

const (
	// BetaInviteTypeEmail is a beta invite type for Email.
	BetaInviteTypeEmail BetaInviteType = "EMAIL"
	// BetaInviteTypePublicLink is a beta invite type for PublicLink.
	BetaInviteTypePublicLink BetaInviteType = "PUBLIC_LINK"
)

// BetaTester defines model for BetaTester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatester
type BetaTester struct {
	Attributes    *BetaTesterAttributes    `json:"attributes,omitempty"`
	ID            string                   `json:"id"`
	Links         ResourceLinks            `json:"links"`
	Relationships *BetaTesterRelationships `json:"relationships,omitempty"`
	Type          string                   `json:"type"`
}

// BetaTesterAttributes defines model for BetaTester.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatester/attributes
type BetaTesterAttributes struct {
	Email      *Email          `json:"email,omitempty"`
	FirstName  *string         `json:"firstName,omitempty"`
	InviteType *BetaInviteType `json:"inviteType,omitempty"`
	LastName   *string         `json:"lastName,omitempty"`
}

// BetaTesterRelationships defines model for BetaTester.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatester/relationships
type BetaTesterRelationships struct {
	Apps       *PagedRelationship `json:"apps,omitempty"`
	BetaGroups *PagedRelationship `json:"betaGroups,omitempty"`
	Builds     *PagedRelationship `json:"builds,omitempty"`
}

// BetaTesterAppsLinkagesResponse defines model for BetaTesterAppsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterappslinkagesresponse
type BetaTesterAppsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// BetaTesterBetaGroupsLinkagesResponse defines model for BetaTesterBetaGroupsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterbetagroupslinkagesresponse
type BetaTesterBetaGroupsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// BetaTesterBuildsLinkagesResponse defines model for BetaTesterBuildsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterbuildslinkagesresponse
type BetaTesterBuildsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// betaTesterCreateRequest defines model for betaTesterCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatestercreaterequest/data
type betaTesterCreateRequest struct {
	Attributes    BetaTesterCreateRequestAttributes     `json:"attributes"`
	Relationships *betaTesterCreateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                `json:"type"`
}

// BetaTesterCreateRequestAttributes are attributes for BetaTesterCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatestercreaterequest/data/attributes
type BetaTesterCreateRequestAttributes struct {
	Email     Email   `json:"email"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// betaTesterCreateRequestRelationships are relationships for BetaTesterCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatestercreaterequest/data/relationships
type betaTesterCreateRequestRelationships struct {
	BetaGroups *pagedRelationshipDeclaration `json:"betaGroups,omitempty"`
	Builds     *pagedRelationshipDeclaration `json:"builds,omitempty"`
}

// BetaTesterResponse defines model for BetaTesterResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatesterresponse
type BetaTesterResponse struct {
	Data     BetaTester                   `json:"data"`
	Included []BetaTesterResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                `json:"links"`
}

// BetaTestersResponse defines model for BetaTestersResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/betatestersresponse
type BetaTestersResponse struct {
	Data     []BetaTester                 `json:"data"`
	Included []BetaTesterResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks           `json:"links"`
	Meta     *PagingInformation           `json:"meta,omitempty"`
}

// BetaTesterResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a BetaTesterResponse or BetaTestersResponse.
type BetaTesterResponseIncluded included

// ListBetaTestersQuery defines model for ListBetaTesters
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_testers
type ListBetaTestersQuery struct {
	FieldsApps        []string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  []string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      []string `url:"fields[builds],omitempty"`
	FilterApps        []string `url:"filter[apps],omitempty"`
	FilterBetaGroups  []string `url:"filter[betaGroups],omitempty"`
	FilterBuilds      []string `url:"filter[builds],omitempty"`
	FilterEmail       []string `url:"filter[email],omitempty"`
	FilterFirstName   []string `url:"filter[firstName],omitempty"`
	FilterInviteType  []string `url:"filter[inviteType],omitempty"`
	FilterLastName    []string `url:"filter[lastName],omitempty"`
	Include           []string `url:"include,omitempty"`
	Sort              []string `url:"sort,omitempty"`
	Limit             int      `url:"limit,omitempty"`
	LimitApps         []string `url:"limit[apps],omitempty"`
	LimitBetaGroups   []string `url:"limit[betaGroups],omitempty"`
	LimitBuilds       []string `url:"limit[builds],omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// GetBetaTesterQuery defines model for GetBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_tester_information
type GetBetaTesterQuery struct {
	FieldsApps        []string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  []string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      []string `url:"fields[builds],omitempty"`
	Include           []string `url:"include,omitempty"`
	LimitApps         []string `url:"limit[apps],omitempty"`
	LimitBetaGroups   []string `url:"limit[betaGroups],omitempty"`
	LimitBuilds       []string `url:"limit[builds],omitempty"`
}

// ListAppsForBetaTesterQuery defines model for ListAppsForBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_apps_for_a_beta_tester
type ListAppsForBetaTesterQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
	Limit      int      `url:"limit,omitempty"`
	Cursor     string   `url:"cursor,omitempty"`
}

// ListAppIDsForBetaTesterQuery defines model for ListAppIDsForBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_resource_ids_for_a_beta_tester
type ListAppIDsForBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListBuildsIndividuallyAssignedToBetaTesterQuery defines model for ListBuildsIndividuallyAssignedToBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_individually_assigned_to_a_beta_tester
type ListBuildsIndividuallyAssignedToBetaTesterQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// ListBuildIDsIndividuallyAssignedToBetaTesterQuery defines model for ListBuildIDsIndividuallyAssignedToBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_ids_of_builds_individually_assigned_to_a_beta_tester
type ListBuildIDsIndividuallyAssignedToBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListIndividualTestersForBuildQuery defines model for ListIndividualTestersForBuild
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_individual_testers_for_a_build
type ListIndividualTestersForBuildQuery struct {
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// ListBetaGroupsForBetaTesterQuery defines model for ListBetaGroupsForBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_groups_to_which_a_beta_tester_belongs
type ListBetaGroupsForBetaTesterQuery struct {
	FieldsBetaGroups []string `url:"fields[betaGroups],omitempty"`
	Limit            int      `url:"limit,omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListBetaGroupIDsForBetaTesterQuery defines model for ListBetaGroupIDsForBetaTester
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_beta_group_ids_of_a_beta_tester_s_groups
type ListBetaGroupIDsForBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// CreateBetaTester creates a beta tester assigned to a group, a build, or an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_beta_tester
func (s *TestflightService) CreateBetaTester(ctx context.Context, attributes BetaTesterCreateRequestAttributes, betaGroupIDs []string, buildIDs []string) (*BetaTesterResponse, *Response, error) {
	req := betaTesterCreateRequest{
		Attributes: attributes,
		Type:       "betaTesters",
	}

	anyBetaGroups := len(betaGroupIDs) > 0
	anyBuilds := len(buildIDs) > 0 // nolint: ifshort

	if anyBetaGroups || anyBuilds {
		req.Relationships = &betaTesterCreateRequestRelationships{}

		if anyBetaGroups {
			relationships := newPagedRelationshipDeclaration(betaGroupIDs, "betaGroups")
			req.Relationships.BetaGroups = &relationships
		}

		if anyBuilds {
			relationships := newPagedRelationshipDeclaration(buildIDs, "builds")
			req.Relationships.Builds = &relationships
		}
	}

	res := new(BetaTesterResponse)
	resp, err := s.client.post(ctx, "betaTesters", newRequestBody(req), res)

	return res, resp, err
}

// DeleteBetaTester removes a beta tester's ability to test all apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_beta_tester
func (s *TestflightService) DeleteBetaTester(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListBetaTesters finds and lists beta testers for all apps, builds, and beta groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_beta_testers
func (s *TestflightService) ListBetaTesters(ctx context.Context, params *ListBetaTestersQuery) (*BetaTestersResponse, *Response, error) {
	res := new(BetaTestersResponse)
	resp, err := s.client.get(ctx, "betaTesters", params, res)

	return res, resp, err
}

// GetBetaTester gets a specific beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_beta_tester_information
func (s *TestflightService) GetBetaTester(ctx context.Context, id string, params *GetBetaTesterQuery) (*BetaTesterResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	res := new(BetaTesterResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// AddBetaTesterToBetaGroups adds one or more beta testers to a specific beta group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_a_beta_tester_to_beta_groups
func (s *TestflightService) AddBetaTesterToBetaGroups(ctx context.Context, id string, betaGroupIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(betaGroupIDs, "betaGroups")
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)

	return s.client.post(ctx, url, newRequestBody(linkages.Data), nil)
}

// RemoveBetaTesterFromBetaGroups removes a specific beta tester from one or more beta groups, revoking their access to test builds associated with those groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_a_beta_tester_from_beta_groups
func (s *TestflightService) RemoveBetaTesterFromBetaGroups(ctx context.Context, id string, betaGroupIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(betaGroupIDs, "betaGroups")
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)

	return s.client.delete(ctx, url, newRequestBody(linkages.Data))
}

// AssignSingleBetaTesterToBuilds individually assign a beta tester to a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/individually_assign_a_beta_tester_to_builds
func (s *TestflightService) AssignSingleBetaTesterToBuilds(ctx context.Context, id string, buildIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(buildIDs, "builds")
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)

	return s.client.post(ctx, url, newRequestBody(linkages.Data), nil)
}

// UnassignSingleBetaTesterFromBuilds removes an individually assigned beta tester's ability to test a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/individually_unassign_a_beta_tester_from_builds
func (s *TestflightService) UnassignSingleBetaTesterFromBuilds(ctx context.Context, id string, buildIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(buildIDs, "builds")
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)

	return s.client.delete(ctx, url, newRequestBody(linkages.Data))
}

// RemoveSingleBetaTesterAccessApps removes a specific beta tester's access to test any builds of one or more apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_a_beta_tester_s_access_to_apps
func (s *TestflightService) RemoveSingleBetaTesterAccessApps(ctx context.Context, id string, appIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appIDs, "apps")
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)

	return s.client.delete(ctx, url, newRequestBody(linkages.Data))
}

// ListAppsForBetaTester gets a list of apps that a beta tester can test.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_apps_for_a_beta_tester
func (s *TestflightService) ListAppsForBetaTester(ctx context.Context, id string, params *ListAppsForBetaTesterQuery) (*AppsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/apps", id)
	res := new(AppsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListAppIDsForBetaTester gets a list of app resource IDs associated with a beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_app_resource_ids_for_a_beta_tester
func (s *TestflightService) ListAppIDsForBetaTester(ctx context.Context, id string, params *ListAppIDsForBetaTesterQuery) (*BetaTesterAppsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	res := new(BetaTesterAppsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBuildsIndividuallyAssignedToBetaTester gets a list of builds individually assigned to a specific beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_builds_individually_assigned_to_a_beta_tester
func (s *TestflightService) ListBuildsIndividuallyAssignedToBetaTester(ctx context.Context, id string, params *ListBuildsIndividuallyAssignedToBetaTesterQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBuildIDsIndividuallyAssignedToBetaTester gets a list of build resource IDs individually assigned to a specific beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_ids_of_builds_individually_assigned_to_a_beta_tester
func (s *TestflightService) ListBuildIDsIndividuallyAssignedToBetaTester(ctx context.Context, id string, params *ListBuildIDsIndividuallyAssignedToBetaTesterQuery) (*BetaTesterBuildsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	res := new(BetaTesterBuildsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListIndividualTestersForBuild gets a list of beta testers individually assigned to a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_individual_testers_for_a_build
func (s *TestflightService) ListIndividualTestersForBuild(ctx context.Context, id string, params *ListIndividualTestersForBuildQuery) (*BetaTestersResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/individualTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBetaGroupsForBetaTester gets a list of beta groups that contain a specific beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_beta_groups_to_which_a_beta_tester_belongs
func (s *TestflightService) ListBetaGroupsForBetaTester(ctx context.Context, id string, params *ListBetaGroupsForBetaTesterQuery) (*BetaGroupsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListBetaGroupIDsForBetaTester gets a list of group resource IDs associated with a beta tester.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_beta_group_ids_of_a_beta_tester_s_groups
func (s *TestflightService) ListBetaGroupIDsForBetaTester(ctx context.Context, id string, params *ListBetaGroupIDsForBetaTesterQuery) (*BetaTesterBetaGroupsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	res := new(BetaTesterBetaGroupsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in BetaTesterResponseIncluded.
func (i *BetaTesterResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// App returns the App stored within, if one is present.
func (i *BetaTesterResponseIncluded) App() *App {
	return extractIncludedApp(i.inner)
}

// BetaGroup returns the BetaGroup stored within, if one is present.
func (i *BetaTesterResponseIncluded) BetaGroup() *BetaGroup {
	return extractIncludedBetaGroup(i.inner)
}

// Build returns the Build stored within, if one is present.
func (i *BetaTesterResponseIncluded) Build() *Build {
	return extractIncludedBuild(i.inner)
}
