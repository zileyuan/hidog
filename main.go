package main

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
)

func newInstance() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("static"))
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		IndentJSON: macaron.Env != macaron.PROD,
		IndentXML:  macaron.Env != macaron.PROD,
	}))

	//DoXXX 表示GET请求；
	//OnXXX 表示POST请求；
	//AnyXXX 表示GET、POST混合请求
	m.Any("/", AnyValidate)
	m.Get("/dog", DoDog)
	m.Get("/pup", DoPup)
	m.Get("/about", DoAbout)
	m.Get("/comments", DoComments)
	m.Get("/signin", DoSignin)
	m.Get("/detail", DoDetail)

	return m
}

func main() {
	CreateMenu()
	m := newInstance()
	listenAddr := fmt.Sprintf("0.0.0.0:%d", 8070)
	http.ListenAndServe(listenAddr, m)
}
