package codesdk

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/ecode"
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

func (c *Cli) Run(p *Params) (res *result, err error) {
	ext, ok := c.compMap[p.Compiler]
	if !ok {
		//err = errors.New("不支持的编译语言")
		err = ecode.New(100)
		return
	}
	cli := NewCli(c.conf).params(p, ext)
	//fmt.Println(cli.toString())
	st := time.Now().UnixNano() / 1e6
	var (
		stdout = "this is test"
		stderr = ""
	)
	if os.Getenv("DEPLOY_ENV") != "uat" {
		stdout, stderr, err = shell(cli.toString())
	} else {
		fmt.Println(cli.toString())
	}
	et := time.Now().UnixNano() / 1e6
	if err != nil || stderr != "" {
		//err = errors.New("运行出错")
		err = ecode.New(101)
		return
	}
	res = newResult(stdout, et-st)
	return
}
