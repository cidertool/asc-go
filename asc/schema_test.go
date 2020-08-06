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
	want := dateContainerJSON("2020-04-01")
	b := newDateContainer(2020, 4, 1)
	got, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))
}

func TestDateUnmarshal(t *testing.T) {
	want := time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)
	jsonStr := dateContainerJSON("2020-04-01")
	var b dateContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, want, b.Field.Time)
}

func TestDateUnmarshalWrongType(t *testing.T) {
	jsonStr := `{"date":-1}`
	var b dateContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestDateUnmarshalInvalidDate(t *testing.T) {
	jsonStr := dateContainerJSON("TEST")
	var b dateContainer
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
	email := "my@email.com"
	want := emailContainerJSON(email)
	b := newEmailContainer(email)
	got, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, want, string(got))
}

func TestEmailMarshalInvalidEmail(t *testing.T) {
	b := newEmailContainer("TEST")
	_, err := json.Marshal(b)
	assert.Error(t, err)
}

func TestEmailUnmarshal(t *testing.T) {
	want := "my@email.com"
	jsonStr := emailContainerJSON(want)
	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, Email(want), b.Field)
}

func TestEmailUnmarshalWrongType(t *testing.T) {
	jsonStr := `{"email":-1}`
	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestEmailUnmarshalInvalidEmail(t *testing.T) {
	jsonStr := emailContainerJSON("TEST")
	var b emailContainer
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.Error(t, err)
}

func TestBool(t *testing.T) {
	got := true
	want := Bool(got)
	assert.Equal(t, want, &got, "Bool returned same *bool, but they should differ")
}

func TestInt(t *testing.T) {
	got := 100
	want := Int(got)
	assert.Equal(t, want, &got, "Int returned same *int, but they should differ")
}

func TestFloat(t *testing.T) {
	got := 100.5
	want := Float(got)
	assert.Equal(t, want, &got, "Float returned same *float64, but they should differ")
}

func TestString(t *testing.T) {
	got := "App Store Connect"
	want := String(got)
	assert.Equal(t, want, &got, "String returned same *string, but they should differ")
}
