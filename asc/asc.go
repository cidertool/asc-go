package asc

import (
	"net/http"
	"net/url"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
	"github.com/aaronsky/asc-go/v1/asc/pricing"
	"github.com/aaronsky/asc-go/v1/asc/provisioning"
	"github.com/aaronsky/asc-go/v1/asc/publishing"
	"github.com/aaronsky/asc-go/v1/asc/reporting"
	"github.com/aaronsky/asc-go/v1/asc/submission"
	"github.com/aaronsky/asc-go/v1/asc/testflight"
	"github.com/aaronsky/asc-go/v1/asc/users"
)

const (
	defaultBaseURL = "https://api.appstoreconnect.apple.com/v1/"
	userAgent      = "asc-go"
)

// Client is the root instance of the App Store Connect API
type Client struct {
	// common.Client
	common common.Service

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

	c.common.Client = common.Client{
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
