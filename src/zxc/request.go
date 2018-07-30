package zxc

import "net/http"

type Request struct {
    Request *http.Request
    Params  *Params
}
