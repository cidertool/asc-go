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

// CapabilityType defines model for CapabilityType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilitytype
type CapabilityType string

const (
	// CapabilityTypeAccessWifiInformation is a capability type for AccessWifiInformation.
	CapabilityTypeAccessWifiInformation CapabilityType = "ACCESS_WIFI_INFORMATION"
	// CapabilityTypeAppleIDAuth is a capability type for AppleIDAuth.
	CapabilityTypeAppleIDAuth CapabilityType = "APPLE_ID_AUTH"
	// CapabilityTypeApplePay is a capability type for ApplePay.
	CapabilityTypeApplePay CapabilityType = "APPLE_PAY"
	// CapabilityTypeAppGroups is a capability type for AppGroups.
	CapabilityTypeAppGroups CapabilityType = "APP_GROUPS"
	// CapabilityTypeAssociatedDomains is a capability type for AssociatedDomains.
	CapabilityTypeAssociatedDomains CapabilityType = "ASSOCIATED_DOMAINS"
	// CapabilityTypeAutoFillCredentialProvider is a capability type for AutoFillCredentialProvider.
	CapabilityTypeAutoFillCredentialProvider CapabilityType = "AUTOFILL_CREDENTIAL_PROVIDER"
	// CapabilityTypeClassKit is a capability type for ClassKit.
	CapabilityTypeClassKit CapabilityType = "CLASSKIT"
	// CapabilityTypeCoreMediaHLSLowLatency is a capability type for CoreMediaHLSLowLatency.
	CapabilityTypeCoreMediaHLSLowLatency CapabilityType = "COREMEDIA_HLS_LOW_LATENCY"
	// CapabilityTypeDataProtection is a capability type for DataProtection.
	CapabilityTypeDataProtection CapabilityType = "DATA_PROTECTION"
	// CapabilityTypeGameCenter is a capability type for GameCenter.
	CapabilityTypeGameCenter CapabilityType = "GAME_CENTER"
	// CapabilityTypeHealthKit is a capability type for HealthKit.
	CapabilityTypeHealthKit CapabilityType = "HEALTHKIT"
	// CapabilityTypeHomeKit is a capability type for HomeKit.
	CapabilityTypeHomeKit CapabilityType = "HOMEKIT"
	// CapabilityTypeHotSpot is a capability type for HotSpot.
	CapabilityTypeHotSpot CapabilityType = "HOT_SPOT"
	// CapabilityTypeiCloud is a capability type for iCloud.
	CapabilityTypeiCloud CapabilityType = "ICLOUD"
	// CapabilityTypeInterAppAudio is a capability type for InterAppAudio.
	CapabilityTypeInterAppAudio CapabilityType = "INTER_APP_AUDIO"
	// CapabilityTypeInAppPurchase is a capability type for InAppPurchase.
	CapabilityTypeInAppPurchase CapabilityType = "IN_APP_PURCHASE"
	// CapabilityTypeMaps is a capability type for Maps.
	CapabilityTypeMaps CapabilityType = "MAPS"
	// CapabilityTypeMultipath is a capability type for Multipath.
	CapabilityTypeMultipath CapabilityType = "MULTIPATH"
	// CapabilityTypeNetworkCustomProtocol is a capability type for NetworkCustomProtocol.
	CapabilityTypeNetworkCustomProtocol CapabilityType = "NETWORK_CUSTOM_PROTOCOL"
	// CapabilityTypeNetworkExtensions is a capability type for NetworkExtensions.
	CapabilityTypeNetworkExtensions CapabilityType = "NETWORK_EXTENSIONS"
	// CapabilityTypeNFCTagReading is a capability type for NFCTagReading.
	CapabilityTypeNFCTagReading CapabilityType = "NFC_TAG_READING"
	// CapabilityTypePersonalVPN is a capability type for PersonalVPN.
	CapabilityTypePersonalVPN CapabilityType = "PERSONAL_VPN"
	// CapabilityTypePushNotifications is a capability type for PushNotifications.
	CapabilityTypePushNotifications CapabilityType = "PUSH_NOTIFICATIONS"
	// CapabilityTypeSiriKit is a capability type for SiriKit.
	CapabilityTypeSiriKit CapabilityType = "SIRIKIT"
	// CapabilityTypeSystemExtensionInstall is a capability type for SystemExtensionInstall.
	CapabilityTypeSystemExtensionInstall CapabilityType = "SYSTEM_EXTENSION_INSTALL"
	// CapabilityTypeUserManagement is a capability type for UserManagement.
	CapabilityTypeUserManagement CapabilityType = "USER_MANAGEMENT"
	// CapabilityTypeWallet is a capability type for Wallet.
	CapabilityTypeWallet CapabilityType = "WALLET"
	// CapabilityTypeWirelessAccessoryConfiguration is a capability type for WirelessAccessoryConfiguration.
	CapabilityTypeWirelessAccessoryConfiguration CapabilityType = "WIRELESS_ACCESSORY_CONFIGURATION"
)

