package apps

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/internal/services"
)

// AppInfo defines model for AppInfo.
type AppInfo struct {
	Attributes *struct {
		AppStoreAgeRating *AppStoreAgeRating    `json:"appStoreAgeRating,omitempty"`
		AppStoreState     *AppStoreVersionState `json:"appStoreState,omitempty"`
		BrazilAgeRating   *BrazilAgeRating      `json:"brazilAgeRating,omitempty"`
		KidsAgeBand       *KidsAgeBand          `json:"kidsAgeBand,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         services.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppInfoLocalizations *struct {
			Data  *[]services.RelationshipsData `json:"data,omitempty"`
			Links *services.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *services.PagingInformation   `json:"meta,omitempty"`
		} `json:"appInfoLocalizations,omitempty"`
		PrimaryCategory *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primaryCategory,omitempty"`
		PrimarySubcategoryOne *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryOne,omitempty"`
		PrimarySubcategoryTwo *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryTwo,omitempty"`
		SecondaryCategory *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondaryCategory,omitempty"`
		SecondarySubcategoryOne *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryOne,omitempty"`
		SecondarySubcategoryTwo *struct {
			Data  *services.RelationshipsData  `json:"data,omitempty"`
			Links *services.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryTwo,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppInfoResponse defines model for AppInfoResponse.
type AppInfoResponse struct {
	Data     AppInfo                `json:"data"`
	Included *[]interface{}         `json:"included,omitempty"`
	Links    services.DocumentLinks `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
type AppInfosResponse struct {
	Data     []AppInfo                   `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    services.PagedDocumentLinks `json:"links"`
	Meta     *services.PagingInformation `json:"meta,omitempty"`
}

// AppInfoUpdateRequest defines model for AppInfoUpdateRequest.
type AppInfoUpdateRequest struct {
	ID            string                             `json:"id"`
	Relationships *AppInfoUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// AppInfoUpdateRequestRelationships are relationships for AppInfoUpdateRequest
type AppInfoUpdateRequestRelationships struct {
	PrimaryCategory *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"primaryCategory,omitempty"`
	PrimarySubcategoryOne *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"primarySubcategoryOne,omitempty"`
	PrimarySubcategoryTwo *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"primarySubcategoryTwo,omitempty"`
	SecondaryCategory *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"secondaryCategory,omitempty"`
	SecondarySubcategoryOne *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"secondarySubcategoryOne,omitempty"`
	SecondarySubcategoryTwo *struct {
		Data *services.RelationshipsData `json:"data,omitempty"`
	} `json:"secondarySubcategoryTwo,omitempty"`
}

// GetAppInfoQuery are query options for GetAppInfo
type GetAppInfoQuery struct {
	FieldsAppInfos             *[]string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        *[]string `url:"fields[appCategories],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	LimitAppInfoLocalizations  *int      `url:"limit[appInfoLocalizations],omitempty"`
}

// ListAppInfosForAppQuery are query options for ListAppInfosForApp
type ListAppInfosForAppQuery struct {
	FieldsAppInfos             *[]string `url:"fields[appInfos],omitempty"`
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        *[]string `url:"fields[appCategories],omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// GetAppInfo reads App Store information including your App Store state, age ratings, Brazil age rating, and kids' age band.
func (s *Service) GetAppInfo(id string, params *GetAppInfoQuery) (*AppInfoResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppInfosForApp gets information about an app that is currently live on App Store, or that goes live with the next version.
func (s *Service) ListAppInfosForApp(id string, params *ListAppInfosForAppQuery) (*AppInfosResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/appInfos", id)
	res := new(AppInfosResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateAppInfo updates the App Store categories and sub-categories for your app.
func (s *Service) UpdateAppInfo(id string, body *AppInfoUpdateRequest) (*AppInfoResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
