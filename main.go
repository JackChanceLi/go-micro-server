package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handler () *httprouter.Router {
	router := httprouter.New()
	router.POST("/",Handle)
	router.POST("/user/login", Login)
	router.POST("/user/register", Register)

	return router
}

func main() {
	r := handler()
	http.ListenAndServe(":9090",r)

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
