package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Unknwon/macaron"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
	"github.com/macaron-contrib/session"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	//测试平台
	//	ORIID     = "gh_f8ad776f4ba1"                             //微信公众平台的ID
	//	APPID     = "wxe444ec9abad6329e"                          //微信公众平台的AppID
	//	APPSECRET = "eb063d0872cd5bd78ae65c686592093a"            //微信公众平台的AppSecret
	//	TOKEN     = "0badfdd13de84ed6be82db2fdef3331b"            //微信公众平台的Token
	//	AESKEY    = "8whYoPHztw5Ju9mvhJtfX1owkYOWjqsc32ScjqQDacM" //微信公众平台的AESKey

	//苑子乐服务号
	ORIID     = "gh_51b409399a82"                             //微信公众平台的ID
	APPID     = "wx4f7d9dc0554abecb"                          //微信公众平台的AppID
	APPSECRET = "1c656d57cc4fd1fdb0bb5ece126f2da1"            //微信公众平台的AppSecret
	TOKEN     = "0badfdd13de84ed6be82db2fdef3331b"            //微信公众平台的Token
	AESKEY    = "8whYoPHztw5Ju9mvhJtfX1owkYOWjqsc32ScjqQDacM" //微信公众平台的AESKey

	//灵睿服务号
	ORIID     = "gh_51b409399a82"                             //微信公众平台的ID
	APPID     = "wx6221edf20f78c0e6"                          //微信公众平台的AppID
	APPSECRET = "303df10f05564b30d4f51c1967b9dca9"            //微信公众平台的AppSecret
	TOKEN     = "0badfdd13de84ed6be82db2fdef3331b"            //微信公众平台的Token
	AESKEY    = "8whYoPHztw5Ju9mvhJtfX1owkYOWjqsc32ScjqQDacM" //微信公众平台的AESKey

	//DB URL
	DBURL = "mongodb://218.244.128.58:27017" //
)

var (
	db *mgo.Database
)

func getDogs() []Dog {
	c := db.C("Dog")
	dogs := []Dog{}
	err := c.Find(nil).All(&dogs)
	if err != nil {
		panic(err)
	}
	return dogs
}

func findDog(id string) *Dog {
	c := db.C("Dog")
	objid := bson.ObjectIdHex(id)
	dog := Dog{}
	err := c.FindId(objid).One(&dog)
	if err != nil {
		panic(err)
	}
	return &dog
}

func getPups() []Pup {
	c := db.C("Pup")
	pups := []Pup{}
	err := c.Find(nil).All(&pups)
	if err != nil {
		panic(err)
	}
	return pups
}

func findPup(id string) *Pup {
	c := db.C("Pup")
	objid := bson.ObjectIdHex(id)
	pup := Pup{}
	err := c.FindId(objid).One(&pup)
	if err != nil {
		panic(err)
	}
	return &pup
}

func getGlobal() *Global {
	c := db.C("Global")
	global := Global{}
	err := c.Find(bson.M{"Key": "Introduce"}).One(&global)
	if err != nil {
		panic(err)
	}
	return &global
}

func getComments() []Comment {
	c := db.C("Comment")
	comments := []Comment{}
	err := c.Find(nil).Sort("-DateTime").All(&comments)
	if err != nil {
		panic(err)
	}
	return comments
}

func DoPups(ctx *macaron.Context) {
	ctx.Data["Title"] = "待售幼犬"
	ctx.Data["Pups"] = getPups()
	ctx.HTML(200, "showList")
}

func DoDogs(ctx *macaron.Context) {
	ctx.Data["Title"] = "种犬展示"
	ctx.Data["Dogs"] = getDogs()
	ctx.HTML(200, "showList")
}

func DoAbout(ctx *macaron.Context) {
	ctx.Data["Global"] = getGlobal()
	ctx.HTML(200, "about")
}

func DoComment(ctx *macaron.Context) {
	ctx.HTML(200, "comment")
}

