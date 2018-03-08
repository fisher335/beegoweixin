package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	code := this.GetString("code")
	fmt.Println(code)
	if len(code) > 0 {
		UserId := GetUserInfo(code)
		fmt.Println(UserId)
		url := "http://disk.bjsasc.com:8180/NetDisk/rest/wechat?method=wx_login&UserId=" + UserId
		this.Redirect(url, 302)
		this.StopRun()
	} else
	{
		fmt.Println("好像获取用户信息失败，不知道该怎么办")
	}

}
