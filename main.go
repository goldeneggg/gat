package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
)

func main() {
	handleSigint()
}

func handleSigint() {
	var sts int
	defer finalize(sts)

	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt, syscall.SIGTERM)

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

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		sts = 1
	}

	chSts <- sts
}

func finalize(sts int) {
	os.Exit(sts)
}
