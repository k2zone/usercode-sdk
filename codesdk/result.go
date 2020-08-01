package codesdk

import (
	"bytes"
	"strings"
)

const _SPLIT_FLAG = "~~>|kTWO|<~~"

type Result interface {
	Out() string
	Err() string
}

type result struct {
	out string
	err string
	spent int64
}

func newResult(out string, spent int64) (r *result) {
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

	r = &result{
		out: outBuffer.String(),
		err: errs,
		spent: spent,
	}
	return
}

func (r *result) Output() string {
	return r.out
}

func (r *result) Err() string {
	return r.err
}
