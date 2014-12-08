package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	h "github.com/goldeneggg/gat/client/http"
)

const (
	NAME_SLACK   = "slack"
	slackTimeout = 10
)

type slack struct {
	WebhookUrl      string `json:"webhook-url"`
	UserName        string `json:"username"`
	Icon            string `json:"icon"`
	Channel         string `json:"channel"`
	Timeout         int    `json:"timeout"`
	WithoutMarkdown bool   `json:"without-markdown"`
	WithoutUnfUrl   bool   `json:"without-unfurl"`
	Linkfy          bool   `json:"linkfy"`
}

func newSlack() *slack {
	return &slack{}
}

func (s *slack) CheckConf() error {
	if len(s.WebhookUrl) == 0 {
		return fmt.Errorf("webhook-url is empty")
	}

	if s.Timeout < 0 {
		s.Timeout = slackTimeout
	}

	return nil
}

func (s *slack) Cat(catInf *CatInfo) (string, error) {
	if res, err := s.postSlack(catInf.Files); err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (s *slack) CatP(catInf *CatInfo, chOut chan string, chErr chan error) {
	if res, err := s.Cat(catInf); err != nil {
		chErr <- err
	} else {
		chOut <- res
	}
}

func (s *slack) postSlack(files map[string][]byte) (string, error) {
	content := make([]byte, 0, 128)
	for _, in := range files {
		content = bytes.Join([][]byte{content, in}, []byte(""))
	}

	pl, err := s.getPayload(content)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", fmt.Errorf("payload is empty")
	}
	L.Debug("payload: ", string(pl))

	hr := &h.HttpReq{
		Body: pl,
	}

	respBody, err := hr.Post(s.WebhookUrl)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

type slackPayload struct {
	Text        string `json:"text"`
	UserName    string `json:"username,omitempty"`
	IconUrl     string `json:"icon_url,omitempty"`
	IconEmoji   string `json:"icon_emoji,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Mrkdwn      bool   `json:"mrkdwn"`
	UnfurlLinks bool   `json:"unfurl_links"`
	LinkNames   int    `json:"link_names,omitempty"`
}

func (s *slack) getPayload(content []byte) ([]byte, error) {
	pl := slackPayload{
		Text:        string(content),
		UserName:    s.UserName,
		Channel:     s.Channel,
		Mrkdwn:      !s.WithoutMarkdown,
		UnfurlLinks: !s.WithoutUnfUrl,
	}

	matched, err := regexp.MatchString(":[^:]+:", s.Icon)
	if err != nil {
		return []byte{}, err
	} else if matched {
		pl.IconEmoji = s.Icon
	} else {
		pl.IconUrl = s.Icon
	}

	if s.Linkfy {
		pl.LinkNames = 1
	}

	if b, err := json.Marshal(pl); err != nil {
		return []byte{}, err
	} else {
		return b, nil
	}
}
