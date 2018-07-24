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
    routes := r.Routes
    for index, value := range routes {
        reg := regexp.MustCompile(value.Regex)
        match := reg.MatchString(req.URL.Path)

        fmt.Println("Path", req.URL.Path)
        fmt.Println("Reg", index)
        fmt.Println("Match", match)
        fmt.Println("Match", reg.FindAllString(req.URL.Path, -1))
        fmt.Println("-----------")

    }

    //result := strings.Split(req.URL.Path, "/")
    //fmt.Println("Hello")
    //fmt.Println(result)
}
func (r *Router) RegisterRoute(route *Route) {
    r.Routes[route.Regex] = route
}
func (r *Router) GetRegexpFromPath(path string) string {
    //validPathRegexp := regexp.MustCompile(`\/([a-zA-Z0-9]+)|(\/:)([a-zA-Z]+)|^\/`)
    //fmt.Println(validPathRegexp.MatchString(path))
    //
    //validPath := validPathRegexp.MatchString(path)
    //if !validPath {
    //   return ""
    //}
    lit := `[a-zA-Z0-9\_\-]+`
    allowedParamChars := `:(` + lit + `)`
    //sample := `(?<$1>` + lit + `)`
    sample := `(?P<$1>` + lit + `)`

    //var re = regexp.MustCompile(`\/([a-zA-Z]+)|(\/:)([a-zA-Z]+)`)
    var re = regexp.MustCompile(allowedParamChars)
    s := re.ReplaceAllString(path, sample)
    //a := `@^` + s + `$@D`
    a := /*`^` +*/ s + ``
    //fmt.Println(s)

    return a
}
func (r *Router) isValidPath(path string) bool {
    return true
}
