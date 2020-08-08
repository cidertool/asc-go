package asc

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.Equal(t, defaultBaseURL, c.BaseURL.String())
	assert.Equal(t, userAgent, c.UserAgent)

	c2 := NewClient(nil)
	assert.NotSame(t, c.client, c2.client, "NewClient returned same http.Clients, but they should differ")
}

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
	resp, err := client.get(context.Background(), "test", nil, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestGetWithQuery(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	params := mockParams{"TEST"}

	var unmarshaled mockPayload
	// Testing absolute URL path
	resp, err := client.get(context.Background(), server.URL+"/test", &params, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestGetError(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.get(context.Background(), "test", nil, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestGetBadRequest(t *testing.T) {
}

func TestPost(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.post(context.Background(), "test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestPostRespondingNoContent(t *testing.T) {
}

func TestPatch(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.patch(context.Background(), "test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestPatchRespondingNoContent(t *testing.T) {
}

func TestDelete(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	resp, err := client.delete(context.Background(), "test", &body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func newServer(raw string) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, raw)
	}))
	base, _ := url.Parse(server.URL)
	client := &Client{
		client:    server.Client(),
		BaseURL:   base,
		UserAgent: "asc-go",
	}
	return client, server
}

func TestCheckGoodResponse(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	err := checkResponse(resp)
	assert.NoError(t, err)
}

func TestCheckBadResponse(t *testing.T) {
	resp := &http.Response{
		StatusCode: 400,
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{},
		},
		Body: ioutil.NopCloser(strings.NewReader(`{"errors":[{"code":"","status":"","title":"","detail":""}]}`)),
	}
	err := checkResponse(resp)
	assert.Error(t, err)
	assert.IsType(t, new(ErrorResponse), err)
	assert.Equal(t, resp, err.(*ErrorResponse).Response)
	assert.NotZero(t, len(err.Error()))
}
