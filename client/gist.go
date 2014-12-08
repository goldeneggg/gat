package client

import (
	"encoding/json"
	"fmt"
	"path"

	h "github.com/goldeneggg/gat/client/http"
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

var _ Client = &gist{}

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

func (g *gist) postGist(files map[string][]byte) (string, error) {
	pl, err := g.getPayload(files)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", fmt.Errorf("payload is empty")
	}
	L.Debug("payload: ", string(pl))

	hr := &h.HttpReq{
		Body:    pl,
		Headers: map[string]string{"Authorization": "token " + g.AccessToken},
	}

	respBody, err := hr.Post(g.ApiDomain + "/gists")
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

func (g *gist) parseGistResp(respBody []byte) (string, error) {
	var respMap map[string]interface{}
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return "", err
	}

	return respMap["html_url"].(string), nil
}
