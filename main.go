package main

import (
	"os"
	"os/signal"
	"syscall"
	"vartul14/locus/server"
)

func main() {
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
}
