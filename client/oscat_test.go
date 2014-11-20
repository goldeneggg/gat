package client

import "testing"

type oscatTest struct {
	*oscat
	*CatInfo
}

const (
	testOscatDir = "../test/oscat_test"
)

var oTests = []oscatTest{
	oscatTest{
		&oscat{Number: false, NumberNonBlank: false},
		NewCatInfo(map[string][]byte{testOscatDir + "/test.txt": []byte("dummy1")}),
	},
	oscatTest{
		&oscat{Number: true, NumberNonBlank: false},
		NewCatInfo(map[string][]byte{testOscatDir + "/test2.txt": []byte("dummy2")}),
	},
	oscatTest{
		&oscat{Number: false, NumberNonBlank: true},
		NewCatInfo(map[string][]byte{testOscatDir + "/test3.txt": []byte("dummy3")}),
	},
	oscatTest{
		&oscat{Number: true, NumberNonBlank: true},
		NewCatInfo(map[string][]byte{testOscatDir + "/test_empty.txt": []byte("dummy4")}),
	},
	oscatTest{
		&oscat{Number: false, NumberNonBlank: false},
		NewCatInfo(map[string][]byte{"/dev/stdin": []byte("5")}),
	},
}

func TestCheckConfNormal(t *testing.T) {
	for _, te := range oTests {
		if err := te.oscat.CheckConf(); err != nil {
			t.Errorf("CheckConf error: %v, oscat: %#v, CatInfo: %#v", err, te.oscat, te.CatInfo)
		}
	}
}

func TestCatNormal(t *testing.T) {
	for _, te := range oTests {
		if _, err := te.oscat.Cat(te.CatInfo); err != nil {
			t.Errorf("Cat error: %v, oscat: %#v, CatInfo: %#v", err, te.oscat, te.CatInfo)
		}
	}
}
