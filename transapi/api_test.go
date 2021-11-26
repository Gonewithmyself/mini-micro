package transapi

import (
	"strconv"
	"testing"
	"time"
)

func Test_baduTrans_Trans(t *testing.T) {
	confPath = "../"
	Init()
	ts := baduTrans{
		appID:  "20211125001009725",
		secret: conf.Secret,
		salt:   strconv.Itoa(int(time.Now().Unix())),
		from:   "en",
		to:     "zh",
		uri:    "http://api.fanyi.baidu.com/api/trans/vip/translate",
	}

	t.Log(ts.Trans("go to school"))
}

func Test_baduTrans_sign(t *testing.T) {
	ts := baduTrans{
		appID:  "20211125001009725",
		secret: conf.Secret,
		salt:   time.Now().String(),
	}

	t.Log(ts.sign("1"))
}

func Test_baduTrans_form(t *testing.T) {
	ts := baduTrans{
		appID:  "20211125001009725",
		secret: conf.Secret,
		salt:   strconv.Itoa(int(time.Now().Unix())),
		from:   "en",
		to:     "zh",
		uri:    "http://api.fanyi.baidu.com/api/trans/vip/translate",
	}

	t.Log(ts.form("go"))
}
