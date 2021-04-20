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
	"testing"
)

func TestCreateCertificate(t *testing.T) {
	t.Parallel()

	testEndpointExpectingError(t, "{}", func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateCertificate(ctx, CertificateTypeDevelopment, nil)
	})
	testEndpointWithResponse(t, "{}", &CertificateResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.CreateCertificate(ctx, CertificateTypeDevelopment, &bytes.Buffer{})
	})
}

func TestListCertificates(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CertificatesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.ListCertificates(ctx, &ListCertificatesQuery{})
	})
}

func TestGetCertificate(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &CertificateResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Provisioning.GetCertificate(ctx, "10", &GetCertificateQuery{})
	})
}

func TestRevokeCertificate(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Provisioning.RevokeCertificate(ctx, "10")
	})
}
