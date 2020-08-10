package asc

import (
	"context"
	"fmt"
)

// AppsService handles communication with build-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/apps
// https://developer.apple.com/documentation/appstoreconnectapi/app_metadata
type AppsService service

// Platform defines model for Platform.
//
// https://developer.apple.com/documentation/appstoreconnectapi/platform
type Platform string

// List of Platform
const (
	PlatformIOS   Platform = "IOS"
	PlatformMACOS Platform = "MAC_OS"
	PlatformTVOS  Platform = "TV_OS"
)

// App defines model for App.
//
// https://developer.apple.com/documentation/appstoreconnectapi/app
type App struct {
	Attributes    *AppAttributes    `json:"attributes,omitempty"`
	ID            string            `json:"id"`
	Links         ResourceLinks     `json:"links"`
	Relationships *AppRelationships `json:"relationships,omitempty"`
	Type          string            `json:"type"`
}

// AppAttributes defines model for App.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/app/attributes
type AppAttributes struct {
	AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
	BundleID                  *string `json:"bundleId,omitempty"`
	ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
	IsOrEverWasMadeForKids    *bool   `json:"isOrEverWasMadeForKids,omitempty"`
	Name                      *string `json:"name,omitempty"`
	PrimaryLocale             *string `json:"primaryLocale,omitempty"`
	Sku                       *string `json:"sku,omitempty"`
}

// AppRelationships defines model for App.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/app/relationships
type AppRelationships struct {
	AppInfos                  *PagedRelationship `json:"appInfos,omitempty"`
	AppStoreVersions          *PagedRelationship `json:"appStoreVersions,omitempty"`
	AvailableTerritories      *PagedRelationship `json:"availableTerritories,omitempty"`
	BetaAppLocalizations      *PagedRelationship `json:"betaAppLocalizations,omitempty"`
	BetaAppReviewDetail       *Relationship      `json:"betaAppReviewDetail,omitempty"`
	BetaGroups                *PagedRelationship `json:"betaGroups,omitempty"`
	BetaLicenseAgreement      *Relationship      `json:"betaLicenseAgreement,omitempty"`
	Builds                    *PagedRelationship `json:"builds,omitempty"`
	EndUserLicenseAgreement   *Relationship      `json:"endUserLicenseAgreement,omitempty"`
	GameCenterEnabledVersions *PagedRelationship `json:"gameCenterEnabledVersions,omitempty"`
	InAppPurchases            *PagedRelationship `json:"inAppPurchases,omitempty"`
	PreOrder                  *Relationship      `json:"preOrder,omitempty"`
	PreReleaseVersions        *PagedRelationship `json:"preReleaseVersions,omitempty"`
	Prices                    *PagedRelationship `json:"prices,omitempty"`
}

