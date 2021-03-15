package codesdk

import (
	"context"
	"github.com/go-kratos/kratos/pkg/log"
	"usercode-sdk/command"
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

func (c *Cli) Run(ctx context.Context, p *Params) (res *Result, err error) {
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
		cmd    = command.NewCmd(cli.toString())
	)
	st := time.Now().UnixNano() / 1e6
	select {
	case sRet := <-cmd.Run():
		stdout = sRet.Stdout
		stderr = sRet.Stderr
		err = sRet.Err

	case <-time.After(time.Second * time.Duration(c.conf.Timeout+2)): // +1秒的编译时间
		cmd.Kill()
		err = runTimeout

	case <-ctx.Done():
		cmd.Kill()
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
