package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

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

	m.Get("/dogs", DoDogs)
	m.Get("/pups", DoPups)
	m.Get("/about", DoAbout)
	m.Get("/comment", DoComment)
	m.Get("/signin", DoSignin)
	m.Get("/dogDetail", DoDogDetail)
	m.Get("/pupDetail", DoPupDetail)

	m.Post("/onComment", OnComment)
	return m
}

func main() {
	CreateMenu()
	m := newInstance()

	session, err := mgo.Dial(DBURL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db = session.DB("hidog") //数据库名称
	listenAddr := fmt.Sprintf("0.0.0.0:%d", 8070)
	http.ListenAndServe(listenAddr, m)
}
