package services

import "net/http"

// Response contains the HTTP response returned by the API
type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}
