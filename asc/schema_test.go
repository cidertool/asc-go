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
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type dateContainer struct {
	Field Date `json:"date"`
}

func newDateContainer(year int, month time.Month, date int) dateContainer {
	return dateContainer{
		Date{
			time.Date(year, month, date, 0, 0, 0, 0, time.UTC),
		},
	}
}

func dateContainerJSON(date string) string {
	return fmt.Sprintf(`{"date":"%s"}`, date)
}

func TestDateMarshal(t *testing.T) {
	t.Parallel()

	want := dateContainerJSON("2020-04-01")
	b := newDateContainer(2020, 4, 1)
	got, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))
}

func TestDateUnmarshal(t *testing.T) {
	t.Parallel()

	want := time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)
	jsonStr := dateContainerJSON("2020-04-01")

	var b dateContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, want, b.Field.Time)
}

func TestDateUnmarshalWrongType(t *testing.T) {
	t.Parallel()

	jsonStr := `{"date":-1}`

	var b dateContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestDateUnmarshalInvalidDate(t *testing.T) {
	t.Parallel()

	jsonStr := dateContainerJSON("TEST")

	var b dateContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

type dateTimeContainer struct {
	Field DateTime `json:"time"`
}

func newDateTimeContainer(year int, month time.Month, date int, hour int, min int, sec int, nsec int) dateTimeContainer {
	return dateTimeContainer{
		DateTime{
			time.Date(year, month, date, hour, min, sec, nsec, time.UTC),
		},
	}
}

func dateTimeContainerJSON(dateTime string) string {
	return fmt.Sprintf(`{"time":"%s"}`, dateTime)
}

func TestDateTimeMarshal(t *testing.T) {
	t.Parallel()

	want := dateTimeContainerJSON("2020-04-01T05:16:48.915+0000")
	b := newDateTimeContainer(2020, 4, 1, 5, 16, 48, 915000000)
	got, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))
}

func TestDateTimeUnmarshal(t *testing.T) {
	t.Parallel()

	want := time.Date(2020, 4, 1, 5, 16, 48, 915000000, time.UTC)
	jsonStr := dateTimeContainerJSON("2020-04-01T05:16:48.915Z")

	var b dateTimeContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, b.Field.Time, want)
}

func TestDateTimeUnmarshalWrongType(t *testing.T) {
	t.Parallel()

	jsonStr := `{"time":-1}`

	var b dateTimeContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestDateTimeUnmarshalInvalidDate(t *testing.T) {
	t.Parallel()

	jsonStr := dateTimeContainerJSON("TEST")

	var b dateTimeContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

type emailContainer struct {
	Field Email `json:"email"`
}

func newEmailContainer(email string) emailContainer {
	return emailContainer{
		Email(email),
	}
}

func emailContainerJSON(email string) string {
	return fmt.Sprintf(`{"email":"%s"}`, email)
}

func TestEmailMarshal(t *testing.T) {
	t.Parallel()

	email := "my@email.com"
	want := emailContainerJSON(email)
	b := newEmailContainer(email)
	got, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))
}

func TestEmailMarshalInvalidEmail(t *testing.T) {
	t.Parallel()

	b := newEmailContainer("TEST")
	_, err := json.Marshal(b)
	assert.Error(t, err)
}

func TestEmailUnmarshal(t *testing.T) {
	t.Parallel()

	want := "my@email.com"
	jsonStr := emailContainerJSON(want)

	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, Email(want), b.Field)
}

func TestEmailUnmarshalWrongType(t *testing.T) {
	t.Parallel()

	jsonStr := `{"email":-1}`

	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestEmailUnmarshalInvalidEmail(t *testing.T) {
	t.Parallel()

	jsonStr := emailContainerJSON("TEST")

	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestBool(t *testing.T) {
	t.Parallel()

	got := true
	want := Bool(got)
	assert.Equal(t, want, &got, "Bool returned same *bool, but they should differ")
}

func TestInt(t *testing.T) {
	t.Parallel()

	got := 100
	want := Int(got)
	assert.Equal(t, want, &got, "Int returned same *int, but they should differ")
}

func TestFloat(t *testing.T) {
	t.Parallel()

	got := 100.5
	want := Float(got)
	assert.Equal(t, want, &got, "Float returned same *float64, but they should differ")
}

func TestString(t *testing.T) {
	t.Parallel()

	got := "App Store Connect"
	want := String(got)
	assert.Equal(t, want, &got, "String returned same *string, but they should differ")
}
