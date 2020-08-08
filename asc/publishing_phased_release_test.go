package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePhasedRelease(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppStoreVersionPhasedReleaseResponse{}
	got, resp, err := client.Publishing.CreatePhasedRelease(context.Background(), &AppStoreVersionPhasedReleaseCreateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestUpdatePhasedRelease(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppStoreVersionPhasedReleaseResponse{}
	got, resp, err := client.Publishing.UpdatePhasedRelease(context.Background(), "10", &AppStoreVersionPhasedReleaseUpdateRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestDeletePhasedRelease(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Publishing.DeletePhasedRelease(context.Background(), "10")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetAppStoreVersionPhasedReleaseForAppStoreVersion(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppStoreVersionPhasedReleaseResponse{}
	got, resp, err := client.Publishing.GetAppStoreVersionPhasedReleaseForAppStoreVersion(context.Background(), "10", &GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}
