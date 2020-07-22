package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
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
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
		AppInfoLocalizations *struct {
			Data  *[]internal.RelationshipsData `json:"data,omitempty"`
			Links *internal.RelationshipsLinks  `json:"links,omitempty"`
			Meta  *internal.PagingInformation   `json:"meta,omitempty"`
		} `json:"appInfoLocalizations,omitempty"`
		PrimaryCategory *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primaryCategory,omitempty"`
		PrimarySubcategoryOne *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryOne,omitempty"`
		PrimarySubcategoryTwo *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"primarySubcategoryTwo,omitempty"`
		SecondaryCategory *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondaryCategory,omitempty"`
		SecondarySubcategoryOne *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryOne,omitempty"`
		SecondarySubcategoryTwo *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"secondarySubcategoryTwo,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppInfoResponse defines model for AppInfoResponse.
type AppInfoResponse struct {
	Data     AppInfo                `json:"data"`
	Included *[]interface{}         `json:"included,omitempty"`
	Links    internal.DocumentLinks `json:"links"`
}

// AppInfosResponse defines model for AppInfosResponse.
type AppInfosResponse struct {
	Data     []AppInfo                   `json:"data"`
	Included *[]interface{}              `json:"included,omitempty"`
	Links    internal.PagedDocumentLinks `json:"links"`
	Meta     *internal.PagingInformation `json:"meta,omitempty"`
}

// AppInfoUpdateRequest defines model for AppInfoUpdateRequest.
type AppInfoUpdateRequest struct {
	Data struct {
		ID            string `json:"id"`
		Relationships *struct {
			PrimaryCategory *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"primaryCategory,omitempty"`
			PrimarySubcategoryOne *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"primarySubcategoryOne,omitempty"`
			PrimarySubcategoryTwo *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"primarySubcategoryTwo,omitempty"`
			SecondaryCategory *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"secondaryCategory,omitempty"`
			SecondarySubcategoryOne *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"secondarySubcategoryOne,omitempty"`
			SecondarySubcategoryTwo *struct {
				Data *internal.RelationshipsData `json:"data,omitempty"`
			} `json:"secondarySubcategoryTwo,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

type GetAppInfoQuery struct {
	FieldsAppInfos             *[]string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        *[]string `url:"fields[appCategories],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	LimitAppInfoLocalizations  *int      `url:"limit[appInfoLocalizations],omitempty"`
}

type ListAppInfosForAppQuery struct {
	FieldsAppInfos             *[]string `url:"fields[appInfos],omitempty"`
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	FieldsAppCategories        *[]string `url:"fields[appCategories],omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// Read App Store information including your App Store state, age ratings, Brazil age rating, and kids' age band.
func (s *Service) GetAppInfo(id string, params *GetAppInfoQuery) (*AppInfoResponse, *internal.Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// ListAppInfosForApp gets information about an app that is currently live on App Store, or that goes live with the next version.
func (s *Service) ListAppInfosForApp(id string, params *ListAppInfosForAppQuery) (*AppInfosResponse, *internal.Response, error) {
	url := fmt.Sprintf("apps/%s/appInfos", id)
	res := new(AppInfosResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// Update the App Store categories and sub-categories for your app.
func (s *Service) UpdateAppInfo(id string, body *AppInfoUpdateRequest) (*AppInfoResponse, *internal.Response, error) {
	url := fmt.Sprintf("appInfos/%s", id)
	res := new(AppInfoResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
