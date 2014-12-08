package client

import "testing"

type SlackTest struct {
	s          *slack
	content    []byte
	expPayload string
}

var sTests = []SlackTest{
	SlackTest{
		s: &slack{
			WebhookURL: "hook1",
		},
		content:    []byte("content1"),
		expPayload: `{"text":"content1","mrkdwn":true,"unfurl_links":true}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL: "hook2",
			UserName:   "user2",
		},
		content:    []byte("content2"),
		expPayload: `{"text":"content2","username":"user2","mrkdwn":true,"unfurl_links":true}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL: "hook3",
			UserName:   "user3",
			Icon:       "icon3",
		},
		content:    []byte("content3"),
		expPayload: `{"text":"content3","username":"user3","icon_url":"icon3","mrkdwn":true,"unfurl_links":true}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL: "hook4",
			UserName:   "user4",
			Icon:       "icon4",
			Channel:    "channel4",
		},
		content:    []byte("content4"),
		expPayload: `{"text":"content4","username":"user4","icon_url":"icon4","channel":"channel4","mrkdwn":true,"unfurl_links":true}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL:      "hook5",
			UserName:        "user5",
			Icon:            "icon5",
			Channel:         "channel5",
			WithoutMarkdown: true,
		},
		content:    []byte("content5"),
		expPayload: `{"text":"content5","username":"user5","icon_url":"icon5","channel":"channel5","mrkdwn":false,"unfurl_links":true}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL:      "hook6",
			UserName:        "user6",
			Icon:            "icon6",
			Channel:         "channel6",
			WithoutMarkdown: true,
			WithoutUnfURL:   true,
		},
		content:    []byte("content6"),
		expPayload: `{"text":"content6","username":"user6","icon_url":"icon6","channel":"channel6","mrkdwn":false,"unfurl_links":false}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL:      "hook7",
			UserName:        "user7",
			Icon:            ":icon7:",
			Channel:         "channel7",
			WithoutMarkdown: true,
			WithoutUnfURL:   true,
		},
		content:    []byte("content7"),
		expPayload: `{"text":"content7","username":"user7","icon_emoji":":icon7:","channel":"channel7","mrkdwn":false,"unfurl_links":false}`,
	},
	SlackTest{
		s: &slack{
			WebhookURL:      "hook8",
			UserName:        "user8",
			Icon:            "::",
			Channel:         "channel8",
			WithoutMarkdown: true,
			WithoutUnfURL:   true,
			Linkfy:          true,
		},
		content:    []byte("content8"),
		expPayload: `{"text":"content8","username":"user8","icon_url":"::","channel":"channel8","mrkdwn":false,"unfurl_links":false,"link_names":1}`,
	},
}

func TestSlack(t *testing.T) {
	for _, te := range sTests {
		if err := te.s.CheckConf(); err != nil {
			t.Errorf("CheckConf error: %#v, slack: %#v", err, te.s)
		}

		pl, errP := te.s.getPayload(te.content)
		if errP != nil {
			t.Errorf("getPayload error: %#v", errP)
		}
		if string(pl) != te.expPayload {
			t.Errorf("getPayload unexpected payload: %s, expected: %s", pl, te.expPayload)
		}
	}
}

var testsSlackError = []SlackTest{
	SlackTest{
		s: &slack{
			WebhookURL: "",
		},
	},
	SlackTest{
		s: &slack{},
	},
}

func TestSlackCheckConfError(t *testing.T) {
	for _, te := range testsSlackError {
		if err := te.s.CheckConf(); err == nil {
			t.Errorf("CheckConf not occured expected error: slack: %#v", te.s)
		}
	}
}
