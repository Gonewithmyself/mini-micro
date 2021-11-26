package transapi

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Secret string `json:"secret"`
}

var (
	conf     *config
	confPath = "../"
)

func Init() {
	d, er := ioutil.ReadFile(confPath + "config.json")
	if er != nil {
		panic(er)
	}

	var c config
	er = json.Unmarshal(d, &c)
	if er != nil {
		panic(er)
	}

	conf = &c
	defaultTranser.setConf(conf)
}
