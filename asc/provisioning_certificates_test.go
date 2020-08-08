package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCertificate(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &CertificateResponse{}
	got, resp, err := client.Provisioning.CreateCertificate(context.Background(), &CertificateCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListCertificates(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &CertificatesResponse{}
	got, resp, err := client.Provisioning.ListCertificates(context.Background(), &ListCertificatesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetCertificate(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &CertificateResponse{}
	got, resp, err := client.Provisioning.GetCertificate(context.Background(), "10", &GetCertificateQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestRevokeCertificate(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Provisioning.RevokeCertificate(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
