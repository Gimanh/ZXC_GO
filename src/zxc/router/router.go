package router

import (
    "net/http"
    "fmt"
    "regexp"
)

type Router struct {
    Routes map[string]*Route
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Hello")
    fmt.Println(req)
}
func (r *Router) RegisterRoute(route *Route) {
    r.Routes[route.Regex] = route
}
func (r *Router) GetRegexpFromPath(path string) string {
    validPathRegexp := regexp.MustCompile(`[^-:\/_{}()a-zA-Z\d]`)
    fmt.Println(validPathRegexp.MatchString(path))

    validPath := validPathRegexp.MatchString(path)
    if !validPath {
       return ""
    }
    return path
}
