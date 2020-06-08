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
	sv := &fasthttp.Server{
		Handler: router.R.Handler,
	}
	go sv.ListenAndServe(":8080")
	log.Println("start server")

	sigProc(sv)
}

func sigProc(sv *fasthttp.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	defer log.Println("stop server")
	for sig := range ch {
		log.Println("recv signal:", sig)
		switch sig {
		case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
			if err := sv.Shutdown(); nil != err {
				log.Println("shutdown error", err)
			}
			return
		}
	}
}
