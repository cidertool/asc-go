package provisioning

import (
	"fmt"
	"net/http"

	"github.com/aaronsky/asc-go/internal/services"
)

// CapabilityType defines model for CapabilityType.
type CapabilityType string

// List of CapabilityType
const (
	AccessWifiInformation          CapabilityType = "ACCESS_WIFI_INFORMATION"
	AppleIDAuth                    CapabilityType = "APPLE_ID_AUTH"
	ApplePay                       CapabilityType = "APPLE_PAY"
	AppGroups                      CapabilityType = "APP_GROUPS"
	AssociatedDomains              CapabilityType = "ASSOCIATED_DOMAINS"
	AutoFillCredentialProvider     CapabilityType = "AUTOFILL_CREDENTIAL_PROVIDER"
	ClassKit                       CapabilityType = "CLASSKIT"
	CoreMediaHLSLowLatency         CapabilityType = "COREMEDIA_HLS_LOW_LATENCY"
	DataProtection                 CapabilityType = "DATA_PROTECTION"
	GameCenter                     CapabilityType = "GAME_CENTER"
	HealthKit                      CapabilityType = "HEALTHKIT"
	HomeKit                        CapabilityType = "HOMEKIT"
	HotSpot                        CapabilityType = "HOT_SPOT"
	ICloud                         CapabilityType = "ICLOUD"
	InterAppAudio                  CapabilityType = "INTER_APP_AUDIO"
	InAppPurchase                  CapabilityType = "IN_APP_PURCHASE"
	Maps                           CapabilityType = "MAPS"
	Multipath                      CapabilityType = "MULTIPATH"
	NetworkCustomProtocol          CapabilityType = "NETWORK_CUSTOM_PROTOCOL"
	NetworkExtensions              CapabilityType = "NETWORK_EXTENSIONS"
	NFCTagReading                  CapabilityType = "NFC_TAG_READING"
	PersonalVPN                    CapabilityType = "PERSONAL_VPN"
	PushNotifications              CapabilityType = "PUSH_NOTIFICATIONS"
	SiriKit                        CapabilityType = "SIRIKIT"
	SystemExtensionInstall         CapabilityType = "SYSTEM_EXTENSION_INSTALL"
	UserManagement                 CapabilityType = "USER_MANAGEMENT"
	Wallet                         CapabilityType = "WALLET"
	WirelessAccessoryConfiguration CapabilityType = "WIRELESS_ACCESSORY_CONFIGURATION"
)

// BundleIDCapability defines model for BundleIdCapability.
type BundleIDCapability struct {
	Attributes *struct {
		CapabilityType *CapabilityType      `json:"capabilityType,omitempty"`
		Settings       *[]CapabilitySetting `json:"settings,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string                 `json:"id"`
	Links services.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// BundleIDCapabilityCreateRequest defines model for BundleIdCapabilityCreateRequest.
type BundleIDCapabilityCreateRequest struct {
	Attributes    BundleIDCapabilityCreateRequestAttributes    `json:"attributes"`
	Relationships BundleIDCapabilityCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// BundleIDCapabilityCreateRequestAttributes are attributes for BundleIDCapabilityCreateRequest
type BundleIDCapabilityCreateRequestAttributes struct {
	CapabilityType CapabilityType       `json:"capabilityType"`
	Settings       *[]CapabilitySetting `json:"settings,omitempty"`
}

// BundleIDCapabilityCreateRequestRelationships are attributes for BundleIDCapabilityCreateRequest
type BundleIDCapabilityCreateRequestRelationships struct {
	BundleID struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"bundleId"`
}

// BundleIDCapabilityUpdateRequest defines model for BundleIdCapabilityUpdateRequest.
type BundleIDCapabilityUpdateRequest struct {
	Attributes *BundleIDCapabilityUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// BundleIDCapabilityUpdateRequestAttributes are attributes for BundleIDCapabilityUpdateRequest
type BundleIDCapabilityUpdateRequestAttributes struct {
	CapabilityType *CapabilityType      `json:"capabilityType,omitempty"`
	Settings       *[]CapabilitySetting `json:"settings,omitempty"`
}

// BundleIDCapabilityResponse defines model for BundleIdCapabilityResponse.
type BundleIDCapabilityResponse struct {
	Data  BundleIDCapability     `json:"data"`
	Links services.DocumentLinks `json:"links"`
}

// BundleIDCapabilitiesResponse defines model for BundleIdCapabilitiesResponse.
type BundleIDCapabilitiesResponse struct {
	Data  []BundleIDCapability        `json:"data"`
	Links services.PagedDocumentLinks `json:"links"`
	Meta  *services.PagingInformation `json:"meta,omitempty"`
}

// CapabilityOption defines model for CapabilityOption.
type CapabilityOption struct {
	Description      *string `json:"description,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
	EnabledByDefault *bool   `json:"enabledByDefault,omitempty"`
	Key              *string `json:"key,omitempty"`
	Name             *string `json:"name,omitempty"`
	SupportsWildcard *bool   `json:"supportsWildcard,omitempty"`
}

// CapabilitySetting defines model for CapabilitySetting.
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
func (s *Service) EnableCapability(body *BundleIDCapabilityCreateRequest) (*BundleIDCapabilityResponse, *http.Response, error) {
	res := new(BundleIDCapabilityResponse)
	resp, err := s.Patch("bundleIdCapabilities", body, res)
	return res, resp, err
}

// DisableCapability disables a capability for a bundle ID.
func (s *Service) DisableCapability(id string) (*http.Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	return s.Delete(url, nil)
}

// UpdateCapability updates the configuration of a specific capability.
func (s *Service) UpdateCapability(id string, body *BundleIDCapabilityUpdateRequest) (*BundleIDCapabilityResponse, *http.Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	res := new(BundleIDCapabilityResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
