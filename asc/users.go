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

// UsersService handles communication with user and role-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/users
// https://developer.apple.com/documentation/appstoreconnectapi/user_invitations
type UsersService service

// UserRole defines model for UserRole.
//
// https://developer.apple.com/documentation/appstoreconnectapi/userrole
type UserRole string

const (
	// UserRoleAccessToReports allows download access to reports associated with a role. The Access To Reports role is an additional permission for users with the App Manager, Developer, Marketing, or Sales role. If this permission is added, the user has access to all of your apps.
	UserRoleAccessToReports UserRole = "ACCESS_TO_REPORTS"
	// UserRoleAccountHolder is responsible for entering into legal agreements with Apple. The person who completes program enrollment is assigned the Account Holder role in both the Apple Developer account and App Store Connect.
	UserRoleAccountHolder UserRole = "ACCOUNT_HOLDER"
	// UserRoleAdmin serves as a secondary contact for teams and has many of the same responsibilities as the Account Holder role. Admins have access to all apps.
	UserRoleAdmin UserRole = "ADMIN"
	// UserRoleAppManager manages all aspects of an app, such as pricing, App Store information, and app development and delivery.
	UserRoleAppManager UserRole = "APP_MANAGER"
	// UserRoleCloudManagedAppDistribution is a service role for managing Cloud Signing of apps for distribution.
	UserRoleCloudManagedAppDistribution UserRole = "CLOUD_MANAGED_APP_DISTRIBUTION"
	// UserRoleCloudManagedDeveloperID is a service role covering the developer account ID for an app enabled for Cloud Signing.
	UserRoleCloudManagedDeveloperID UserRole = "CLOUD_MANAGED_DEVELOPER_ID"
	// UserRoleCreateApps is a service role for automatic creation of apps.
	UserRoleCreateApps UserRole = "CREATE_APPS"
	// UserRoleCustomerSupport analyzes and responds to customer reviews on the App Store. If a user has only the Customer Support role, they'll go straight to the Ratings and Reviews section when they click on an app in My Apps.
	UserRoleCustomerSupport UserRole = "CUSTOMER_SUPPORT"
	// UserRoleDeveloper manages development and delivery of an app.
	UserRoleDeveloper UserRole = "DEVELOPER"
	// UserRoleFinance manages financial information, including reports and tax forms. A user assigned this role can view all apps in Payments and Financial Reports, Sales and Trends, and App Analytics.
	UserRoleFinance UserRole = "FINANCE"
	// UserRoleMarketing manages marketing materials and promotional artwork. A user assigned this role will be contacted by Apple if the app is in consideration to be featured on the App Store.
	UserRoleMarketing UserRole = "MARKETING"
	// UserRoleReadOnly represents a user with limited access and no write access.
	UserRoleReadOnly UserRole = "READ_ONLY"
	// UserRoleSales analyzes sales, downloads, and other analytics for the app.
	UserRoleSales UserRole = "SALES"
	// UserRoleTechnical role is no longer assignable to new users in App Store Connect. Existing users with the Technical role can manage all the aspects of an app, such as pricing, App Store information, and app development and delivery. Techncial users have access to all apps.
	UserRoleTechnical UserRole = "TECHNICAL"
)

// User defines model for User.
//
// https://developer.apple.com/documentation/appstoreconnectapi/user
type User struct {
	Attributes    *UserAttributes    `json:"attributes,omitempty"`
	ID            string             `json:"id"`
	Links         ResourceLinks      `json:"links"`
	Relationships *UserRelationships `json:"relationships,omitempty"`
	Type          string             `json:"type"`
}

// UserAttributes defines model for User.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/user/attributes
type UserAttributes struct {
	AllAppsVisible      *bool      `json:"allAppsVisible,omitempty"`
	FirstName           *string    `json:"firstName,omitempty"`
	LastName            *string    `json:"lastName,omitempty"`
	ProvisioningAllowed *bool      `json:"provisioningAllowed,omitempty"`
	Roles               []UserRole `json:"roles,omitempty"`
	Username            *string    `json:"username,omitempty"`
}

// UserRelationships defines model for User.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/user/relationships
type UserRelationships struct {
	VisibleApps *PagedRelationship `json:"visibleApps,omitempty"`
}

// UserUpdateRequest defines model for UserUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/userupdaterequest/data
type userUpdateRequest struct {
	Attributes    *UserUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                          `json:"id"`
	Relationships *userUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                          `json:"type"`
}

// UserUpdateRequestAttributes are attributes for UserUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/userupdaterequest/data/attributes
type UserUpdateRequestAttributes struct {
	AllAppsVisible      *bool      `json:"allAppsVisible,omitempty"`
	ProvisioningAllowed *bool      `json:"provisioningAllowed,omitempty"`
	Roles               []UserRole `json:"roles,omitempty"`
}

// userUpdateRequestRelationships are relationships for UserUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/userupdaterequest/data/relationships
type userUpdateRequestRelationships struct {
	VisibleApps *pagedRelationshipDeclaration `json:"visibleApps,omitempty"`
}

// UserResponse defines model for UserResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/userresponse
type UserResponse struct {
	Data     User          `json:"data"`
	Included []App         `json:"included,omitempty"`
	Links    DocumentLinks `json:"links"`
}

