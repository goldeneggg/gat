package client

import (
	"fmt"
	"strings"

	h "github.com/goldeneggg/gat/client/http"
)

const (
	// NamePlaygo represents a factory key for "playgo"
	NamePlaygo = "playgo"

	playgoURL      = "https://play.golang.org/share"
	shareURLPrefix = "https://play.golang.org/p/"
)

var (
	_ Client = new(playgo)

	errPlaygoNothing = func(l int) error { return fmt.Errorf("Only 1 go file is available, but not 1: %d", l) }
	errPlaygoNotGo   = func(f string) error { return fmt.Errorf("file %s is not .go file", f) }
)

type playgo struct{}

func newPlaygo() *playgo {
	return new(playgo)
}

func (p *playgo) CheckConf() error {
	return nil
}

func (p *playgo) Cat(catInf *CatInfo) (string, error) {
	res, err := p.postPlaygo(catInf.Files)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (p *playgo) postPlaygo(files map[string][]byte) (string, error) {
	content, err := p.getContent(files)
	if err != nil {
		return "", err
	}

	hr := &h.Req{
		Body: content,
	}

	respBody, err := hr.Post(playgoURL)
	if err != nil {
		return "", err
	}

	return shareURLPrefix + string(respBody), nil
}

func (p *playgo) getContent(files map[string][]byte) ([]byte, error) {
	if len(files) > 1 {
		return []byte{}, errPlaygoNothing(len(files))
	}

	var content []byte
	for f, in := range files {
		if !strings.HasSuffix(f, ".go") {
			return []byte{}, errPlaygoNotGo(f)
		}
		content = in
	}

	return content, nil
}
