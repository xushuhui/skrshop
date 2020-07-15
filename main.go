package main

import (
	"os"
	"os/signal"
	"skrshop-api/core"
	"skrshop-api/route"
	"syscall"
)

func main() {
	core.StartModule()

	route.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	route.HttpServerStop()
}
