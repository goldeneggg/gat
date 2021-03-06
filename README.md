gat
==========

[![Workflow Status](https://github.com/goldeneggg/gat/workflows/CI/badge.svg)](https://github.com/goldeneggg/gat/actions)
[![MIT License](http://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/goldeneggg/gat/blob/master/LICENSE)

__gat__ is utility tool of concatnating and printing file to various services.

Target services
* Gist
* Slack
* Hipchat
* play.golang.org


## Getting Started

### Using Homebrew for OS X

```
$ brew install goldeneggg/tap/gat
```

### Or direct download

Download from [latest release](https://github.com/goldeneggg/gat/releases/latest)

## Usage

* If you'd like to read usage, type `gat` on your terminal.
* Put your setting file at __`~/.gat/conf.json`__ (or indicated path by `-c` global option).

### gist
`gat gist FILE` command is upload your file to gist

Edit `~/.gat/conf.json`, and write `gist` key, `api-domain` and `access-token`

```json
{
  "gist" : {
    "api-domain" : "https://api.github.com",
    "access-token" : "YOUR_GITHUB_TOKEN"
  }
}
```

Result of `gat gist FILE` command is __auto generated gist URL (ex. `https://gist.github.com/goldeneggg/4727d6c712dc6f3528f3`)__

```bash
# output file contents to your gist, (default private mode)
$ gat gist hoge.go
https://gist.github.com/164b687d8d7f7cd9083f

# "-p" option switch mode to public
$ gat gist -p hoge.go
https://gist.github.com/164b687d8d7f7cd9083f

# "-d <description>" option add description
$ gat gist -d "description" hoge.go
https://gist.github.com/164b687d8d7f7cd9083f

# output command result to your gist using pipe. filename of this case is "stdin"
$ sh huga.sh | gat gist
https://gist.github.com/164b687d8d7f7cd9083f
```

If you'd like to post to Github Enterprise on your internal network, you should run with another config json for GHE specidied by `-c` global option.
Edit config json for GHE (ex. `~/.gat/conf_ghe.json`)

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

### slack
`gat slack FILE` command is post your file to slack channel

Edit `~/.gat/conf.json`, and write `slack` key,`api-token` and `channel` values
You need to get or generate a test token for your team on https://api.slack.com/docs/oauth-test-tokens, and set `api-token` key

```json
{
  "gist" : {

  },
  "slack" : {
    "api-token" : "xoxp-xxxxxxxxxx-xxxxxxxxx-xxxxxxxxx",
    "channel" : "#general",
  }
}
```

Result of `gat slack FILE` command is `ok` or error message.

```bash
# output file contents to your slack room
$ gat slack hoge.txt
ok

# output command result to your slack using pipe.
$ echo 'Foo <!everyone> bar http://test.com' | gat slack  # output format is "Foo <!everyone> bar <http://test.com>"
```

___Note: I might correspond to oauth2 authentication flow someday...___

### hipchat
`gat hipchat -r ROOMID FILE` command is upload your file contents to hipchat

Edit `~/.gat/conf.json`, and write `hipchat` key, `api-root` and `access-token`

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

You need to use `-r ROOMID` option for sending message.

```bash
# output file contents to your hipchat room
$ gat hipchat -r YOUR_ROOM_ID hoge.txt
```

### playgo
`gat playgo GOLANG_SOURCE_FILE` command is upload your golang source file to play.golang.org

You don't need to edit `~/.gat/conf.json`

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
