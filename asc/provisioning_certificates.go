package asc

import (
	"fmt"
	"net/http"
	"time"
)

// CertificateType defines model for CertificateType.
type CertificateType string

// List of CertificateType
const (
	CertificateTypeDeveloperIDApplication   CertificateType = "DEVELOPER_ID_APPLICATION"
	CertificateTypeDeveloperIDKext          CertificateType = "DEVELOPER_ID_KEXT"
	CertificateTypeDevelopment              CertificateType = "DEVELOPMENT"
	CertificateTypeDistribution             CertificateType = "DISTRIBUTION"
	CertificateTypeiOSDevelopment           CertificateType = "IOS_DEVELOPMENT"
	CertificateTypeiOSDistribution          CertificateType = "IOS_DISTRIBUTION"
	CertificateTypeMacAppDevelopment        CertificateType = "MAC_APP_DEVELOPMENT"
	CertificateTypeMacAppDistribution       CertificateType = "MAC_APP_DISTRIBUTION"
	CertificateTypeMacInstallerDistribution CertificateType = "MAC_INSTALLER_DISTRIBUTION"
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
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// CertificateCreateRequest defines model for CertificateCreateRequest.
type CertificateCreateRequest struct {
	Attributes CertificateCreateRequestAttributes `json:"attributes"`
	Type       string                             `json:"type"`
}

// CertificateCreateRequestAttributes are attributes for CertificateCreateRequest
type CertificateCreateRequestAttributes struct {
	CertificateType CertificateType `json:"certificateType"`
	CsrContent      string          `json:"csrContent"`
}

// CertificateResponse defines model for CertificateResponse.
type CertificateResponse struct {
	Data  Certificate   `json:"data"`
	Links DocumentLinks `json:"links"`
}

// CertificatesResponse defines model for CertificatesResponse.
type CertificatesResponse struct {
	Data  []Certificate      `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListCertificatesQuery are query options for ListCertificates
type ListCertificatesQuery struct {
	FieldsCertificates    []string `url:"fields[certificates],omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	Include               []string `url:"include,omitempty"`
	Sort                  []string `url:"sort,omitempty"`
	FilterID              []string `url:"filter[id],omitempty"`
	FilterSerialNumber    []string `url:"filter[serialNumber],omitempty"`
	FilterCertificateType []string `url:"filter[certificateType],omitempty"`
	FilterDisplayName     []string `url:"filter[displayName],omitempty"`
	Cursor                string   `url:"cursor,omitempty"`
}

// GetCertificateQuery are query options for GetCertificate
type GetCertificateQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
}

// CreateCertificate creates a new certificate using a certificate signing request.
func (s *ProvisioningService) CreateCertificate(body *CertificateCreateRequest) (*CertificateResponse, *http.Response, error) {
	res := new(CertificateResponse)
	resp, err := s.client.post("certificates", body, res)
	return res, resp, err
}

// ListCertificates finds and lists certificates and download their data.
func (s *ProvisioningService) ListCertificates(params *ListCertificatesQuery) (*CertificatesResponse, *http.Response, error) {
	res := new(CertificatesResponse)
	resp, err := s.client.get("certificates", params, res)
	return res, resp, err
}

// GetCertificate gets information about a certificate and download the certificate data.
func (s *ProvisioningService) GetCertificate(id string, params *GetCertificateQuery) (*CertificateResponse, *http.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	res := new(CertificateResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// RevokeCertificate revokes a lost, stolen, compromised, or expiring signing certificate.
func (s *ProvisioningService) RevokeCertificate(id string) (*http.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	return s.client.delete(url, nil)
}
