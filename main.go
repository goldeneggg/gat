package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
)

const (
	VERSION = "0.1.0"
)

func main() {
	handleSigint()
}

func handleSigint() {
	// handler for return
	var sts int
	defer finalize(sts)

	// channel as SIGINT handler
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT)

	// run app
	chSts := make(chan int)
	go run(chSts)

	select {
	case <-chSig:
		fmt.Fprintln(os.Stderr, "CTRL-C; exiting")
		sts = 1
	case sts = <-chSts:
	}
}

func run(chSts chan int) {
	var sts int

	app := cli.NewApp()
	app.Name = "gat"
	app.Version = VERSION
	app.Usage = "Cat to anywhere"
	app.Author = "@goldeneggg"
	app.Email = "jpshadowapps@gmail.com"
	app.Flags = globalFlags
	app.Commands = commands

	//app.RunAndExitOnError()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		sts = 1
	}

	chSts <- sts
}

func finalize(sts int) {
	os.Exit(sts)
}
