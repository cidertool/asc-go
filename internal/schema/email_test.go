package schema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
