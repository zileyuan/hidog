package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//装载Json文件
func loadJson(jsonFile string) []byte {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(fmt.Errorf("error to read json file!"))
	}
	return content
}

func getDogs() []Dog {
	bytes := loadJson("data/dog.json")
	dogs := []Dog{}
	err := json.Unmarshal(bytes, &dogs)
	if err != nil {
		panic(fmt.Errorf("error to unmarshal json! %v", err))
	}
	return dogs
}

func findDog(uuid string) *Dog {
	dogs := getDogs()
	for _, dog := range dogs {
		if uuid == dog.Uuid {
			return &dog
		}
	}
	return nil
}

func getPups() []Pup {
	bytes := loadJson("data/pup.json")
	pups := []Pup{}
	err := json.Unmarshal(bytes, &pups)
	if err != nil {
		panic(fmt.Errorf("error to unmarshal json!"))
	}
	return pups
}

func findPup(uuid string) *Pup {
	pups := getPups()
	for _, pup := range pups {
		if uuid == pup.Uuid {
			return &pup
		}
	}
	return nil
}

func getComments() []Comment {
	bytes := loadJson("data/comment.json")
	comments := []Comment{}
	err := json.Unmarshal(bytes, comments)
	if err != nil {
		panic(fmt.Errorf("error to unmarshal json!"))
	}
	return comments
}

func getGlobal() *Global {
	bytes := loadJson("data/global.json")
	global := &Global{}
	err := json.Unmarshal(bytes, global)
	if err != nil {
		panic(fmt.Errorf("error to unmarshal json!"))
	}
	return global
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

func DoComments(ctx *macaron.Context) {
	ctx.Data["Comments"] = getComments()
	ctx.HTML(200, "comments")
}

func DoSignin(ctx *macaron.Context) {
	ctx.HTML(200, "signin")
}

func DoDogDetail(ctx *macaron.Context) {
	uuid := ctx.Query("Uuid")
	dog := findDog(uuid)
	if dog != nil {
		ctx.Data["IsDog"] = true
		ctx.Data["Dog"] = dog
		ctx.HTML(200, "showDetail")
	}
}

func DoPupDetail(ctx *macaron.Context) {
	uuid := ctx.Query("Uuid")
	pup := findPup(uuid)
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
	mn.Buttons[0].SetAsViewButton("种犬展示", "http://test.lichengsoft.com/dogs")
	mn.Buttons[1].SetAsViewButton("待售幼犬", "http://test.lichengsoft.com/pups")

	var subButtons = make([]menu.Button, 2)
	subButtons[0].SetAsViewButton("我要留言", "http://test.lichengsoft.com/comments")
	subButtons[1].SetAsViewButton("关于灵睿", "http://test.lichengsoft.com/about")

	mn.Buttons[2].SetAsSubMenuButton("更多信息", subButtons)

	if err := clt.CreateMenu(mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("menu reset success !")
}
