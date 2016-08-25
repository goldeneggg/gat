package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

import (
	"github.com/goldeneggg/gat/client"
	"github.com/urfave/cli"
)

// Commands are executable commands.
var (
	Commands = []cli.Command{
		cli.Command{
			Name:  client.NameGist,
			Usage: "Upload file to gist",
			Flags: []cli.Flag{
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
			},
			Action: exec,
		},
		cli.Command{
			Name:  client.NameSlack,
			Usage: "Send file contents to slack",
			Flags: []cli.Flag{
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
			},
			Action: exec,
		},
		cli.Command{
			Name:   client.NamePlaygo,
			Usage:  "Upload go code to play.golang.org",
			Action: exec,
		},
		cli.Command{
			Name:  client.NameHipchat,
			Usage: "Send file contents to hipchat",
			Flags: []cli.Flag{
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
			},
			Action: exec,
		},
		/*
			cli.Command{
				Name:  client.NameOscat,
				Usage: "Cat using os cat",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "n, number",
						Usage: "Number all output lines",
					},
					cli.BoolFlag{
						Name:  "b, number-nonblank",
						Usage: "Number nonempty output lines",
					},
				},
				Action: exec,
			},
		*/
		cli.Command{
			Name:   "list",
			Usage:  "Show target service list",
			Action: list,
		},
	}

	errInputIsEmpty = func(n string) error { return fmt.Errorf("input %s is empty", n) }
	errInputIsDir   = func(n string) error { return fmt.Errorf("input %s is directory", n) }
)

func exec(c *cli.Context) {
	isDebug := c.GlobalBool("debug")
	client.PrepareLogger(isDebug)

	start := time.Now()
	client.L.Debug("START")

	attr := client.Attr{
		Name:       c.Command.Name,
		ConfPath:   c.GlobalString("confpath"),
		Overwrites: flags2map(c),
	}
	client.L.DebugF("attr: %v", attr)

	clnt, errN := client.NewClient(attr)
	if errN != nil {
		fmt.Fprintf(os.Stderr, "%v\n", errN)
		exitSts = 1
		return
	}
	client.L.DebugF("clnt: %v", clnt)

	catInf, errB := buildCatInfo(c)
	if errB != nil {
		fmt.Fprintf(os.Stderr, "%v\n", errB)
		exitSts = 1
		return
	}
	for k := range catInf.Files {
		client.L.DebugF("catInf files: %v", k)
	}

	res, errC := clnt.Cat(catInf)
	if errC != nil {
		fmt.Fprintf(os.Stderr, "%v", errC)
		exitSts = 1
		return
	}
	client.L.DebugF("res: %v", res)

	elapsed := time.Since(start)
	client.L.Debug("END. elapsed: ", elapsed)

	fmt.Fprintf(os.Stdout, "%s\n", res)
}

func flags2map(c *cli.Context) map[string]interface{} {
	m := make(map[string]interface{})

	for _, fName := range c.FlagNames() {
		if c.IsSet(fName) {
			m[fName] = c.Generic(fName)
		}
	}

	return m
}

func buildCatInfo(c *cli.Context) (*client.CatInfo, error) {
	files, err := getInputFiles(c)
	if err != nil {
		return nil, err
	}

	fm := make(map[string][]byte)

	for _, file := range files {
		byteIn, err := readInput(file)
		if err != nil {
			return nil, err
		}
		fm[file.Name()] = byteIn
	}

	return client.NewCatInfo(fm), nil
}

func getInputFiles(c *cli.Context) ([]*os.File, error) {
	args := c.Args()

	if len(args) == 0 {
		return []*os.File{os.Stdin}, nil
	}

	var fs []*os.File
	for _, a := range args {
		if f, err := os.Open(a); err == nil {
			fs = append(fs, f)
		} else {
			return nil, err
		}
	}

	return fs, nil
}

func readInput(file *os.File) ([]byte, error) {
	stat, _ := file.Stat()
	client.L.DebugF("file stat Name: %v", stat.Name())
	client.L.DebugF("file stat Size: %v", stat.Size())
	client.L.DebugF("file stat Mode: %v", stat.Mode())
	client.L.DebugF("file stat ModTime: %v", stat.ModTime())
	client.L.DebugF("file stat IsDir: %v", stat.IsDir())
	client.L.DebugF("file stat Sys: %v", stat.Sys())
	if stat.Size() == 0 {
		return nil, errInputIsEmpty(stat.Name())
	}
	if stat.IsDir() {
		return nil, errInputIsDir(stat.Name())
	}

	return ioutil.ReadAll(file)
}

func list(c *cli.Context) {
	fmt.Println("Supported gat commands are:")
	for _, command := range c.App.Commands {
		name := command.Name
		switch name {
		case "list":
		case "help":
			break
		default:
			usage := command.Usage
			sName := command.ShortName
			if len(sName) > 0 {
				sName = "(" + sName + ")"
			}
			fmt.Println(" ", name, sName, "-", usage)
		}
	}
}
