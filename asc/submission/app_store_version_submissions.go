package submission

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppStoreVersionSubmission defines model for AppStoreVersionSubmission.
type AppStoreVersionSubmission struct {
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppStoreVersionSubmissionCreateRequest defines model for AppStoreVersionSubmissionCreateRequest.
type AppStoreVersionSubmissionCreateRequest struct {
	Data struct {
		Relationships struct {
			AppStoreVersion struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppStoreVersionSubmissionResponse defines model for AppStoreVersionSubmissionResponse.
type AppStoreVersionSubmissionResponse struct {
	Data  AppStoreVersionSubmission `json:"data"`
	Links common.DocumentLinks      `json:"links"`
}

type GetAppStoreVersionSubmissionForAppStoreVersionQuery struct {
	FieldsAppStoreVersions           *[]string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions *[]string `url:"fields[appStoreVersionSubmissions],omitempty"`
	Include                          *[]string `url:"include,omitempty"`
}

// CreateSubmission submits an App Store version to App Review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_submission
func (s *Service) CreateSubmission(body *AppStoreVersionSubmissionCreateRequest) (*AppStoreVersionSubmissionResponse, *common.Response, error) {
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.Post("appStoreVersionSubmissions", body, res)
	return res, resp, err
}

// DeleteSubmission removes a version from App Store review.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_submission
func (s *Service) DeleteSubmission(id string) (*common.Response, error) {
	url := fmt.Sprintf("appStoreVersionSubmissions/%s", id)
	return s.Delete(url, nil)
}

// GetAppStoreVersionSubmissionForAppStoreVersion reads the App Store Version Submission Information of an App Store Version
func (s *Service) GetAppStoreVersionSubmissionForAppStoreVersion(id string, params *GetAppStoreVersionSubmissionForAppStoreVersionQuery) (*AppStoreVersionSubmissionResponse, *common.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionSubmission", id)
	res := new(AppStoreVersionSubmissionResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}
