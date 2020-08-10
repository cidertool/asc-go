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

func (b *mockBody) applyTypes() {}

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

func TestDelete(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	resp, err := client.delete(context.Background(), "test", &body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCheckGoodResponse(t *testing.T) {
	resp := &Response{&http.Response{StatusCode: 200}, Rate{}}
	err := checkResponse(resp)
	assert.NoError(t, err)
}

func TestCheckBadResponse(t *testing.T) {
	resp := &Response{
		Response: &http.Response{
			StatusCode: 400,
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{},
			},
			Body: ioutil.NopCloser(strings.NewReader(`{"errors":[{"code":"","status":"","title":"","detail":""}]}`)),
		},
	}
	err := checkResponse(resp)
	assert.Error(t, err)
	assert.IsType(t, new(ErrorResponse), err)
	assert.Equal(t, resp.Response, err.(*ErrorResponse).Response)
	assert.NotZero(t, len(err.Error()))
}

func mustParseURL(t *testing.T, u string) url.URL {
	parsed, err := url.Parse(u)
	if err != nil {
		t.Fatalf("Parse(%q) got err %v", u, err)
	}
	return *parsed
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

	client.common.client = client

	client.Apps = (*AppsService)(&client.common)
	client.Builds = (*BuildsService)(&client.common)
	client.Pricing = (*PricingService)(&client.common)
	client.Provisioning = (*ProvisioningService)(&client.common)
	client.Publishing = (*PublishingService)(&client.common)
	client.Reporting = (*ReportingService)(&client.common)
	client.Submission = (*SubmissionService)(&client.common)
	client.TestFlight = (*TestflightService)(&client.common)
	client.Users = (*UsersService)(&client.common)

	return client, server
}

func testEndpointWithResponse(t *testing.T, marshalledGot string, want interface{}, endpoint func(ctx context.Context, client *Client) (interface{}, *Response, error)) {
	client, server := newServer(marshalledGot)
	defer server.Close()

	got, resp, err := endpoint(context.Background(), client)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if want != nil {
		assert.Equal(t, want, got)
	}
}

func testEndpointWithNoContent(t *testing.T, endpoint func(ctx context.Context, client *Client) (*Response, error)) {
	client, server := newServer("")
	defer server.Close()

	resp, err := endpoint(context.Background(), client)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
