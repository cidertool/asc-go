package asc

import (
	"context"
	"fmt"
)

// CapabilityType defines model for CapabilityType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilitytype
type CapabilityType string

// List of CapabilityType
const (
	CapabilityTypeAccessWifiInformation          CapabilityType = "ACCESS_WIFI_INFORMATION"
	CapabilityTypeAppleIDAuth                    CapabilityType = "APPLE_ID_AUTH"
	CapabilityTypeApplePay                       CapabilityType = "APPLE_PAY"
	CapabilityTypeAppGroups                      CapabilityType = "APP_GROUPS"
	CapabilityTypeAssociatedDomains              CapabilityType = "ASSOCIATED_DOMAINS"
	CapabilityTypeAutoFillCredentialProvider     CapabilityType = "AUTOFILL_CREDENTIAL_PROVIDER"
	CapabilityTypeClassKit                       CapabilityType = "CLASSKIT"
	CapabilityTypeCoreMediaHLSLowLatency         CapabilityType = "COREMEDIA_HLS_LOW_LATENCY"
	CapabilityTypeDataProtection                 CapabilityType = "DATA_PROTECTION"
	CapabilityTypeGameCenter                     CapabilityType = "GAME_CENTER"
	CapabilityTypeHealthKit                      CapabilityType = "HEALTHKIT"
	CapabilityTypeHomeKit                        CapabilityType = "HOMEKIT"
	CapabilityTypeHotSpot                        CapabilityType = "HOT_SPOT"
	CapabilityTypeiCloud                         CapabilityType = "ICLOUD"
	CapabilityTypeInterAppAudio                  CapabilityType = "INTER_APP_AUDIO"
	CapabilityTypeInAppPurchase                  CapabilityType = "IN_APP_PURCHASE"
	CapabilityTypeMaps                           CapabilityType = "MAPS"
	CapabilityTypeMultipath                      CapabilityType = "MULTIPATH"
	CapabilityTypeNetworkCustomProtocol          CapabilityType = "NETWORK_CUSTOM_PROTOCOL"
	CapabilityTypeNetworkExtensions              CapabilityType = "NETWORK_EXTENSIONS"
	CapabilityTypeNFCTagReading                  CapabilityType = "NFC_TAG_READING"
	CapabilityTypePersonalVPN                    CapabilityType = "PERSONAL_VPN"
	CapabilityTypePushNotifications              CapabilityType = "PUSH_NOTIFICATIONS"
	CapabilityTypeSiriKit                        CapabilityType = "SIRIKIT"
	CapabilityTypeSystemExtensionInstall         CapabilityType = "SYSTEM_EXTENSION_INSTALL"
	CapabilityTypeUserManagement                 CapabilityType = "USER_MANAGEMENT"
	CapabilityTypeWallet                         CapabilityType = "WALLET"
	CapabilityTypeWirelessAccessoryConfiguration CapabilityType = "WIRELESS_ACCESSORY_CONFIGURATION"
)

// BundleIDCapability defines model for BundleIdCapability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapability
type BundleIDCapability struct {
	Attributes *struct {
		CapabilityType *CapabilityType      `json:"capabilityType,omitempty"`
		Settings       *[]CapabilitySetting `json:"settings,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// BundleIDCapabilityCreateRequest defines model for BundleIdCapabilityCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest
type BundleIDCapabilityCreateRequest struct {
	Attributes    BundleIDCapabilityCreateRequestAttributes    `json:"attributes"`
	Relationships BundleIDCapabilityCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// BundleIDCapabilityCreateRequestAttributes are attributes for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/attributes
type BundleIDCapabilityCreateRequestAttributes struct {
	CapabilityType CapabilityType       `json:"capabilityType"`
	Settings       *[]CapabilitySetting `json:"settings,omitempty"`
}

// BundleIDCapabilityCreateRequestRelationships are relationships for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/relationships
type BundleIDCapabilityCreateRequestRelationships struct {
	BundleID struct {
		Data RelationshipsData `json:"data"`
	} `json:"bundleId"`
}

// BundleIDCapabilityUpdateRequest defines model for BundleIdCapabilityUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest
type BundleIDCapabilityUpdateRequest struct {
	Attributes *BundleIDCapabilityUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// BundleIDCapabilityUpdateRequestAttributes are attributes for BundleIDCapabilityUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest/data/attributes
type BundleIDCapabilityUpdateRequestAttributes struct {
	CapabilityType *CapabilityType      `json:"capabilityType,omitempty"`
	Settings       *[]CapabilitySetting `json:"settings,omitempty"`
}

// BundleIDCapabilityResponse defines model for BundleIdCapabilityResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityresponse
type BundleIDCapabilityResponse struct {
	Data  BundleIDCapability `json:"data"`
	Links DocumentLinks      `json:"links"`
}

// BundleIDCapabilitiesResponse defines model for BundleIdCapabilitiesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitiesresponse
type BundleIDCapabilitiesResponse struct {
	Data  []BundleIDCapability `json:"data"`
	Links PagedDocumentLinks   `json:"links"`
	Meta  *PagingInformation   `json:"meta,omitempty"`
}

// CapabilityOption defines model for CapabilityOption.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilityoption
type CapabilityOption struct {
	Description      *string `json:"description,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
	EnabledByDefault *bool   `json:"enabledByDefault,omitempty"`
	Key              *string `json:"key,omitempty"`
	Name             *string `json:"name,omitempty"`
	SupportsWildcard *bool   `json:"supportsWildcard,omitempty"`
}

// CapabilitySetting defines model for CapabilitySetting.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilitysetting
type CapabilitySetting struct {
	AllowedInstances *string             `json:"allowedInstances,omitempty"`
	Description      *string             `json:"description,omitempty"`
	EnabledByDefault *bool               `json:"enabledByDefault,omitempty"`
	Key              *string             `json:"key,omitempty"`
	MinInstances     *int                `json:"minInstances,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Options          *[]CapabilityOption `json:"options,omitempty"`
	Visible          *bool               `json:"visible,omitempty"`
}

// EnableCapability enables a capability for a bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enable_a_capability
func (s *ProvisioningService) EnableCapability(ctx context.Context, body *BundleIDCapabilityCreateRequest) (*BundleIDCapabilityResponse, *Response, error) {
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.patch(ctx, "bundleIdCapabilities", body, res)
	return res, resp, err
}

// DisableCapability disables a capability for a bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/disable_a_capability
func (s *ProvisioningService) DisableCapability(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	return s.client.delete(ctx, url, nil)
}

// UpdateCapability updates the configuration of a specific capability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_capability_configuration
func (s *ProvisioningService) UpdateCapability(ctx context.Context, id string, body *BundleIDCapabilityUpdateRequest) (*BundleIDCapabilityResponse, *Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}
