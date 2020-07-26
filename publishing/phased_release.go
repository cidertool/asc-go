package publishing

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aaronsky/asc-go/internal/services"
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
	Links services.ResourceLinks `json:"links"`
	Type  string                 `json:"type"`
}

// AppStoreVersionPhasedReleaseCreateRequest defines model for AppStoreVersionPhasedReleaseCreateRequest.
type AppStoreVersionPhasedReleaseCreateRequest struct {
	Attributes    *AppStoreVersionPhasedReleaseCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships AppStoreVersionPhasedReleaseCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// AppStoreVersionPhasedReleaseCreateRequestAttributes are attributes for AppStoreVersionPhasedReleaseCreateRequest
type AppStoreVersionPhasedReleaseCreateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// AppStoreVersionPhasedReleaseCreateRequestRelationships are relationships for AppStoreVersionPhasedReleaseCreateRequest
type AppStoreVersionPhasedReleaseCreateRequestRelationships struct {
	AppStoreVersion struct {
		Data services.RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// AppStoreVersionPhasedReleaseResponse defines model for AppStoreVersionPhasedReleaseResponse.
type AppStoreVersionPhasedReleaseResponse struct {
	Data  AppStoreVersionPhasedRelease `json:"data"`
	Links services.DocumentLinks       `json:"links"`
}

// AppStoreVersionPhasedReleaseUpdateRequest defines model for AppStoreVersionPhasedReleaseUpdateRequest.
type AppStoreVersionPhasedReleaseUpdateRequest struct {
	Attributes *AppStoreVersionPhasedReleaseUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                               `json:"id"`
	Type       string                                               `json:"type"`
}

// AppStoreVersionPhasedReleaseUpdateRequestAttributes are attributes for AppStoreVersionPhasedReleaseUpdateRequest
type AppStoreVersionPhasedReleaseUpdateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery are query options for GetAppStoreVersionPhasedReleaseForAppStoreVersion
type GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery struct {
	FieldsAppStoreVersionPhasedReleases *[]string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
}

// CreatePhasedRelease enables phased release for an App Store version.
func (s *Service) CreatePhasedRelease(body *AppStoreVersionPhasedReleaseCreateRequest) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.Post("appStoreVersionPhasedReleases", body, res)
	return res, resp, err
}

// UpdatePhasedRelease pauses or resumes a phased release, or immediately release an app.
func (s *Service) UpdatePhasedRelease(id string, body *AppStoreVersionPhasedReleaseUpdateRequest) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeletePhasedRelease cancels a planned phased release that has not been started.
func (s *Service) DeletePhasedRelease(id string) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	return s.Delete(url, nil)
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersion reads the phased release status and configuration for a version with phased release enabled.
func (s *Service) GetAppStoreVersionPhasedReleaseForAppStoreVersion(id string, params *GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionPhasedRelease", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
