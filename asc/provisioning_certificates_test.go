package asc

import (
	"context"
	"testing"
)

func TestCreateCertificate(t *testing.T) {
	testEndpointWithResponse(t, "{}", &CertificateResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateCertificate(ctx, &CertificateCreateRequest{})
	})
}

func TestListCertificates(t *testing.T) {
	testEndpointWithResponse(t, "{}", &CertificatesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCertificates(ctx, &ListCertificatesQuery{})
	})
}

func TestGetCertificate(t *testing.T) {
	testEndpointWithResponse(t, "{}", &CertificateResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetCertificate(ctx, "10", &GetCertificateQuery{})
	})
}

func TestRevokeCertificate(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.RevokeCertificate(ctx, "10")
	})
}
