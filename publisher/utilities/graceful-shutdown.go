package utilities

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("Shutdown")
		os.Exit(1)
	}()
}
