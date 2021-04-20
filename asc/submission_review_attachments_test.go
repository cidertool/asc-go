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

package asc // nolint: dupl

import (
	"context"
	"testing"
)

func TestGetAttachment(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetAttachment(ctx, "10", &GetAttachmentQuery{})
	})
}

func TestListAttachmentsForReviewDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.ListAttachmentsForReviewDetail(ctx, "10", &ListAttachmentQuery{})
	})
}

func TestCreateAttachment(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CreateAttachment(ctx, "", 0, "")
	})
}

func TestCommitAttachment(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CommitAttachment(ctx, "10", Bool(true), String("10"))
	})
}

func TestDeleteAttachment(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Submission.DeleteAttachment(ctx, "10")
	})
}
