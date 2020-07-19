package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaInviteType defines model for BetaInviteType.
type BetaInviteType string

// List of BetaInviteType
const (
	Email      BetaInviteType = "EMAIL"
	PublicLink BetaInviteType = "PUBLIC_LINK"
)

// BetaTester defines model for BetaTester.
type BetaTester struct {
	Attributes *struct {
		Email      *common.Email   `json:"email,omitempty"`
		FirstName  *string         `json:"firstName,omitempty"`
		InviteType *BetaInviteType `json:"inviteType,omitempty"`
		LastName   *string         `json:"lastName,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		Apps *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"apps,omitempty"`
		BetaGroups *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"betaGroups,omitempty"`
		Builds *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaTesterAppsLinkagesRequest defines model for BetaTesterAppsLinkagesRequest.
type BetaTesterAppsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterAppsLinkagesResponse defines model for BetaTesterAppsLinkagesResponse.
type BetaTesterAppsLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BetaTesterBetaGroupsLinkagesRequest defines model for BetaTesterBetaGroupsLinkagesRequest.
type BetaTesterBetaGroupsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterBetaGroupsLinkagesResponse defines model for BetaTesterBetaGroupsLinkagesResponse.
type BetaTesterBetaGroupsLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BetaTesterBuildsLinkagesRequest defines model for BetaTesterBuildsLinkagesRequest.
type BetaTesterBuildsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterBuildsLinkagesResponse defines model for BetaTesterBuildsLinkagesResponse.
type BetaTesterBuildsLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BetaTesterCreateRequest defines model for BetaTesterCreateRequest.
type BetaTesterCreateRequest struct {
	Data struct {
		Attributes struct {
			Email     common.Email `json:"email"`
			FirstName *string      `json:"firstName,omitempty"`
			LastName  *string      `json:"lastName,omitempty"`
		} `json:"attributes"`
		Relationships *struct {
			BetaGroups *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"betaGroups,omitempty"`
			Builds *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"builds,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaTesterResponse defines model for BetaTesterResponse.
type BetaTesterResponse struct {
	Data     BetaTester           `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BetaTestersResponse defines model for BetaTestersResponse.
type BetaTestersResponse struct {
	Data     []BetaTester              `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaTestersQuery struct {
	FieldsApps        *[]string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  *[]string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      *[]string `url:"fields[builds],omitempty"`
	FilterApps        *[]string `url:"filter[apps],omitempty"`
	FilterBetaGroups  *[]string `url:"filter[betaGroups],omitempty"`
	FilterBuilds      *[]string `url:"filter[builds],omitempty"`
	FilterEmail       *[]string `url:"filter[email],omitempty"`
	FilterFirstName   *[]string `url:"filter[firstName],omitempty"`
	FilterInviteType  *[]string `url:"filter[inviteType],omitempty"`
	FilterLastName    *[]string `url:"filter[lastName],omitempty"`
	Include           *[]string `url:"include,omitempty"`
	Sort              *[]string `url:"sort,omitempty"`
	Limit             *int      `url:"limit,omitempty"`
	LimitApps         *[]string `url:"limit[apps],omitempty"`
	LimitBetaGroups   *[]string `url:"limit[betaGroups],omitempty"`
	LimitBuilds       *[]string `url:"limit[builds],omitempty"`
	Cursor            *string   `url:"cursor,omitempty"`
}

type GetBetaTesterQuery struct {
	FieldsApps        *[]string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  *[]string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      *[]string `url:"fields[builds],omitempty"`
	Include           *[]string `url:"include,omitempty"`
	LimitApps         *[]string `url:"limit[apps],omitempty"`
	LimitBetaGroups   *[]string `url:"limit[betaGroups],omitempty"`
	LimitBuilds       *[]string `url:"limit[builds],omitempty"`
}

type ListAppsForBetaTesterQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
	Limit      *int      `url:"limit,omitempty"`
	Cursor     *string   `url:"cursor,omitempty"`
}

type ListAppIDsForBetaTesterQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListBuildsIndividuallyAssignedToBetaTesterQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
	Limit        *int      `url:"limit,omitempty"`
	Cursor       *string   `url:"cursor,omitempty"`
}

type ListBuildIDsIndividuallyAssignedToBetaTesterQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

type ListIndividualTestersForBuildQuery struct {
	FieldsBetaTesters *[]string `url:"fields[betaTesters],omitempty"`
	Limit             *int      `url:"limit,omitempty"`
	Cursor            *string   `url:"cursor,omitempty"`
}

type ListBetaGroupsForBetaTesterQuery struct {
	FieldsBetaGroups *[]string `url:"fields[betaGroups],omitempty"`
	Limit            *int      `url:"limit,omitempty"`
	Cursor           *string   `url:"cursor,omitempty"`
}

type ListBetaGroupIDsForBetaTesterQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// CreateBetaTester creates a beta tester assigned to a group, a build, or an app.
func (s *Service) CreateBetaTester(body *BetaTesterCreateRequest) (*BetaTesterResponse, *common.Response, error) {
	res := new(BetaTesterResponse)
	resp, err := s.Post("betaTesters", body, res)
	return res, resp, err
}

// DeleteBetaTester removes a beta tester's ability to test all apps.
func (s *Service) DeleteBetaTester(id string) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	return s.Delete(url, nil)
}

