package asc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppEncryptionDeclarations(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppEncryptionDeclarationsResponse{}
	got, resp, err := client.Builds.ListAppEncryptionDeclarations(context.Background(), &ListAppEncryptionDeclarationsQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppEncryptionDeclaration(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppEncryptionDeclarationResponse{}
	got, resp, err := client.Builds.GetAppEncryptionDeclaration(context.Background(), "10", &GetAppEncryptionDeclarationQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestGetAppForAppEncryptionDeclaration(t *testing.T) {
	client, server := newServer("{}")
	defer server.Close()

	want := &AppResponse{}
	got, resp, err := client.Builds.GetAppForAppEncryptionDeclaration(context.Background(), "10", &GetAppForEncryptionDeclarationQuery{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, want, got)
}

func TestAssignBuildsToAppEncryptionDeclaration(t *testing.T) {
	client, server := newServer("")
	defer server.Close()

	resp, err := client.Builds.AssignBuildsToAppEncryptionDeclaration(context.Background(), "10", &[]RelationshipsData{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
