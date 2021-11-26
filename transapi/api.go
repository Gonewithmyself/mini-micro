package transapi

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var defaultTranser ITranser

const (
	badiFanyiURI = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

func init() {
	defaultTranser = &baduTrans{
		uri:   badiFanyiURI,
		from:  "en",
		to:    "zh",
		appID: "20211125001009725",
		salt:  strconv.Itoa(int(time.Now().Unix())),
	}
}

func Trans(word string) string {
	res, er := defaultTranser.Trans(word)
	if er != nil {
		return er.Error()
	}
	return res
}

type ITranser interface {
	Trans(word string) (string, error)
	setConf(*config)
}

type baduTrans struct {
	from   string
	to     string
	appID  string
	secret string
	salt   string
	uri    string
}

type result struct {
	Dst string `json:"dst"`
}

type results struct {
	Res []*result `json:"trans_result"`
}

func replaceSpace(word string) string {
	return strings.ReplaceAll(word, " ", "%20")
}

func (b *baduTrans) Trans(word string) (res string, err error) {
	httpGet(b.form(word), func(body io.Reader, er error) {
		var r results
		er = json.NewDecoder(body).Decode(&r)
		if er != nil {
			err = er
			return
		}
		var strs = make([]string, len(r.Res))
		for i := range r.Res {
			strs[i] = r.Res[i].Dst
		}
		res = strings.Join(strs, "|")
	})
	return
}

func httpGet(url string, fn func(body io.Reader, er error)) {
	req := fasthttp.AcquireRequest()
	rsp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(rsp)
	}()
	req.SetRequestURI(url)

	er := fasthttp.Do(req, rsp)
	var bufer bytes.Buffer
	if er == nil {
		rsp.BodyWriteTo(&bufer)
	}
	fn(&bufer, er)
}

func (b *baduTrans) setConf(c *config) {
	b.secret = c.Secret
}

func (b *baduTrans) sign(word string) string {
	toSign := b.appID + word + b.salt + b.secret
	h := md5.New()
	h.Write([]byte(toSign))
	return hex.EncodeToString(h.Sum(nil))
}

func (b *baduTrans) form(word string) string {
	var buf strings.Builder
	buf.WriteString(b.uri)
	buf.WriteString("?")

	sign := b.sign(word)
	word = replaceSpace(word)
	ss := [][2]string{
		{"q=", word},
		{"from=", b.from},
		{"to=", b.to},
		{"appid=", b.appID},
		{"salt=", b.salt},
		{"sign=", sign},
	}

	for i := range ss {
		buf.WriteString(ss[i][0])
		buf.WriteString(ss[i][1])
		buf.WriteString("&")
	}

	res := buf.String()
	fmt.Println(res)
	return res
}
