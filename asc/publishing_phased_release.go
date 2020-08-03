package asc

import (
	"fmt"
	"net/http"
	"time"
)

// PhasedReleaseState defines model for PhasedReleaseState.
type PhasedReleaseState string

const (
	// PhasedReleaseStateInactive is a representation of the INACTIVE state
	PhasedReleaseStateInactive PhasedReleaseState = "INACTIVE"
	// PhasedReleaseStateActive is a representation of the ACTIVE state
	PhasedReleaseStateActive PhasedReleaseState = "ACTIVE"
	// PhasedReleaseStatePaused is a representation of the PAUSED state
	PhasedReleaseStatePaused PhasedReleaseState = "PAUSED"
	// PhasedReleaseStateComplete is a representation of the COMPLETE state
	PhasedReleaseStateComplete PhasedReleaseState = "COMPLETE"
)

// AppStoreVersionPhasedRelease defines model for AppStoreVersionPhasedRelease.
type AppStoreVersionPhasedRelease struct {
	Attributes *struct {
		CurrentDayNumber   *int                `json:"currentDayNumber,omitempty"`
		PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
		StartDate          *time.Time          `json:"startDate,omitempty"`
		TotalPauseDuration *int                `json:"totalPauseDuration,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
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
		Data RelationshipsData `json:"data"`
	} `json:"appStoreVersion"`
}

// AppStoreVersionPhasedReleaseResponse defines model for AppStoreVersionPhasedReleaseResponse.
type AppStoreVersionPhasedReleaseResponse struct {
	Data  AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks                `json:"links"`
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
func (s *PublishingService) CreatePhasedRelease(body *AppStoreVersionPhasedReleaseCreateRequest) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.Post("appStoreVersionPhasedReleases", body, res)
	return res, resp, err
}

// UpdatePhasedRelease pauses or resumes a phased release, or immediately release an app.
func (s *PublishingService) UpdatePhasedRelease(id string, body *AppStoreVersionPhasedReleaseUpdateRequest) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.Patch(url, body, res)
	return res, resp, err
}

// DeletePhasedRelease cancels a planned phased release that has not been started.
func (s *PublishingService) DeletePhasedRelease(id string) (*http.Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	return s.client.Delete(url, nil)
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersion reads the phased release status and configuration for a version with phased release enabled.
func (s *PublishingService) GetAppStoreVersionPhasedReleaseForAppStoreVersion(id string, params *GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery) (*AppStoreVersionPhasedReleaseResponse, *http.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionPhasedRelease", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}
