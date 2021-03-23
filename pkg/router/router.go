package router

import (
	"handsongolang/pkg/hello"
	"net/http"
	"github.com/gorilla/mux"
    "handsongolang/pkg/user"
    "handsongolang/pkg/handlers"
)

func Routers(userService user.Service) http.Handler {
	r := mux.NewRouter()
    uh := handlers.NewUserHandler(userService)

	r.HandleFunc("/hello/ok", hello.Hello)
	r.HandleFunc("/hello/error", hello.HelloError)
	r.HandleFunc("/headers", hello.Headers)
    r.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)


	return r
	//         http.Handle("/", r)
}
