package codesdk

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"os"
	"time"
)

type Cli struct {
	compMap compMap
	conf    *Config
}

func New(conf *Config) (cli *Cli) {
	cli = &Cli{
		conf: conf,
	}
	cli.init()
	return
}

func (c *Cli) init() {
	makeCompMap(&c.compMap)
}

func (c *Cli) Run(p *Params) (res *Result, err error) {
	ext, ok := c.compMap[p.Compiler]
	if !ok {
		err = notSupport
		return
	}
	cli := NewCli(c.conf).params(p, ext)
	//fmt.Println(cli.toString())
	var (
		stdout string
		stderr string
	)

	st := time.Now().UnixNano() / 1e6
	select {
	case sRet := <-runCli(cli.toString()):
		stdout = sRet.stdout
		stderr = sRet.stderr
		err = sRet.err

	case <-time.After(time.Second * time.Duration(c.conf.Timeout+1)): // +1秒的编译时间
		err = runTimeout
	}
	et := time.Now().UnixNano() / 1e6

	if err != nil || stderr != "" {
		log.Error("code run err(%v) stderr(%s)", err, stderr)
		err = runError
		return
	}
	res = newResult(stdout, et-st)
	return
}

func runCli(cli string) <-chan *shellResult {
	ret := make(chan *shellResult)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic err(%v)", err)
				ret <- &shellResult{
					err: runError,
				}
			}
		}()
		if os.Getenv("DEPLOY_ENV") == "uat" {
			ret <- &shellResult{
				stdout: "this is test~~>|kTWO|<~~warning",
			}
			fmt.Println(cli)
			return
		}
		stdout, stderr, err := shell(cli)
		ret <- &shellResult{
			stdout: stdout,
			stderr: stderr,
			err:    err,
		}
	}()

	return ret
}
