package client

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	NAME_OSCAT = "os"
	STDIN      = "/dev/stdin"
)

type oscat struct {
	Number         bool `json:"n"`
	NumberNonBlank bool `json:"b"`
}

func newOs() *oscat {
	return &oscat{}
}

func (c *oscat) CheckConf() error {
	return nil
}

func (c *oscat) Cat(catInf *CatInfo) (string, error) {
	if out, err := c.catUsingExec(catInf.Files); err != nil {
		return "", err
	} else {
		return out, nil
	}
}

func (c *oscat) CatP(catInf *CatInfo, chOut chan string, chErr chan error) {
	if out, err := c.Cat(catInf); err != nil {
		chErr <- err
	} else {
		chOut <- out
	}
}

func (c *oscat) catUsingExec(files map[string][]byte) (string, error) {
	opts := make([]string, 0, 3)
	if c.Number {
		opts = append(opts, "-n")
	}
	if c.NumberNonBlank {
		opts = append(opts, "-b")
	}

	isStdin := false
	for f, _ := range files {
		if f == STDIN {
			isStdin = true
		} else {
			opts = append(opts, f)
		}
	}

	cmd := exec.Command("cat", opts...)

	var out bytes.Buffer
	cmd.Stdout = &out
	if isStdin {
		cmd.Stdin = bytes.NewReader(files[STDIN])
	}

	// XXX
	// if input is multiple file, result is not sometimes sequential
	if err := cmd.Run(); err == nil {
		return out.String(), err
	} else {
		return "", fmt.Errorf("run error: %v", err)
	}

	/*
			for f, in := range files {
				L.Debug("Start cat", f, cmd.Process.Pid)
				cmd.Stdout = &out
				cmd.Stdin = bytes.NewReader(in)

				if err := cmd.Run(); err == nil {
					str += out.String()
				} else {
					return "", fmt.Errorf("run error: %v", err)
				}
			}
		return str, nil
	*/
}

//func (c *oscat) catUsingPipe(catInf *CatInfo) (string, error) {
//	// コマンド構築
//	opts := make([]string, 0, 2)
//	if c.Number {
//		opts = append(opts, "-n")
//	}
//	if c.NumberNonBlank {
//		opts = append(opts, "-b")
//	}
//	cmd := exec.Command("cat", opts...)
//
//	// 入力へのパイプを取得
//	pipeIn, errIn := cmd.StdinPipe()
//	if errIn != nil {
//		return "", fmt.Errorf("stdinPipe error: %v", errIn)
//	}
//
//	// 出力へのパイプを取得
//	pipeOut, errOut := cmd.StdoutPipe()
//	if errOut != nil {
//		return "", fmt.Errorf("stdoutPipe error: %v", errOut)
//	}
//
//	// コマンド開始
//	err := cmd.Start()
//	if err != nil {
//		return "", fmt.Errorf("cmd start error: %v", err)
//	}
//
//	// stdinへ書き込み
//	_, errW := pipeIn.Write(catInf.In)
//	if errW != nil {
//		return "", fmt.Errorf("write to stdin pipe error: %v", errW)
//	}
//
//	// stdinがcloseされないとexitしないコマンドもあるので(catがそう)、明示的にcloseを呼ぶ
//	pipeIn.Close()
//
//	// コマンドの出力を確認
//	out, errRead := ioutil.ReadAll(pipeOut)
//	if errRead != nil {
//		return "", fmt.Errorf("read from stdout pipe error: %v", errRead)
//	}
//
//	// Wait で実行完了を待つ
//	err = cmd.Wait()
//	if err != nil {
//		return "", err
//	} else {
//		return string(out), nil
//	}
//}
