/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const marshaledMockPayload = `{"value":"TEST"}`

func TestNewClient(t *testing.T) {
	t.Parallel()

	c := NewClient(nil)

	assert.Equal(t, defaultBaseURL, c.baseURL.String())
	assert.Equal(t, userAgent, c.UserAgent)

	c2 := NewClient(nil)
	assert.NotSame(t, c.client, c2.client, "NewClient returned same http.Clients, but they should differ")
}

func TestSetHTTPDebug(t *testing.T) {
	t.Parallel()

	client := NewClient(nil)
	client.SetHTTPDebug(true)
	assert.True(t, client.httpDebug)
	client.SetHTTPDebug(false)
	assert.False(t, client.httpDebug)
}

type mockPayload struct {
	Value string `json:"value"`
}

type mockParams struct {
	Field    string  `url:"field,omitempty"`
	NilField *string `url:"nil_field,omitempty"`
}

type mockBody struct {
	Field string `url:"field,omitempty"`
}

func TestNilContextReturnsError(t *testing.T) {
	t.Parallel()

	client, server := newServer("", http.StatusOK, false)

	defer server.Close()

	_, err := client.get(nil, "test", nil, nil) // nolint: staticcheck
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)
	client.SetHTTPDebug(true)

	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.get(context.Background(), "test", nil, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestGetWithQuery(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)

	defer server.Close()

	params := mockParams{
		Field:    "TEST",
		NilField: nil,
	}

	var unmarshaled mockPayload
	// Testing absolute URL path
	resp, err := client.get(context.Background(), server.URL+"/test", &params, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestGetWithQuery_Error(t *testing.T) {
	t.Parallel()

	client, server := newServer("", http.StatusOK, true)
	defer server.Close()

	badQueryValue := []string{"horses"}
	resp, err := client.get(context.Background(), server.URL, &badQueryValue, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetError(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.get(context.Background(), "test", nil, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestPost(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.post(context.Background(), "test", newRequestBody(body), &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestPatch(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.patch(context.Background(), "test", newRequestBody(body), &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockPayload{"TEST"}, unmarshaled)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	client, server := newServer(marshaledMockPayload, http.StatusOK, true)
	defer server.Close()

	body := mockBody{"TEST"}

	resp, err := client.delete(context.Background(), "test", newRequestBody(body))

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCheckGoodResponse(t *testing.T) {
	t.Parallel()

	resp := &Response{&http.Response{StatusCode: 200}, Rate{}}
	err := checkResponse(resp)
	assert.NoError(t, err)
}

func TestCheckBadResponse(t *testing.T) {
	t.Parallel()

	resp := &Response{
		Response: &http.Response{
			StatusCode: 400,
			Request: &http.Request{
				Method: "GET",
				URL:    &url.URL{},
			},
			Body: io.NopCloser(strings.NewReader(`{
				"errors": [
				  	{
						"code": "",
						"status": "",
						"title": "",
						"detail": "",
						"meta": {
					  		"associatedErrors": {
								"/v1/route/": [
						  			{
										"code": "",
										"status": "",
										"title": "",
										"detail": ""
						  			}
								]
					  		}
						}
				  	}
				]
			}`)),
		},
	}
	err := checkResponse(resp)
	assert.Error(t, err)
	assert.IsType(t, new(ErrorResponse), err)

	var respErr *ErrorResponse

	ok := errors.As(err, &respErr)
	assert.True(t, ok)
	assert.Equal(t, resp.Response, respErr.Response)
	assert.NotEmpty(t, err.Error())
}

func TestAppendingQueryOptions(t *testing.T) {
	t.Parallel()

	got, err := appendingQueryOptions("dog", nil)
	assert.NoError(t, err)
	assert.Equal(t, "dog", got)

	got, err = appendingQueryOptions("dog", &mockParams{})
	assert.NoError(t, err)
	assert.Equal(t, "dog", got)

	got, err = appendingQueryOptions("dog", &mockParams{Field: "cat"})
	assert.NoError(t, err)
	assert.Equal(t, "dog?field=cat", got)

	// bad url
	_, err = appendingQueryOptions(":", nil)
	assert.Error(t, err)

	// invalid input to query.Values
	badQueryValue := []string{"horses"}
	_, err = appendingQueryOptions("dog", badQueryValue)
	assert.Error(t, err)
}

func newServer(raw string, status int, addRateLimit bool) (*Client, *httptest.Server) { // nolint: unparam
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if addRateLimit {
			w.Header().Add("X-Rate-Limit", "user-hour-lim:2500;user-hour-rem:10;hank:dean:venture;rusty:jr;;")
		}
		w.WriteHeader(status)
		fmt.Fprintln(w, raw)
	}))

	base, _ := url.Parse(server.URL)
	client := NewClient(server.Client())
	client.baseURL = base

	return client, server
}

func testEndpointWithResponse(t *testing.T, marshalledGot string, want interface{}, endpoint func(ctx context.Context, client *Client) (interface{}, *Response, error)) {
	t.Helper()

	client, server := newServer(marshalledGot, http.StatusOK, true)
	defer server.Close()

	got, resp, err := endpoint(context.Background(), client)

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	if want != nil {
		assert.Equal(t, want, got)
	}
}

func testEndpointWithNoContent(t *testing.T, endpoint func(ctx context.Context, client *Client) (*Response, error)) {
	t.Helper()

	client, server := newServer("", http.StatusOK, false)
	defer server.Close()

	resp, err := endpoint(context.Background(), client)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func testEndpointExpectingError(t *testing.T, marshalledGot string, endpoint func(ctx context.Context, client *Client) (interface{}, *Response, error)) {
	t.Helper()

	client, server := newServer(marshalledGot, http.StatusOK, true)
	defer server.Close()

	got, _, err := endpoint(context.Background(), client)

	assert.Error(t, err)
	assert.Nil(t, got)
}

func testEndpointCustomBehavior(marshalledGot string, behavior func(ctx context.Context, client *Client)) {
	client, server := newServer(marshalledGot, http.StatusOK, true)
	defer server.Close()
	behavior(context.Background(), client)
}
