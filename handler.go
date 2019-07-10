package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Login success!\n")
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Request success!\n")
}

func Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "No http router\n")
}