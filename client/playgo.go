package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	NAME_PLAYGO    = "playgo"
	playgoURL      = "https://play.golang.org/share"
	shareURLPrefix = "https://play.golang.org/p/"
)

type playgo struct{}

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
	content, err := p.getContent(files)
	if err != nil {
		return "", err
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

func (p *playgo) getContent(files map[string][]byte) ([]byte, error) {
	if len(files) > 1 {
		return []byte{}, fmt.Errorf("Only 1 go file is available, but not 1: %d", len(files))
	}

	var content []byte
	for f, in := range files {
		if !strings.HasSuffix(f, ".go") {
			return []byte{}, fmt.Errorf("file %s is not .go file", f)
		}
		content = in
	}

	return content, nil
}

func (p *playgo) getRequest(content []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", playgoURL, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	return req, nil
}
