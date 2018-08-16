# ZXC_GO

## ZXC Router v0.0.2 Usages (It's that simple)

```go
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

    router.Add("GET", "/user/:id", func(writer http.ResponseWriter, request *http.Request, route *zxc.Route) {
        writer.Write([]byte(route.Get("id")))
    })

    router.Add("GET", "/:user/:key", func(writer http.ResponseWriter, request *http.Request, route *zxc.Route) {
        writer.Write([]byte("User dynamic name" + route.Get("user") + " \n"))
        writer.Write([]byte("User dynamic key" + route.Get("key") + " \n" ))
    })

    router.Add("GET", "/:user/article/:articleId/comment/:commentId", func(writer http.ResponseWriter, request *http.Request, route *zxc.Route) {
        writer.Write([]byte("User id " + route.Get("user") + " \n"))
        writer.Write([]byte("Article id " + route.Get("articleId") + " \n"))
        writer.Write([]byte("Comment id " + route.Get("commentId") + " \n"))
    })

    http.ListenAndServe(":8080", router)
```

## CONTRIBUTORS
Giman Nikolay

## LICENSE
MIT License

Copyright (c) 2018 by Giman Nikolay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.