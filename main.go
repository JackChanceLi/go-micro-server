package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func handler () *httprouter.Router {
	router := httprouter.New()
	router.POST("/",Handle)
	router.POST("/user/register", Register)
	router.POST("/user/login",LoginByMail)
	router.POST("/user/create_room",CrteateLiveRoom)

	return router
}

func main() {
	r := handler()
	mh := NewMiddleWareHandler(r)
	log.Printf("Server start1\n")
	http.ListenAndServe(":9090",mh)

}
//func main() {
//	fmt.Println("nihaoa")
//	http.HandleFunc("/", handler) // each request calls handler
//	err := http.ListenAndServe(":9090", nil) //设置监听的端口
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//	//log.Fatal(http.ListenAndServe(":8000", nil))
//}
//
//// handler echoes the Path component of the requested URL.
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}
//
////!-
