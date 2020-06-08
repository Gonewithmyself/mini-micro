package rpc

import (
	"io/ioutil"
	"log"
	"net/rpc"
	"strings"
)

const (
	ConfFile = "config/conf.ini"
)

var client *rpc.Client

func init() {
	readConf()
}

func dial() error {
	rpcHost := confMap["rpchost"]
	var err error
	client, err = rpc.DialHTTP("tcp", rpcHost)
	if err != nil {
		log.Println("dial", rpcHost, err)
	}
	return err
}

func Trans(word string) string {
	var (
		in  = &Request{Word: word}
		out = &Response{}
	)

	if client == nil {
		if err := dial(); err != nil {
			return "no available backend"
		}
	}

	err := client.Call("SpiderService.Crawl", in, out)
	if nil != err {
		log.Println(err)
		out.Means = "error: " + err.Error()
		client = nil
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
