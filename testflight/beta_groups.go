package testflight

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal"
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
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		BetaTesters *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"betaTesters,omitempty"`
		Builds *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaGroupResponse defines model for BetaGroupResponse.
type BetaGroupResponse struct {
	Data     BetaGroup              `json:"data"`
	Included *[]interface{}         `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
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
				Data internal.RelationshipsData `json:"data"`
			} `json:"app"`
			BetaTesters *struct {
				Data *[]internal.RelationshipsData `json:"data,omitempty"`
			} `json:"betaTesters,omitempty"`
			Builds *struct {
				Data *[]internal.RelationshipsData `json:"data,omitempty"`
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
	Data []internal.RelationshipsData `json:"data"`
}

// BetaGroupBetaTestersLinkagesRequest defines model for BetaGroupBetaTestersLinkagesRequest.
type BetaGroupBetaTestersLinkagesRequest struct {
	Data []internal.RelationshipsData `json:"data"`
}

// BetaGroupBetaTestersLinkagesResponse defines model for BetaGroupBetaTestersLinkagesResponse.
type BetaGroupBetaTestersLinkagesResponse struct {
	Data  []internal.RelationshipsData `json:"data"`
	Links internal.PagedDocumentLinks  `json:"links"`
	Meta  *internal.PagingInformation  `json:"meta,omitempty"`
}

// BetaGroupBuildsLinkagesResponse defines model for BetaGroupBuildsLinkagesResponse.
type BetaGroupBuildsLinkagesResponse struct {
	Data  []internal.RelationshipsData `json:"data"`
	Links internal.PagedDocumentLinks  `json:"links"`
	Meta  *internal.PagingInformation  `json:"meta,omitempty"`
}

