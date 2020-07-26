package schema

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
