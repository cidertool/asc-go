package provisioning

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// Device defines model for Device.
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
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// DeviceCreateRequest defines model for DeviceCreateRequest.
type DeviceCreateRequest struct {
	Data struct {
		Attributes struct {
			Name     string           `json:"name"`
			Platform BundleIDPlatform `json:"platform"`
			UDID     string           `json:"udid"`
		} `json:"attributes"`
		Type string `json:"type"`
	} `json:"data"`
}

// DeviceUpdateRequest defines model for DeviceUpdateRequest.
type DeviceUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Name   *string `json:"name,omitempty"`
			Status *string `json:"status,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// DeviceResponse defines model for DeviceResponse.
type DeviceResponse struct {
	Data  Device               `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// DevicesResponse defines model for DevicesResponse.
type DevicesResponse struct {
	Data  []Device                  `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type ListDevicesQuery struct {
	Fields *struct {
		Devices *[]string `url:"devices,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		ID       *[]string `url:"id,omitempty"`
		Name     *[]string `url:"name,omitempty"`
		Platform *[]string `url:"platform,omitempty"`
		Status   *[]string `url:"status,omitempty"`
		UDID     *[]string `url:"udid,omitempty"`
	} `url:"filter,omitempty"`
	Limit  *int      `url:"limit,omitempty"`
	Sort   *[]string `url:"sort,omitempty"`
	Cursor *string   `url:"cursor,omitempty"`
}

type GetDeviceQuery struct {
	Fields *struct {
		Devices *[]string `url:"devices,omitempty"`
	} `url:"fields,omitempty"`
}

// CreateDevice registers a new device for app development.
func (s *Service) CreateDevice(body *DeviceCreateRequest) (*DeviceResponse, *common.Response, error) {
	res := new(DeviceResponse)
	resp, err := s.Post("devices", body, res)
	return res, resp, err
}

// ListDevices finds and lists devices registered to your team.
func (s *Service) ListDevices(params *ListDevicesQuery) (*DevicesResponse, *common.Response, error) {
	res := new(DevicesResponse)
	resp, err := s.GetWithQuery("devices", params, res)
	return res, resp, err
}

// GetDevice gets information for a specific device registered to your team.
func (s *Service) GetDevice(id string, params *GetDeviceQuery) (*DeviceResponse, *common.Response, error) {
	url := fmt.Sprintf("devices/%s", id)
	res := new(DeviceResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateDevice updates the name or status of a specific device.
func (s *Service) UpdateDevice(id string, body *DeviceUpdateRequest) (*DeviceResponse, *common.Response, error) {
	url := fmt.Sprintf("devices/%s", id)
	res := new(DeviceResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