// BetaGroupsResponse defines model for BetaGroupsResponse.
type BetaGroupsResponse struct {
	Data     []BetaGroup                 `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// ListBetaGroupsQuery defines model for ListBetaGroups
type ListBetaGroupsQuery struct {
	FieldsApps                   *[]string `url:"fields[apps],omitempty"`
	FieldsBetaGroups             *[]string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters            *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds                 *[]string `url:"fields[builds],omitempty"`
	FilterApp                    *[]string `url:"filter[app],omitempty"`
	FilterBuilds                 *[]string `url:"filter[builds],omitempty"`
	FilterID                     *[]string `url:"filter[id],omitempty"`
	FilterIsInternalGroup        *[]string `url:"filter[isInternalGroup],omitempty"`
	FilterName                   *[]string `url:"filter[name],omitempty"`
	FilterPublicLinkEnabled      *[]string `url:"filter[publicLinkEnabled],omitempty"`
	FilterPublicLinkLimitEnabled *[]string `url:"filter[publicLinkLimitEnabled],omitempty"`
	FilterPublicLink             *[]string `url:"filter[publicLink],omitempty"`
	Include                      *[]string `url:"include,omitempty"`
	Sort                         *[]string `url:"sort,omitempty"`
	Limit                        *int      `url:"limit,omitempty"`
	LimitBuilds                  *int      `url:"limit[builds],omitempty"`
	LimitBetaTesters             *int      `url:"limit[betaTesters],omitempty"`
	Cursor                       *string   `url:"cursor,omitempty"`
}

// GetBetaGroupQuery defines model for GetBetaGroup
type GetBetaGroupQuery struct {
	FieldsApps        *[]string `url:"fields[apps],omitempty"`
	FieldsBetaGroups  *[]string `url:"fields[betaGroups],omitempty"`
	FieldsBetaTesters *[]string `url:"fields[betaTesters],omitempty"`
	FieldsBuilds      *[]string `url:"fields[builds],omitempty"`
	Include           *[]string `url:"include,omitempty"`
	LimitBuilds       *int      `url:"limit[builds],omitempty"`
	LimitBetaTesters  *int      `url:"limit[betaTesters],omitempty"`
}

// GetAppForBetaGroupQuery defines model for GetAppForBetaGroup
type GetAppForBetaGroupQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// ListBetaGroupsForAppQuery defines model for ListBetaGroupsForApp
type ListBetaGroupsForAppQuery struct {
	FieldsBetaGroups *[]string `url:"fields[betaGroups],omitempty"`
	Limit            *int      `url:"limit,omitempty"`
	Cursor           *string   `url:"cursor,omitempty"`
}

// ListBuildsForBetaGroupQuery defines model for ListBuildsForBetaGroup
type ListBuildsForBetaGroupQuery struct {
	FieldsBuilds *[]string `url:"fields[builds],omitempty"`
	Limit        *int      `url:"limit,omitempty"`
	Cursor       *string   `url:"cursor,omitempty"`
}

// ListBuildIDsForBetaGroupQuery defines model for ListBuildIDsForBetaGroup
type ListBuildIDsForBetaGroupQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// ListBetaTestersForBetaGroupQuery defines model for ListBetaTestersForBetaGroup
type ListBetaTestersForBetaGroupQuery struct {
	FieldsBetaTesters *[]string `url:"fields[betaTesters],omitempty"`
	Limit             *int      `url:"limit,omitempty"`
	Cursor            *string   `url:"cursor,omitempty"`
}

// ListBetaTesterIDsForBetaGroupQuery defines model for ListBetaTesterIDsForBetaGroup
type ListBetaTesterIDsForBetaGroupQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// CreateBetaGroup creates a beta group associated with an app, optionally enabling TestFlight public links.
func (s *Service) CreateBetaGroup(body *BetaGroupCreateRequest) (*BetaGroupResponse, *internal.Response, error) {
	res := new(BetaGroupResponse)
	resp, err := s.Post("betaGroups", body, res)
	return res, resp, err
}

// UpdateBetaGroup modifies a beta group's metadata, including changing its Testflight public link status.
func (s *Service) UpdateBetaGroup(id string, body *BetaGroupUpdateRequest) (*BetaGroupResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteBetaGroup deletes a beta group and remove beta tester access to associated builds.
func (s *Service) DeleteBetaGroup(id string) (*internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	return s.Delete(url, nil)
}

// ListBetaGroups finds and lists beta groups for all apps.
func (s *Service) ListBetaGroups(params *ListBetaGroupsQuery) (*BetaGroupsResponse, *internal.Response, error) {
	res := new(BetaGroupsResponse)
	resp, err := s.GetWithQuery("betaGroups", params, res)
	return res, resp, err
}

// GetBetaGroup gets a specific beta group.
func (s *Service) GetBetaGroup(id string, params *GetBetaGroupQuery) (*BetaGroupResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s", id)
	res := new(BetaGroupResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaGroup gets the app information for a specific beta group.
func (s *Service) GetAppForBetaGroup(id string, params *GetAppForBetaGroupQuery) (*apps.AppResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaGroupsForApp gets a list of beta groups associated with a specific app.
func (s *Service) ListBetaGroupsForApp(id string, params *ListBetaGroupsForAppQuery) (*BetaGroupsResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/betaGroups", id)
	res := new(BetaGroupsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// AddBetaTestersToBetaGroup adds a specific beta tester to one or more beta groups for beta testing.
func (s *Service) AddBetaTestersToBetaGroup(id string, body *BetaGroupBetaTestersLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.Post(url, body, nil)
}

// RemoveBetaTestersFromBetaGroup removes a specific beta tester from a one or more beta groups, revoking their access to test builds associated with those groups.
func (s *Service) RemoveBetaTestersFromBetaGroup(id string, body *BetaGroupBetaTestersLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	return s.Delete(url, body)
}

// AddBuildsToBetaGroup associates builds with a beta group to enable the group to test the builds.
func (s *Service) AddBuildsToBetaGroup(id string, body *BetaGroupBuildsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.Post(url, body, nil)
}

// RemoveBuildsFromBetaGroup removes access to test one or more builds from beta testers in a specific beta group.
func (s *Service) RemoveBuildsFromBetaGroup(id string, body *BetaGroupBuildsLinkagesRequest) (*internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	return s.Delete(url, body)
}

// ListBuildsForBetaGroup gets a list of builds associated with a specific beta group.
func (s *Service) ListBuildsForBetaGroup(id string, params *ListBuildsForBetaGroupQuery) (*builds.BuildsResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/builds", id)
	res := new(builds.BuildsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBuildIDsForBetaGroup gets a list of build resource IDs in a specific beta group.
func (s *Service) ListBuildIDsForBetaGroup(id string, params *ListBuildIDsForBetaGroupQuery) (*BetaGroupBuildsLinkagesResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/builds", id)
	res := new(BetaGroupBuildsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaTestersForBetaGroup gets a list of beta testers contained in a specific beta group.
func (s *Service) ListBetaTestersForBetaGroup(id string, params *ListBetaTestersForBetaGroupQuery) (*BetaTestersResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/betaTesters", id)
	res := new(BetaTestersResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListBetaTesterIDsForBetaGroup gets a list of the beta tester resource IDs in a specific beta group.
func (s *Service) ListBetaTesterIDsForBetaGroup(id string, params *ListBetaTesterIDsForBetaGroupQuery) (*BetaGroupBetaTestersLinkagesResponse, *internal.Response, error) {
	url := fmt.Sprintf("betaGroups/%s/relationships/betaTesters", id)
	res := new(BetaGroupBetaTestersLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
