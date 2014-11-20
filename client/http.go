package client

import "net/http"

type respHandler func(http.Response) (string, error)

type httpClient interface {
	getPayload(map[string][]byte) (string, error)
}

func (rh respHandler) post(files map[string][]byte) (string, error) {
	var resp http.Response
	return rh(resp)
}
