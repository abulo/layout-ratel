package main

import (
	"fmt"
	"os"

	"github.com/abulo/ratel/v3/watch"
)

var (
	cmd     string
	runArgs string
)

func main() {
	//解析参数
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("路径参数不存在", "参考:./gorun cmd/server.go")
		os.Exit(1)
	}
	cmd = args[1]
	watch.Run(cmd, runArgs)
}
