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
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipartUpload(t *testing.T) {
	t.Parallel()

	file, err := os.CreateTemp("", "big_file")
	if err != nil {
		assert.FailNow(t, "temp file creation produced an error", err)
	}

	defer rmFile(file)

	contents := make([]byte, 64)

	_, err = rand.Read(contents)
	if err != nil {
		assert.FailNow(t, "random bytes produced an error", err)
	}

	_, err = file.Write(contents)
	if err != nil {
		assert.FailNow(t, "writing the temp file produced an error", err)
	}

	client, server := newServer("", http.StatusOK, false)

	defer server.Close()

	operations := []UploadOperation{
		{
			URL:    String(client.baseURL.String()),
			Offset: Int(0),
			Length: Int(10),
			Method: String("PATCH"),
			RequestHeaders: []UploadOperationHeader{
				{
					Name: String("Hank"),
				},
				{
					Value: String("Dean"),
				},
				{
					Name:  String("Hank"),
					Value: String("Dean"),
				},
			},
		},
		{
			URL:            String(client.baseURL.String()),
			Offset:         Int(10),
			Length:         Int(10),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.baseURL.String()),
			Offset:         Int(20),
			Length:         Int(30),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.baseURL.String()),
			Offset:         Int(50),
			Length:         Int(10),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.baseURL.String()),
			Offset:         Int(60),
			Length:         Int(4),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
	}

	err = client.Upload(context.Background(), operations, file)
	assert.NoError(t, err)
}

func TestUploadOperationChunk(t *testing.T) {
	t.Parallel()

	file, err := os.CreateTemp("", "small_file")
	if err != nil {
		assert.FailNow(t, "temp file creation produced an error", err)
	}

	defer rmFile(file)

	contents := make([]byte, 20)

	_, err = rand.Read(contents)
	if err != nil {
		assert.FailNow(t, "random bytes produced an error", err)
	}

	_, err = file.Write(contents)
	if err != nil {
		assert.FailNow(t, "writing the temp file produced an error", err)
	}

	op := UploadOperation{
		URL:            String("test"),
		Offset:         Int(0),
		Method:         String("PATCH"),
		RequestHeaders: []UploadOperationHeader{},
	}

	_, err = op.chunk(file)
	assert.Error(t, err)

	op.Length = Int(10)

	chunk, err := op.chunk(file)
	assert.NoError(t, err)

	buf := &bytes.Buffer{}
	written, err := io.Copy(buf, chunk)
	assert.NoError(t, err)
	assert.EqualValues(t, written, *op.Length)
}

func TestUploadOperationUploadError_InvalidOperation(t *testing.T) {
	t.Parallel()

	file, err := os.CreateTemp("", "big_file")
	if err != nil {
		assert.FailNow(t, "temp file creation produced an error", err)
	}

	defer rmFile(file)

	contents := make([]byte, 64)

	_, err = rand.Read(contents)
	if err != nil {
		assert.FailNow(t, "random bytes produced an error", err)
	}

	_, err = file.Write(contents)
	if err != nil {
		assert.FailNow(t, "writing the temp file produced an error", err)
	}

	client, server := newServer("", http.StatusOK, false)
	defer server.Close()

	operations := []UploadOperation{
		{
			URL:            String(client.baseURL.String()),
			Offset:         Int(0),
			Length:         Int(64),
			RequestHeaders: []UploadOperationHeader{},
		},
	}

	err = client.Upload(context.Background(), operations, file)
	assert.Error(t, err)
}

// rmFile closes an open descriptor.
func rmFile(f *os.File) {
	if err := os.Remove(f.Name()); err != nil {
		fmt.Println(err) // nolint: forbidigo
	}
}
