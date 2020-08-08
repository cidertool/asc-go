package asc

import (
	"context"
	"fmt"
	"time"
)

// Device defines model for Device.
//
// https://developer.apple.com/documentation/appstoreconnectapi/device
type Device struct {
	Attributes *struct {
		AddedDate   *time.Time        `json:"addedDate,omitempty"`
		DeviceClass *string           `json:"deviceClass,omitempty"`
		Model       *string           `json:"model,omitempty"`
		Name        *string           `json:"name,omitempty"`
		Platform    *BundleIDPlatform `json:"platform,omitempty"`
		Status      *string           `json:"status,omitempty"`
		UDID        *string           `json:"udid,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// DeviceCreateRequest defines model for DeviceCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/devicecreaterequest
type DeviceCreateRequest struct {
	Attributes DeviceCreateRequestAttributes `json:"attributes"`
	Type       string                        `json:"type"`
}

// DeviceCreateRequestAttributes are attributes for DeviceCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/devicecreaterequest/data/attributes
type DeviceCreateRequestAttributes struct {
	Name     string           `json:"name"`
	Platform BundleIDPlatform `json:"platform"`
	UDID     string           `json:"udid"`
}

// DeviceUpdateRequest defines model for DeviceUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/deviceupdaterequest
type DeviceUpdateRequest struct {
	Attributes *DeviceUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                         `json:"id"`
	Type       string                         `json:"type"`
}

// DeviceUpdateRequestAttributes are attributes for DeviceUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/deviceupdaterequest/attributes
type DeviceUpdateRequestAttributes struct {
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// DeviceResponse defines model for DeviceResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/deviceresponse
type DeviceResponse struct {
	Data  Device        `json:"data"`
	Links DocumentLinks `json:"links"`
}

// DevicesResponse defines model for DevicesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/devicesresponse
type DevicesResponse struct {
	Data  []Device           `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListDevicesQuery are query options for ListDevices
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_devices
type ListDevicesQuery struct {
	FieldsDevices  []string `url:"fields[devices],omitempty"`
	FilterID       []string `url:"filter[id],omitempty"`
	FilterName     []string `url:"filter[name],omitempty"`
	FilterPlatform []string `url:"filter[platform],omitempty"`
	FilterStatus   []string `url:"filter[status],omitempty"`
	FilterUDID     []string `url:"filter[udid],omitempty"`
	Limit          int      `url:"limit,omitempty"`
	Sort           []string `url:"sort,omitempty"`
	Cursor         string   `url:"cursor,omitempty"`
}

// GetDeviceQuery are query options for GetDevice
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_device_information
type GetDeviceQuery struct {
	FieldsDevices []string `url:"fields[devices],omitempty"`
}

// CreateDevice registers a new device for app development.
//
// https://developer.apple.com/documentation/appstoreconnectapi/register_a_new_device
func (s *ProvisioningService) CreateDevice(ctx context.Context, body *DeviceCreateRequest) (*DeviceResponse, *Response, error) {
	res := new(DeviceResponse)
	resp, err := s.client.post(ctx, "devices", body, res)
	return res, resp, err
}

// ListDevices finds and lists devices registered to your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_devices
func (s *ProvisioningService) ListDevices(ctx context.Context, params *ListDevicesQuery) (*DevicesResponse, *Response, error) {
	res := new(DevicesResponse)
	resp, err := s.client.get(ctx, "devices", params, res)
	return res, resp, err
}

// GetDevice gets information for a specific device registered to your team.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_device_information
func (s *ProvisioningService) GetDevice(ctx context.Context, id string, params *GetDeviceQuery) (*DeviceResponse, *Response, error) {
	url := fmt.Sprintf("devices/%s", id)
	res := new(DeviceResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// UpdateDevice updates the name or status of a specific device.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_registered_device
func (s *ProvisioningService) UpdateDevice(ctx context.Context, id string, body *DeviceUpdateRequest) (*DeviceResponse, *Response, error) {
	url := fmt.Sprintf("devices/%s", id)
	res := new(DeviceResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}
