package command

import (
	"bytes"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/log"
	"os/exec"
)

type Command struct {
	cmd    *exec.Cmd
	result chan (*shellResult)
}

type shellResult struct {
	Stdout string
	Stderr string
	Err    error
}

func NewCmd(shell string) *Command {
	return &Command{
		cmd: exec.Command("sh", "-c", shell),
	}
}

func (c *Command) Run() <-chan *shellResult {
	var ret = make(chan *shellResult)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic err(%v)", err)
				ret <- &shellResult{
					Err: fmt.Errorf("code run err"),
				}
			}
		}()
		if env.DeployEnv == env.DeployEnvUat {
			log.Info("%+v", c.cmd.Args)

			ret <- &shellResult{
				Stdout: "this is test~~>|kTWO|<~~warning",
			}
			return
		}

		var stdout, stderr bytes.Buffer
		// 设置接收
		c.cmd.Stdout = &stdout
		// err
		c.cmd.Stderr = &stderr
		// 执行
		err := c.cmd.Run()
		ret <- &shellResult{
			Stdout: stdout.String(),
			Stderr: stderr.String(),
			Err:    err,
		}
	}()
	return ret
}

func (c *Command) Kill() {
	if c.cmd.Process != nil {
		c.cmd.Process.Kill()
	}
}
