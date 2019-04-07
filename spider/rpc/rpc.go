package rpc

import (
	"mini-micro/spider/spider"
	"net"
	"net/http"
	"net/rpc"
)

func (s *SpiderService) Crawl(in *Request, out *Response) error {
	out.Means = spider.Trans(in.Word)
	return nil
}

func Serve() {
	svc := &SpiderService{}
	rpc.Register(svc)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":19999")
	if nil != err {
		panic(err)
	}

	http.Serve(l, nil)
}
