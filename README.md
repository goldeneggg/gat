gat [![Build Status](https://travis-ci.org/goldeneggg/gat.svg?branch=master)](https://travis-ci.org/goldeneggg/gat) [![Build Status](http://drone.io/github.com/goldeneggg/gat/status.png)](https://drone.io/github.com/goldeneggg/gat/latest) [![GoDoc](https://godoc.org/github.com/goldeneggg/gat?status.png)](https://godoc.org/github.com/goldeneggg/gat) [![MIT License](http://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/goldeneggg/gat/blob/master/LICENSE)
==========
__gat__ is utility tool of concatnating and printing file to various services.

Target services
* Gist
* Slack
* play.golang.org


## Getting Started

### for Mac using homebrew

```
$ brew tap goldeneggg/gat
$ brew install gat
```

### Direct download link
* [linux amd64](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/linux_amd64/gat)
* [linux 386](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/linux_386/gat)
* [darwin amd64](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/darwin_amd64/gat)
* [darwin 386](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/darwin_386/gat)
* [windows amd64](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/windows_amd64/gat.exe)
* [windows 386](https://drone.io/github.com/goldeneggg/gat/files/artifacts/bin/windows_386/gat.exe)


## Usage

```bash
NAME:
   gat - Utility tool of concatnating and printing file to various services

USAGE:
   gat [global options] command [command options] [arguments...]

VERSION:
   0.3.0

AUTHOR:
  goldeneggg - <jpshadowapps@gmail.com>

COMMANDS:
   gist         Cat to gist
   slack        Cat to slack
   playgo       Cat to play.golang.org
   list         Show target service list
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --confpath, -c       Your original config json path
   --debug, -d          Debug detail information
   --help, -h           show help
   --version, -v        print the version
```

* Put your setting file at `~/.gat/conf.json` (or indicated path by `-c` global option)


### Supported commands

#### "gist"

```bash
NAME:
   gist - Upload file to gist

USAGE:
   command gist [command options] [arguments...]

OPTIONS:
   --api-domain 	Github api domain
   --access-token 	Github api access token
   --timeout "0"	Timeout for connection
   --description, -d 	A description of the gist
   --public, -p		Indicates whether the gist is public. Default: false
```

* Edit `~/.gat/conf.json`
    * All settings are possible overwriting by commandline option (ex. `--access-token`)

    ```json
    {
      "gist" : {
        "api-domain" : "https://api.github.com",
        "access-token" : "YOUR_GITHUB_TOKEN",
        "timeout" : 10,
      }
    }
    ```

* Result of `gat gist` command is __auto generated gist URL (ex. `https://gist.github.com/goldeneggg/4727d6c712dc6f3528f3`)

* examples

```bash
### output file contents to your gist, (default private mode)
$ gat gist hoge.go
https://gist.github.com/164b687d8d7f7cd9083f


### "-p" option switch mode to public
$ gat gist -p hoge.go
https://gist.github.com/164b687d8d7f7cd9083f


### "-d <description>" option add description
$ gat gist -d "description" hoge.go
https://gist.github.com/164b687d8d7f7cd9083f


### output command result to your gist using pipe. filename of this case is "stdin"
$ sh huga.sh | gat gist
https://gist.github.com/164b687d8d7f7cd9083f
```

*  If you'd like to post to Github Enterprise on your internal network, you should run with another config json for GHE specidied by `-c` global option.
    * Edit config json for GHE (ex. `~/.gat/conf_ghe.json`)

    ```json
    {
      "gist" : {
        "api-domain" : "https://YOUR_GHE_DOMAIN/api/v3",
        "access-token" : "YOUR_GITHUB_TOKEN_ON_GHE"
      }
    }
    ```

    ```
    $ gat -c ~/.gat/conf_ghe.json gist -p -d "post to GHE" hoge.go
    https://YOUR_GHE_DOMAIN/xxxxxxxxxxxxxxxxxxx
    ```

#### "slack"

```bash
NAME:
   slack - Send file contents to slack

USAGE:
   command slack [command options] [arguments...]

OPTIONS:
   --webhook-url 	Webhook URL
   --channel, -c 	Target channel
   --username, -u 	Username
   --icon, -i 		Icon url or emoji format text (:EMOJI_NAME:)
   --timeout "0"	Timeout for connection
   --without-markdown	Not format slack's markdown
   --without-unfurl	Not unfurl media links
   --linkfy, -l		Linkify channel names (starting with a '#') and usernames (starting with an '@')
```

* [Setup "Incoming Webhooks" of your team](https://my.slack.com/services/new/incoming-webhook)
* Edit `~/.gat/conf.json`
    * All settings are possible overwriting by commandline option (ex. `--webhook-url`)

    ```json
    {
      "slack" : {
        "webhook-url" : "YOUR_WEBHOOK_URL"
      }
    }
    ```

* Result of `gat slack` command is `ok` or error message.
* example

```bash
### output file contents to your gist, (default private mode)
$ gat slack "test slack"
ok


### output command result to your slack using pipe.
$ echo 'Foo <!everyone> bar http://test.com' | gat slack  # output format is "Foo <!everyone> bar <http://test.com>"
```

* [Show more information of slack's message formatting](https://api.slack.com/docs/formatting)

#### "playgo"

```bash

NAME:
   playgo - Upload go code to play.golang.org

USAGE:
   command playgo [arguments...]
```


### Confirm supported service list

* `list` command

```
$ gat list

gist  - Cat to gist
slack  - Cat to slack
playgo  - Cat to play.golang.org
```

### Run debug mode

* You can specify debug flag by `-d` `--debug`


## Contact

* Bugs: [issues](https://github.com/goldeneggg/gat/issues)


## ChangeLog
[CHANGELOG](CHANGELOG) file for details.


## License

[LICENSE](LICENSE) file for details.
