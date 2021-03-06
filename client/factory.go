package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	configPath = ".gat/conf.json"
)

var (
	errConfJSON = func(path string, orgError error) error {
		return fmt.Errorf("conf json %s is invalid: %v", path, orgError)
	}
	errCheckConf = func(client Client, orgError error) error {
		return fmt.Errorf("%v conf includes invalid setting: %v", client, orgError)
	}
	errFactoryName = func(name string) error { return fmt.Errorf("invalid name: %s", name) }
)

// Attr represents the configuration for NewClient
type Attr struct {
	// Target service name
	Name string

	// Config file path
	ConfPath string

	// Key-Value to overwrite attributes of config json file
	Overwrites map[string]interface{}
}

// NewClient returns a new Client,
func NewClient(attr Attr) (Client, error) {
	var client Client

	name := attr.Name
	switch name {
	case NameOscat:
		client = newOs()
	case NameGist:
		client = newGist()
	case NameSlack:
		client = newSlack()
	case NamePlaygo:
		client = newPlaygo()
	case NameHipchat:
		client = newHipchat()
	default:
		return client, errFactoryName(name)
	}

	confPath := os.Getenv("HOME") + "/" + configPath
	if len(attr.ConfPath) > 0 {
		confPath = attr.ConfPath
	}

	if err := configure(name, confPath, attr.Overwrites, client); err != nil {
		return client, errConfJSON(confPath, err)
	}

	if err := client.CheckConf(); err != nil {
		return client, errCheckConf(client, err)
	}

	return client, nil
}

func configure(name string, confPath string, ows map[string]interface{}, client interface{}) error {

	f, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	// unmarshal to map
	var m map[string]map[string]interface{}
	if err := json.Unmarshal(f, &m); err != nil {
		return err
	}

	tMap, ok := m[name]
	if ok {
		// overwrite by original attributes
		for k, v := range ows {
			tMap[k] = v
		}

		// marshal from map to bytes(json)
		b, err := json.Marshal(tMap)
		if err != nil {
			return err
		}

		// unmarshal to target client struct
		if err := json.Unmarshal(b, client); err != nil {
			return err
		}
	}

	return nil
}