// BundleIDCapability defines model for BundleIdCapability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapability
type BundleIDCapability struct {
	Attributes *BundleIDCapabilityAttributes `json:"attributes,omitempty"`
	ID         string                        `json:"id"`
	Links      ResourceLinks                 `json:"links"`
	Type       string                        `json:"type"`
}

// BundleIDCapabilityAttributes defines model for BundleIdCapability.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapability/attributes
type BundleIDCapabilityAttributes struct {
	CapabilityType *CapabilityType     `json:"capabilityType,omitempty"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
}

// bundleIDCapabilityCreateRequest defines model for BundleIdCapabilityCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data
type bundleIDCapabilityCreateRequest struct {
	Attributes    bundleIDCapabilityCreateRequestAttributes    `json:"attributes"`
	Relationships bundleIDCapabilityCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// bundleIDCapabilityCreateRequestAttributes are attributes for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/attributes
type bundleIDCapabilityCreateRequestAttributes struct {
	CapabilityType CapabilityType      `json:"capabilityType"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
}

// bundleIDCapabilityCreateRequestRelationships are relationships for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/relationships
type bundleIDCapabilityCreateRequestRelationships struct {
	BundleID relationshipDeclaration `json:"bundleId"`
}

// BundleIDCapabilityUpdateRequest defines model for BundleIdCapabilityUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest/data
type bundleIDCapabilityUpdateRequest struct {
	Attributes *bundleIDCapabilityUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// BundleIDCapabilityUpdateRequestAttributes are attributes for BundleIDCapabilityUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest/data/attributes
type bundleIDCapabilityUpdateRequestAttributes struct {
	CapabilityType *CapabilityType     `json:"capabilityType,omitempty"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
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
	AllowedInstances *string            `json:"allowedInstances,omitempty"`
	Description      *string            `json:"description,omitempty"`
	EnabledByDefault *bool              `json:"enabledByDefault,omitempty"`
	Key              *string            `json:"key,omitempty"`
	MinInstances     *int               `json:"minInstances,omitempty"`
	Name             *string            `json:"name,omitempty"`
	Options          []CapabilityOption `json:"options,omitempty"`
	Visible          *bool              `json:"visible,omitempty"`
}

// EnableCapability enables a capability for a bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enable_a_capability
func (s *ProvisioningService) EnableCapability(ctx context.Context, capabilityType CapabilityType, capabilitySettings []CapabilitySetting, bundleIDRelationship string) (*BundleIDCapabilityResponse, *Response, error) {
	req := bundleIDCapabilityCreateRequest{
		Attributes: bundleIDCapabilityCreateRequestAttributes{
			CapabilityType: capabilityType,
			Settings:       capabilitySettings,
		},
		Relationships: bundleIDCapabilityCreateRequestRelationships{
			BundleID: relationshipDeclaration{
				Data: RelationshipData{
					ID:   bundleIDRelationship,
					Type: "bundleIds",
				},
			},
		},
		Type: "bundleIdCapabilities",
	}
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.patch(ctx, "bundleIdCapabilities", newRequestBody(req), res)

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
func (s *ProvisioningService) UpdateCapability(ctx context.Context, id string, capabilityType *CapabilityType, settings []CapabilitySetting) (*BundleIDCapabilityResponse, *Response, error) {
	req := bundleIDCapabilityUpdateRequest{
		ID:   id,
		Type: "bundleIdCapabilities",
	}

	if capabilityType != nil || settings != nil {
		req.Attributes = &bundleIDCapabilityUpdateRequestAttributes{
			CapabilityType: capabilityType,
			Settings:       settings,
		}
	}

	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
