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
	DEFAULT_TIMEOUT = 10
)

var ToSec = DEFAULT_TIMEOUT

type HttpReq struct {
	Body    []byte
	Queries map[string]string
	Headers map[string]string
}

func (hr *HttpReq) Post(u string) ([]byte, error) {
	return hr.req("POST", u)
}

func (hr *HttpReq) Get(u string) ([]byte, error) {
	return hr.req("GET", u)
}

func (hr *HttpReq) req(m string, u string) ([]byte, error) {
	req, err := http.NewRequest(m, u, bytes.NewReader(hr.Body))
	if err != nil {
		return nil, err
	}

	hr.buildQuery(req)

	hr.buildHeader(req)

	c := &http.Client{Timeout: time.Duration(ToSec) * time.Second}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("invalid status: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (hr *HttpReq) buildQuery(r *http.Request) {
	values := url.Values{}
	for k, v := range hr.Queries {
		values.Set(k, v)
	}

	r.URL.RawQuery = values.Encode()
}

func (hr *HttpReq) buildHeader(r *http.Request) {
	for k, v := range hr.Headers {
		r.Header.Set(k, v)
	}
}
