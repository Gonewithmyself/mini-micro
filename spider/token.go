package spider

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/parnurzeal/gorequest"
	"github.com/robertkrimen/otto"
)

const (
	Agent  = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36"
	cookie = `BAIDUID=26573F942C2978C7C6FAA992BD1C75C8:FG=1;locale=zh;path=/;domain=.baidu.com`
	SignJs = "../spider/sign.js"
)

func getSign(word string) string {
	jsWord, err := vm.ToValue(word)
	if nil != err {
		panic(err)
	}

	res, err := vm.Call("token", nil, jsWord, gtkV)
	if nil != err {
		panic(err)
	}

	sign, _ := res.ToString()
	return sign
}

func getMeans(src string) string {
	res, err := vm.Call("parse", nil, jsArg(src))
	if nil != err {
		fmt.Println(err)
	}
	means, _ := res.ToString()
	return means
}

func jsArg(v interface{}) otto.Value {
	arg, err := vm.ToValue(v)
	if nil != err {
		panic(err)
	}
	return arg
}

// init vm
var vm *otto.Otto
var gtkV otto.Value

func init() {
	prepareToken()
	vm = otto.New()
	src, err := ioutil.ReadFile(SignJs)
	if nil != err {
		src, err = ioutil.ReadFile("sign.js")
		if nil != err {
			panic(err)
		}

	}

	_, err = vm.Run(src)
	if nil != err {
		panic(err)
	}

	gtkV, err = vm.ToValue(gtk)
	if nil != err {
		panic(err)
	}
}

var token, gtk string

func prepareToken() {
	url := `https://fanyi.baidu.com/`
	agent := gorequest.New().Get(url).Set("User-Agent", Agent).Set("Connection", "keep-alive").
		Set("Cookie", cookie)

	resp, body, err := agent.End()
	if nil != err {
		fmt.Println(err)
	}

	// token: '05301f0daa555e723f477ec3b63f3638',
	tkPatt := regexp.MustCompile(`token: '([\S]+?)'`)
	token = tkPatt.FindStringSubmatch(body)[1]

	// gtk: '320305.131321201'
	gtkPatt := regexp.MustCompile(`gtk: '([\S]+?)'`)
	gtk = gtkPatt.FindStringSubmatch(body)[1]
	// fmt.Println(tk, gtk)
	_, _ = resp, body
}
