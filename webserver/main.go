package main

import (
	"log"
	"mini-micro/transapi"
	"mini-micro/webserver/router"

	"github.com/valyala/fasthttp"
)

func main() {
	transapi.Init()
	log.Fatal(fasthttp.ListenAndServe(":8080", router.R.Handler))
}
