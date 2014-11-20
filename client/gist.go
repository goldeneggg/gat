package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

const (
	NAME_GIST   = "gist"
	gistTimeout = 10
)

type gist struct {
	ApiDomain   string `json:"api-domain"`
	AccessToken string `json:"access-token"`
	Timeout     int    `json:"timeout"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

func newGist() *gist {
	return &gist{}
}

func (g *gist) CheckConf() error {
	if len(g.ApiDomain) == 0 {
		return fmt.Errorf("api-domain is empty")
	}

	if len(g.AccessToken) == 0 {
		return fmt.Errorf("access-token is empty")
	}

	if g.Timeout < 0 {
		g.Timeout = gistTimeout
	}

	return nil
}

func (g *gist) Cat(catInf *CatInfo) (string, error) {
	if createdUrl, err := g.postGist(catInf.Files); err != nil {
		return "", err
	} else {
		return createdUrl, nil
	}
}

func (g *gist) CatP(catInf *CatInfo, chOut chan string, chErr chan error) {
	if createdUrl, err := g.Cat(catInf); err != nil {
		chErr <- err
	} else {
		chOut <- createdUrl
	}
}

func (g *gist) postGist(files map[string][]byte) (string, error) {
	pl, err := g.getPayload(files)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", fmt.Errorf("payload is empty")
	}
	L.Debug("payload: ", string(pl))

	req, err := g.getRequest(pl)
	if err != nil {
		return "", err
	}

	c := &http.Client{Timeout: time.Duration(g.Timeout) * time.Second}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("invalid status: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return g.parseGistResp(respBody)
}

type gistPayload struct {
	Description string                       `json:"description,omitempty"`
	Public      bool                         `json:"public,omitempty"`
	Files       map[string]map[string]string `json:"files"`
}

func (g *gist) getPayload(files map[string][]byte) ([]byte, error) {
	plFiles := make(map[string]map[string]string)
	for file, in := range files {
		// Note:
		// If file includes directory separator like "/" (ex. /dev/stdin),
		// this name will occur github api error "422 Unprocessable Entity"
		plFiles[path.Base(file)] = map[string]string{"content": string(in)}
	}

	pl := gistPayload{
		Description: g.Description,
		Public:      g.Public,
		Files:       plFiles,
	}

	if b, err := json.Marshal(pl); err != nil {
		return []byte{}, err
	} else {
		return b, nil
	}
}

func (g *gist) getRequest(payload []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", g.ApiDomain+"/gists", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+g.AccessToken)

	return req, nil
}

func (g *gist) parseGistResp(respBody []byte) (string, error) {
	var respMap map[string]interface{}
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return "", err
	}

	return respMap["html_url"].(string), nil
}
