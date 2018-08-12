package main

import (
    "net/http"
    "zxc"
    "fmt"
)

func main() {
    z := zxc.GO()
    router := z.Router

    router.NotFoundHandler = func(writer http.ResponseWriter, request *http.Request) {
        writer.WriteHeader(404)
        fmt.Fprint(writer, "HTTP Not Found: 404")
    }

    router.Add("GET", "/user/:id", func(writer http.ResponseWriter, request *http.Request, p *zxc.Route) {
        writer.Write([]byte(p.Get("id")))
    })

    router.Add("GET", "/:user/:key", func(writer http.ResponseWriter, request *http.Request, p *zxc.Route) {
        writer.Write([]byte("User dynamic name" + p.Get("user") + " \n"))
        writer.Write([]byte("User dynamic key" + p.Get("key") + " \n" ))
    })

    router.Add("GET", "/:user/article/:articleId/comment/:commentId", func(writer http.ResponseWriter, request *http.Request, p *zxc.Route) {
        writer.Write([]byte("User id " + p.Get("user") + " \n"))
        writer.Write([]byte("Article id " + p.Get("articleId") + " \n"))
        writer.Write([]byte("Comment id " + p.Get("commentId") + " \n"))
    })

    http.ListenAndServe(":8080", router)
}
