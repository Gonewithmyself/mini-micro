package main

import (
	"log"
	"mini-micro/webserver/router"

	"github.com/valyala/fasthttp"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":8080", router.R.Handler))
}
