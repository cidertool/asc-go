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

type ListAppsOptions struct {
	Fields *struct {
		Apps                     *[]string `url:"apps,omitempty"`
		BetaLicenseAgreements    *[]string `url:"betaLicenseAgreements,omitempty"`
		PreReleaseVersions       *[]string `url:"preReleaseVersions,omitempty"`
		BetaAppReviewDetails     *[]string `url:"betaAppReviewDetails,omitempty"`
		BetaAppLocalizations     *[]string `url:"betaAppLocalizations,omitempty"`
		Builds                   *[]string `url:"builds,omitempty"`
		BetaGroups               *[]string `url:"betaGroups,omitempty"`
		EndUserLicenseAgreements *[]string `url:"endUserLicenseAgreements,omitempty"`
		AppStoreVersions         *[]string `url:"appStoreVersions,omitempty"`
		Territories              *[]string `url:"territories,omitempty"`
		AppPrices                *[]string `url:"appPrices,omitempty"`
		AppPreOrders             *[]string `url:"appPreOrders,omitempty"`
		AppInfos                 *[]string `url:"appInfos,omitempty"`
		PerfPowerMetrics         *[]string `url:"perfPowerMetrics,omitempty"`
		InAppPurchases           *[]string `url:"inAppPurchases,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		BundleID                      *[]string `url:"bundleId,omitempty"`
		ID                            *[]string `url:"id,omitempty"`
		Name                          *[]string `url:"name,omitempty"`
		SKU                           *[]string `url:"sku,omitempty"`
		AppStoreVersions              *[]string `url:"appStoreVersions,omitempty"`
		AppStoreVersionsPlatform      *[]string `url:"appStoreVersionsPlatform,omitempty"`
		AppStoreVersionsAppStoreState *[]string `url:"appStoreVersionsAppStoreState,omitempty"`
		GameCenterEnabledVersions     *[]string `url:"gameCenterEnabledVersions,omitempty"`
	} `url:"filter,omitempty"`
	Include  *[]string `url:"include,omitempty"`
	LimitAll *int      `url:"limit,omitempty"`
	Limit    *struct {
		PreReleaseVersions        *int `url:"preReleaseVersions,omitempty"`
		Builds                    *int `url:"builds,omitempty"`
		BetaGroups                *int `url:"betaGroups,omitempty"`
		BetaAppLocalizations      *int `url:"betaAppLocalizations,omitempty"`
		Prices                    *int `url:"prices,omitempty"`
		AvailableTerritories      *int `url:"availableTerritories,omitempty"`
		AppStoreVersions          *int `url:"appStoreVersions,omitempty"`
		AppInfos                  *int `url:"appInfos,omitempty"`
		GameCenterEnabledVersions *int `url:"gameCenterEnabledVersions,omitempty"`
		InAppPurchases            *int `url:"inAppPurchases,omitempty"`
	} `url:"limit,omitempty"`
	Sort   *[]string `url:"sort,omitempty"`
	Exists *struct {
		GameCenterEnabledVersions *[]string `url:"gameCenterEnabledVersions,omitempty"`
	} `url:"exists,omitempty"`
}

type GetAppOptions struct {
	Fields *struct {
		Apps                      *[]string `url:"apps,omitempty"`
		BetaLicenseAgreements     *[]string `url:"betaLicenseAgreements,omitempty"`
		PreReleaseVersions        *[]string `url:"preReleaseVersions,omitempty"`
		BetaAppReviewDetails      *[]string `url:"betaAppReviewDetails,omitempty"`
		BetaAppLocalizations      *[]string `url:"betaAppLocalizations,omitempty"`
		Builds                    *[]string `url:"builds,omitempty"`
		BetaGroups                *[]string `url:"betaGroups,omitempty"`
		EndUserLicenseAgreements  *[]string `url:"endUserLicenseAgreements,omitempty"`
		AppStoreVersions          *[]string `url:"appStoreVersions,omitempty"`
		Territories               *[]string `url:"territories,omitempty"`
		AppPrices                 *[]string `url:"appPrices,omitempty"`
		AppPreOrders              *[]string `url:"appPreOrders,omitempty"`
		AppInfos                  *[]string `url:"appInfos,omitempty"`
		PerfPowerMetrics          *[]string `url:"perfPowerMetrics,omitempty"`
		GameCenterEnabledVersions *[]string `url:"gameCenterEnabledVersions,omitempty"`
		InAppPurchases            *[]string `url:"inAppPurchases,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		PreReleaseVersions        *int `url:"preReleaseVersions,omitempty"`
		Builds                    *int `url:"builds,omitempty"`
		BetaGroups                *int `url:"betaGroups,omitempty"`
		BetaAppLocalizations      *int `url:"betaAppLocalizations,omitempty"`
		Prices                    *int `url:"prices,omitempty"`
		AvailableTerritories      *int `url:"availableTerritories,omitempty"`
		AppStoreVersions          *int `url:"appStoreVersions,omitempty"`
		AppInfos                  *int `url:"appInfos,omitempty"`
		GameCenterEnabledVersions *int `url:"gameCenterEnabledVersions,omitempty"`
		InAppPurchases            *int `url:"inAppPurchases,omitempty"`
	} `url:"limit,omitempty"`
}

type ListInAppPurchasesOptions struct {
	Fields *struct {
		Apps           *[]string `url:"apps,omitempty"`
		InAppPurchases *[]string `url:"inAppPurchases,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		CanBeSubmitted    *[]string `url:"canBeSubmitted,omitempty"`
		InAppPurchaseType *[]string `url:"inAppPurchaseType,omitempty"`
	} `url:"filter,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Sort    *[]string `url:"sort,omitempty"`
}

type GetInAppPurchaseOptions struct {
	Fields *struct {
		InAppPurchases *[]string `url:"inAppPurchases,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *struct {
		Apps *int `url:"apps,omitempty"`
	} `url:"limit,omitempty"`
}

// ListApps finds and lists apps added in App Store Connect.
func (s *Service) ListApps(params *ListAppsOptions) (*AppsResponse, *common.Response, error) {
	res := new(AppsResponse)
	resp, err := s.Get("apps", params, res)
	return res, resp, err
}

// GetApp gets information about a specific app.
func (s *Service) GetApp(id string, params *GetAppOptions) (*AppResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.Get(url, params, res)
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
func (s *Service) ListInAppPurchasesForApp(id string, params *ListInAppPurchasesOptions) (*InAppPurchasesResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/inAppPurchases", id)
	res := new(InAppPurchasesResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// GetInAppPurchase gets information about an in-app purchase.
func (s *Service) GetInAppPurchase(id string, params *GetInAppPurchaseOptions) (*InAppPurchaseResponse, *common.Response, error) {
	url := fmt.Sprintf("inAppPurchases/%s", id)
	res := new(InAppPurchaseResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}
