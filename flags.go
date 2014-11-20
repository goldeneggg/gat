package main

import "github.com/codegangsta/cli"

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "confpath, c",
		Usage: "Your original config json path",
	},
	cli.BoolFlag{
		Name:  "debug, d",
		Usage: "Debug detail information",
	},
}
