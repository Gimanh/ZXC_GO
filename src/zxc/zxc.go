package zxc

import (
    "io/ioutil"
    "log"
    "os"
    "encoding/json"
    "net/http"
    "zxc/config"
    "zxc/router"
    "fmt"
)

type ZXC struct {
    config config.Config
    Router *router.Router
}

func (zxc *ZXC) GO(configPath string) {
    dir := zxc.getRootDir()
    fullConfigPath := dir + configPath
    configFile := zxc.readConfigFile(fullConfigPath)

    err := json.Unmarshal(configFile, &zxc.config)
    if err != nil {
        log.Fatal(err)
        return
    }

    if zxc.config.Server.Port == "" {
        log.Fatal("Porn is not defined")
        return
    }
    zxc.initializeRouter()
    zxc.runServer()
    //userVar2, _ := json.Marshal(zxc.config)
    //fmt.Println(string(userVar2))
    return
}

func (zxc *ZXC) getRootDir() string {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    return dir
}

func (zxc *ZXC) readConfigFile(fullConfigPath string) ([]byte) {
    configFile, err := ioutil.ReadFile(fullConfigPath)
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return configFile
}

func (zxc *ZXC) initializeRouter() {
    zxc.Router = new(router.Router)
    zxc.Router.Routes = make(map[string]*router.Route)
    for index, value := range zxc.config.Router {
        route := &router.Route{
            Method:  value.Type,
            Path:    value.Path,
            Handler: value.Handler,
            Regex:   zxc.Router.GetRegexpFromPath(value.Path),
        }
        zxc.Router.RegisterRoute(route)
        fmt.Println(index)
        //fmt.Println(value)
    }
    fmt.Println(zxc.config.Router)
}

func (zxc *ZXC) runServer() {
    //routerI := &router.Router{}
    port := ":" + zxc.config.Server.Port
    s := &http.Server{
        Addr: port,
        //Handler: routerI,
    }
    zxc.registerRoutes()
    log.Fatal(s.ListenAndServe())
}
func (zxc *ZXC) registerRoutes() {

    http.HandleFunc("/:user", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Println("Hello")
    })

    //http.HandleFunc("/:user", uio.Handler)
}
