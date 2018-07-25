package zxc

import (
    "io/ioutil"
    "log"
    "os"
    "encoding/json"
    "zxc/config"
    "zxc/router"
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
    zxc.Router = router.New(zxc.config.Router, zxc.config.Server.Port)
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
