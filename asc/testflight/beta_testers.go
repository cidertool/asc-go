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

type ListBetaTestersOptions struct {
	Fields *struct {
		Apps        *[]string `url:"apps,omitempty"`
		BetaGroups  *[]string `url:"betaGroups,omitempty"`
		BetaTesters *[]string `url:"betaTesters,omitempty"`
		Builds      *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		Apps       *[]string `url:"apps,omitempty"`
		BetaGroups *[]string `url:"betaGroups,omitempty"`
		Builds     *[]string `url:"builds,omitempty"`
		Email      *[]string `url:"email,omitempty"`
		FirstName  *[]string `url:"firstName,omitempty"`
		InviteType *[]string `url:"inviteType,omitempty"`
		LastName   *[]string `url:"lastName,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	Sort     *[]string `url:"sort,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		Apps       *int `url:"apps,omitempty"`
		BetaGroups *int `url:"betaGroups,omitempty"`
		Builds     *int `url:"builds,omitempty"`
	} `url:"limit,omitempty"`
}

type GetBetaTesterOptions struct {
	Fields *struct {
		Apps        *[]string `url:"apps,omitempty"`
		BetaGroups  *[]string `url:"betaGroups,omitempty"`
		BetaTesters *[]string `url:"betaTesters,omitempty"`
		Builds      *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		Apps       *int `url:"apps,omitempty"`
		BetaGroups *int `url:"betaGroups,omitempty"`
		Builds     *int `url:"builds,omitempty"`
	} `url:"limit,omitempty"`
}

type ListAppsForBetaTesterOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListAppIDsForBetaTesterOptions struct {
	Limit *int `url:"limit,omitempty"`
}

type ListBuildsIndividuallyAssignedToBetaTesterOptions struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBuildIDsIndividuallyAssignedToBetaTesterOptions struct {
	Limit *int `url:"limit,omitempty"`
}

type ListIndividualTestersForBuildOptions struct {
	Fields *struct {
		BetaTesters *[]string `url:"betaTesters,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBetaGroupsForBetaTesterOptions struct {
	Fields *struct {
		BetaGroups *[]string `url:"betaGroups,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBetaGroupIDsForBetaTesterOptions struct {
	Limit *int `url:"limit,omitempty"`
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
func (s *Service) ListBetaTesters(id string, params *ListBetaTestersOptions) (*BetaTestersResponse, *common.Response, error) {
	res := new(BetaTestersResponse)
	resp, err := s.Get("betaTesters", params, res)
	return res, resp, err
}

// GetBetaTester gets a specific beta tester.
func (s *Service) GetBetaTester(id string, params *GetBetaTesterOptions) (*BetaTesterResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s", id)
	res := new(BetaTesterResponse)
	resp, err := s.Get(url, params, res)
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
func (s *Service) ListAppsForBetaTester(id string, params *ListAppsForBetaTesterOptions) (*apps.AppsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/apps", id)
	res := new(apps.AppsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListAppIDsForBetaTester gets a list of app resource IDs associated with a beta tester.
func (s *Service) ListAppIDsForBetaTester(id string, params *ListAppIDsForBetaTesterOptions) (*BetaTesterAppsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/apps", id)
	res := new(BetaTesterAppsLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBuildsIndividuallyAssignedToBetaTester gets a list of builds individually assigned to a specific beta tester.
func (s *Service) ListBuildsIndividuallyAssignedToBetaTester(id string, params *ListBuildsIndividuallyAssignedToBetaTesterOptions) (*builds.BuildsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBuildIDsIndividuallyAssignedToBetaTester gets a list of build resource IDs individually assigned to a specific beta tester.
func (s *Service) ListBuildIDsIndividuallyAssignedToBetaTester(id string, params *ListBuildIDsIndividuallyAssignedToBetaTesterOptions) (*BetaTesterBuildsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/builds", id)
	res := new(BetaTesterBuildsLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListIndividualTestersForBuild gets a list of beta testers individually assigned to a build.
func (s *Service) ListIndividualTestersForBuild(id string, params *ListIndividualTestersForBuildOptions) (*BetaTestersResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/individualTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaGroupsForBetaTester gets a list of beta groups that contain a specific beta tester.
func (s *Service) ListBetaGroupsForBetaTester(id string, params *ListBetaGroupsForBetaTesterOptions) (*BetaGroupsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaGroupIDsForBetaTester gets a list of group resource IDs associated with a beta tester.
func (s *Service) ListBetaGroupIDsForBetaTester(id string, params *ListBetaGroupIDsForBetaTesterOptions) (*BetaTesterBetaGroupsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaTesters/%s/relationships/betaGroups", id)
	res := new(BetaTesterBetaGroupsLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
