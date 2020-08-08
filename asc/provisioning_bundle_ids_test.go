package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDResponse{}
	got, resp, err := client.Provisioning.CreateBundleID(context.Background(), &BundleIDCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdateBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDResponse{}
	got, resp, err := client.Provisioning.UpdateBundleID(context.Background(), "10", &BundleIDUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestDeleteBundleID(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Provisioning.DeleteBundleID(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListBundleIDs(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDsResponse{}
	got, resp, err := client.Provisioning.ListBundleIDs(context.Background(), &ListBundleIDsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDResponse{}
	got, resp, err := client.Provisioning.GetBundleID(context.Background(), "10", &GetBundleIDQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppForBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppResponse{}
	got, resp, err := client.Provisioning.GetAppForBundleID(context.Background(), "10", &GetAppForBundleIDQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListProfilesForBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &ProfilesResponse{}
	got, resp, err := client.Provisioning.ListProfilesForBundleID(context.Background(), "10", &ListProfilesForBundleIDQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListCapabilitiesForBundleID(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BundleIDCapabilitiesResponse{}
	got, resp, err := client.Provisioning.ListCapabilitiesForBundleID(context.Background(), "10", &ListCapabilitiesForBundleIDQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
