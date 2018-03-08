package controllers

import (
	"strings"
	"fmt"
	"net/http"

	"io/ioutil"
	"encoding/json"
	"time"
)

var access_token string = ""
var token_time time.Time = time.Date(
	2017, 06, 21, 20, 34, 58, 0, time.UTC)
var corpid string = "wxc13f6ccc5b51910a"
var corpsecret string = "kNSLkjUfkRSXNsCGvYbCpuE9SKNuU2WuLKDRn0OWk8o"

type Token struct {
	//Errcode      int `json:"errcode"`
	Access_token string `json:"access_token"`
	Errmsg       string `json:"errmsg"`
	//Expires_in int `json"expires_in"`
}

type UserInfo struct {
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`
}

func GetAccessToken() (string) {

	now := time.Now()
	diff := now.Sub(token_time)
	//fmt.Println(diff.Seconds())
	if access_token == "" || diff.Seconds() > 7000 {

		url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=CORPID&corpsecret=SECRET"
		url = strings.Replace(url, "CORPID", corpid, -1)
		url = strings.Replace(url, "SECRET", corpsecret, -1)
		//fmt.Println(url)
		resp, err := http.Get(url)
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("read body err, %v\n", err)

		}
		defer resp.Body.Close()
		var a Token
		if err = json.Unmarshal(body, &a); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
		}
		//fmt.Printf("%+v", a)

		access_token = a.Access_token
		token_time = time.Now()

		fmt.Println("token 超时，重新获取token")
	}
	//fmt.Println(access_token)
	//fmt.Println(token_time)

	return access_token
}

func GetUserInfo(code string) (string) {

	url := "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=ACCESS_TOKEN&code=CODE"
	url = strings.Replace(url, "ACCESS_TOKEN", GetAccessToken(), -1)
	url = strings.Replace(url, "CODE", code, -1)
	//fmt.Println(url)
	resp, err := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("read body err, %v\n", err)

	}
	var a UserInfo
	if err = json.Unmarshal(body, &a); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
	}

	return a.UserId
}
