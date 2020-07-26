package services

// Request is a common structure for a request body sent to the API
type request struct {
	Data interface{} `json:"data"`
}

func newRequest(v interface{}) *request {
	req := &request{Data: v}
	return req
}
