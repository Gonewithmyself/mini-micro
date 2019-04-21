package rpc

import (
	"fmt"
	"io/ioutil"
	"net/rpc"
	"strings"
)

const (
	ConfFile = "config/conf.ini"
)

var client *rpc.Client

func init() {
	readConf()
	var err error

	rpcHost := confMap["rpchost"]
	fmt.Println(rpcHost)
	client, err = rpc.DialHTTP("tcp", rpcHost)
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

var confMap = make(map[string]string)

func readConf() {
	file, err := ioutil.ReadFile(ConfFile)
	if nil != err {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		pair := strings.Split(line, "=")

		if len(pair) < 2 {
			continue
		}

		k, v := pair[0], pair[1]
		strings.Trim(k, " ")
		strings.Trim(v, " ")
		strings.Trim(v, "\"")
		confMap[k] = v
	}
}
