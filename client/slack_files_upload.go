package client

import (
	"path/filepath"

	h "github.com/goldeneggg/gat/client/http"
)

const (
	apiRoot = "https://slack.com/api"
)

func (s *slack) filesUpload(files map[string][]byte) (string, error) {
	var resps string

	for fileName, content := range files {
		queries := map[string]string{
			"token":    s.APIToken,
			"filetype": filepath.Ext(fileName),
			"filename": fileName,
			"title":    fileName,
			"content":  string(content),
			"channels": s.Channel,
		}

		hr := &h.Req{
			Queries: queries,
		}

		respBody, err := hr.Post(apiRoot + "/files.upload")
		if err != nil {
			return "", err
		}

		resps += string(respBody)
	}

	return resps, nil
}
