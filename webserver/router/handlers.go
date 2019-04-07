package router

import (
	"encoding/json"
	"io/ioutil"
	"mini-micro/spider"

	"github.com/valyala/fasthttp"
)

type response struct {
	Code int32  `json:"status"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func Index(ctx *fasthttp.RequestCtx) {
	data, _ := ioutil.ReadFile("views/index.html")
	ctx.Response.Header.Set("Content-Type", "text/html")
	ctx.Write(data)
}

func Post(ctx *fasthttp.RequestCtx) {
	word := ctx.PostArgs().Peek("ctx")

	res := &response{
		Data: spider.Trans(string(word)),
	}

	data, _ := json.Marshal(res)
	ctx.Write(data)
}
