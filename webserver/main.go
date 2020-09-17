package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webserver/router"

	"github.com/valyala/fasthttp"
)

func main() {
	sv := &fasthttp.Server{
		Handler:     router.R.Handler,
		ReadTimeout: time.Second * 60,
	}
	go sv.ListenAndServe(":8081")
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
			ctx, cn := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
			go func() {
				if err := sv.Shutdown(); nil != err {
					log.Println("shutdown error", err)
				}
				log.Println("shutdown ok")
				cn()
			}()

			<-ctx.Done()
			return
		}
	}
}
