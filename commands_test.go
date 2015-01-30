package main_test

import (
	"github.com/codegangsta/cli"

	. "github.com/goldeneggg/gat"
)

const (
	testCommandsDir = "./test/commands_test"
	testTextFile    = testCommandsDir + "/test.txt"
	testTextFile2   = testCommandsDir + "/test2.txt"
	testEmptyFile   = testCommandsDir + "/test_empty.txt"
)

func getApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gatTest"
	app.Version = Version
	app.Usage = "Test gat"
	app.Author = "@goldeneggg"
	app.Email = "jpshadowapps@gmail.com"

	app.Flags = GlobalFlags
	app.Commands = Commands

	return app
}

// global flags
func ExampleHelp() {
	getApp().Run([]string{"", "-h"})
}

func ExampleVersion() {
	getApp().Run([]string{"", "-v"})
	// Output:
	// gatTest version 0.3.0
}

func ExampleVersionRunningCommand() {
	getApp().Run([]string{"", "-v", "gist", testTextFile})
	// Output:
	// gatTest version 0.3.0
}

// "gist" Command
func ExampleRunGistEmptyConf() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_empty.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistEmptyDomain() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistEmptyToken() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_token.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNullDomain() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_null_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNullToken() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_null_token.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistInvalidDomain() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_i_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNotFound() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_notfound.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistHelp() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_domain.json", "gist", "-h"})
	// Output:
	// NAME:
	//    gist - Cat to gist
	//
	// USAGE:
	//    command gist [command options] [arguments...]
	//
	// OPTIONS:
	//    --api-domain 	Github api domain
	//    --access-token 	Github api access token
	//    --timeout "0"	Timeout for connection
	//    --description, -d 	A description of the gist
	//    --public, -p		Indicates whether the gist is public. Default: false
}

// "slack" command
func ExampleRunSlackEmptyUrl() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "slack", testTextFile})
	// Output:
}

func ExampleRunSlackNullUrl() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_null_domain.json", "slack", testTextFile})
	// Output:
}

func ExampleRunSlackHelp() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "slack", "-h"})
	// Output:
	// NAME:
	//    slack - Cat to slack
	//
	// USAGE:
	//    command slack [command options] [arguments...]
	//
	// OPTIONS:
	//    --webhook-url 	Webhook URL
	//    --channel, -c 	Target channel
	//    --username, -u 	Username
	//    --icon, -i 		Icon url or emoji format text (:EMOJI_NAME:)
	//    --timeout "0"	Timeout for connection
	//    --without-markdown	Not format slack's markdown
	//    --without-unfurl	Not unfurl media links
	//    --linkfy, -l		Linkify channel names (starting with a '#') and usernames (starting with an '@')
}

func ExampleRunPlaygoHelp() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "playgo", "-h"})
	// Output:
	// NAME:
	//    playgo - Cat to play.golang.org
	//
	// USAGE:
	//    command playgo [arguments...]
}

// "list" Command
func ExampleRunListCommand() {
	getApp().Run([]string{"", "list"})
	// Output:
	// Supported gat commands are:
	//   gist  - Cat to gist
	//   slack  - Cat to slack
	//   playgo  - Cat to play.golang.org
}

func ExampleRunListCommandWithInput() {
	getApp().Run([]string{"", "list", testTextFile})
	// Output:
	// Supported gat commands are:
	//   gist  - Cat to gist
	//   slack  - Cat to slack
	//   playgo  - Cat to play.golang.org
}

/*
// "os" Command
func ExampleRunOs() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testTextFile})
	// Output:
	// test1
	// test2
	// test3
}
*/

/* XXX
func ExampleRunOsMultiInput() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testTextFile, testTextFile2})
	// Output:
	// test1
	// test2
	// test3
	// TEST1
	// TEST2
	// TEST3
}
*/

/*
func ExampleRunOsN() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os_n.json", "os", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsB() {
	getApp().Run([]string{"", "--confpath", testCommandsDir + "/test_conf_os_b.json", "os", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

// use conf.json that does not have keys.
func ExampleRunOsNoKey() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os_nokey.json", "os", testTextFile})
	// Output:
	// test1
	// test2
	// test3
}

func ExampleRunOsNoKeyOptN() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os_nokey.json", "os", "-n", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsConfNFalseOptN() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os_n_false.json", "os", "-n", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsEmptyTarget() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testEmptyFile})
	// Output:
	//
}

func ExampleRunOsDirTarget() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testCommandsDir})
	// Output:
	//
}

func ExampleRunOsNoTarget() {
	getApp().Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os"})
	// Output:
	//
}
*/

// abnormal cases
func ExampleInvalidCommand() {
	getApp().Run([]string{"", "invalid", testTextFile})
	// Output:
	// No help topic for 'invalid'
}

func ExampleEmptyCommand() {
	getApp().Run([]string{"", testTextFile})
	// Output:
	// No help topic for './test/commands_test/test.txt'
}

func ExampleNotExistFile() {
	getApp().Run([]string{"", "gist", testCommandsDir + "/notexist.txt"})
	// Output:
	//
}
