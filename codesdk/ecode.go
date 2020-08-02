package codesdk

import "github.com/go-kratos/kratos/pkg/ecode"

var (
	notSupport = ecode.New(100) // 不支持的语言
	runError   = ecode.New(101) // 运行失败
)
