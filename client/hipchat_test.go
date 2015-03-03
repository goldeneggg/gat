package client

import "testing"

type HipchatTest struct {
	hc         *hipchat
	content    []byte
	expURL     string
	expPayload string
}

var hcTests = []HipchatTest{
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctest1/v2",
			AccessToken: "token1",
			Room:        "room1",
		},
		content:    []byte("content1"),
		expURL:     "https://hctest1/v2/room/room1/notification",
		expPayload: `{"message":"content1"}`,
	},
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctest2/v2",
			AccessToken: "token2",
			Room:        "room2",
			Color:       "color2",
		},
		content:    []byte("content2"),
		expURL:     "https://hctest2/v2/room/room2/notification",
		expPayload: `{"message":"content2","color":"color2"}`,
	},
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctest3/v2",
			AccessToken: "token3",
			Room:        "room3",
			Color:       "color3",
			Notify:      true,
		},
		content:    []byte("content3"),
		expURL:     "https://hctest3/v2/room/room3/notification",
		expPayload: `{"message":"content3","color":"color3","notify":true}`,
	},
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctest4/v2",
			AccessToken: "token4",
			Room:        "room4",
			Color:       "color4",
			Notify:      false,
			Format:      "text4",
		},
		content:    []byte("content4"),
		expURL:     "https://hctest4/v2/room/room4/notification",
		expPayload: `{"message":"content4","color":"color4","message_format":"text4"}`,
	},
}

func TestHipchat(t *testing.T) {
	for _, te := range hcTests {
		if err := te.hc.CheckConf(); err != nil {
			t.Errorf("CheckConf error: %#v, hipchat: %#v", err, te.hc)
		}

		u := te.hc.getAPIURL()
		if u != te.expURL {
			t.Errorf("getAPIURL unexpected url: %s, expected: %s", u, te.expURL)
		}
		pl, errP := te.hc.getPayload(te.content)
		if errP != nil {
			t.Errorf("getPayload error: %#v", errP)
		}
		if string(pl) != te.expPayload {
			t.Errorf("getPayload unexpected payload: %s, expected: %s", pl, te.expPayload)
		}
	}
}

var testsHipchatError = []HipchatTest{
	HipchatTest{
		hc: &hipchat{
			APIRoot: "",
		},
	},
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctesterr2/v2",
			AccessToken: "",
		},
	},
	HipchatTest{
		hc: &hipchat{
			APIRoot:     "https://hctesterr3/v2",
			AccessToken: "token3",
			Room:        "",
		},
	},
	HipchatTest{
		hc: &hipchat{},
	},
}

func TestHipchatCheckConfError(t *testing.T) {
	for _, te := range testsHipchatError {
		if err := te.hc.CheckConf(); err == nil {
			t.Errorf("CheckConf not occured expected error: hipchat: %#v", te.hc)
		}
	}
}
