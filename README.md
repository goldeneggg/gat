gat [![Build Status](https://travis-ci.org/goldeneggg/gat.svg?branch=master)](https://travis-ci.org/goldeneggg/gat) [![Build Status](http://drone.io/github.com/goldeneggg/gat/status.png)](https://drone.io/github.com/goldeneggg/gat/latest) [![Go Report Card](https://goreportcard.com/badge/github.com/goldeneggg/gat)](https://goreportcard.com/report/github.com/goldeneggg/gat) [![GoDoc](https://godoc.org/github.com/goldeneggg/gat?status.png)](https://godoc.org/github.com/goldeneggg/gat) [![MIT License](http://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/goldeneggg/gat/blob/master/LICENSE)
==========
__gat__ is utility tool of concatnating and printing file to various services.

Target services
* Gist
* Slack
* Hipchat
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

* If you'd like to read usage, type `gat` on your terminal.
* Put your setting file at __`~/.gat/conf.json`__ (or indicated path by `-c` global option).


### Supported commands

#### `gat gist FILE`
Upload your file to gist

* Edit `~/.gat/conf.json`, and write `gist` key, `api-domain` and `access-token`

    ```json
    {
      "gist" : {
        "api-domain" : "https://api.github.com",
        "access-token" : "YOUR_GITHUB_TOKEN"
      }
    }
    ```

    * (All settings are possible overwriting by commandline option (ex. `--access-token`))

* Result of `gat gist FILE` command is __auto generated gist URL (ex. `https://gist.github.com/goldeneggg/4727d6c712dc6f3528f3`)__
* Do you want to read more information? type `gat help gist` on your terminal.


##### examples

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


#### `gat slack FILE`
Send your file contents to slack as message

* You need to add "incoming webhook" integration on `https://YOURTEAM.slack.com/services`
    * And get your incoming webhook URL on `https://YOURTEAM.slack.com/services/INTEGRATION_ID#service_setup`
        * like `https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX`
* Edit `~/.gat/conf.json`, and write `slack` key and `webhook-url`

    ```json
    {
      "gist" : {

      },
      "slack" : {
        "webhook-url" : "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
      }
    }
    ```

* Result of `gat slack` command is `ok` or error message.
* If you'd like to format message, use `--icon`, `--without-markdown` or other option.
    * You can read more information if type `gat help slack` on your terminal.
    * [Show more information of slack's message formatting](https://api.slack.com/docs/formatting)

##### examples

```bash
### output file contents to your slack room
$ gat slack hoge.txt
ok


### output command result to your slack using pipe.
$ echo 'Foo <!everyone> bar http://test.com' | gat slack  # output format is "Foo <!everyone> bar <http://test.com>"
```


#### `gat hipchat -r ROOMID FILE`
Send your file contents to hipchat as message

* You need to confirm your access token for hipchat API.
* Edit `~/.gat/conf.json`, and write `hipchat` key, `api-root` and `access-token`

    ```json
    {
      "gist" : {

      },
      "hipchat" : {
        "api-root" : "https://api.hipchat.com/v2",
        "access-token" : "YOUR_ACCESS_TOKEN"
      }
    }
    ```

    * (All settings are possible overwriting by commandline option (ex. `--access-token`))

* You need to use `-r ROOMID` option for sending message.
* Do you want to read more information? type `gat help hipchat` on your terminal.

##### examples

```bash
### output file contents to your hipchat room
$ gat hipchat -r YOUR_ROOM_ID hoge.txt
```

#### `gat playgo GOLANG_SOURCE_FILE`
Upload your golang source file to playgo.org

* You don't need to edit `~/.gat/conf.json`
* Result of `gat playgo GOLANG_SOURCE_FILE` command is __auto generated playgo.org URL (ex. `https://play.golang.org/p/BrhRIGnmEY`)__
* Do you want to read more information? type `gat help playgo` on your terminal.

##### examples

```bash
$ gat playgo main.go
https://play.golang.org/p/BrhRIGnmEY
```


## Contact

* Bugs: [issues](https://github.com/goldeneggg/gat/issues)


## ChangeLog
[CHANGELOG.md](CHANGELOG.md) file for details.


## License

[LICENSE](LICENSE) file for details.
