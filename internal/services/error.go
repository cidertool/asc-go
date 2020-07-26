package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ErrorResponse contains information with error details that an API returns in the response body whenever the API request is not successful.
type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Errors   *[]struct {
		Code   string       `json:"code"`
		Detail string       `json:"detail"`
		ID     *string      `json:"id,omitempty"`
		Source *interface{} `json:"source,omitempty"`
		Status string       `json:"status"`
		Title  string       `json:"title"`
	} `json:"errors,omitempty"`
}

func (e *ErrorResponse) Error() string {
	report := strings.Builder{}
	if e.Errors != nil {
		for _, err := range *e.Errors {
			report.WriteString(fmt.Sprintf("* %s %s – %s\n\t%s\n", err.Status, err.Code, err.Title, err.Detail))
		}
	}
	return fmt.Sprintf(
		"%v %v: %d\n%v",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		report.String(),
	)
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	data, err := ioutil.ReadAll(r.Body)
	erro := new(ErrorResponse)
	if err == nil && data != nil {
		json.Unmarshal(data, erro)
	}
	erro.Response = r
	return erro
}
