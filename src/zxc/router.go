package zxc

import (
    "net/http"
    "regexp"
    "strings"
    "fmt"
)

type Router struct {
    Methods         map[string]Methods
    NotFoundHandler Handler
}

type Handler func(writer http.ResponseWriter, request *Request)

type Methods struct {
    Routes       map[string]*Route
    SectionCount map[int]string
}

func NewRouter() *Router {
    router := &Router{
        Methods:         make(map[string]Methods),
    }
    return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
    routeParams := r.GetRouteParamsByPath(request.URL.Path, request.Method)
    if routeParams == nil {
        r.NotFound(w, request)
        return
    }
    req := &Request{
        Request: request,
        Params:  routeParams,
    }
    routeParams.Route.Handler(w, req)
}

func (r *Router) GetUserHandler(path string, method string) Handler {
    if !r.MethodExist(method) {
        return nil
    }
    routeParams := r.GetRouteParamsByPath(path, method)
    return routeParams.Route.Handler
}

func (r *Router) NotFound(w http.ResponseWriter, req *http.Request) {
    if r.NotFoundHandler == nil {
        w.WriteHeader(404)
        fmt.Fprint(w, "HTTP: 404")
    }
    r.NotFoundHandler(w, nil)
}

func (r *Router) RegisterRoute(route *Route, method string) {
    r.Methods[method].Routes[route.Regex] = route
}

func (r *Router) Add(method string, path string, handler Handler) {
    if !r.MethodExist(method) {
        r.CreateMethod(method)
    }
    method = strings.ToUpper(method)
    route := &Route{
        Method:  strings.ToUpper(method),
        Path:    path,
        Handler: handler,
        Regex:   r.GetRegexpFromPath(path),
    }
    r.RegisterRoute(route, method)
}

func (r *Router) GetRouteParamsByPath(path string, method string) *Params {
    for index, value := range r.Methods[method].Routes {
        reg := regexp.MustCompile(value.Regex)
        match := reg.MatchString(path)
        if match {
            regSub := regexp.MustCompile(index)
            params := regSub.FindStringSubmatch(path)
            names := regSub.SubexpNames()
            par := &Params{Route: value, Params: make(map[string]string)}
            if len(params) > 1 {
                for i := 1; i < len(params); i++ {
                    par.Params[names[i]] = params[i]
                }
            }
            return par
        }
    }
    return nil
}

func (r *Router) GetRegexpFromPath(path string) string {
    allowedParamChars := `[a-zA-Z0-9]+`
    rxCompile := `:(` + allowedParamChars + `)`
    replace := `(?P<$1>` + allowedParamChars + `)`
    var re = regexp.MustCompile(rxCompile)
    s := re.ReplaceAllString(path, replace)
    a := s + `$`
    return a
}

func (r *Router) IsValidPath(path string) bool {
    var re = regexp.MustCompile(`^[\/a-zA-Z0-9:.]*$`)
    match := re.MatchString(path)
    return match
}

func (r *Router) RouteExist(path string, method string) bool {
    sections := strings.Split(path, "/")
    sections = sections[1:]
    sectionCount := len(sections)
    _, exist := r.Methods[method].SectionCount[sectionCount]
    return exist
}

func (r *Router) MethodExist(method string) bool {
    _, exists := r.Methods[method]
    return exists
}

func (r *Router) CreateMethod(method string) bool {
    r.Methods[method] = Methods{
        Routes:       make(map[string]*Route),
        SectionCount: make(map[int]string),
    }
    return true
}
