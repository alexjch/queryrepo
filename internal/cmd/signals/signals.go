package signals

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForShutdown(server *http.Server) {
	waitSignal := make(chan os.Signal, 1)
	signal.Notify(waitSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-waitSignal
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shuttting down")
	os.Exit(0)
}
