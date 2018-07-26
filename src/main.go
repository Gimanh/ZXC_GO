package main

import (
    "zxc"
    "fmt"
    "net/http"
)

func Hello() {
    fmt.Println("Hello")
}
func main() {
    index := []zxc.Routes{
        zxc.Routes {},
        zxc.Routes {},
    }

    z := new(zxc.ZXC)
    z.GO("/src/config/config.json")

    http.ListenAndServe(":8080", z.Router)
}
