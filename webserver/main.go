package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"webserver/router"

	"github.com/valyala/fasthttp"
)

func main() {
	sv := fasthttp.Server{
		Handler: router.R.Handler,
	}
	go sv.ListenAndServe(":8080")
	log.Println("start server")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	if err := sv.Shutdown(); nil != err {
		log.Println("shutdown error", err)
	}
	log.Println("stop server")
}
