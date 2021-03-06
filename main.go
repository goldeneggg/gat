package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

import (
	"github.com/urfave/cli"
)

var exitSts int

func main() {
	defer finalize()

	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)

	ch := make(chan struct{})

	// FIXME: parallel post using goroutine
	go run(ch)

	select {
	case <-chSig:
		fmt.Fprintln(os.Stderr, "CTRL-C; exiting")
		exitSts = 1
	case <-ch:
	}
}

func run(ch chan struct{}) {
	defer close(ch)

	app := cli.NewApp()
	app.Name = "gat"
	app.Version = VERSION
	app.Usage = "Utility tool of concatnating and printing file to various services"
	app.Author = "goldeneggg"
	app.Email = "jpshadowapps@gmail.com"
	app.Flags = GlobalFlags
	app.Commands = Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitSts = 1
	}
}

func finalize() {
	os.Exit(exitSts)
}
