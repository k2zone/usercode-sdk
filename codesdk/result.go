package codesdk

import (
	"bytes"
	"strings"
)

const _SPLIT_FLAG = "~~>|kTWO|<~~"

type Result struct {
	Out   string
	Err   string
	Spent int64
}

func newResult(out string, spent int64) (r *Result) {
	var (
		outBuffer bytes.Buffer
		errs      string
	)

	sp := strings.Split(out, _SPLIT_FLAG)
	splen := len(sp)
	output := sp[0 : splen-1]
	errlog := sp[splen-1 : splen]
	for _, o := range output {
		outBuffer.WriteString(o)
	}
	if len(errlog) > 0 {
		errs = errlog[0]
	}

	r = &Result{
		Out:   outBuffer.String(),
		Err:   errs,
		Spent: spent,
	}
	return
}
