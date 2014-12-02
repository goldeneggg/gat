package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	NAME_PLAYGO    = "playgo"
	playgoURL      = "https://play.golang.org/share"
	shareURLPrefix = "https://play.golang.org/p/"
)

type playgo struct {
	WithRun bool `json:"with-run"`
}

func newPlaygo() *playgo {
	return &playgo{}
}

func (p *playgo) CheckConf() error {
	return nil
}

func (p *playgo) Cat(catInf *CatInfo) (string, error) {
	if res, err := p.postPlaygo(catInf.Files); err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (p *playgo) CatP(catInf *CatInfo, chOut chan string, chErr chan error) {
	if res, err := p.Cat(catInf); err != nil {
		chErr <- err
	} else {
		chOut <- res
	}
}

func (p *playgo) postPlaygo(files map[string][]byte) (string, error) {
	content := make([]byte, 0, 128)
	for _, in := range files {
		content = bytes.Join([][]byte{content, in}, []byte(""))
	}

	req, err := p.getRequest(content)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("invalid status: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return shareURLPrefix + string(respBody), nil
}

func (p *playgo) getRequest(content []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", playgoURL, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	return req, nil
}
