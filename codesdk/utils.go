package codesdk

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"reflect"
)

func rangeStruct(in interface{}, h func(k, v string)) {
	rType, rVal := reflect.TypeOf(in), reflect.ValueOf(in)
	if rType.Kind() == reflect.Ptr { // 传入的in是指针,需要.Elem()取得指针指向的value
		rType, rVal = rType.Elem(), rVal.Elem()
	}
	if rType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rType.NumField(); i++ { // 遍历结构体
		t, f := rType.Field(i), rVal.Field(i)
		// 此处可以参照f.String(),f.Int(),f.Float()源码,处理不同类型,我这里统一转成string类型了
		if f.Kind() != reflect.Struct { // 不深入遍历结构体了,有需要自己实现吧
			h(t.Name, fmt.Sprint(f))
		}
	}
}

func shell(shell string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", shell)
	// 设置接收
	cmd.Stdout = &stdout
	// err
	cmd.Stderr = &stderr
	// 执行
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func b64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
