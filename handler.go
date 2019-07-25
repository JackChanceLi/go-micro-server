package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-micro-server/dbop"
	"go-micro-server/defs"
	"go-micro-server/session"
	"io/ioutil"
	"log"
	"net/http"
)
//通过用户名、密码登录验证
func LoginByMail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.UserIdentity{}
	w.Header().Set("Access-Control-Allow-Origin","*")  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	//解析包
	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	userInfo, password, err := dbop.UserLogin(ubody.Email)
	fmt.Println(password)
	fmt.Println(userInfo)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	if password == ubody.Password {
		//fmt.Println("Login successfully!")

		id := session.GenerateNewSessionID(userInfo.Cid)
		su := &defs.SignedUp{Success:true, SessionId:id, Cid:userInfo.Cid, Name:userInfo.Name, Email:userInfo.Email, Auth:userInfo.Auth}

		if resp, err := json.Marshal(su); err != nil {
			sendErrorResponse(w, defs.ErrorInternalFaults)
			return
		} else {
			sendNormalResponse(w, string(resp),201)
		}
	} else {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return
	}
}

//用户注册的实现，包含用户名、密码、邮箱、权限等信息
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	fmt.Println(res)
	if err != nil {
		log.Printf("Http body read failed")
	}
	ubody := &defs.UserIdentity{}

	if err := json.Unmarshal(res, ubody); err != nil {
		fmt.Println(ubody)
		sendErrorResponse(w, defs.ErrRequestBodyParseFailed)
		return
	}
	fmt.Println(ubody)
	if err := dbop.UserRegister(ubody.UserName, ubody.Email, ubody.Password); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	su := defs.Register{Success:true, Username:ubody.UserName}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
	//fmt.Fprintf(w, "Request success!\n")
}
//路由测试函数，目前没有太大用处
func Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "111111111111111111111111111111\n")
}


func CrteateLiveRoom(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	liveInfo := &defs.LiveRoomIdentity{}
	liveInfo.Lid = "000001"
	liveInfo.Name = "ljc"
	liveInfo.Kind = "1"
	liveInfo.Size = "100"
	liveInfo.StartTime = "2006-01-02 15:04:05"
	liveInfo.EndTime = "2006-01-02 15:04:05"
	liveInfo.PushUrl = "www.baidu.com"
	liveInfo.PullHlsUrl = "www.google.com"
	liveInfo.PullRtmpUrl = "rtmp://www.ljc.com"
	liveInfo.PullHttpFlvUrl = "www.hlv.com"
	liveInfo.DisplayUrl = "www.display.com"
	liveInfo.Status = "5"
	liveInfo.Permission = "Auth users"

	if resp, err := json.Marshal(liveInfo); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),200)
	}
}