func OnComment(ctx *macaron.Context) {
	title := ctx.Query("title")
	content := ctx.Query("content")

	c := db.C("Comment")
	comment := Comment{}
	comment.Id = bson.NewObjectId()
	comment.Title = title
	comment.Content = content
	comment.DateTime = time.Now()
	err := c.Insert(comment)
	if err != nil {
		panic(err)
	}
	//	resp := response.NewText("oMl6fs9C4x583NvZJfTcJxqvcomw", "", comment.DateTime, "["+comment.Title+"]"+comment.Content)
	//	mp.WriteRawResponse(ctx.Resp, nil, resp)
	ctx.Data["Title"] = "成功啦！"
	ctx.Data["Info"] = "您的留言已经第一时间发送出去啦！"
	ctx.HTML(200, "info")
}

func DoSignin(ctx *macaron.Context) {
	ctx.HTML(200, "signin")
}

func OnSignin(ctx *macaron.Context, f *session.Flash) {
	userName := ctx.Query("username")
	Password := ctx.Query("password")

	c := db.C("Account")
	account := Account{}
	err := c.Find(bson.M{"UserName": userName, "Password": Password}).One(&account)
	if err != nil {
		ctx.Data["Title"] = "出错啦！"
		ctx.Data["Info"] = "您的账户或密码错误，登录失败啦！"
		ctx.HTML(200, "info")
	}
	ctx.Data["Admin"] = account.Role == 1
	ctx.Data["Comments"] = getComments()
	ctx.HTML(200, "profile")
}

func DoDogDetail(ctx *macaron.Context) {
	id := ctx.Query("Id")
	dog := findDog(id)
	if dog != nil {
		ctx.Data["IsDog"] = true
		ctx.Data["Dog"] = dog
		ctx.HTML(200, "showDetail")
	}
}

func DoPupDetail(ctx *macaron.Context) {
	id := ctx.Query("Id")
	pup := findPup(id)
	if pup != nil {
		ctx.Data["IsPup"] = true
		ctx.Data["Pup"] = pup
		ctx.HTML(200, "showDetail")
	}
}

func AnyValidate(ctx *macaron.Context) {
	aesKey, err := util.AESKeyDecode(AESKEY)
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler
	wechatServer := mp.NewDefaultServer(ORIID, TOKEN, APPID, aesKey, messageServeMux)
	wechatServerFrontend := mp.NewServerFrontend(wechatServer, mp.ErrorHandlerFunc(ErrorHandler), nil)
	wechatServerFrontend.ServeHTTP(ctx.Resp, ctx.Req.Request)
}

// 非法请求的 Handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略...
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	mp.WriteRawResponse(w, r, resp) // 明文模式
	//mp.WriteAESResponse(w, r, resp) // 安全模式
}

func CreateMenu() {
	AccessTokenServer := mp.NewDefaultAccessTokenServer(APPID, APPSECRET, nil) // 一個應用只能有一個實例
	WechatClient := mp.NewClient(AccessTokenServer, nil)
	clt := menu.NewClient(WechatClient.AccessTokenServer, WechatClient.HttpClient)
	clt.DeleteMenu()

	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3)
	mn.Buttons[0].SetAsViewButton("种犬展示", "http://test.lichengsoft.com/dogs")
	mn.Buttons[1].SetAsViewButton("待售幼犬", "http://test.lichengsoft.com/pups")

	var subButtons = make([]menu.Button, 3)
	subButtons[0].SetAsViewButton("我要留言", "http://test.lichengsoft.com/comment")
	subButtons[1].SetAsViewButton("关于灵睿", "http://test.lichengsoft.com/about")
	subButtons[2].SetAsViewButton("登录灵睿", "http://test.lichengsoft.com/signin")

	mn.Buttons[2].SetAsSubMenuButton("更多信息", subButtons)

	if err := clt.CreateMenu(mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("menu reset success !")
}