// UsersResponse defines model for UsersResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/usersresponse
type UsersResponse struct {
	Data     []User             `json:"data"`
	Included []App              `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// UserVisibleAppsLinkagesResponse defines model for UserVisibleAppsLinkagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/uservisibleappslinkagesresponse
type UserVisibleAppsLinkagesResponse struct {
	Data  []RelationshipData `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListUsersQuery is query options for ListUsers
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_users
type ListUsersQuery struct {
	FieldsApps        []string `url:"fields[apps],omitempty"`
	FieldsUsers       []string `url:"fields[users],omitempty"`
	FilterRoles       []string `url:"filter[roles],omitempty"`
	FilterVisibleApps []string `url:"filter[visibleApps],omitempty"`
	FilterUsername    []string `url:"filter[username],omitempty"`
	Limit             int      `url:"limit,omitempty"`
	LimitVisibleApps  int      `url:"limit[visibleApps],omitempty"`
	Include           []string `url:"include,omitempty"`
	Sort              []string `url:"sort,omitempty"`
	Cursor            string   `url:"cursor,omitempty"`
}

// GetUserQuery is query options for GetUser
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_user_information
type GetUserQuery struct {
	FieldsApps       []string `url:"fields[apps],omitempty"`
	FieldsUsers      []string `url:"fields[users],omitempty"`
	Include          []string `url:"include,omitempty"`
	Limit            int      `url:"limit,omitempty"`
	LimitVisibleApps int      `url:"limit[visibleApps],omitempty"`
}

// ListVisibleAppsQuery is query options for ListVisibleAppsForUser
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_apps_visible_to_a_user
type ListVisibleAppsQuery struct {
	FieldsApps []string `url:"fields[apps],omitempty"`
	Limit      int      `url:"limit,omitempty"`
	Cursor     string   `url:"cursor,omitempty"`
}

// ListVisibleAppsByResourceIDQuery is query options for ListVisibleAppsByResourceIDForUser
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_visible_app_resource_ids_for_a_user
type ListVisibleAppsByResourceIDQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// ListUsers gets a list of the users on your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_users
func (s *UsersService) ListUsers(ctx context.Context, params *ListUsersQuery) (*UsersResponse, *Response, error) {
	res := new(UsersResponse)
	resp, err := s.client.get(ctx, "users", params, res)

	return res, resp, err
}

// GetUser gets information about a user on your team, such as name, roles, and app visibility.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_user_information
func (s *UsersService) GetUser(ctx context.Context, id string, params *GetUserQuery) (*UserResponse, *Response, error) {
	url := fmt.Sprintf("users/%s", id)
	res := new(UserResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateUser changes a user's role, app visibility information, or other account details.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_user_account
func (s *UsersService) UpdateUser(ctx context.Context, id string, attributes *UserUpdateRequestAttributes, visibleAppIDs []string) (*UserResponse, *Response, error) {
	req := userUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "users",
	}

	if len(visibleAppIDs) > 0 {
		relationships := newPagedRelationshipDeclaration(visibleAppIDs, "apps")
		req.Relationships = &userUpdateRequestRelationships{
			VisibleApps: &relationships,
		}
	}

	url := fmt.Sprintf("users/%s", id)
	res := new(UserResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// RemoveUser removes a user from your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_a_user_account
func (s *UsersService) RemoveUser(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("users/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListVisibleAppsForUser gets a list of apps that a user on your team can view.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_apps_visible_to_a_user
func (s *UsersService) ListVisibleAppsForUser(ctx context.Context, id string, params *ListVisibleAppsQuery) (*AppsResponse, *Response, error) {
	url := fmt.Sprintf("users/%s/visibleApps", id)
	res := new(AppsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListVisibleAppsByResourceIDForUser gets a list of app resource IDs to which a user on your team has access.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_all_visible_app_resource_ids_for_a_user
func (s *UsersService) ListVisibleAppsByResourceIDForUser(ctx context.Context, id string, params *ListVisibleAppsByResourceIDQuery) (*UserVisibleAppsLinkagesResponse, *Response, error) {
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)
	res := new(UserVisibleAppsLinkagesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// AddVisibleAppsForUser gives a user on your team access to one or more apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/add_visible_apps_to_a_user
func (s *UsersService) AddVisibleAppsForUser(ctx context.Context, id string, appIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appIDs, "apps")

	return s.client.post(ctx, "appStoreReviewDetails", newRequestBody(linkages.Data), nil)
}

// UpdateVisibleAppsForUser replaces the list of apps a user on your team can see.
//
// https://developer.apple.com/documentation/appstoreconnectapi/replace_the_list_of_visible_apps_for_a_user
func (s *UsersService) UpdateVisibleAppsForUser(ctx context.Context, id string, appIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appIDs, "apps")
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}

// RemoveVisibleAppsFromUser removes a user on your team’s access to one or more apps.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_visible_apps_from_a_user
func (s *UsersService) RemoveVisibleAppsFromUser(ctx context.Context, id string, appIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(appIDs, "apps")
	url := fmt.Sprintf("users/%s/relationships/visibleApps", id)

	return s.client.delete(ctx, url, newRequestBody(linkages.Data))
}
