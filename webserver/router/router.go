package router

import (
	"github.com/buaazp/fasthttprouter"
)

var R *fasthttprouter.Router

func init() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.POST("/", Post)
	router.ServeFiles("/static/*filepath", "static/")

	R = router
}
