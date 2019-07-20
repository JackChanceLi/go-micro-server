package main

import (
	"go-micro-server/defs"
	"go-micro-server/session"
	"net/http"
)

//http协议中自定义的header字段
var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	//将自定义的header字段加入鉴权过程
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
