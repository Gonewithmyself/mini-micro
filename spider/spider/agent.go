package spider

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	JSONquery = `from=en&to=zh&query=%s&token=%s&sign=%s`
	url       = `https://fanyi.baidu.com/basetrans`
)

func Post(word string) (body string) {
	data := fmt.Sprintf(JSONquery, word, token, getSign(word))
	// setQuery(word)
	resp, body, err := agent.SendString(data).End()
	if nil != err {
		fmt.Println("error", err)
		return
	}
	_ = resp
	return body
}

var agent *gorequest.SuperAgent

func init() {
	agent = gorequest.New().Post(url).Set("User-Agent", Agent).Set("X-Requested-With", "XMLHttpRequest").Set("Connection", "keep-alive").
		Set("Cookie", cookie).Set("Referer", "https://fanyi.baidu.com/")
}
