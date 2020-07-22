package publishing

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/internal"
)

// PhasedReleaseState defines model for PhasedReleaseState.
type PhasedReleaseState string

const (
	// Inactive is a representation of the INACTIVE state
	Inactive PhasedReleaseState = "INACTIVE"
	// Active is a representation of the ACTIVE state
	Active PhasedReleaseState = "ACTIVE"
	// Paused is a representation of the PAUSED state
	Paused PhasedReleaseState = "PAUSED"
	// Complete is a representation of the COMPLETE state
	Complete PhasedReleaseState = "COMPLETE"
)

// AppStoreVersionPhasedRelease defines model for AppStoreVersionPhasedRelease.
type AppStoreVersionPhasedRelease struct {
	Attributes *struct {
		CurrentDayNumber   *int                `json:"currentDayNumber,omitempty"`
		PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
		StartDate          *time.Time          `json:"startDate,omitempty"`
		TotalPauseDuration *int                `json:"totalPauseDuration,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string                 `json:"id"`
	Links internal.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// AppStoreVersionPhasedReleaseCreateRequest defines model for AppStoreVersionPhasedReleaseCreateRequest.
type AppStoreVersionPhasedReleaseCreateRequest struct {
	Data struct {
		Attributes *struct {
			PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			AppStoreVersion struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionPhasedReleaseResponse defines model for AppStoreVersionPhasedReleaseResponse.
type AppStoreVersionPhasedReleaseResponse struct {
	Data  AppStoreVersionPhasedRelease `json:"data"`
	Links internal.DocumentLinks       `json:"links"`
}

// AppStoreVersionPhasedReleaseUpdateRequest defines model for AppStoreVersionPhasedReleaseUpdateRequest.
type AppStoreVersionPhasedReleaseUpdateRequest struct {
	Data struct {
		Attributes *struct {
			PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery are query options for GetAppStoreVersionPhasedReleaseForAppStoreVersion
type GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery struct {
	FieldsAppStoreVersionPhasedReleases *[]string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
}

// CreatePhasedRelease enables phased release for an App Store version.
func (s *Service) CreatePhasedRelease(body *AppStoreVersionPhasedReleaseCreateRequest) (*AppStoreVersionPhasedReleaseResponse, *internal.Response, error) {
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.Post("appStoreVersionPhasedReleases", body, res)
	return res, resp, err
}

// UpdatePhasedRelease pauses or resumes a phased release, or immediately release an app.
func (s *Service) UpdatePhasedRelease(id string, body *AppStoreVersionPhasedReleaseUpdateRequest) (*AppStoreVersionPhasedReleaseResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeletePhasedRelease cancels a planned phased release that has not been started.
func (s *Service) DeletePhasedRelease(id string) (*internal.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	return s.Delete(url, nil)
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersion reads the phased release status and configuration for a version with phased release enabled.
func (s *Service) GetAppStoreVersionPhasedReleaseForAppStoreVersion(id string, params *GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery) (*AppStoreVersionPhasedReleaseResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionPhasedRelease", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
