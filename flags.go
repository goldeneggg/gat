package main

import (
	"github.com/urfave/cli"
)

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

// GistFlags are gist subcommand flags
var GistFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "api-domain",
		Usage: "Github api domain",
	},
	cli.StringFlag{
		Name:  "access-token",
		Usage: "Github api access token",
	},
	cli.IntFlag{
		Name:  "timeout",
		Usage: "Timeout for connection",
	},
	cli.StringFlag{
		Name:  "description, d",
		Usage: "A description of the gist",
	},
	cli.BoolFlag{
		Name:  "public, p",
		Usage: "Indicates whether the gist is public. Default: false",
	},
}

// SlackFlags are slack subcommand flags
var SlackFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "webhook-url",
		Usage: "Webhook URL",
	},
	cli.StringFlag{
		Name:  "channel, c",
		Usage: "Target channel",
	},
	cli.StringFlag{
		Name:  "username, u",
		Usage: "Username",
	},
	cli.StringFlag{
		Name:  "icon, i",
		Usage: "Icon url or emoji format text (:EMOJI_NAME:)",
	},
	cli.IntFlag{
		Name:  "timeout",
		Usage: "Timeout for connection",
	},
	cli.BoolFlag{
		Name:  "without-markdown",
		Usage: "Not format slack's markdown",
	},
	cli.BoolFlag{
		Name:  "without-unfurl",
		Usage: "Not unfurl media links",
	},
	cli.BoolFlag{
		Name:  "linkfy, l",
		Usage: "Linkify channel names (starting with a '#') and usernames (starting with an '@')",
	},
	cli.StringFlag{
		Name:  "color",
		Usage: "Color name or any hex color code",
	},
}

// HipchatFlags are hipchat subcommand flags
var HipchatFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "api-root",
		Usage: "API root URL",
	},
	cli.StringFlag{
		Name:  "access-token",
		Usage: "Hipchat API access token",
	},
	cli.StringFlag{
		Name:  "room, r",
		Usage: "Target room",
	},
	cli.StringFlag{
		Name:  "color, c",
		Usage: "Message color",
	},
	cli.BoolFlag{
		Name:  "notify, n",
		Usage: "Notify",
	},
	cli.StringFlag{
		Name:  "format, f",
		Usage: "Message format",
	},
}
