package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
)

var exitSts int

func main() {
	handleSigint()
}

func handleSigint() {
	defer finalize()

	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)

	ch := make(chan bool)
	go run(ch)

	select {
	case <-chSig:
		fmt.Fprintln(os.Stderr, "CTRL-C; exiting")
		exitSts = 1
	case <-ch:
	}
}

func run(ch chan bool) {
	app := cli.NewApp()
	app.Name = "gat"
	app.Version = Version
	app.Usage = "Utility tool of concatnating and printing file to various services"
	app.Author = "@goldeneggg"
	app.Email = "jpshadowapps@gmail.com"
	app.Flags = GlobalFlags
	app.Commands = Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitSts = 1
	}

	ch <- true
}

func finalize() {
	os.Exit(exitSts)
}
