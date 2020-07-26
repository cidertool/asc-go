package users

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal/services"
)

// UserRole defines model for UserRole.
type UserRole string

// List of UserRole
const (
	// AccessToReports.Downloads reports associated with a role. The Access To Reports role is an additional permission for users with the App Manager, Developer, Marketing, or Sales role. If this permission is added, the user has access to all of your apps.
	AccessToReports UserRole = "ACCESS_TO_REPORTS"
	// AccountHolder is responsible for entering into legal agreements with Apple. The person who completes program enrollment is assigned the Account Holder role in both the Apple Developer account and App Store Connect.
	AccountHolder UserRole = "ACCOUNT_HOLDER"
	// Admin serves as a secondary contact for teams and has many of the same responsibilities as the Account Holder role. Admins have access to all apps.
	Admin UserRole = "ADMIN"
	// AppManager manages all aspects of an app, such as pricing, App Store information, and app development and delivery.
	AppManager UserRole = "APP_MANAGER"
	// CustomerSupport analyzes and responds to customer reviews on the App Store. If a user has only the Customer Support role, they'll go straight to the Ratings and Reviews section when they click on an app in My Apps.
	CustomerSupport UserRole = "CUSTOMER_SUPPORT"
	// Developer manages development and delivery of an app.
	Developer UserRole = "DEVELOPER"
	// Finance manages financial information, including reports and tax forms. A user assigned this role can view all apps in Payments and Financial Reports, Sales and Trends, and App Analytics.
	Finance UserRole = "FINANCE"
	// Marketing manages marketing materials and promotional artwork. A user assigned this role will be contacted by Apple if the app is in consideration to be featured on the App Store.
	Marketing UserRole = "MARKETING"
	// ReadOnly represents a user with limited access and no write access.
	ReadOnly UserRole = "READ_ONLY"
	// Sales analyzes sales, downloads, and other analytics for the app.
	Sales UserRole = "SALES"
	// Technical role is no longer assignable to new users in App Store Connect. Existing users with the Technical role can manage all the aspects of an app, such as pricing, App Store information, and app development and delivery. Techncial users have access to all apps.
	Technical UserRole = "TECHNICAL"
)

