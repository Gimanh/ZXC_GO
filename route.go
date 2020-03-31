package ZXC_GO

import (
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request, *Route)

type Route struct {
	sections     []string
	handler      Handle
	params       *Params
	url          string
	urlSections  []string
	filledParams bool
}

func (route *Route) Get(key string) string {
	if !route.filledParams {
		route.params.fill(route.sections, route.urlSections)
	}
	return route.params.values[key]
}
