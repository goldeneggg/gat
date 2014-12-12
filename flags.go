package main

import "github.com/codegangsta/cli"

// GlobalFlags are flags of gat command
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "confpath, c",
		Usage: "Your original config json path",
	},
	cli.BoolFlag{
		Name:  "debug, d",
		Usage: "Debug detail information",
	},
}
