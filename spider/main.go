package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"spider/rpc"
	"syscall"
	"time"
)

// import "spider/rpc"

func main() {
	sv := rpc.Serve()
	log.Println("start server")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := sv.Shutdown(ctx)
	if err != nil {
		log.Println("shutdown error", err)
	}
	log.Println("server shutdown")
}
