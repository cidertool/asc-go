package asc

import (
	"net/http"
	"net/url"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/builds"
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/pricing"
	"github.com/aaronsky/asc-go/provisioning"
	"github.com/aaronsky/asc-go/publishing"
	"github.com/aaronsky/asc-go/reporting"
	"github.com/aaronsky/asc-go/submission"
	"github.com/aaronsky/asc-go/testflight"
	"github.com/aaronsky/asc-go/users"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/"
	userAgent      = "asc-go"
)

// Client is the root instance of the App Store Connect API
type Client struct {
	// internal.Client
	common services.Service

	Apps         *apps.Service
	Builds       *builds.Service
	Pricing      *pricing.Service
	Provisioning *provisioning.Service
	Publishing   *publishing.Service
	Reporting    *reporting.Service
	Submission   *submission.Service
	TestFlight   *testflight.Service
	Users        *users.Service
}

// NewClient creates a new Client instance
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{}

	c.common.Client = services.Client{
		Client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.Apps = (*apps.Service)(&c.common)
	c.Builds = (*builds.Service)(&c.common)
	c.Pricing = (*pricing.Service)(&c.common)
	c.Provisioning = (*provisioning.Service)(&c.common)
	c.Publishing = (*publishing.Service)(&c.common)
	c.Reporting = (*reporting.Service)(&c.common)
	c.Submission = (*submission.Service)(&c.common)
	c.TestFlight = (*testflight.Service)(&c.common)
	c.Users = (*users.Service)(&c.common)

	return c
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it, but unlike Int
// its argument value is an int.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it, but unlike Float
// its argument value is a float64.
func Float(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}