// AppUpdateRequest defines model for AppUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appupdaterequest
type AppUpdateRequest struct {
	Attributes    *AppUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Relationships *AppUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

// AppUpdateRequestAttributes are attributes for AppUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appupdaterequest/data/attributes
type AppUpdateRequestAttributes struct {
	AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
	BundleID                  *string `json:"bundleId,omitempty"`
	ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
	PrimaryLocale             *string `json:"primaryLocale,omitempty"`
}

// AppUpdateRequestRelationships are relationships for AppUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appupdaterequest/data/relationships
type AppUpdateRequestRelationships struct {
	AvailableTerritories *PagedRelationshipDeclaration `json:"availableTerritories,omitempty"`
	Prices               *PagedRelationshipDeclaration `json:"prices,omitempty"`
}

// AppResponse defines model for AppResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appresponse
type AppResponse struct {
	Data     App            `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// AppsResponse defines model for AppsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appsresponse
type AppsResponse struct {
	Data     []App              `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// InAppPurchase defines model for InAppPurchase.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchase
type InAppPurchase struct {
	Attributes    *InAppPurchaseAttributes    `json:"attributes,omitempty"`
	ID            string                      `json:"id"`
	Links         ResourceLinks               `json:"links"`
	Relationships *InAppPurchaseRelationships `json:"relationships,omitempty"`
	Type          string                      `json:"type"`
}

// InAppPurchaseAttributes defines model for InAppPurchase.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchase/attributes
type InAppPurchaseAttributes struct {
	InAppPurchaseType *string `json:"inAppPurchaseType,omitempty"`
	ProductID         *string `json:"productId,omitempty"`
	ReferenceName     *string `json:"referenceName,omitempty"`
	State             *string `json:"state,omitempty"`
}

// InAppPurchaseRelationships defines model for InAppPurchase.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchase/relationships
type InAppPurchaseRelationships struct {
	Apps *PagedRelationship `json:"apps,omitempty"`
}

// InAppPurchaseResponse defines model for InAppPurchaseResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchaseresponse
type InAppPurchaseResponse struct {
	Data  InAppPurchase `json:"data"`
	Links DocumentLinks `json:"links"`
}

// InAppPurchasesResponse defines model for InAppPurchasesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/inapppurchasesresponse
type InAppPurchasesResponse struct {
	Data  []InAppPurchase    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// AppBetaTestersLinkagesRequest is a list of relationships to BetaTester objects
//
// https://developer.apple.com/documentation/appstoreconnectapi/appbetatesterslinkagesrequest
type AppBetaTestersLinkagesRequest []RelationshipData

// ListAppsQuery are query options for ListApps
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_apps
type ListAppsQuery struct {
	FieldsApps                          []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements         []string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions            []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails          []string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations          []string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                    []string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements      []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories                   []string `url:"fields[territories],omitempty"`
	FieldsAppPrices                     []string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders                  []string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                      []string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics              []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsInAppPurchases                []string `url:"fields[inAppPurchases],omitempty"`
	FilterBundleID                      []string `url:"filter[bundleId],omitempty"`
	FilterID                            []string `url:"filter[id],omitempty"`
	FilterName                          []string `url:"filter[name],omitempty"`
	FilterSKU                           []string `url:"filter[sku],omitempty"`
	FilterAppStoreVersions              []string `url:"filter[appStoreVersions],omitempty"`
	FilterAppStoreVersionsPlatform      []string `url:"filter[appStoreVersionsPlatform],omitempty"`
	FilterAppStoreVersionsAppStoreState []string `url:"filter[appStoreVersionsAppStoreState],omitempty"`
	FilterGameCenterEnabledVersions     []string `url:"filter[gameCenterEnabledVersions],omitempty"`
	Include                             []string `url:"include,omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	LimitPreReleaseVersions             int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                         int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                     int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations           int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                         int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories           int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions               int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                       int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions      int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases                 int      `url:"limit[inAppPurchases],omitempty"`
	Sort                                []string `url:"sort,omitempty"`
	ExistsGameCenterEnabledVersions     []string `url:"exists[gameCenterEnabledVersions],omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetAppQuery are query options for GetApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_information
type GetAppQuery struct {
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements     []string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions        []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails      []string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations      []string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                    []string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                []string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements  []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions          []string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories               []string `url:"fields[territories],omitempty"`
	FieldsAppPrices                 []string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders              []string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                  []string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics          []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsGameCenterEnabledVersions []string `url:"fields[gameCenterEnabledVersions],omitempty"`
	FieldsInAppPurchases            []string `url:"fields[inAppPurchases],omitempty"`
	Include                         []string `url:"include,omitempty"`
	LimitPreReleaseVersions         int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                     int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                 int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations       int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                     int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories       int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions           int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                   int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions  int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases             int      `url:"limit[inAppPurchases],omitempty"`
}

// ListInAppPurchasesQuery are query options for ListInAppPurchases
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_in-app_purchases_for_an_app
type ListInAppPurchasesQuery struct {
	FieldsApps              []string `url:"fields[apps],omitempty"`
	FieldsInAppPurchases    []string `url:"fields[inAppPurchases],omitempty"`
	FilterCanBeSubmitted    []string `url:"filter[canBeSubmitted],omitempty"`
	FilterInAppPurchaseType []string `url:"filter[inAppPurchaseType],omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Include                 []string `url:"include,omitempty"`
	Sort                    []string `url:"sort,omitempty"`
	Cursor                  string   `url:"cursor,omitempty"`
}

// GetInAppPurchaseQuery are query options for GetInAppPurchase
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_in-app_purchase_information
type GetInAppPurchaseQuery struct {
	FieldsInAppPurchases []string `url:"fields[inAppPurchases],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitApps            int      `url:"limit[apps],omitempty"`
}

func (r *AppUpdateRequest) applyTypes() {
	if r == nil {
		return
	}
	r.Type = "ageRatingDeclarations"
	if r.Relationships == nil {
		return
	}
	r.Relationships.AvailableTerritories.applyType("territories")
	r.Relationships.Prices.applyType("appPrices")
}

func (r *AppBetaTestersLinkagesRequest) applyTypes() {
	if r == nil {
		return
	}
	for _, rel := range *r {
		rel.applyType("betaTesters")
	}
}

// ListApps finds and lists apps added in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_apps
func (s *AppsService) ListApps(ctx context.Context, params *ListAppsQuery) (*AppsResponse, *Response, error) {
	res := new(AppsResponse)
	resp, err := s.client.get(ctx, "apps", params, res)
	return res, resp, err
}

// GetApp gets information about a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_information
func (s *AppsService) GetApp(ctx context.Context, id string, params *GetAppQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateApp updates app information including bundle ID, primary locale, price schedule, and global availability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app
func (s *AppsService) UpdateApp(ctx context.Context, id string, body *AppUpdateRequest) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// RemoveBetaTestersFromApp removes one or more beta testers' access to test any builds of a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_beta_testers_from_all_groups_and_builds_of_an_app
func (s *AppsService) RemoveBetaTestersFromApp(ctx context.Context, id string, linkages *AppBetaTestersLinkagesRequest) (*Response, error) {
	url := fmt.Sprintf("apps/%s/relationships/betaTesters", id)
	return s.client.delete(ctx, url, linkages)
}

// ListInAppPurchasesForApp lists the in-app purchases that are available for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_in-app_purchases_for_an_app
func (s *AppsService) ListInAppPurchasesForApp(ctx context.Context, id string, params *ListInAppPurchasesQuery) (*InAppPurchasesResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/inAppPurchases", id)
	res := new(InAppPurchasesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetInAppPurchase gets information about an in-app purchase.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_in-app_purchase_information
func (s *AppsService) GetInAppPurchase(ctx context.Context, id string, params *GetInAppPurchaseQuery) (*InAppPurchaseResponse, *Response, error) {
	url := fmt.Sprintf("inAppPurchases/%s", id)
	res := new(InAppPurchaseResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
