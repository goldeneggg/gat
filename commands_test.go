package main

import "github.com/codegangsta/cli"

const (
	testCommandsDir = "./test/commands_test"
	testTextFile    = testCommandsDir + "/test.txt"
	testTextFile2   = testCommandsDir + "/test2.txt"
	testEmptyFile   = testCommandsDir + "/test_empty.txt"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "gatTest"
	app.Version = "0.1.0"
	app.Usage = "Test gat"
	app.Author = "@goldeneggg"
	app.Email = "jpshadowapps@gmail.com"

	app.Flags = globalFlags
	app.Commands = commands
}

// global flags
func ExampleHelp() {
	app.Run([]string{"", "-h"})
}

func ExampleVersion() {
	app.Run([]string{"", "-v"})
	// Output:
	// gatTest version 0.1.0
}

func ExampleVersionRunningCommand() {
	app.Run([]string{"", "-v", "os", testTextFile})
	// Output:
	// gatTest version 0.1.0
}

// "gist" Command
func ExampleRunGistEmptyDomain() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistEmptyToken() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_token.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNullDomain() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_null_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNullToken() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_null_token.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistInvalidDomain() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_i_domain.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistNotFound() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_notfound.json", "gist", testTextFile})
	// Output:
}

func ExampleRunGistHelp() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_gist_e_domain.json", "gist", "-h"})
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
	//    --timeout '0'	Timeout for connection
	//    --description, -d 	A description of the gist
	//    --public, -p		Indicates whether the gist is public. Default: false
}

// "slack" command
func ExampleRunSlackEmptyUrl() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "slack", testTextFile})
	// Output:
}

func ExampleRunSlackNullUrl() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_null_domain.json", "slack", testTextFile})
	// Output:
}

func ExampleRunSlackHelp() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "slack", "-h"})
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
	//    --timeout '0'	Timeout for connection
	//    --without-markdown	Not format slack's markdown
	//    --without-unfurl	Not unfurl media links
	//    --linkfy, -l		Linkify channel names (starting with a '#') and usernames (starting with an '@')
}

func ExampleRunPlaygoHelp() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_slack_e_domain.json", "playgo", "-h"})
	// Output:
	// NAME:
	//    playgo - Cat to play.golang.org
	//
	// USAGE:
	//    command playgo [command options] [arguments...]
	//
	// OPTIONS:
	//    --with-run	Share with Run
}

// "list" Command
func ExampleRunListCommand() {
	app.Run([]string{"", "list"})
	// Output:
	// Supported gat commands are:
	//   gist  - Cat to gist
	//   slack  - Cat to slack
	//   playgo  - Cat to play.golang.org
	//   os  - Cat using os cat
}

func ExampleRunListCommandWithInput() {
	app.Run([]string{"", "list", testTextFile})
	// Output:
	// Supported gat commands are:
	//   gist  - Cat to gist
	//   slack  - Cat to slack
	//   playgo  - Cat to play.golang.org
	//   os  - Cat using os cat
}

// "os" Command
func ExampleRunOs() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testTextFile})
	// Output:
	// test1
	// test2
	// test3
}

/* XXX
func ExampleRunOsMultiInput() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testTextFile, testTextFile2})
	// Output:
	// test1
	// test2
	// test3
	// TEST1
	// TEST2
	// TEST3
}
*/

func ExampleRunOsN() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os_n.json", "os", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsB() {
	app.Run([]string{"", "--confpath", testCommandsDir + "/test_conf_os_b.json", "os", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

// use conf.json that does not have keys.
func ExampleRunOsNoKey() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os_nokey.json", "os", testTextFile})
	// Output:
	// test1
	// test2
	// test3
}

func ExampleRunOsNoKeyOptN() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os_nokey.json", "os", "-n", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsConfNFalseOptN() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os_n_false.json", "os", "-n", testTextFile})
	// Output:
	//      1	test1
	//      2	test2
	//      3	test3
}

func ExampleRunOsEmptyTarget() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testEmptyFile})
	// Output:
	//
}

func ExampleRunOsDirTarget() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os", testCommandsDir})
	// Output:
	//
}

func ExampleRunOsNoTarget() {
	app.Run([]string{"", "-c", testCommandsDir + "/test_conf_os.json", "os"})
	// Output:
	//
}

// abnormal cases
func ExampleInvalidCommand() {
	app.Run([]string{"", "invalid", testTextFile})
	// Output:
	// No help topic for 'invalid'
}

func ExampleEmptyCommand() {
	app.Run([]string{"", testTextFile})
	// Output:
	// No help topic for './test/commands_test/test.txt'
}

func ExampleNotExistFile() {
	app.Run([]string{"", "os", testCommandsDir + "/notexist.txt"})
	// Output:
	//
}
