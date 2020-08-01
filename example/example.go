package main

import (
	"fmt"
	"github.com/k2zone/usercode-sdk/codesdk"
)

func main() {
	cli := codesdk.New(&codesdk.Config{
		Timeout:    20,
		DockerId:   "a918b7ce1627",
		MemorySwap: "128M",
		Memory:     "64M",
	})
	res, err := cli.Run(&codesdk.Params{
		Compiler: codesdk.PHP5,
		Stdin:    "",
		Script: `<?php
echo test;`,
	})
	fmt.Printf("res:%+v\n", res)
	fmt.Printf("err:%+v", err)
}