// ListBetaTesters finds and lists beta testers for all apps, builds, and beta groups.
func (s *Service) ListBetaTesters(id string, params *ListBetaTestersQuery) (*BetaTestersResponse, *common.Response, error) {
	res := new(BetaTestersResponse)
	resp, err := s.GetWithQuery("betaTesters", params, res)
	return res, resp, err
}

// GetBetaTester gets a specific beta tester.
func (s *Service) GetBetaTester(id string, params *GetBetaTesterQuery) (*BetaTesterResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	res := new(BetaTesterResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// AddBetaTesterToBetaGroups adds one or more beta testers to a specific beta group.
func (s *Service) AddBetaTesterToBetaGroups(id string, body *BetaTesterBetaGroupsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	return s.Post(url, body, nil)
}

// RemoveBetaTesterFromBetaGroups removes a specific beta tester from one or more beta groups, revoking their access to test builds associated with those groups.
func (s *Service) RemoveBetaTesterFromBetaGroups(id string, body *BetaTesterBetaGroupsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	return s.Delete(url, body)
}

// AssignSingleBetaTesterToBuilds individually assign a beta tester to a build.
func (s *Service) AssignSingleBetaTesterToBuilds(id string, body *BetaTesterBuildsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	return s.Post(url, body, nil)
}

// UnassignSingleBetaTesterFromBuilds removes an individually assigned beta tester's ability to test a build.
func (s *Service) UnassignSingleBetaTesterFromBuilds(id string, body *BetaTesterBuildsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	return s.Delete(url, body)
}

// RemoveSingleBetaTesterAccessApps removes a specific beta tester's access to test any builds of one or more apps.
func (s *Service) RemoveSingleBetaTesterAccessApps(id string, body *BetaTesterAppsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	return s.Delete(url, body)
}

// ListAppsForBetaTester gets a list of apps that a beta tester can test.
func (s *Service) ListAppsForBetaTester(id string, params *ListAppsForBetaTesterQuery) (*apps.AppsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/apps", id)
	res := new(apps.AppsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppIDsForBetaTester gets a list of app resource IDs associated with a beta tester.
func (s *Service) ListAppIDsForBetaTester(id string, params *ListAppIDsForBetaTesterQuery) (*BetaTesterAppsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	res := new(BetaTesterAppsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBuildsIndividuallyAssignedToBetaTester gets a list of builds individually assigned to a specific beta tester.
func (s *Service) ListBuildsIndividuallyAssignedToBetaTester(id string, params *ListBuildsIndividuallyAssignedToBetaTesterQuery) (*builds.BuildsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBuildIDsIndividuallyAssignedToBetaTester gets a list of build resource IDs individually assigned to a specific beta tester.
func (s *Service) ListBuildIDsIndividuallyAssignedToBetaTester(id string, params *ListBuildIDsIndividuallyAssignedToBetaTesterQuery) (*BetaTesterBuildsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	res := new(BetaTesterBuildsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListIndividualTestersForBuild gets a list of beta testers individually assigned to a build.
func (s *Service) ListIndividualTestersForBuild(id string, params *ListIndividualTestersForBuildQuery) (*BetaTestersResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/individualTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaGroupsForBetaTester gets a list of beta groups that contain a specific beta tester.
func (s *Service) ListBetaGroupsForBetaTester(id string, params *ListBetaGroupsForBetaTesterQuery) (*BetaGroupsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaGroupIDsForBetaTester gets a list of group resource IDs associated with a beta tester.
func (s *Service) ListBetaGroupIDsForBetaTester(id string, params *ListBetaGroupIDsForBetaTesterQuery) (*BetaTesterBetaGroupsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	res := new(BetaTesterBetaGroupsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
