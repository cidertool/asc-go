package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBuilds(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildsResponse{}
	got, resp, err := client.Builds.ListBuilds(context.Background(), &ListBuildsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestListBuildsForApp(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildsResponse{}
	got, resp, err := client.Builds.ListBuildsForApp(context.Background(), "10", &ListBuildsForAppQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildResponse{}
	got, resp, err := client.Builds.GetBuild(context.Background(), "10", &GetBuildQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppResponse{}
	got, resp, err := client.Builds.GetAppForBuild(context.Background(), "10", &GetAppForBuildQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppStoreVersionForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppStoreVersionResponse{}
	got, resp, err := client.Builds.GetAppStoreVersionForBuild(context.Background(), "10", &GetAppStoreVersionForBuildQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetBuildForAppStoreVersion(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildResponse{}
	got, resp, err := client.Builds.GetBuildForAppStoreVersion(context.Background(), "10", &GetBuildForAppStoreVersionQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdateBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildResponse{}
	got, resp, err := client.Builds.UpdateBuild(context.Background(), "10", &BuildUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdateAppEncryptionDeclarationForBuild(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.UpdateAppEncryptionDeclarationForBuild(context.Background(), "10", &RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCreateAccessForBetaGroupsToBuild(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.CreateAccessForBetaGroupsToBuild(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRemoveAccessForBetaGroupsFromBuild(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.RemoveAccessForBetaGroupsFromBuild(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCreateAccessForIndividualTestersToBuild(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.CreateAccessForIndividualTestersToBuild(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRemoveAccessForIndividualTestersFromBuild(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.RemoveAccessForIndividualTestersFromBuild(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListResourceIDsForIndividualTestersForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildIndividualTestersLinkagesResponse{}
	got, resp, err := client.Builds.ListResourceIDsForIndividualTestersForBuild(context.Background(), "10", &ListResourceIDsForIndividualTestersForBuildQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppEncryptionDeclarationForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppEncryptionDeclarationResponse{}
	got, resp, err := client.Builds.GetAppEncryptionDeclarationForBuild(context.Background(), "10", &GetAppEncryptionDeclarationForBuildQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppEncryptionDeclarationIDForBuild(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &BuildAppEncryptionDeclarationLinkageResponse{}
	got, resp, err := client.Builds.GetAppEncryptionDeclarationIDForBuild(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
