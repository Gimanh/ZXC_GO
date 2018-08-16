package zxc

import "log"

type Params struct {
    values map[string]string
}

func (p *Params) Get(key string) string {
    return p.values[key]
}

func (p *Params) fill(routeSections []string, urlPathSections []string) {
    if len(routeSections) != len(urlPathSections) {
        log.Fatal("Sections length error ", routeSections, urlPathSections)
    }
    for key, value := range routeSections {
        if value[0] == 58 {
            p.values[value[1:]] = urlPathSections[key]
        }
    }
}
