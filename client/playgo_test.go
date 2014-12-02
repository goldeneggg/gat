package client

import (
	"reflect"
	"testing"
)

type PlaygoTest struct {
	p          *playgo
	files      map[string][]byte
	expContent []byte
}

var pTests = []PlaygoTest{
	PlaygoTest{
		p:          &playgo{},
		files:      map[string][]byte{"a.go": []byte("content1")},
		expContent: []byte("content1"),
	},
}

func TestPlaygo(t *testing.T) {
	for _, te := range pTests {
		c, err := te.p.getContent(te.files)
		if err != nil {
			t.Errorf("getContent error: %#v, playgo: %#v", err, te.p)
		}
		if !reflect.DeepEqual(c, te.expContent) {
			t.Errorf("Content expected: %#v, but want: %#v", c, te.expContent)
		}
	}
}

var testsPlaygoError = []PlaygoTest{
	PlaygoTest{
		p:     &playgo{},
		files: map[string][]byte{"z.java": []byte("content1")},
	},
	PlaygoTest{
		p:     &playgo{},
		files: map[string][]byte{"x.go": []byte("content1"), "y.go": []byte("content2")},
	},
}

func TestPlaygoError(t *testing.T) {
	for _, te := range testsPlaygoError {
		if _, err := te.p.getContent(te.files); err == nil {
			t.Errorf("getContent not occured expected error: files: %#v", te.files)
		}
	}
}
