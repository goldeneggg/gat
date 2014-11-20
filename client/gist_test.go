package client

import "testing"

type GistTest struct {
	g          *gist
	files      map[string][]byte
	expPayload string
	respBody   []byte
	expUrl     string
}

var gTests = []GistTest{
	GistTest{
		g: &gist{
			ApiDomain:   "domain1",
			AccessToken: "token1",
			Timeout:     0,
			Description: "",
			Public:      false,
		},
		files: map[string][]byte{
			"test1.txt": []byte("text content"),
		},
		expPayload: `{"files":{"test1.txt":{"content":"text content"}}}`,
		respBody:   []byte(`{"html_url":"http://gist.github.com/goldeneggg/1"}`),
		expUrl:     "http://gist.github.com/goldeneggg/1",
	},
	GistTest{
		g: &gist{
			ApiDomain:   "domain2",
			AccessToken: "token2",
			Timeout:     15,
			Description: "desc",
			Public:      false,
		},
		files: map[string][]byte{
			"subdir/test2.java": []byte("java content"),
		},
		expPayload: `{"description":"desc","files":{"test2.java":{"content":"java content"}}}`,
		respBody:   []byte(`{"html_url":"http://gist.github.com/goldeneggg/2"}`),
		expUrl:     "http://gist.github.com/goldeneggg/2",
	},
	GistTest{
		g: &gist{
			ApiDomain:   "domain3",
			AccessToken: "token3",
			Timeout:     -1,
			Description: "",
			Public:      true,
		},
		files: map[string][]byte{
			"test31.rb": []byte("ruby content 1"),
			"test32.rb": []byte("ruby content 2"),
		},
		expPayload: `{"public":true,"files":{"test31.rb":{"content":"ruby content 1"},"test32.rb":{"content":"ruby content 2"}}}`,
		respBody:   []byte(`{"html_url":"http://gist.github.com/goldeneggg/3"}`),
		expUrl:     "http://gist.github.com/goldeneggg/3",
	},
	GistTest{
		g: &gist{
			ApiDomain:   "domain4",
			AccessToken: "token4",
			Timeout:     30,
			Description: "desc",
			Public:      true,
		},
		files: map[string][]byte{
			"/dev/stdin": []byte("stdin content"),
		},
		expPayload: `{"description":"desc","public":true,"files":{"stdin":{"content":"stdin content"}}}`,
		respBody:   []byte(`{"html_url":"http://gist.github.com/goldeneggg/4"}`),
		expUrl:     "http://gist.github.com/goldeneggg/4",
	},
}

func TestGist(t *testing.T) {
	for _, te := range gTests {
		if err := te.g.CheckConf(); err != nil {
			t.Errorf("CheckConf error: %#v, gist: %#v", err, te.g)
		}

		pl, errP := te.g.getPayload(te.files)
		if errP != nil {
			t.Errorf("getPayload error: %#v", errP)
		}
		if string(pl) != te.expPayload {
			t.Errorf("getPayload unexpected payload: %s, expected: %s", pl, te.expPayload)
		}

		u, errU := te.g.parseGistResp(te.respBody)
		if errU != nil {
			t.Errorf("parseGistResp error: %#v", errU)
		}
		if u != te.expUrl {
			t.Errorf("parseGistResp unexpected url: %s, expected: %s", u, te.expUrl)
		}
	}
}

var gErrTests = []GistTest{
	GistTest{
		g: &gist{
			ApiDomain:   "",
			AccessToken: "tokenE1",
			Description: "",
			Public:      false,
		},
		files: map[string][]byte{
			"": []byte(""),
		},
	},
	GistTest{
		g: &gist{
			ApiDomain:   "domainE2",
			AccessToken: "",
			Timeout:     300,
			Description: "desc",
			Public:      false,
		},
		files: map[string][]byte{
			"": []byte(""),
		},
	},
	GistTest{
		g: &gist{
			ApiDomain:   "",
			AccessToken: "",
			Description: "",
			Public:      true,
		},
		files: map[string][]byte{
			"": []byte(""),
		},
	},
	GistTest{
		g: &gist{
			ApiDomain:   "domainE4",
			Description: "desc",
			Public:      true,
		},
		files: map[string][]byte{
			"": []byte(""),
		},
	},
	GistTest{
		g: &gist{
			AccessToken: "tokenE5",
			Description: "",
			Public:      false,
		},
		files: map[string][]byte{
			"": []byte(""),
		},
	},
}

func TestGistErrorConf(t *testing.T) {
	for _, te := range gErrTests {
		if err := te.g.CheckConf(); err == nil {
			t.Errorf("CheckConf not occured expected error: gist: %#v",
				te.g)
		}
	}
}
