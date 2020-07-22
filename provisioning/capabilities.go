package provisioning

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
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
	Links internal.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// BundleIDCapabilityCreateRequest defines model for BundleIdCapabilityCreateRequest.
type BundleIDCapabilityCreateRequest struct {
	Data struct {
		Attributes struct {
			CapabilityType CapabilityType       `json:"capabilityType"`
			Settings       *[]CapabilitySetting `json:"settings,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			BundleID struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"bundleId"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// BundleIDCapabilityUpdateRequest defines model for BundleIdCapabilityUpdateRequest.
type BundleIDCapabilityUpdateRequest struct {
	Data struct {
		Attributes *struct {
			CapabilityType *CapabilityType      `json:"capabilityType,omitempty"`
			Settings       *[]CapabilitySetting `json:"settings,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BundleIDCapabilityResponse defines model for BundleIdCapabilityResponse.
type BundleIDCapabilityResponse struct {
	Data  BundleIDCapability     `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// BundleIDCapabilitiesResponse defines model for BundleIdCapabilitiesResponse.
type BundleIDCapabilitiesResponse struct {
	Data  []BundleIDCapability        `json:"data"`
	Links internal.PagedDocumentLinks `json:"links"`
	Meta  *internal.PagingInformation `json:"meta,omitempty"`
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
func (s *Service) EnableCapability(body *BundleIDCapabilityCreateRequest) (*BundleIDCapabilityResponse, *internal.Response, error) {
	res := new(BundleIDCapabilityResponse)
	resp, err := s.Patch("bundleIdCapabilities", body, res)
	return res, resp, err
}

// DisableCapability disables a capability for a bundle ID.
func (s *Service) DisableCapability(id string) (*internal.Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	return s.Delete(url, nil)
}

// UpdateCapability updates the configuration of a specific capability.
func (s *Service) UpdateCapability(id string, body *BundleIDCapabilityUpdateRequest) (*BundleIDCapabilityResponse, *internal.Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	res := new(BundleIDCapabilityResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
