package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	h "github.com/goldeneggg/gat/client/http"
)

const (
	// NameSlack represents a factory key for "slack"
	NameSlack = "slack"

	slackTimeout = 10
)

var (
	_ Client = new(slack)

	errSlackToken   = fmt.Errorf("api-token is empty")
	errSlackPayload = fmt.Errorf("payload is empty")
)

type slack struct {
	WebhookURL      string `json:"webhook-url"`
	UserName        string `json:"username"`
	Icon            string `json:"icon"`
	Channel         string `json:"channel"`
	Timeout         int    `json:"timeout"`
	WithoutMarkdown bool   `json:"without-markdown"`
	WithoutUnfURL   bool   `json:"without-unfurl"`
	Linkfy          bool   `json:"linkfy"`
	Color           string `json:"color"`
	APIToken        string `json:"api-token"`
}

func newSlack() *slack {
	return new(slack)
}

func (s *slack) CheckConf() error {
	if len(s.APIToken) == 0 {
		return errSlackToken
	}

	if s.Timeout < 0 {
		s.Timeout = slackTimeout
	}

	return nil
}

func (s *slack) Cat(catInf *CatInfo) (string, error) {
	// migrated to web API
	//res, err := s.postSlack(catInf.Files)
	res, err := s.filesUpload(catInf.Files)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *slack) postSlack(files map[string][]byte) (string, error) {
	content := filesToSlackContent(files)

	pl, err := s.getPayload(content)
	if err != nil {
		return "", err
	} else if len(pl) == 2 {
		return "", errSlackPayload
	}
	L.Debug("payload: ", string(pl))

	hr := &h.Req{
		Body: pl,
	}

	respBody, err := hr.Post(s.WebhookURL)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func filesToSlackContent(files map[string][]byte) []byte {
	content := make([]byte, 0, 128)
	for _, in := range files {
		content = bytes.Join([][]byte{content, in}, []byte(""))
	}

	return content
}

// https://api.slack.com/docs/formatting
// https://api.slack.com/docs/attachments
// https://api.slack.com/docs/unfurling
type slackPayload struct {
	Text        string `json:"text"`
	UserName    string `json:"username,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	IconEmoji   string `json:"icon_emoji,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Mrkdwn      bool   `json:"mrkdwn"`
	UnfurlLinks bool   `json:"unfurl_links"`
	LinkNames   int    `json:"link_names,omitempty"`
	Color       string `json:"color,omitempty"`
}

func (s *slack) getPayload(content []byte) ([]byte, error) {
	pl := slackPayload{
		Text:        string(content),
		UserName:    s.UserName,
		Channel:     s.Channel,
		Mrkdwn:      !s.WithoutMarkdown,
		UnfurlLinks: !s.WithoutUnfURL,
		Color:       s.Color,
	}

	matched, err := regexp.MatchString(":[^:]+:", s.Icon)
	if err != nil {
		return []byte{}, err
	} else if matched {
		pl.IconEmoji = s.Icon
	} else {
		pl.IconURL = s.Icon
	}

	if s.Linkfy {
		pl.LinkNames = 1
	}

	b, err := json.Marshal(pl)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
