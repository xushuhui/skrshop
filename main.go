/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://xushuhui.github.io
 */
package main

import (
	"goal/internal/router"
	"goal/pkg/core"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	core.StartModule()

	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()

}
