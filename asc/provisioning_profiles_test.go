package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &ProfileResponse{}
	got, resp, err := client.Provisioning.CreateProfile(context.Background(), &ProfileCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestDeleteProfile(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Provisioning.DeleteProfile(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListProfiles(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &ProfilesResponse{}
	got, resp, err := client.Provisioning.ListProfiles(context.Background(), &ListProfilesQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetProfile(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &ProfileResponse{}
	got, resp, err := client.Provisioning.GetProfile(context.Background(), "10", &GetProfileQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetBundleIDForProfile(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDResponse{}
	got, resp, err := client.Provisioning.GetBundleIDForProfile(context.Background(), "10", &GetBundleIDForProfileQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListCertificatesInProfile(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &CertificatesResponse{}
	got, resp, err := client.Provisioning.ListCertificatesInProfile(context.Background(), "10", &ListCertificatesForProfileQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListDevicesInProfile(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &DevicesResponse{}
	got, resp, err := client.Provisioning.ListDevicesInProfile(context.Background(), "10", &ListDevicesInProfileQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
