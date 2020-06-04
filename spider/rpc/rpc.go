package rpc

import (
	"net"
	"net/http"
	"net/rpc"
	"spider/spider"
)

func (s *SpiderService) Crawl(in *Request, out *Response) error {
	out.Means = spider.Trans(in.Word)
	return nil
}

func Serve() *http.Server {
	svc := &SpiderService{}
	rpc.Register(svc)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":19999")
	if nil != err {
		panic(err)
	}

	sv := &http.Server{
		Handler: rpc.DefaultServer,
	}
	go sv.Serve(l)
	return sv
}
