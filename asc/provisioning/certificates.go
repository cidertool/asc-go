package provisioning

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// CertificateType defines model for CertificateType.
type CertificateType string

// List of CertificateType
const (
	DeveloperIDApplication   CertificateType = "DEVELOPER_ID_APPLICATION"
	DeveloperIDKext          CertificateType = "DEVELOPER_ID_KEXT"
	Development              CertificateType = "DEVELOPMENT"
	Distribution             CertificateType = "DISTRIBUTION"
	IOSDevelopment           CertificateType = "IOS_DEVELOPMENT"
	IOSDistribution          CertificateType = "IOS_DISTRIBUTION"
	MacAppDevelopment        CertificateType = "MAC_APP_DEVELOPMENT"
	MacAppDistribution       CertificateType = "MAC_APP_DISTRIBUTION"
	MacInstallerDistribution CertificateType = "MAC_INSTALLER_DISTRIBUTION"
)

// Certificate defines model for Certificate.
type Certificate struct {
	Attributes *struct {
		CertificateContent *string           `json:"certificateContent,omitempty"`
		CertificateType    *CertificateType  `json:"certificateType,omitempty"`
		DisplayName        *string           `json:"displayName,omitempty"`
		ExpirationDate     *time.Time        `json:"expirationDate,omitempty"`
		Name               *string           `json:"name,omitempty"`
		Platform           *BundleIDPlatform `json:"platform,omitempty"`
		SerialNumber       *string           `json:"serialNumber,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string               `json:"id"`
	Links common.ResourceLinks `json:"links"`
	Type  string               `json:"type"`
}

// CertificateCreateRequest defines model for CertificateCreateRequest.
type CertificateCreateRequest struct {
	Data struct {
		Attributes struct {
			CertificateType CertificateType `json:"certificateType"`
			CsrContent      string          `json:"csrContent"`
		} `json:"attributes"`
		Type string `json:"type"`
	} `json:"data"`
}

// CertificateResponse defines model for CertificateResponse.
type CertificateResponse struct {
	Data  Certificate          `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// CertificatesResponse defines model for CertificatesResponse.
type CertificatesResponse struct {
	Data  []Certificate             `json:"data"`
	Links common.PagedDocumentLinks `json:"links"`
	Meta  *common.PagingInformation `json:"meta,omitempty"`
}

type ListCertificatesOptions struct {
	Fields *struct {
		Certificates *[]string `url:"certificates,omitempty"`
	} `url:"fields,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Sort    *[]string `url:"sort,omitempty"`
	Filter  *struct {
		ID              *[]string `url:"id,omitempty"`
		SerialNumber    *[]string `url:"serialNumber,omitempty"`
		CertificateType *[]string `url:"certificateType,omitempty"`
		DisplayName     *[]string `url:"displayName,omitempty"`
	} `url:"filter,omitempty"`
}

type GetCertificateOptions struct {
	Fields *struct {
		Certificates *[]string `url:"certificates,omitempty"`
	} `url:"fields,omitempty"`
}

// CreateCertificate creates a new certificate using a certificate signing request.
func (s *Service) CreateCertificate(body *CertificateCreateRequest) (*CertificateResponse, *common.Response, error) {
	res := new(CertificateResponse)
	resp, err := s.Post("certificates", body, res)
	return res, resp, err
}

// ListCertificates finds and lists certificates and download their data.
func (s *Service) ListCertificates(params *ListCertificatesOptions) (*CertificatesResponse, *common.Response, error) {
	res := new(CertificatesResponse)
	resp, err := s.Get("certificates", params, res)
	return res, resp, err
}

// GetCertificate gets information about a certificate and download the certificate data.
func (s *Service) GetCertificate(id string, params *GetCertificateOptions) (*CertificateResponse, *common.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	res := new(CertificateResponse)
	resp, err := s.Get(url, params, res)
	return res, resp, err
}

// RevokeCertificate revokes a lost, stolen, compromised, or expiring signing certificate.
func (s *Service) RevokeCertificate(id string) (*common.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	return s.Delete(url, nil)
}
