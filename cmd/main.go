package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func main() {

	// 读取配置
	// 初始化项目组件
	// 启动服务
	// 启动debug端口 用于pprof分析

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %s\n\n", r)
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
	select {
	case _ = <-c:
		// do something
	}
}
