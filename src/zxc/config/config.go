package config

type RoutesConfig []struct {
    Type    string `json:"type, omitempty"`
    Path    string `json:"path"`
    Handler string
}

type Config struct {
    Server struct {
        Port string `json:"port"`
    } `json:"server"`
    Router RoutesConfig
}
