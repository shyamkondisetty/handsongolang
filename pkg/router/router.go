package router

import (
    "net/http"
    "handsongolang/pkg/hello"
    "github.com/gorilla/mux"
)

func Routers() http.Handler{
    r := mux.NewRouter()
        r.HandleFunc("/hello", hello.Hello)
        r.HandleFunc("/headers", hello.Headers)

    return r;
//         http.Handle("/", r)
}