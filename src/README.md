# ZXC_GO

## ZXC Router Usage

``` package main

import (
    "net/http"
    "zxc"
    "fmt"
)

func main() {
    z := zxc.GO()
    router := z.Router

    router.NotFoundHandler = func(w http.ResponseWriter, r *zxc.Request) {
        w.WriteHeader(404)
        fmt.Fprint(w, "HTTP Not Found: 404")
    }

    router.Add("GET", "/user/:id", func(w http.ResponseWriter, r *zxc.Request) {
        w.Write([]byte(r.Params.Params["id"]))
    })

    router.Add("GET", "/user/articles/:userid/:articleid", func(w http.ResponseWriter, r *zxc.Request) {
        w.Write([]byte("User id " + r.Params.Params["userid"] + " \n"))
        w.Write([]byte("Article id " + r.Params.Params["articleid"]))
    })

    http.ListenAndServe(":8080", router)
}```