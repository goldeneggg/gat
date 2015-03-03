package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

import (
	h "github.com/goldeneggg/gat/client/http"
)

const (
	// NameHipchat represents a factory key for "hipchat"
	NameHipchat = "hipchat"

	hipchatTimeout = 10
)

var (
	_ Client = new(hipchat)

	errHipchatAPIRoot = fmt.Errorf("api-root is empty")
	errHipchatToken   = fmt.Errorf("access-token is empty")
	errHipchatRoom    = fmt.Errorf("room is empty")
	errHipchatPayload = fmt.Errorf("payload is empty")
)

type hipchat struct {
	APIRoot     string `json:"api-root"`
	AccessToken string `json:"access-token"`
	Room        string `json:"room"`
	Color       string `json:"color"`
	Notify      bool   `json:"notify"`
	Format      string `json:"format"`
	Timeout     int    `json:"timeout"`
}

func newHipchat() *hipchat {
	return new(hipchat)
}

func (hc *hipchat) CheckConf() error {
	if len(hc.APIRoot) == 0 {
		return errHipchatAPIRoot
	}

	if len(hc.AccessToken) == 0 {
		return errHipchatToken
	}

	if len(hc.Room) == 0 {
		return errHipchatRoom
	}

	if hc.Timeout < 0 {
		hc.Timeout = hipchatTimeout
	}

	return nil
}

func (hc *hipchat) Cat(catInf *CatInfo) (string, error) {
	res, err := hc.postHipchat(catInf.Files)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (hc *hipchat) postHipchat(files map[string][]byte) (string, error) {
	content := make([]byte, 0, 128)
	for _, in := range files {
		content = bytes.Join([][]byte{content, in}, []byte(""))
	}

	pl, err := hc.getPayload(content)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", errHipchatPayload
	}
	L.Debug("payload: ", string(pl))

	hr := &h.Req{
		Body: pl,
		Headers: map[string]string{
			"content-type":  "application/json",
			"Authorization": "Bearer " + hc.AccessToken,
		},
	}

	respBody, err := hr.Post(hc.getAPIURL())
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

type hipchatPayload struct {
	Message       string `json:"message"`
	Color         string `json:"color,omitempty"`
	Notify        bool   `json:"notify,omitempty"`
	MessageFormat string `json:"message_format,omitempty"`
}

func (hc *hipchat) getPayload(content []byte) ([]byte, error) {
	pl := hipchatPayload{
		Message:       string(content),
		Color:         hc.Color,
		Notify:        hc.Notify,
		MessageFormat: hc.Format,
	}

	b, err := json.Marshal(pl)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

func (hc *hipchat) getAPIURL() string {
	return fmt.Sprintf(hc.APIRoot+"/room/%s/notification", hc.Room)
}
