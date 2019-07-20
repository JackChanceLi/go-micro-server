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

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Login success!\n")
}

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
	if err := dbop.UserRegister(ubody.UserName, ubody.Email, ubody.Passwd, ubody.Role); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionID(ubody.UserName)
	su := &defs.SignedUp{Success:true, SessionId:id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp),201)
	}
	//fmt.Fprintf(w, "Request success!\n")
}

func Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "No http router\n")
}