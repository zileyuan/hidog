package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Unknwon/macaron"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
)

const (
	//	//	//测试平台
	WECHATID  = "gh_f8ad776f4ba1"                             //微信公众平台的ID
	APPID     = "wxe444ec9abad6329e"                          //微信公众平台的AppID
	APPSECRET = "eb063d0872cd5bd78ae65c686592093a"            //微信公众平台的AppSecret
	TOKEN     = "0badfdd13de84ed6be82db2fdef3331b"            //微信公众平台的Token
	AESKEY    = "8whYoPHztw5Ju9mvhJtfX1owkYOWjqsc32ScjqQDacM" //微信公众平台的AESKey
)

func DoPup(ctx *macaron.Context) {
	ctx.HTML(200, "pup")
}

func DoDog(ctx *macaron.Context) {
	ctx.HTML(200, "dog")
}

func DoAbout(ctx *macaron.Context) {
	ctx.HTML(200, "about")
}

func DoComments(ctx *macaron.Context) {
	ctx.HTML(200, "comments")
}

func DoSignin(ctx *macaron.Context) {
	ctx.HTML(200, "signin")
}

func DoDetail(ctx *macaron.Context) {
	ctx.HTML(200, "detail")
}

func AnyValidate(ctx *macaron.Context) {
	aesKey, err := util.AESKeyDecode(AESKEY)
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	wechatServer := mp.NewDefaultServer(WECHATID, TOKEN, APPID, aesKey, messageServeMux)

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
	mn.Buttons[0].SetAsViewButton("种犬展示", "http://test.lichengsoft.com/dog")
	mn.Buttons[1].SetAsViewButton("待售幼犬", "http://test.lichengsoft.com/pup")

	var subButtons = make([]menu.Button, 3)
	subButtons[0].SetAsViewButton("明日之星", "http://test.lichengsoft.com/tomorrow")
	subButtons[1].SetAsViewButton("我要留言", "http://test.lichengsoft.com/comments")
	subButtons[2].SetAsViewButton("关于灵睿", "http://test.lichengsoft.com/about")

	mn.Buttons[2].SetAsSubMenuButton("更多信息", subButtons)

	if err := clt.CreateMenu(mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("menu reset success !")
}
