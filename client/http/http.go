package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 10
)

var (
	// ToSec is timeout duration second
	ToSec = defaultTimeout

	errInvalidStatus = func(s int) error { return fmt.Errorf("invalid status: %d", s) }
)

// A Req is http request attributes
type Req struct {
	Body    []byte
	Queries map[string]string
	Headers map[string]string
}

// Post requests contents to argument URL by POST method
func (req *Req) Post(u string) ([]byte, error) {
	return req.request("POST", u)
}

// Get requests contents to argument URL by GET method
func (req *Req) Get(u string) ([]byte, error) {
	return req.request("GET", u)
}

func (req *Req) request(m string, u string) ([]byte, error) {
	newRequest, err := http.NewRequest(m, u, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	req.buildRequestQuery(newRequest)

	req.buildRequestHeader(newRequest)

	c := &http.Client{Timeout: time.Duration(ToSec) * time.Second}
	resp, err := c.Do(newRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errInvalidStatus(resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (req *Req) buildRequestQuery(r *http.Request) {
	values := url.Values{}
	for k, v := range req.Queries {
		values.Set(k, v)
	}

	r.URL.RawQuery = values.Encode()
}

func (req *Req) buildRequestHeader(r *http.Request) {
	for k, v := range req.Headers {
		r.Header.Set(k, v)
	}
}
