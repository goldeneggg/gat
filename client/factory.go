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
	var clnt Client

	name := attr.Name
	switch name {
	case NameOscat:
		clnt = newOs()
	case NameGist:
		clnt = newGist()
	case NameSlack:
		clnt = newSlack()
	case NamePlaygo:
		clnt = newPlaygo()
	default:
		return clnt, fmt.Errorf("invalid service name: " + name)
	}

	confPath := os.Getenv("HOME") + "/" + configPath
	if len(attr.ConfPath) > 0 {
		confPath = attr.ConfPath
	}

	if err := configure(name, confPath, attr.Overwrites, clnt); err != nil {
		return clnt, err
	}

	if err := clnt.CheckConf(); err != nil {
		return clnt, err
	}

	return clnt, nil
}

func configure(name string,
	confPath string,
	attrs map[string]interface{},
	client interface{}) error {

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
	if !ok {
		return fmt.Errorf("not exist config name: " + name)
	}

	// overwrite by original attributes
	for k, v := range attrs {
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

	return nil
}
