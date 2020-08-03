package asc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/v1/"
	userAgent      = "asc-go"

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

type service struct {
	client *Client
}

// request is a common structure for a request body sent to the API
type request struct {
	Data interface{} `json:"data"`
}

func newRequest(v interface{}) *request {
	req := &request{Data: v}
	return req
}

// Get sends a GET request to the API as configured
func (c *Client) Get(url string, v interface{}) (*http.Response, error) {
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

// GetWithQuery sends a GET request with a query to the API as configured
func (c *Client) GetWithQuery(url string, query interface{}, v interface{}) (*http.Response, error) {
	var err error
	if query != nil {
		url, err = c.addOptions(url, query)
		if err != nil {
			return nil, err
		}
	}
	return c.Get(url, v)
}

// Post sends a POST request to the API as configured
func (c *Client) Post(url string, body interface{}, v interface{}) (*http.Response, error) {
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

// Patch sends a PATCH request to the API as configured
func (c *Client) Patch(url string, body interface{}, v interface{}) (*http.Response, error) {
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

// Delete sends a DELETE request to the API as configured
func (c *Client) Delete(url string, body interface{}) (*http.Response, error) {
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
		err := json.NewEncoder(buf).Encode(newRequest(body))
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

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
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
	defer io.Copy(ioutil.Discard, resp.Body)

	if err := checkResponse(resp); err != nil {
		return resp, err
	}

	var err error

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// Reference is a wrapper type for a URL that contains a cursor parameter
type Reference struct {
	url.URL
}

// Cursor returns the cursor parameter on the Reference's internal URL.
func (r *Reference) Cursor() string {
	return r.Query().Get("cursor")
}

// MarshalJSON marshals the Reference into a JSON fragment
func (r Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON unmarshals the JSON fragment into a Reference
func (r *Reference) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	if u != nil {
		r.URL = *u
	}
	return nil
}

// PagedDocumentLinks defines model for PagedDocumentLinks.
type PagedDocumentLinks struct {
	First *Reference `json:"first,omitempty"`
	Next  *Reference `json:"next,omitempty"`
	Self  Reference  `json:"self"`
}

// PagingInformation defines model for PagingInformation.
type PagingInformation struct {
	Paging struct {
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"paging"`
}

// ResourceLinks defines model for ResourceLinks.
type ResourceLinks struct {
	Self Reference `json:"self"`
}

// DocumentLinks defines model for DocumentLinks.
type DocumentLinks struct {
	Self Reference `json:"self"`
}

// RelationshipsData contains data on the given relationship
type RelationshipsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// RelationshipsLinks contains links on the given relationship
type RelationshipsLinks struct {
	Related *string `json:"related,omitempty"`
	Self    *string `json:"self,omitempty"`
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

func (e *ErrorResponse) Error() string {
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

// Date represents a date with no time component
type Date struct {
	time.Time
}

// MarshalJSON is a custom marshaller for time-less dates
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(dateFormat))
}

// UnmarshalJSON is a custom unmarshaller for time-less dates
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// Email is a validated email address string
type Email string

// MarshalJSON is a custom marshaler for email addresses
func (e Email) MarshalJSON() ([]byte, error) {
	if !emailRegex.MatchString(string(e)) {
		return nil, errors.New("email: failed to pass regex validation")
	}
	return json.Marshal(string(e))
}

// UnmarshalJSON is a custom unmarshaller for email addresses
func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !emailRegex.MatchString(s) {
		return errors.New("email: failed to pass regex validation")
	}
	*e = Email(s)
	return nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	return &v
}

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float(v float64) *float64 {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}
