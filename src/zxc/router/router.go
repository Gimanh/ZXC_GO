package router

import (
    "net/http"
    "fmt"
    "regexp"
    "zxc/config"
    "log"
    "strings"
)

type Router struct {
    Routes       map[string]*Route
    sectionCount map[int]string
}

func New(config config.RoutesConfig) *Router {
    router := &Router{
        Routes:       make(map[string]*Route),
        sectionCount: make(map[int]string),
    }
    router.RegisterRoutesFromConfig(config)
    return router
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    params := r.GetRouteParamsByPath(req.URL.Path)
    fmt.Println(params.params["id"])
}

func (r *Router) RegisterRoute(route *Route) {
    r.Routes[route.Regex] = route
}

func (r *Router) RegisterRoutesFromConfig(config config.RoutesConfig) {
    for value := range config {
        registeredRoutePath, sectionsCount := r.wasRouteRegistered(config[value].Path)
        if registeredRoutePath == "" {
            if r.IsValidPath(config[value].Path) {
                route := &Route{
                    Method:  config[value].Type,
                    Path:    config[value].Path,
                    Handler: config[value].Handler,
                    Regex:   r.GetRegexpFromPath(config[value].Path),
                }
                r.RegisterRoute(route)
                r.sectionCount[sectionsCount] = config[value].Path
            } else {
                log.Println("Invalid route path", config[value].Path)
            }
        } else {
            log.Println("Route with the same sections count was registered \n" +
                "Was registered route with the next path \"path\":\"" + registeredRoutePath + "\"")
        }
    }
}

func (r *Router) GetRouteParamsByPath(path string) *Params {
    for index, value := range r.Routes {
        reg := regexp.MustCompile(value.Regex)
        match := reg.MatchString(path)
        if match {
            regSub := regexp.MustCompile(index)
            params := regSub.FindStringSubmatch(path)
            names := regSub.SubexpNames()
            par := &Params{route: value, params: make(map[string]string)}
            if len(params) > 1 {
                for i := 1; i < len(params); i++ {
                    par.params[names[i]] = params[i]
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

func (r *Router) wasRouteRegistered(path string) (string, int) {
    sections := strings.Split(path, "/")
    sectionCount := len(sections)
    return r.sectionCount[sectionCount], sectionCount
}
