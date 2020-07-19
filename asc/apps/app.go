package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// Platform defines model for Platform.
type Platform string

// List of Platform
const (
	IOS   Platform = "IOS"
	MACOS Platform = "MAC_OS"
	TVOS  Platform = "TV_OS"
)

// App defines model for App.
type App struct {
	Attributes *struct {
		AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
		BundleID                  *string `json:"bundleId,omitempty"`
		ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
		IsOrEverWasMadeForKids    *bool   `json:"isOrEverWasMadeForKids,omitempty"`
		Name                      *string `json:"name,omitempty"`
		PrimaryLocale             *string `json:"primaryLocale,omitempty"`
		Sku                       *string `json:"sku,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppInfos *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"appInfos,omitempty"`
		AppStoreVersions *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"appStoreVersions,omitempty"`
		AvailableTerritories *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"availableTerritories,omitempty"`
		BetaAppLocalizations *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"betaAppLocalizations,omitempty"`
		BetaAppReviewDetail *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"betaAppReviewDetail,omitempty"`
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
		BetaLicenseAgreement *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"betaLicenseAgreement,omitempty"`
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
		EndUserLicenseAgreement *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"endUserLicenseAgreement,omitempty"`
		GameCenterEnabledVersions *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"gameCenterEnabledVersions,omitempty"`
		InAppPurchases *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"inAppPurchases,omitempty"`
		PreOrder *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"preOrder,omitempty"`
		PreReleaseVersions *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"preReleaseVersions,omitempty"`
		Prices *struct {
			Data *[]struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
			Meta *common.PagingInformation `json:"meta,omitempty"`
		} `json:"prices,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppUpdateRequest defines model for AppUpdateRequest.
type AppUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
			BundleID                  *string `json:"bundleId,omitempty"`
			ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
			PrimaryLocale             *string `json:"primaryLocale,omitempty"`
		} `json:"attributes,omitempty"`
		ID            string `json:"id"`
		Relationships *struct {
			AvailableTerritories *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"availableTerritories,omitempty"`
			Prices *struct {
				Data *[]struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data,omitempty"`
			} `json:"prices,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppResponse defines model for AppResponse.
type AppResponse struct {
	Data     App                  `json:"data"`
	Included *[]interface{}       `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// AppsResponse defines model for AppsResponse.
type AppsResponse struct {
	Data     []App                     `json:"data"`
	Included *[]interface{}            `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

// InAppPurchase defines model for InAppPurchase.
type InAppPurchase struct {
	Attributes *struct {
		InAppPurchaseType *string `json:"inAppPurchaseType,omitempty"`
		ProductID         *string `json:"productId,omitempty"`
		ReferenceName     *string `json:"referenceName,omitempty"`
		State             *string `json:"state,omitempty"`
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
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// InAppPurchaseResponse defines model for InAppPurchaseResponse.
type InAppPurchaseResponse struct {
	Data  InAppPurchase        `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// InAppPurchasesResponse defines model for InAppPurchasesResponse.
type InAppPurchasesResponse struct {
	Data  []InAppPurchase           `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

// AppBetaTestersLinkagesRequest defines model for AppBetaTestersLinkagesRequest.
type AppBetaTestersLinkagesRequest struct {
	Data []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

type ListAppsQuery struct {
	FieldsApps                          *[]string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements         *[]string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions            *[]string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails          *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations          *[]string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                        *[]string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                    *[]string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements      *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions              *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories                   *[]string `url:"fields[territories],omitempty"`
	FieldsAppPrices                     *[]string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders                  *[]string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                      *[]string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics              *[]string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsInAppPurchases                *[]string `url:"fields[inAppPurchases],omitempty"`
	FilterBundleID                      *[]string `url:"filter[bundleId],omitempty"`
	FilterID                            *[]string `url:"filter[id],omitempty"`
	FilterName                          *[]string `url:"filter[name],omitempty"`
	FilterSKU                           *[]string `url:"filter[sku],omitempty"`
	FilterAppStoreVersions              *[]string `url:"filter[appStoreVersions],omitempty"`
	FilterAppStoreVersionsPlatform      *[]string `url:"filter[appStoreVersionsPlatform],omitempty"`
	FilterAppStoreVersionsAppStoreState *[]string `url:"filter[appStoreVersionsAppStoreState],omitempty"`
	FilterGameCenterEnabledVersions     *[]string `url:"filter[gameCenterEnabledVersions],omitempty"`
	Include                             *[]string `url:"include,omitempty"`
	Limit                               *int      `url:"limit,omitempty"`
	LimitPreReleaseVersions             *int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                         *int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                     *int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations           *int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                         *int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories           *int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions               *int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                       *int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions      *int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases                 *int      `url:"limit[inAppPurchases],omitempty"`
	Sort                                *[]string `url:"sort,omitempty"`
	ExistsGameCenterEnabledVersions     *[]string `url:"exists[gameCenterEnabledVersions],omitempty"`
	Cursor                              *string   `url:"cursor,omitempty"`
}

type GetAppQuery struct {
	FieldsApps                      *[]string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements     *[]string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions        *[]string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails      *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations      *[]string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                    *[]string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                *[]string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements  *[]string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions          *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories               *[]string `url:"fields[territories],omitempty"`
	FieldsAppPrices                 *[]string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders              *[]string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                  *[]string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics          *[]string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsGameCenterEnabledVersions *[]string `url:"fields[gameCenterEnabledVersions],omitempty"`
	FieldsInAppPurchases            *[]string `url:"fields[inAppPurchases],omitempty"`
	Include                         *[]string `url:"include,omitempty"`
	LimitPreReleaseVersions         *int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                     *int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                 *int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations       *int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                     *int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories       *int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions           *int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                   *int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions  *int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases             *int      `url:"limit[inAppPurchases],omitempty"`
}

type ListInAppPurchasesQuery struct {
	FieldsApps              *[]string `url:"fields[apps],omitempty"`
	FieldsInAppPurchases    *[]string `url:"fields[inAppPurchases],omitempty"`
	FilterCanBeSubmitted    *[]string `url:"filter[canBeSubmitted],omitempty"`
	FilterInAppPurchaseType *[]string `url:"filter[inAppPurchaseType],omitempty"`
	Limit                   *int      `url:"limit,omitempty"`
	Include                 *[]string `url:"include,omitempty"`
	Sort                    *[]string `url:"sort,omitempty"`
	Cursor                  *string   `url:"cursor,omitempty"`
}

type GetInAppPurchaseQuery struct {
	FieldsInAppPurchases *[]string `url:"fields[inAppPurchases],omitempty"`
	Include              *[]string `url:"include,omitempty"`
	LimitApps            *int      `url:"limit[apps],omitempty"`
}

// ListApps finds and lists apps added in App Store Connect.
func (s *Service) ListApps(params *ListAppsQuery) (*AppsResponse, *common.Response, error) {
	res := new(AppsResponse)
	resp, err := s.GetWithQuery("apps", params, res)
	return res, resp, err
}

// GetApp gets information about a specific app.
func (s *Service) GetApp(id string, params *GetAppQuery) (*AppResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateApp updates app information including bundle ID, primary locale, price schedule, and global availability.
func (s *Service) UpdateApp(id string, body *AppUpdateRequest) (*AppResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// RemoveBetaTestersFromApp removes one or more beta testers' access to test any builds of a specific app.
func (s *Service) RemoveBetaTestersFromApp(id string, body *AppBetaTestersLinkagesRequest) (*common.Response, error) {
	url := fmt.Sprintf("apps/%s/relationships/betaTesters", id)
	return s.Delete(url, body)
}

// ListInAppPurchasesForApp lists the in-app purchases that are available for your app.
func (s *Service) ListInAppPurchasesForApp(id string, params *ListInAppPurchasesQuery) (*InAppPurchasesResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/inAppPurchases", id)
	res := new(InAppPurchasesResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetInAppPurchase gets information about an in-app purchase.
func (s *Service) GetInAppPurchase(id string, params *GetInAppPurchaseQuery) (*InAppPurchaseResponse, *common.Response, error) {
	url := fmt.Sprintf("inAppPurchases/%s", id)
	res := new(InAppPurchaseResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
