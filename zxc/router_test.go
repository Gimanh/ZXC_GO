package zxc

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
    "fmt"
)

func TestRouterStatic(t *testing.T) {
    routes := make(map[string]bool)
    routes["/"] = false
    routes["/static/route"] = false
    routes["/static/router"] = false
    routes["/a/b/c/d"] = false
    routes["/a/b/c/f"] = false
    routes["/a/b/c/f/q/w/e/r/t/y/u/t/q/f/"] = false

    router := NewRouter()

    for key := range routes {
        request, _ := http.NewRequest("GET", key, nil)
        response := httptest.NewRecorder()
        router.Add("GET", key, func(writer http.ResponseWriter, request *http.Request, p *Route) {
            routes[key] = true
        })
        router.ServeHTTP(response, request)
    }

    for key := range routes {
        if !routes[key] {
            t.Fatal("Static route exec failed " + key)
        }
    }
}

func TestRouterDynamic(t *testing.T) {
    routes := make(map[string]bool)
    routes["/dynamic/:router"] = false
    routes["/dynamic/:route/:module"] = false

    routesPost := make(map[string]bool)
    routesPost["/dynamic/:router"] = false
    routesPost["/dynamic/:route/:module"] = false

    router := NewRouter()

    router.Add("GET", "/dynamic/:router", func(writer http.ResponseWriter, request *http.Request, p *Route) {
        if p.Get("router") != "Qrouter" {
            routes["/dynamic/:router"] = false
        } else {
            routes["/dynamic/:router"] = true
        }
    })
    router.Add("GET", "/dynamic/:route/:module", func(writer http.ResponseWriter, request *http.Request, p *Route) {
        if p.Get("route") != "Qroute" || p.Get("module") != "Qmodule" {
            routes["/dynamic/:route/:module"] = false
        } else {
            routes["/dynamic/:route/:module"] = true
        }
    })

    router.Add("POST", "/dynamic/:router", func(writer http.ResponseWriter, request *http.Request, p *Route) {
        if p.Get("router") != "Qrouter" {
            routesPost["/dynamic/:router"] = false
        } else {
            routesPost["/dynamic/:router"] = true
        }
    })
    router.Add("POST", "/dynamic/:route/:module", func(writer http.ResponseWriter, request *http.Request, p *Route) {
        if p.Get("route") != "Qroute" || p.Get("module") != "Qmodule" {
            routesPost["/dynamic/:route/:module"] = false
        } else {
            routesPost["/dynamic/:route/:module"] = true
        }
    })

    for key := range routes {
        nKey := strings.Replace(key, ":", "Q", -1)
        request, _ := http.NewRequest("GET", nKey, nil)
        response := httptest.NewRecorder()

        router.ServeHTTP(response, request)
    }

    for key := range routesPost {
        nKey := strings.Replace(key, ":", "Q", -1)
        request, _ := http.NewRequest("POST", nKey, nil)
        response := httptest.NewRecorder()

        router.ServeHTTP(response, request)
    }
    for key := range routes {
        if !routes[key] {
            t.Fatal("Static route exec failed " + key)
        }
    }
    for key := range routesPost {
        if !routesPost[key] {
            t.Fatal("Static route exec failed " + key)
        }
    }
}
func TestRouter_NotFound(t *testing.T) {
    routes := make(map[string]bool)
    routes["/"] = false
    routes["/static/route"] = false
    routes["/static/router"] = false
    routes["/a/b/c/d"] = false
    routes["/a/b/c/f"] = false
    routes["/a/b/c/f/q/w/e/r/t/y/u/t/q/f/"] = false

    router := NewRouter()

    for key := range routes {
        router.Add("GET", key, func(writer http.ResponseWriter, request *http.Request, p *Route) {
            routes[key] = true
        })
    }

    request, _ := http.NewRequest("GET", "/undefined/router", nil)
    response := httptest.NewRecorder()
    router.ServeHTTP(response, request)

    if response.Code != 404 {
        t.Fatal("Not found test error code not 404 ")
    }
}

func TestRouter_CustomNotFound(t *testing.T) {
    routes := make(map[string]bool)
    routes["/"] = false
    routes["/static/route"] = false
    routes["/static/router"] = false
    routes["/a/b/c/d"] = false
    routes["/a/b/c/f"] = false
    routes["/a/b/c/f/q/w/e/r/t/y/u/t/q/f/"] = false
    routes["/dynamic/:router"] = false
    routes["/dynamic/:router/:module"] = false

    router := NewRouter()
    router.NotFoundHandler = func(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
        writer.WriteHeader(500)
        fmt.Fprintln(writer, "500 Internal server error")
    }

    for key := range routes {
        router.Add("GET", key, func(writer http.ResponseWriter, request *http.Request, p *Route) {
            routes[key] = true
        })
    }
    request, _ := http.NewRequest("GET", "/undefined/router", nil)
    response := httptest.NewRecorder()
    router.ServeHTTP(response, request)
    if response.Code != 500 {
        t.Fatal("Not found test error code not 404 ")
    }
}
func TestRouter_StartRouteWithDynamicParameters(t *testing.T) {
    var user, articleId, commentId string
    routes := make(map[string]bool)
    routes["/"] = false
    routes["/static/route"] = false
    routes["/static/router"] = false
    routes["/a/b/c/d"] = false
    routes["/a/b/c/f"] = false
    routes["/a/b/c/f/q/w/e/r/t/y/u/t/q/f/"] = false
    routes["/dynamic/:router"] = false
    routes["/dynamic/:router/:module"] = false
    routes["/:user/article/:articleId/comment/:commentId"] = false

    router := NewRouter()
    router.Add("GET", "/:user/article/:articleId/comment/:commentId", func(writer http.ResponseWriter, request *http.Request, p *Route) {
        user = p.Get("user")
        articleId = p.Get("articleId")
        commentId = p.Get("commentId")
    })

    request, _ := http.NewRequest("GET", "/head/article/121987/comment/572389572", nil)
    response := httptest.NewRecorder()
    router.ServeHTTP(response, request)

    if response.Code != 200 || user != "head" || articleId != "121987" || commentId != "572389572" {
        t.Fatal("Not found test error code not 404 ")
    }
}
