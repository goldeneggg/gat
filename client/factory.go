package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// Default config file path (prefix dir is $HOME)
	CONFIG_PATH = ".gat/conf.json"
)

// Information to create new Client
type Attr struct {
	// Target service name
	Name string

	// Config file path
	ConfPath string

	// Key-Value to overwrite attributes of config json file
	Overwrites map[string]interface{}
}

// Create a new Client
func NewClient(attr Attr) (Client, error) {
	var clnt Client

	name := attr.Name
	switch name {
	case NAME_OSCAT:
		clnt = newOs()
	case NAME_GIST:
		clnt = newGist()
	case NAME_SLACK:
		clnt = newSlack()
	case NAME_PLAYGO:
		clnt = newPlaygo()
	default:
		return clnt, fmt.Errorf("invalid service name: " + name)
	}

	confPath := os.Getenv("HOME") + "/" + CONFIG_PATH
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
