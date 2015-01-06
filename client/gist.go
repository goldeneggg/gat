package client

import (
	"encoding/json"
	"fmt"
	"path"

	h "github.com/goldeneggg/gat/client/http"
)

const (
	// NameGist represents a factory key for "gist"
	NameGist = "gist"

	gistTimeout = 10
)

var (
	_ Client = &gist{}

	errGistDomain  = fmt.Errorf("api-domain is empty")
	errGistToken   = fmt.Errorf("access-token is empty")
	errGistPayload = fmt.Errorf("payload is empty")
)

type gist struct {
	APIDomain   string `json:"api-domain"`
	AccessToken string `json:"access-token"`
	Timeout     int    `json:"timeout"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

func newGist() *gist {
	return &gist{}
}

func (g *gist) CheckConf() error {
	if len(g.APIDomain) == 0 {
		return errGistDomain
	}

	if len(g.AccessToken) == 0 {
		return errGistToken
	}

	if g.Timeout < 0 {
		g.Timeout = gistTimeout
	}

	return nil
}

func (g *gist) Cat(catInf *CatInfo) (string, error) {
	createdURL, err := g.postGist(catInf.Files)
	if err != nil {
		return "", err
	}

	return createdURL, nil
}

func (g *gist) postGist(files map[string][]byte) (string, error) {
	pl, err := g.getPayload(files)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", errGistPayload
	}
	L.Debug("payload: ", string(pl))

	hr := &h.Req{
		Body:    pl,
		Headers: map[string]string{"Authorization": "token " + g.AccessToken},
	}

	respBody, err := hr.Post(g.APIDomain + "/gists")
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

	b, err := json.Marshal(pl)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func (g *gist) parseGistResp(respBody []byte) (string, error) {
	var respMap map[string]interface{}
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		return "", err
	}

	return respMap["html_url"].(string), nil
}
