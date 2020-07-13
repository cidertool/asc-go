package testflight

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaGroup defines model for BetaGroup.
type BetaGroup struct {
	Attributes *struct {
		CreatedDate            *time.Time `json:"createdDate,omitempty"`
		FeedbackEnabled        *bool      `json:"feedbackEnabled,omitempty"`
		IsInternalGroup        *bool      `json:"isInternalGroup,omitempty"`
		Name                   *string    `json:"name,omitempty"`
		PublicLink             *string    `json:"publicLink,omitempty"`
		PublicLinkEnabled      *bool      `json:"publicLinkEnabled,omitempty"`
		PublicLinkID           *string    `json:"publicLinkId,omitempty"`
		PublicLinkLimit        *int       `json:"publicLinkLimit,omitempty"`
		PublicLinkLimitEnabled *bool      `json:"publicLinkLimitEnabled,omitempty"`
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
		BetaTesters *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"betaTesters,omitempty"`
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

// BetaGroupResponse defines model for BetaGroupResponse.
type BetaGroupResponse struct {
	Data     BetaGroup            `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BetaGroupCreateRequest defines model for BetaGroupCreateRequest.
type BetaGroupCreateRequest struct {
	Data struct {
		Attributes struct {
			FeedbackEnabled        *bool  `json:"feedbackEnabled,omitempty"`
			Name                   string `json:"name"`
			PublicLinkEnabled      *bool  `json:"publicLinkEnabled,omitempty"`
			PublicLinkLimit        *int   `json:"publicLinkLimit,omitempty"`
			PublicLinkLimitEnabled *bool  `json:"publicLinkLimitEnabled,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			App struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"app"`
			BetaTesters *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"betaTesters,omitempty"`
			Builds *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"builds,omitempty"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaGroupUpdateRequest defines model for BetaGroupUpdateRequest.
type BetaGroupUpdateRequest struct {
	Data struct {
		Attributes *struct {
			FeedbackEnabled        *bool   `json:"feedbackEnabled,omitempty"`
			Name                   *string `json:"name,omitempty"`
			PublicLinkEnabled      *bool   `json:"publicLinkEnabled,omitempty"`
			PublicLinkLimit        *int    `json:"publicLinkLimit,omitempty"`
			PublicLinkLimitEnabled *bool   `json:"publicLinkLimitEnabled,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaGroupBuildsLinkagesRequest defines model for BetaGroupBuildsLinkagesRequest.
type BetaGroupBuildsLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaGroupBetaTestersLinkagesRequest defines model for BetaGroupBetaTestersLinkagesRequest.
type BetaGroupBetaTestersLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaGroupBetaTestersLinkagesResponse defines model for BetaGroupBetaTestersLinkagesResponse.
type BetaGroupBetaTestersLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BetaGroupBuildsLinkagesResponse defines model for BetaGroupBuildsLinkagesResponse.
type BetaGroupBuildsLinkagesResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// BetaGroupsResponse defines model for BetaGroupsResponse.
type BetaGroupsResponse struct {
	Data     []BetaGroup               `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaGroupsOptions struct {
	Fields *struct {
		Apps        *[]string `url:"apps,omitempty"`
		BetaGroups  *[]string `url:"betaGroups,omitempty"`
		BetaTesters *[]string `url:"betaTesters,omitempty"`
		Builds      *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App                    *[]string `url:"app,omitempty"`
		Builds                 *[]string `url:"builds,omitempty"`
		ID                     *[]string `url:"id,omitempty"`
		IsInternalGroup        *[]string `url:"isInternalGroup,omitempty"`
		Name                   *[]string `url:"name,omitempty"`
		PublicLinkEnabled      *[]string `url:"publicLinkEnabled,omitempty"`
		PublicLinkLimitEnabled *[]string `url:"publicLinkLimitEnabled,omitempty"`
		PublicLink             *[]string `url:"publicLink,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	Sort     *[]string `url:"sort,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		Builds      *int `url:"builds,omitempty"`
		BetaTesters *int `url:"betaTesters,omitempty"`
	} `url:"limit,omitempty"`
}

type GetBetaGroupOptions struct {
	Fields *struct {
		Apps        *[]string `url:"apps,omitempty"`
		BetaGroups  *[]string `url:"betaGroups,omitempty"`
		BetaTesters *[]string `url:"betaTesters,omitempty"`
		Builds      *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		Builds      *int `url:"builds,omitempty"`
		BetaTesters *int `url:"betaTesters,omitempty"`
	} `url:"limit,omitempty"`
}

type GetAppForBetaGroupOptions struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type ListBetaGroupsForAppOptions struct {
	Fields *struct {
		BetaGroups *[]string `url:"betaGroups,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBuildsForBetaGroupOptions struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBuildIDsForBetaGroupOptions struct {
	Limit *int `url:"limit,omitempty"`
}

type ListBetaTestersForBetaGroupOptions struct {
	Fields *struct {
		BetaTesters *[]string `url:"betaTesters,omitempty"`
	} `url:"fields,omitempty"`
	Limit *int `url:"limit,omitempty"`
}

type ListBetaTesterIDsForBetaGroupOptions struct {
	Limit *int `url:"limit,omitempty"`
}

// CreateBetaGroup creates a beta group associated with an app, optionally enabling TestFlight public links.
func (s *Service) CreateBetaGroup(body *BetaGroupCreateRequest) (*BetaGroupResponse, *common.Response, error) {
	res := new(BetaGroupResponse)
	resp, err := s.Post("betaGroups", body, res)
	return res, resp, err
}

// UpdateBetaGroup modifies a beta group's metadata, including changing its Testflight public link status.
func (s *Service) UpdateBetaGroup(id string, body *BetaGroupUpdateRequest) (*BetaGroupResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBetaGroup deletes a beta group and remove beta tester access to associated builds.
func (s *Service) DeleteBetaGroup(id string) (*common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	return s.Delete(url, nil)
}

// ListBetaGroups finds and lists beta groups for all apps.
func (s *Service) ListBetaGroups(params *ListBetaGroupsOptions) (*BetaGroupsResponse, *common.Response, error) {
	res := new(BetaGroupsResponse)
	resp, err := s.Get("betaGroups", params, res)
	return res, resp, err
}

// GetBetaGroup gets a specific beta group.
func (s *Service) GetBetaGroup(id string, params *GetBetaGroupOptions) (*BetaGroupResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetAppForBetaGroup gets the app information for a specific beta group.
func (s *Service) GetAppForBetaGroup(id string, params *GetAppForBetaGroupOptions) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaGroupsForApp gets a list of beta groups associated with a specific app.
func (s *Service) ListBetaGroupsForApp(id string, params *ListBetaGroupsForAppOptions) (*BetaGroupsResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// AddBetaTestersToBetaGroup adds a specific beta tester to one or more beta groups for beta testing.
func (s *Service) AddBetaTestersToBetaGroup(id string, body *BetaGroupBetaTestersLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.Post(url, body, nil)
}

// RemoveBetaTestersFromBetaGroup removes a specific beta tester from a one or more beta groups, revoking their access to test builds associated with those groups.
func (s *Service) RemoveBetaTestersFromBetaGroup(id string, body *BetaGroupBetaTestersLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.Delete(url, body)
}

// AddBuildsToBetaGroup associates builds with a beta group to enable the group to test the builds.
func (s *Service) AddBuildsToBetaGroup(id string, body *BetaGroupBuildsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.Post(url, body, nil)
}

// RemoveBuildsFromBetaGroup removes access to test one or more builds from beta testers in a specific beta group.
func (s *Service) RemoveBuildsFromBetaGroup(id string, body *BetaGroupBuildsLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.Delete(url, body)
}

// ListBuildsForBetaGroup gets a list of builds associated with a specific beta group.
func (s *Service) ListBuildsForBetaGroup(id string, params *ListBuildsForBetaGroupOptions) (*builds.BuildsResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBuildIDsForBetaGroup gets a list of build resource IDs in a specific beta group.
func (s *Service) ListBuildIDsForBetaGroup(id string, params *ListBuildIDsForBetaGroupOptions) (*BetaGroupBuildsLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	res := new(BetaGroupBuildsLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaTestersForBetaGroup gets a list of beta testers contained in a specific beta group.
func (s *Service) ListBetaTestersForBetaGroup(id string, params *ListBetaTestersForBetaGroupOptions) (*BetaTestersResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/betaTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// ListBetaTesterIDsForBetaGroup gets a list of the beta tester resource IDs in a specific beta group.
func (s *Service) ListBetaTesterIDsForBetaGroup(id string, params *ListBetaTesterIDsForBetaGroupOptions) (*BetaGroupBetaTestersLinkagesResponse, *common.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	res := new(BetaGroupBetaTestersLinkagesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
