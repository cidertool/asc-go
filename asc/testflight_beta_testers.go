package asc

import (
	"fmt"
)

// BetaInviteType defines model for BetaInviteType.
type BetaInviteType string

// List of BetaInviteType
const (
	BetaInviteTypeEmail      BetaInviteType = "EMAIL"
	BetaInviteTypePublicLink BetaInviteType = "PUBLIC_LINK"
)

// BetaTester defines model for BetaTester.
type BetaTester struct {
	Attributes *struct {
		Email      *Email          `json:"email,omitempty"`
		FirstName  *string         `json:"firstName,omitempty"`
		InviteType *BetaInviteType `json:"inviteType,omitempty"`
		LastName   *string         `json:"lastName,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Apps *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"apps,omitempty"`
		BetaGroups *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"betaGroups,omitempty"`
		Builds *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaTesterAppsLinkagesResponse defines model for BetaTesterAppsLinkagesResponse.
type BetaTesterAppsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// BetaTesterBetaGroupsLinkagesResponse defines model for BetaTesterBetaGroupsLinkagesResponse.
type BetaTesterBetaGroupsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// BetaTesterBuildsLinkagesResponse defines model for BetaTesterBuildsLinkagesResponse.
type BetaTesterBuildsLinkagesResponse struct {
	Data  []RelationshipsData `json:"data"`
	Links PagedDocumentLinks  `json:"links"`
	Meta  *PagingInformation  `json:"meta,omitempty"`
}

// BetaTesterCreateRequest defines model for BetaTesterCreateRequest.
type BetaTesterCreateRequest struct {
	Attributes    BetaTesterCreateRequestAttributes     `json:"attributes"`
	Relationships *BetaTesterCreateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                `json:"type"`
}

// BetaTesterCreateRequestAttributes are attributes for BetaTesterCreateRequest
type BetaTesterCreateRequestAttributes struct {
	Email     Email   `json:"email"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// BetaTesterCreateRequestRelationships are relationships for BetaTesterCreateRequest
type BetaTesterCreateRequestRelationships struct {
	BetaGroups *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"betaGroups,omitempty"`
	Builds *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"builds,omitempty"`
}

// BetaTesterResponse defines model for BetaTesterResponse.
type BetaTesterResponse struct {
	Data     BetaTester     `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// BetaTestersResponse defines model for BetaTestersResponse.
type BetaTestersResponse struct {
	Data     []BetaTester       `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// ListBetaTestersQuery defines model for ListBetaTesters
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
type ListAppsForBetaTesterQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
	Limit      int      `url:"limit,omitempty"`
	Cursor     string   `url:"cursor,omitempty"`
}

// ListAppIDsForBetaTesterQuery defines model for ListAppIDsForBetaTester
type ListAppIDsForBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListBuildsIndividuallyAssignedToBetaTesterQuery defines model for ListBuildsIndividuallyAssignedToBetaTester
type ListBuildsIndividuallyAssignedToBetaTesterQuery struct {
	FieldsBuilds []string `url:"fields[builds],omitempty"`
	Limit        int      `url:"limit,omitempty"`
	Cursor       string   `url:"cursor,omitempty"`
}

// ListBuildIDsIndividuallyAssignedToBetaTesterQuery defines model for ListBuildIDsIndividuallyAssignedToBetaTester
type ListBuildIDsIndividuallyAssignedToBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListIndividualTestersForBuildQuery defines model for ListIndividualTestersForBuild
type ListIndividualTestersForBuildQuery struct {
	FieldsBetaTesters []string `url:"fields[betaTesters],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// ListBetaGroupsForBetaTesterQuery defines model for ListBetaGroupsForBetaTester
type ListBetaGroupsForBetaTesterQuery struct {
	FieldsBetaGroups []string `url:"fields[betaGroups],omitempty"`
	Limit            int      `url:"limit,omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListBetaGroupIDsForBetaTesterQuery defines model for ListBetaGroupIDsForBetaTester
type ListBetaGroupIDsForBetaTesterQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// CreateBetaTester creates a beta tester assigned to a group, a build, or an app.
func (s *TestflightService) CreateBetaTester(body *BetaTesterCreateRequest) (*BetaTesterResponse, *Response, error) {
	res := new(BetaTesterResponse)
	resp, err := s.client.post("betaTesters", body, res)
	return res, resp, err
}

// DeleteBetaTester removes a beta tester's ability to test all apps.
func (s *TestflightService) DeleteBetaTester(id string) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	return s.client.delete(url, nil)
}

// ListBetaTesters finds and lists beta testers for all apps, builds, and beta groups.
func (s *TestflightService) ListBetaTesters(id string, params *ListBetaTestersQuery) (*BetaTestersResponse, *Response, error) {
	res := new(BetaTestersResponse)
	resp, err := s.client.get("betaTesters", params, res)
	return res, resp, err
}

// GetBetaTester gets a specific beta tester.
func (s *TestflightService) GetBetaTester(id string, params *GetBetaTesterQuery) (*BetaTesterResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	res := new(BetaTesterResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// AddBetaTesterToBetaGroups adds one or more beta testers to a specific beta group.
func (s *TestflightService) AddBetaTesterToBetaGroups(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	return s.client.post(url, linkages, nil)
}

// RemoveBetaTesterFromBetaGroups removes a specific beta tester from one or more beta groups, revoking their access to test builds associated with those groups.
func (s *TestflightService) RemoveBetaTesterFromBetaGroups(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	return s.client.delete(url, linkages)
}

// AssignSingleBetaTesterToBuilds individually assign a beta tester to a build.
func (s *TestflightService) AssignSingleBetaTesterToBuilds(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	return s.client.post(url, linkages, nil)
}

// UnassignSingleBetaTesterFromBuilds removes an individually assigned beta tester's ability to test a build.
func (s *TestflightService) UnassignSingleBetaTesterFromBuilds(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	return s.client.delete(url, linkages)
}

// RemoveSingleBetaTesterAccessApps removes a specific beta tester's access to test any builds of one or more apps.
func (s *TestflightService) RemoveSingleBetaTesterAccessApps(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	return s.client.delete(url, linkages)
}

// ListAppsForBetaTester gets a list of apps that a beta tester can test.
func (s *TestflightService) ListAppsForBetaTester(id string, params *ListAppsForBetaTesterQuery) (*AppsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/apps", id)
	res := new(AppsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListAppIDsForBetaTester gets a list of app resource IDs associated with a beta tester.
func (s *TestflightService) ListAppIDsForBetaTester(id string, params *ListAppIDsForBetaTesterQuery) (*BetaTesterAppsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	res := new(BetaTesterAppsLinkagesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBuildsIndividuallyAssignedToBetaTester gets a list of builds individually assigned to a specific beta tester.
func (s *TestflightService) ListBuildsIndividuallyAssignedToBetaTester(id string, params *ListBuildsIndividuallyAssignedToBetaTesterQuery) (*BuildsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/builds", id)
	res := new(BuildsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBuildIDsIndividuallyAssignedToBetaTester gets a list of build resource IDs individually assigned to a specific beta tester.
func (s *TestflightService) ListBuildIDsIndividuallyAssignedToBetaTester(id string, params *ListBuildIDsIndividuallyAssignedToBetaTesterQuery) (*BetaTesterBuildsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	res := new(BetaTesterBuildsLinkagesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListIndividualTestersForBuild gets a list of beta testers individually assigned to a build.
func (s *TestflightService) ListIndividualTestersForBuild(id string, params *ListIndividualTestersForBuildQuery) (*BetaTestersResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/individualTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBetaGroupsForBetaTester gets a list of beta groups that contain a specific beta tester.
func (s *TestflightService) ListBetaGroupsForBetaTester(id string, params *ListBetaGroupsForBetaTesterQuery) (*BetaGroupsResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// ListBetaGroupIDsForBetaTester gets a list of group resource IDs associated with a beta tester.
func (s *TestflightService) ListBetaGroupIDsForBetaTester(id string, params *ListBetaGroupIDsForBetaTesterQuery) (*BetaTesterBetaGroupsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	res := new(BetaTesterBetaGroupsLinkagesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
