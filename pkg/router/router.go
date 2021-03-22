package router

import (
	"handsongolang/pkg/hello"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/ok", hello.Hello)
	r.HandleFunc("/hello/error", hello.HelloError)
	r.HandleFunc("/headers", hello.Headers)

	return r
	//         http.Handle("/", r)
}
