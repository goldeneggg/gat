package client

import (
	"reflect"
	"testing"
)

const (
	ftDir            = "../test/factory_test"
	ftConf           = ftDir + "/test.json"
	ftConfTypoNonReq = ftDir + "/test_typo_non_require.json"
	ftConfEmp        = ftDir + "/test_emp.json"
	ftConfEmpCtnt    = ftDir + "/test_emp_content.json"
	ftConfFmtErr     = ftDir + "/test_format_err.json"
	ftConfTooNest    = ftDir + "/test_toomuch_nest.json"
	ftConfTypoReq    = ftDir + "/test_typo_require.json"
	ftTxt            = ftDir + "/test.txt"
	ftEmp            = ftDir + "/test_emp"
)

type FactoryTest struct {
	attr Attr
	exp  Client
}

var factoryTests = []FactoryTest{
	FactoryTest{
		attr: Attr{
			Name:     NameOscat,
			ConfPath: ftConf,
		},
		exp: &oscat{Number: true, NumberNonBlank: true},
	},
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConf,
		},
		exp: &gist{APIDomain: "https://factory-test.com", AccessToken: "token", Timeout: 5, Description: "desc", Public: true},
	},
	FactoryTest{
		attr: Attr{
			Name:     NameSlack,
			ConfPath: ftConf,
		},
		exp: &slack{WebhookURL: "https://webhook-url.com", UserName: "user", Icon: "icon", Channel: "channel", Timeout: 6, WithoutMarkdown: true, WithoutUnfURL: true, Linkfy: true},
	},
	FactoryTest{
		attr: Attr{
			Name:     NameOscat,
			ConfPath: ftConfTypoNonReq,
		},
		exp: &oscat{Number: false, NumberNonBlank: false},
	},
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfTypoNonReq,
		},
		exp: &gist{APIDomain: "https://factory-test2.com", AccessToken: "token2", Timeout: 0, Description: "", Public: false},
	},
	FactoryTest{
		attr: Attr{
			Name:     NameSlack,
			ConfPath: ftConfTypoNonReq,
		},
		exp: &slack{WebhookURL: "https://webhook-url2.com", UserName: "", Icon: "", Channel: "", Timeout: 0, WithoutMarkdown: false, WithoutUnfURL: false, Linkfy: false},
	},
}

func TestFactory(t *testing.T) {
	for _, te := range factoryTests {
		if c, err := NewClient(te.attr); err != nil {
			t.Errorf("NewClient error: %v, attr: %#v", err, te.attr)
		} else {
			if !reflect.DeepEqual(te.exp, c) {
				t.Errorf("Unexpected client: %#v, expected: %#v", c, te.exp)
			}
		}
	}
}

var factoryErrTests = []FactoryTest{
	// invalid key name
	FactoryTest{
		attr: Attr{
			Name:     "invalid",
			ConfPath: ftConf,
		},
	},
	FactoryTest{
		attr: Attr{
			Name:     "",
			ConfPath: ftConf,
		},
	},
	FactoryTest{
		attr: Attr{
			ConfPath: ftConf,
		},
	},
	// conf not found
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConf + ".notfound",
		},
	},
	// conf is dir
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftDir,
		},
	},
	// empty file
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftEmp,
		},
	},
	// text file
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftTxt,
		},
	},
	// empty json
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfEmp,
		},
	},
	// empty content json
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfEmpCtnt,
		},
	},
	// format error json
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfFmtErr,
		},
	},
	// toomuch nest json
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfTooNest,
		},
	},
	// typo require key json
	FactoryTest{
		attr: Attr{
			Name:     NameGist,
			ConfPath: ftConfTypoReq,
		},
	},
}

func TestFactoryError(t *testing.T) {
	for _, te := range factoryErrTests {
		if _, err := NewClient(te.attr); err == nil {
			t.Errorf("NewClient not occured expected error: %#v", te.attr)
		}
	}
}
