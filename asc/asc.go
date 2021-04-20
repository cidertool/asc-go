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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/v1/"
	userAgent      = "asc-go"
	defaultTimeout = 30 * time.Second

	headerRateLimit = "X-Rate-Limit"
)

// Client is the root instance of the App Store Connect API.
type Client struct {
	client    *http.Client
	baseURL   *url.URL
	UserAgent string
	httpDebug bool

	common service

	Apps         *AppsService
	Builds       *BuildsService
	Pricing      *PricingService
	Provisioning *ProvisioningService
	Publishing   *PublishingService
	Reporting    *ReportingService
	Submission   *SubmissionService
	TestFlight   *TestflightService
	Users        *UsersService
}

// NewClient creates a new Client instance.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				IdleConnTimeout: defaultTimeout,
			},
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		baseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.common.client = c

	c.Apps = (*AppsService)(&c.common)
	c.Builds = (*BuildsService)(&c.common)
	c.Pricing = (*PricingService)(&c.common)
	c.Provisioning = (*ProvisioningService)(&c.common)
	c.Publishing = (*PublishingService)(&c.common)
	c.Reporting = (*ReportingService)(&c.common)
	c.Submission = (*SubmissionService)(&c.common)
	c.TestFlight = (*TestflightService)(&c.common)
	c.Users = (*UsersService)(&c.common)

	return c
}

// SetHTTPDebug this enables global http request/response dumping for this API.
func (c *Client) SetHTTPDebug(flag bool) {
	c.httpDebug = flag
}

// Response is a App Store Connect API response. This wraps the standard http.Response
// returned from Apple and provides convenient access to things like rate limit.
type Response struct {
	*http.Response

	Rate Rate
}

// Rate represents the rate limit for the current client.
//
// https://developer.apple.com/documentation/appstoreconnectapi/identifying_rate_limits
type Rate struct {
	// The number of requests per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`
}

// ErrorResponse contains information with error details that an API returns in the
// response body whenever the API request is not successful.
type ErrorResponse struct {
	Response *http.Response       `json:"-"`
	Errors   []ErrorResponseError `json:"errors,omitempty"`
}

// ErrorResponseError is a model used in ErrorResponse to describe a single error from the API.
type ErrorResponseError struct {
	// Code is a machine-readable indication of the type of error. The code is a hierarchical
	// value with levels of specificity separated by the '.' character. This value is parseable
	// for programmatic error handling in code.
	Code string `json:"code"`
	// Detail is a detailed explanation of the error. Do not use this field for programmatic error handling.
	Detail string `json:"detail"`
	// ID is a unique identifier of a specific instance of an error, request, and response.
	// Use this ID when providing feedback to or debugging issues with Apple.
	ID *string `json:"id,omitempty"`
	// Source wraps one of two possible types of values: source.parameter, provided when a query
	// parameter produced the error, or source.JsonPointer, provided when a problem with the entity
	// produced the error.
	Source *ErrorSource `json:"source,omitempty"`
	// Status is the HTTP status code of the error. This status code usually matches the
	// response's status code; however, if the request produces multiple errors, these two
	// codes may differ.
	Status string `json:"status"`
	// Title is a summary of the error. Do not use this field for programmatic error handling.
	Title string `json:"title"`
	// Meta is an undocumented field associating an error to many other errors.
	Meta *ErrorMeta `json:"meta,omitempty"`
}

// ErrorSource is the union of two API types: `ErrorResponse.Errors.JsonPointer` and `ErrorResponse.Errors.Parameter`.
//
// https://developer.apple.com/documentation/appstoreconnectapi/errorresponse/errors/jsonpointer
// https://developer.apple.com/documentation/appstoreconnectapi/errorresponse/errors/parameter
type ErrorSource struct {
	// A JSON pointer that indicates the location in the request entity where the error originates.
	Pointer string `json:"pointer,omitempty"`
	// The query parameter that produced the error.
	Parameter string `json:"parameter,omitempty"`
}

// ErrorMeta is an undocumented type that contains associations to other errors, grouped by route.
type ErrorMeta struct {
	// AssociatedErrors is a map of routes to array of errors that are associated with the current error.
	AssociatedErrors map[string][]ErrorResponseError `json:"associatedErrors,omitempty"`
}

type service struct {
	client *Client
}

// request is a common structure for a request body sent to the API.
type requestBody struct {
	Data     interface{} `json:"data"`
	Included interface{} `json:"included,omitempty"`
}

func newRequestBody(data interface{}) *requestBody {
	return newRequestBodyWithIncluded(data, nil)
}

func newRequestBodyWithIncluded(data interface{}, included interface{}) *requestBody {
	return &requestBody{Data: data, Included: included}
}

// FollowReference is a convenience method to perform a GET on a relationship link with
// pre-established parameters that you know the response type of.
func (c *Client) FollowReference(ctx context.Context, ref *Reference, v interface{}) (*Response, error) {
	return c.get(ctx, ref.String(), nil, &v)
}

