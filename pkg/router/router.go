package router

import (
	"handsongolang/pkg/hello"
	"net/http"
	"github.com/gorilla/mux"
    "handsongolang/pkg/user"
)

func Routers() http.Handler {
	r := mux.NewRouter()
    uh := user.NewUserHandler()

	r.HandleFunc("/hello/ok", hello.Hello)
	r.HandleFunc("/hello/error", hello.HelloError)
	r.HandleFunc("/headers", hello.Headers)
    r.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)


	return r
	//         http.Handle("/", r)
}
