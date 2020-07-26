package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, err.(*ErrorResponse).Response, resp)
	assert.NotZero(t, len(err.Error()))
}
