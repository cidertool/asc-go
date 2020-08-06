package asc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/v1/"
	userAgent      = "asc-go"

	headerRateLimit = "X-Rate-Limit"

	dateFormat = "2006-01-02"

	emailRegexString = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	emailRegex = regexp.MustCompile(emailRegexString)
)

// Client is the root instance of the App Store Connect API
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

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

// NewClient creates a new Client instance
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
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

// ErrorResponse contains information with error details that an API returns in the response body whenever the API request is not successful.
type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Errors   *[]struct {
		Code   string       `json:"code"`
		Detail string       `json:"detail"`
		ID     *string      `json:"id,omitempty"`
		Source *interface{} `json:"source,omitempty"`
		Status string       `json:"status"`
		Title  string       `json:"title"`
	} `json:"errors,omitempty"`
}

type service struct {
	client *Client
}

// request is a common structure for a request body sent to the API
type request struct {
	Data interface{} `json:"data"`
}

// FollowReference is a convenience method to perform a GET on a relationship link with
// pre-established parameters that you know the response type of.
func (c *Client) FollowReference(ref *Reference, v interface{}) (*Response, error) {
	return c.get(ref.String(), nil, &v)
}

// get sends a GET request to the API as configured
func (c *Client) get(url string, query interface{}, v interface{}) (*Response, error) {
	var err error
	if query != nil {
		url, err = c.addOptions(url, query)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// post sends a POST request to the API as configured
func (c *Client) post(url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.newRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// patch sends a PATCH request to the API as configured
func (c *Client) patch(url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.newRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// delete sends a DELETE request to the API as configured
func (c *Client) delete(url string, body interface{}) (*Response, error) {
	req, err := c.newRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, nil)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var u *url.URL
	if rel.IsAbs() {
		u = rel
	} else {
		u = c.BaseURL.ResolveReference(rel)
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(wrapRequest(body))
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// AddOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func (c *Client) addOptions(s string, opt interface{}) (string, error) {
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

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	respCh := make(chan *http.Response, 1)
	op := func() error {
		resp, err := c.client.Do(req)
		if err != nil {
			return backoff.Permanent(err)
		}
		respCh <- resp
		return nil
	}

	notify := func(err error, delay time.Duration) {}

	if err := backoff.RetryNotify(op, backoff.NewExponentialBackOff(), notify); err != nil {
		return nil, err
	}

	resp := <-respCh

	defer resp.Body.Close()

	response := newResponse(resp)

	if err := checkResponse(resp); err != nil {
		return response, err
	}

	var err error

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

func wrapRequest(v interface{}) *request {
	req := &request{Data: v}
	return req
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRate(r)
	return response
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	erro := new(ErrorResponse)
	if err == nil && data != nil {
		json.Unmarshal(data, erro)
	}
	erro.Response = r
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
		kvp := strings.Split(component, ":")
		if len(kvp) != 2 {
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
		for _, err := range *e.Errors {
			report.WriteString(fmt.Sprintf("* %s %s – %s\n\t%s\n", err.Status, err.Code, err.Title, err.Detail))
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
