package codesdk

import (
	"fmt"
	"strings"
)

type cli struct {
	base     string
	memory   string
	dockerId string
	timeout  string
	stdin    string
	code     string
	cmd      string
	result   string
}

const (
	_BASE = "docker run -i --rm=true"
	//_DIR_NAME         = "/usercode/"
	_SCRIPT_NAME      = "/usr/local/script/coderun"
	//_INPUT_FILE_NAME  = "/usercode/inputFile"
	//_ERROR_FILE_NAME  = "/usercode/errors"
	//_LOG_FILE_NAME    = "/usercode/logfile.txt"
	_ORIGIN_FILE_NAME = "/usercode/file"
)

type template string

const (
	_MEMORY  template = "-m %s --memory-swap=%s"
	_TIMEOUT template = "timeout %d"
	_BASH    template = "/bin/bash -c \"%s&&%s %s\"" // 写入参数&&写入代码&&CMD&&读取结果
	_CODE    template = "echo %s|base64 -d> %s"
	_CMD     template = "%s %s '%s' '%s' '%s'" // 脚本地址 stdin 编译器 文件地址 run
)

func (t template) p(i ...interface{}) string {
	return fmt.Sprintf(string(t), i...)
}

type Config struct {
	DockerId   string
	Timeout    int
	Memory     string
	MemorySwap string
}

type Params struct {
	Compiler Compiler
	Script   string
	Stdin    string
}

func NewCli(conf *Config) (c *cli) {
	c = &cli{
		base:     _BASE,
		memory:   _MEMORY.p(conf.Memory, conf.MemorySwap),
		dockerId: conf.DockerId,
		timeout:  _TIMEOUT.p(conf.Timeout),
	}
	return
}

func (c *cli) params(p *Params, ext *compilerExt) *cli {
	var file = _ORIGIN_FILE_NAME
	if ext.file != "" {
		file = ext.file
	}
	p.Stdin = strings.TrimSpace(p.Stdin)
	if len(p.Stdin) == 0 {
		p.Stdin = "-"
	}
	c.cmd = _CMD.p(_SCRIPT_NAME, b64Encode(p.Stdin), ext.cmd, file, ext.run)
	c.code = _CODE.p(b64Encode(p.Script), file)
	return c
}

func (c *cli) toString() string {
	bash := _BASH.p(c.code, c.timeout, c.cmd)
	return fmt.Sprintf("%s %s %s %s", c.base, c.memory, c.dockerId, bash)
}
