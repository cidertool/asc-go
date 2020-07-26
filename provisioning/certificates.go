package provisioning

import (
	"fmt"
	"time"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
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
	ID    string              `json:"id"`
	Links types.ResourceLinks `json:"links"`
	Type  string              `json:"type"`
}

// CertificateCreateRequest defines model for CertificateCreateRequest.
type CertificateCreateRequest struct {
	Attributes CertificateCreateRequestAttributes `json:"attributes"`
	Type       string                             `json:"type"`
}

type CertificateCreateRequestAttributes struct {
	CertificateType CertificateType `json:"certificateType"`
	CsrContent      string          `json:"csrContent"`
}

// CertificateResponse defines model for CertificateResponse.
type CertificateResponse struct {
	Data  Certificate         `json:"data"`
	Links types.DocumentLinks `json:"links"`
}

// CertificatesResponse defines model for CertificatesResponse.
type CertificatesResponse struct {
	Data  []Certificate            `json:"data"`
	Links types.PagedDocumentLinks `json:"links"`
	Meta  *types.PagingInformation `json:"meta,omitempty"`
}

// ListCertificatesQuery are query options for ListCertificates
type ListCertificatesQuery struct {
	FieldsCertificates    *[]string `url:"fields[certificates],omitempty"`
	Limit                 *int      `url:"limit,omitempty"`
	Include               *[]string `url:"include,omitempty"`
	Sort                  *[]string `url:"sort,omitempty"`
	FilterID              *[]string `url:"filter[id],omitempty"`
	FilterSerialNumber    *[]string `url:"filter[serialNumber],omitempty"`
	FilterCertificateType *[]string `url:"filter[certificateType],omitempty"`
	FilterDisplayName     *[]string `url:"filter[displayName],omitempty"`
	Cursor                *string   `url:"cursor,omitempty"`
}

// GetCertificateQuery are query options for GetCertificate
type GetCertificateQuery struct {
	FieldsCertificates *[]string `url:"fields[certificates],omitempty"`
}

// CreateCertificate creates a new certificate using a certificate signing request.
func (s *Service) CreateCertificate(body *CertificateCreateRequest) (*CertificateResponse, *services.Response, error) {
	res := new(CertificateResponse)
	resp, err := s.Post("certificates", body, res)
	return res, resp, err
}

// ListCertificates finds and lists certificates and download their data.
func (s *Service) ListCertificates(params *ListCertificatesQuery) (*CertificatesResponse, *services.Response, error) {
	res := new(CertificatesResponse)
	resp, err := s.GetWithQuery("certificates", params, res)
	return res, resp, err
}

// GetCertificate gets information about a certificate and download the certificate data.
func (s *Service) GetCertificate(id string, params *GetCertificateQuery) (*CertificateResponse, *services.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	res := new(CertificateResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// RevokeCertificate revokes a lost, stolen, compromised, or expiring signing certificate.
func (s *Service) RevokeCertificate(id string) (*services.Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	return s.Delete(url, nil)
}
