package types

import (
	"encoding/json"
	"time"
)

const dateFormat = "2006-01-02"

// Date represents a date with no time component
type Date struct {
	time.Time
}

// MarshalJSON is a custom marshaller for time-less dates
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(dateFormat))
}

// UnmarshalJSON is a custom unmarshaller for time-less dates
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}
