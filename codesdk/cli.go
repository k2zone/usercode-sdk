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
	st := time.Now().UnixNano() / 1e6
	var (
		stdout = "this is test~~>|kTWO|<~~warning"
		stderr = ""
	)
	if os.Getenv("DEPLOY_ENV") != "uat" {
		stdout, stderr, err = shell(cli.toString())
	} else {
		fmt.Println(cli.toString())
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
