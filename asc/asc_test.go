package asc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}

	c2 := NewClient(nil)
	if c.client == c2.client {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

type mockPayload struct {
	Value string `json:"value"`
}

type mockParams struct {
	Field string `url:"field,omitempty"`
}

type mockBody struct {
	Field string `url:"field,omitempty"`
}

func TestGet(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.Get("test", &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetWithQuery(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	params := mockParams{"TEST"}

	var unmarshaled mockPayload
	// Testing absolute URL path
	resp, err := client.GetWithQuery(server.URL+"/test", &params, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetError(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	var unmarshaled mockPayload
	resp, err := client.Get("test", &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestGetBadRequest(t *testing.T) {
}

func TestPost(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.Post("test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestPostRespondingNoContent(t *testing.T) {
}

func TestPatch(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	var unmarshaled mockPayload
	resp, err := client.Patch("test", &body, &unmarshaled)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, unmarshaled, mockPayload{"TEST"})
}

func TestPatchRespondingNoContent(t *testing.T) {
}

func TestDelete(t *testing.T) {
	marshaled := `{"value":"TEST"}`
	client, server := newServer(marshaled)
	defer server.Close()

	body := mockBody{"TEST"}

	resp, err := client.Delete("test", &body)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func newServer(raw string) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, raw)
	}))
	base, _ := url.Parse(server.URL)
	client := &Client{
		client:    server.Client(),
		BaseURL:   base,
		UserAgent: "asc-go",
	}
	return client, server
}

func TestReference(t *testing.T) {
	marshaled := `{"self":"https://api.appstoreconnect.apple.com/me?cursor=TEST"}`
	var links DocumentLinks
	err := json.Unmarshal([]byte(marshaled), &links)
	assert.NoError(t, err)
	assert.Equal(t, "TEST", links.Self.Cursor())
	remarshaled, err := json.Marshal(links)
	assert.NoError(t, err)
	assert.Equal(t, marshaled, string(remarshaled))
}

func TestReferenceNoCursor(t *testing.T) {
	ref := new(Reference)
	assert.Empty(t, ref.Cursor())
	ref = &Reference{url.URL{}}
	assert.Empty(t, ref.Cursor())
}

func TestReferenceBadUnmarshal(t *testing.T) {
	marshaledNoString := `{"self":0}`
	marshaledNoURL := `{"self":":"}`
	var links DocumentLinks
	var err error
	err = json.Unmarshal([]byte(marshaledNoString), &links)
	assert.Error(t, err)
	err = json.Unmarshal([]byte(marshaledNoURL), &links)
	assert.Error(t, err)
}

func TestCheckGoodResponse(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	err := checkResponse(resp)
	assert.NoError(t, err)
}

func TestCheckBadResponse(t *testing.T) {
	resp := &http.Response{
		StatusCode: 400,
		Request: &http.Request{
			Method: "GET",
			URL:    &url.URL{},
		},
		Body: ioutil.NopCloser(strings.NewReader(`{"errors":[{"code":"","status":"","title":"","detail":""}]}`)),
	}
	err := checkResponse(resp)
	assert.Error(t, err)
	assert.IsType(t, new(ErrorResponse), err)
	assert.Equal(t, err.(*ErrorResponse).Response, resp)
	assert.NotZero(t, len(err.Error()))
}

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
	want := emailContainerJSON("my@email.com")
	b := newEmailContainer("my@email.com")
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
	jsonStr := emailContainerJSON("my@email.com")
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
	if &got == want {
		t.Error("Bool returned same *bool, but they should differ")
	}
}

func TestInt(t *testing.T) {
	got := 100
	want := Int(got)
	if &got == want {
		t.Error("Int returned same *int, but they should differ")
	}
}

func TestFloat(t *testing.T) {
	got := 100.5
	want := Float(got)
	if &got == want {
		t.Error("Float returned same *float64, but they should differ")
	}
}

func TestString(t *testing.T) {
	got := "App Store Connect"
	want := String(got)
	if &got == want {
		t.Error("String returned same *string, but they should differ")
	}
}