type requestOption func(*http.Request)

func withAccept(typ string) requestOption {
	return func(req *http.Request) {
		req.Header.Set("Accept", typ)
	}
}

func withContentType(typ string) requestOption {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", typ)
	}
}

// AddOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func appendingQueryOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// get sends a GET request to the API as configured.
func (c *Client) get(ctx context.Context, url string, query interface{}, v interface{}, options ...requestOption) (*Response, error) {
	var err error
	if query != nil {
		url, err = appendingQueryOptions(url, query)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.newRequest(ctx, "GET", url, nil, options...)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// post sends a POST request to the API as configured.
func (c *Client) post(ctx context.Context, url string, body *requestBody, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, "POST", url, body, withContentType("application/json"))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// patch sends a PATCH request to the API as configured.
func (c *Client) patch(ctx context.Context, url string, body *requestBody, v interface{}) (*Response, error) {
	req, err := c.newRequest(ctx, "PATCH", url, body, withContentType("application/json"))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// delete sends a DELETE request to the API as configured.
func (c *Client) delete(ctx context.Context, url string, body *requestBody) (*Response, error) {
	req, err := c.newRequest(ctx, "DELETE", url, body, withContentType("application/json"))
	if err != nil {
		return nil, err
	}

	return c.do(ctx, req, nil)
}

func (c *Client) newRequest(ctx context.Context, method string, path string, body *requestBody, options ...requestOption) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var u *url.URL
	if rel.IsAbs() {
		u = rel
	} else {
		u = c.baseURL.ResolveReference(rel)
	}

	buf := new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	for _, option := range options {
		option(req)
	}

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	respCh := make(chan *http.Response, 1)
	op := func() error {
		if c.httpDebug {
			if dump, err := httputil.DumpRequest(req, true); err == nil {
				fmt.Printf("DEBUG request uri=%s\n%s\n", req.URL, dump) // nolint: forbidigo
			}
		}

		resp, err := c.client.Do(req) // nolint: bodyclose
		if err != nil {
			select {
			case <-ctx.Done():
				return backoff.Permanent(ctx.Err())
			default:
				return backoff.Permanent(err)
			}
		}

		if c.httpDebug {
			if dump, err := httputil.DumpResponse(resp, true); err == nil {
				fmt.Printf("DEBUG response uri=%s\n%s\n", req.URL, dump) // nolint: forbidigo
			}
		}

		respCh <- resp

		return nil
	}

	notify := func(err error, delay time.Duration) {
		if c.httpDebug {
			fmt.Printf("DEBUG error %v, retry in %v\n", err, delay) // nolint: forbidigo
		}
	}

	err := backoff.RetryNotify(op, backoff.NewExponentialBackOff(), notify)

	resp := <-respCh

	defer closeDesc(resp.Body)

	response := newResponse(resp)

	if err != nil {
		return response, err
	}

	if err := checkResponse(response); err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRate(r)

	return response
}

func checkResponse(r *Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := io.ReadAll(r.Body)
	erro := new(ErrorResponse)

	if err == nil && data != nil {
		err := json.Unmarshal(data, erro)
		if err != nil {
			return err
		}
	}

	erro.Response = r.Response

	return erro
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate

	header := r.Header.Get(headerRateLimit)

	if header == "" {
		return rate
	}

	for _, component := range strings.Split(header, ";") {
		if component == "" {
			continue
		}

		var kvpLen = 2

		kvp := strings.Split(component, ":")
		if len(kvp) != kvpLen {
			continue
		}

		key := kvp[0]
		value, err := strconv.Atoi(kvp[1])

		if err != nil {
			continue
		}

		switch key {
		case "user-hour-lim":
			rate.Limit = value
		case "user-hour-rem":
			rate.Remaining = value
		}
	}

	return rate
}

func (e ErrorResponse) Error() string {
	report := strings.Builder{}

	if e.Errors != nil {
		for _, err := range e.Errors {
			report.WriteString(fmt.Sprintf("* %s", err.String(1)))
		}
	}

	return fmt.Sprintf(
		"%v %v: %d\n%v",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		report.String(),
	)
}

func (e ErrorResponseError) String(level int) string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s %s – %s\n%s%s\n", e.Status, e.Code, e.Title, strings.Repeat("\t", level), e.Detail))

	if e.Meta == nil {
		return str.String()
	}

	for route, errs := range e.Meta.AssociatedErrors {
		str.WriteString(fmt.Sprintf("\t%s:\n", route))

		for _, err := range errs {
			var (
				tabLen     = 1
				contentLen = 2
			)

			str.WriteString(fmt.Sprintf("%s%s", strings.Repeat("\t", level+tabLen), err.String(level+contentLen)))
		}
	}

	return str.String()
}

// Close closes an open descriptor.
func closeDesc(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
