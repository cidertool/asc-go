package common

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-querystring/query"
)

type Service struct {
	Client
}

// Client is the root instance of the App Store Connect API
type Client struct {
	Client    *http.Client
	BaseURL   *url.URL
	UserAgent string
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var u *url.URL
	if rel.IsAbs() {
		u = rel
	} else {
		u = c.BaseURL.ResolveReference(rel)
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	respCh := make(chan *http.Response, 1)
	op := func() error {
		resp, err := c.Client.Do(req)
		if err != nil {
			return backoff.Permanent(err)
		}
		respCh <- resp
		return nil
	}

	notify := func(err error, delay time.Duration) {

	}

	if err := backoff.RetryNotify(op, backoff.NewExponentialBackOff(), notify); err != nil {
		return nil, err
	}

	resp := <-respCh

	defer resp.Body.Close()
	defer io.Copy(ioutil.Discard, resp.Body)

	response := newResponse(resp)

	if err := checkResponse(resp); err != nil {
		return response, err
	}

	var err error

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return response, err
}

func (c *Client) Get(url string, v interface{}) (*Response, error) {
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *Client) GetWithQuery(url string, query interface{}, v interface{}) (*Response, error) {
	var err error
	if query != nil {
		url, err = c.addOptions(url, query)
		if err != nil {
			return nil, err
		}
	}
	return c.Get(url, v)
}

func (c *Client) Post(url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.newRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *Client) Patch(url string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.newRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req, v)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *Client) Delete(url string, body interface{}) (*Response, error) {
	req, err := c.newRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, nil)
}

// AddOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func (c *Client) addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
