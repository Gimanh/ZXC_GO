package zxc_go

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"path"
	"strings"
)

const SlashByte = 47
const Colon = 58

type sectionsCountInt int
type httpMethodName string
type routeStartWithDynamic bool
type hashForFirstSection uint64

type Router struct {
	methods         map[httpMethodName]*Methods
	NotFoundHandler http.HandlerFunc
}

type Methods struct {
	sections map[sectionsCountInt]map[routeStartWithDynamic]map[hashForFirstSection]*Route
}

func (r *Router) Add(method string, pathUrl string, handler Handle) {
	pathUrl = path.Clean(pathUrl)
	sections, sectionsCount := r.split(pathUrl)
	httpMethod := httpMethodName(strings.ToUpper(method))
	params := &Params{values: make(map[string]string)}
	route := &Route{
		sections: sections,
		handler:  handler,
		params:   params,
	}

	if _, exist := r.methods[httpMethod]; !exist {
		r.methods[httpMethod] = &Methods{sections: make(map[sectionsCountInt]map[routeStartWithDynamic]map[hashForFirstSection]*Route)}
	}

	hasDynamic := routeStartWithDynamic(r.hasDynamicParams(pathUrl))
	var sec hashForFirstSection
	if !hasDynamic {
		sec = hashForFirstSection(r.getUintHash(pathUrl))
	} else {
		sec = hashForFirstSection(r.getUintHash(sections[0]))
	}

	if _, exist := r.methods[httpMethod].sections[sectionsCount]; exist {
		if _, existDynamic := r.methods[httpMethod].sections[sectionsCount][hasDynamic]; !existDynamic {
			r.methods[httpMethod].sections[sectionsCount][hasDynamic] = make(map[hashForFirstSection]*Route)
		}
		r.methods[httpMethod].sections[sectionsCount][hasDynamic][sec] = route
	} else {
		r.methods[httpMethod].sections[sectionsCount] = make(map[routeStartWithDynamic]map[hashForFirstSection]*Route)
		r.methods[httpMethod].sections[sectionsCount][hasDynamic] = make(map[hashForFirstSection]*Route)
		r.methods[httpMethod].sections[sectionsCount][hasDynamic][sec] = route
	}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	sections, sectionsCount := r.split(path.Clean(request.URL.Path))
	if method, methodExist := r.methods[httpMethodName(request.Method)]; methodExist {
		if routerSections, secCountExist := method.sections[sectionsCount]; secCountExist {
			if dynSections, dynExist := routerSections[true]; dynExist {
				hash := hashForFirstSection(r.getUintHash(sections[0]))
				route := dynSections[hash]
				if route != nil {
					route.url = path.Clean(request.URL.Path)
					route.urlSections = sections
					route.handler(writer, request, route)
				} else {
					for _, value := range dynSections {
						if value.sections[0][0] == Colon {
							route = value
							break
						}
					}
					if route != nil {
						route.url = path.Clean(request.URL.Path)
						route.urlSections = sections
						route.handler(writer, request, route)
					} else {
						r.NotFound(writer, request)
					}
				}
			} else {
				if staticSections, staticExist := routerSections[false]; staticExist {
					hash := hashForFirstSection(r.getUintHash(path.Clean(request.URL.Path)))
					route := staticSections[hash]
					if route != nil {
						route.handler(writer, request, route)
					} else {
						r.NotFound(writer, request)
					}
				}
			}
		}
	} else {
		r.NotFound(writer, request)
	}
}

func (r *Router) NotFound(writer http.ResponseWriter, request *http.Request) {
	if r.NotFoundHandler == nil {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(404)
		fmt.Fprintln(writer, "404 page not found")
	} else {
		r.NotFoundHandler(writer, request)
	}
}

func NewRouter() *Router {
	router := &Router{methods: make(map[httpMethodName]*Methods)}
	return router
}

func (r *Router) split(s string) ([]string, sectionsCountInt) {
	begin := 1
	length := len(s)
	var sliceCapacity sectionsCountInt
	sliceCapacity = 0

	for j := 0; j < length; j++ {
		if s[j] == SlashByte {
			sliceCapacity++
		}
	}

	sliceStrings := make([]string, sliceCapacity)
	indexInSlice := 0
	for i := 0; i < length; i++ {
		if s[i] == SlashByte && i != 0 {
			sliceStrings[indexInSlice] = s[begin:i]
			begin = i + 1
			indexInSlice++
		}
	}
	sliceStrings[indexInSlice] = s[begin:length]
	return sliceStrings, sliceCapacity
}

func (r *Router) hasDynamicParams(s string) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == Colon {
			return true
		}
	}

	return false
}

func (r *Router) getUintHash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
