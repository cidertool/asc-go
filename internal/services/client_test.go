package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPayload struct {
	Value string `json:"value"`
}

type mockParams struct {
	Field string `url:"field,omitempty"`
}

type mockBody struct {
	Field string `url:"field,omitempty"`
}

func TestGet(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.Get("test", &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetWithQuery(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	params := mockParams{"TEST"}

	var unmarshaled mockPayload
	// Testing absolute URL path
	resp, err := client.GetWithQuery(server.URL+"/test", &params, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetError(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.Get("test", &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetBadRequest(t *testing.T) {
}

func TestPost(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.Post("test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestPostRespondingNoContent(t *testing.T) {
}

func TestPatch(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.Patch("test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestPatchRespondingNoContent(t *testing.T) {
}

func TestDelete(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	resp, err := client.Delete("test", &body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func newServer(raw string) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, raw)
	}))
	base, _ := url.Parse(server.URL)
	client := &Client{
		Client:    server.Client(),
		BaseURL:   base,
		UserAgent: "asc-go",
	}
	return client, server
}