// User defines model for User.
type User struct {
	Attributes *struct {
		AllAppsVisible      *bool       `json:"allAppsVisible,omitempty"`
		FirstName           *string     `json:"firstName,omitempty"`
		LastName            *string     `json:"lastName,omitempty"`
		ProvisioningAllowed *bool       `json:"provisioningAllowed,omitempty"`
		Roles               *[]UserRole `json:"roles,omitempty"`
		Username            *string     `json:"username,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		VisibleApps *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"visibleApps,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// UserUpdateRequest defines model for UserUpdateRequest.
type UserUpdateRequest struct {
	Attributes    *UserUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                          `json:"id"`
	Relationships *UserUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                          `json:"type"`
}

// UserUpdateRequestAttributes are attributes for UserUpdateRequest
type UserUpdateRequestAttributes struct {
	AllAppsVisible      *bool       `json:"allAppsVisible,omitempty"`
	ProvisioningAllowed *bool       `json:"provisioningAllowed,omitempty"`
	Roles               *[]UserRole `json:"roles,omitempty"`
}

// UserUpdateRequestRelationships are relationships for UserUpdateRequest
type UserUpdateRequestRelationships struct {
	VisibleApps *struct {
		Data *[]services.RelationshipsData `json:"data,omitempty"`
	} `json:"visibleApps,omitempty"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Data     User                   `json:"data"`
	Included *[]apps.App            `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// UsersResponse defines model for UsersResponse.
type UsersResponse struct {
	Data     []User                      `json:"data"`
	Included *[]apps.App                 `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// UserVisibleAppsLinkagesResponse defines model for UserVisibleAppsLinkagesResponse.
type UserVisibleAppsLinkagesResponse struct {
	Data  []services.RelationshipsData `json:"data"`
	Links services.PagedDocumentLinks  `json:"links"`
	Meta  *services.PagingInformation  `json:"meta,omitempty"`
}

// ListUsersQuery is the query params structure for ListUsers
type ListUsersQuery struct {
	FieldsApps        *[]string `url:"fields[apps],omitempty"`
	FieldsUsers       *[]string `url:"fields[users],omitempty"`
	FilterRoles       *[]string `url:"filter[roles],omitempty"`
	FilterVisibleApps *[]string `url:"filter[visibleApps],omitempty"`
	FilterUsername    *[]string `url:"filter[username],omitempty"`
	Limit             *int      `url:"limit,omitempty"`
	LimitVisibleApps  *int      `url:"limit[visibleApps],omitempty"`
	Include           *[]string `url:"include,omitempty"`
	Sort              *[]string `url:"sort,omitempty"`
	Cursor            *string   `url:"cursor,omitempty"`
}

// GetUserQuery is the query params structure for GetUser
type GetUserQuery struct {
	FieldsApps       *[]string `url:"fields[apps],omitempty"`
	FieldsUsers      *[]string `url:"fields[users],omitempty"`
	Include          *[]string `url:"include,omitempty"`
	Limit            *int      `url:"limit,omitempty"`
	LimitVisibleApps *int      `url:"limit[visibleApps],omitempty"`
}

// ListVisibleAppsQuery is the query params structure for ListVisibleAppsForUser
type ListVisibleAppsQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
	Limit      *int      `url:"limit,omitempty"`
	Cursor     *string   `url:"cursor,omitempty"`
}

// ListVisibleAppsByResourceIDQuery is the query params structure for ListVisibleAppsByResourceIDForUser
type ListVisibleAppsByResourceIDQuery struct {
	Limit  *int    `url:"limit,omitempty"`
	Cursor *string `url:"cursor,omitempty"`
}

// ListUsers gets a list of the users on your team.
func (s *Service) ListUsers(params *ListUsersQuery) (*UsersResponse, *http.Response, error) {
	res := new(UsersResponse)
	resp, err := s.GetWithQuery("users", params, res)
	return res, resp, err
}

// GetUser gets information about a user on your team, such as name, roles, and app visibility.
func (s *Service) GetUser(id string, params *GetUserQuery) (*UserResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%s", id)
	res := new(UserResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateUser changes a user's role, app visibility information, or other account details.
func (s *Service) UpdateUser(id string, body *UserUpdateRequest) (*UserResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%s", id)
	res := new(UserResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// RemoveUser removes a user from your team.
func (s *Service) RemoveUser(id string) (*http.Response, error) {
	url := fmt.Sprintf("users/%s", id)
	return s.Delete(url, nil)
}

// ListVisibleAppsForUser gets a list of apps that a user on your team can view.
func (s *Service) ListVisibleAppsForUser(id string, params *ListVisibleAppsQuery) (*apps.AppsResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%s/visibleApps", id)
	res := new(apps.AppsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListVisibleAppsByResourceIDForUser gets a list of app resource IDs to which a user on your team has access.
func (s *Service) ListVisibleAppsByResourceIDForUser(id string, params *ListVisibleAppsByResourceIDQuery) (*UserVisibleAppsLinkagesResponse, *http.Response, error) {
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)
	res := new(UserVisibleAppsLinkagesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// AddVisibleAppsForUser gives a user on your team access to one or more apps.
func (s *Service) AddVisibleAppsForUser(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	return s.Post("appStoreReviewDetails", linkages, nil)
}

// UpdateVisibleAppsForUser replaces the list of apps a user on your team can see.
func (s *Service) UpdateVisibleAppsForUser(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)
	return s.Patch(url, linkages, nil)
}

// RemoveVisibleAppsFromUser removes a user on your team’s access to one or more apps.
func (s *Service) RemoveVisibleAppsFromUser(id string, linkages *[]services.RelationshipsData) (*http.Response, error) {
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)
	return s.Delete(url, linkages)
}
