package rpc

import (
	"fmt"
	"net/rpc"
)

var client *rpc.Client

func init() {
	var err error
	client, err = rpc.DialHTTP("tcp", "localhost:19999")
	if nil != err {
		panic(err)
	}
}

func Trans(word string) string {
	var (
		in  = &Request{Word: word}
		out = &Response{}
	)
	err := client.Call("SpiderService.Crawl", in, out)
	if nil != err {
		fmt.Println(err)
	}

	return out.Means
}
