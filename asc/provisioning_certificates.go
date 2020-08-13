package asc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

// CertificateType defines model for CertificateType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificatetype
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificate
type Certificate struct {
	Attributes *CertificateAttributes `json:"attributes,omitempty"`
	ID         string                 `json:"id"`
	Links      ResourceLinks          `json:"links"`
	Type       string                 `json:"type"`
}

// CertificateAttributes defines model for Certificate.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificate/attributes
type CertificateAttributes struct {
	CertificateContent *string           `json:"certificateContent,omitempty"`
	CertificateType    *CertificateType  `json:"certificateType,omitempty"`
	DisplayName        *string           `json:"displayName,omitempty"`
	ExpirationDate     *time.Time        `json:"expirationDate,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Platform           *BundleIDPlatform `json:"platform,omitempty"`
	SerialNumber       *string           `json:"serialNumber,omitempty"`
}

// certificateCreateRequest defines model for certificateCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificatecreaterequest/data
type certificateCreateRequest struct {
	Attributes certificateCreateRequestAttributes `json:"attributes"`
	Type       string                             `json:"type"`
}

// certificateCreateRequestAttributes are attributes for CertificateCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificatecreaterequest/data/attributes
type certificateCreateRequestAttributes struct {
	CertificateType CertificateType `json:"certificateType"`
	CsrContent      string          `json:"csrContent"`
}

// CertificateResponse defines model for CertificateResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificateresponse
type CertificateResponse struct {
	Data  Certificate   `json:"data"`
	Links DocumentLinks `json:"links"`
}

// CertificatesResponse defines model for CertificatesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/certificatesresponse
type CertificatesResponse struct {
	Data  []Certificate      `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListCertificatesQuery are query options for ListCertificates
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_and_download_certificates
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
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_and_download_certificate_information
type GetCertificateQuery struct {
	FieldsCertificates []string `url:"fields[certificates],omitempty"`
}

// CreateCertificate creates a new certificate using a certificate signing request.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_certificate
func (s *ProvisioningService) CreateCertificate(ctx context.Context, certificateType CertificateType, csrContent io.Reader) (*CertificateResponse, *Response, error) {
	if csrContent == nil {
		return nil, nil, errors.New("no csr content provided")
	}
	csrBytes, err := ioutil.ReadAll(csrContent)
	if err != nil {
		return nil, nil, err
	}
	req := certificateCreateRequest{
		Attributes: certificateCreateRequestAttributes{
			CertificateType: certificateType,
			CsrContent:      string(csrBytes),
		},
		Type: "certificates",
	}
	res := new(CertificateResponse)
	resp, err := s.client.post(ctx, "certificates", req, res)
	return res, resp, err
}

// ListCertificates finds and lists certificates and download their data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_and_download_certificates
func (s *ProvisioningService) ListCertificates(ctx context.Context, params *ListCertificatesQuery) (*CertificatesResponse, *Response, error) {
	res := new(CertificatesResponse)
	resp, err := s.client.get(ctx, "certificates", params, res)
	return res, resp, err
}

// GetCertificate gets information about a certificate and download the certificate data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_and_download_certificate_information
func (s *ProvisioningService) GetCertificate(ctx context.Context, id string, params *GetCertificateQuery) (*CertificateResponse, *Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	res := new(CertificateResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// RevokeCertificate revokes a lost, stolen, compromised, or expiring signing certificate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/revoke_a_certificate
func (s *ProvisioningService) RevokeCertificate(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("certificates/%s", id)
	return s.client.delete(ctx, url, nil)
}
