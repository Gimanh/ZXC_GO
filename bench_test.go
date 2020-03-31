package ZXC_GO

import (
	"net/http"
	"testing"
)

func BenchmarkPatternDynamic(b *testing.B) {
	p := NewRouter()
	p.Add("GET", "/user/:name", func(writer http.ResponseWriter, request *http.Request, route *Route) {

	})
	req, err := http.NewRequest("GET", "/user/john", nil)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		p.ServeHTTP(nil, req)
	}
}
func BenchmarkPatternStatic(b *testing.B) {
	p := NewRouter()
	p.Add("GET", "/user/name", func(writer http.ResponseWriter, request *http.Request, route *Route) {

	})
	req, err := http.NewRequest("GET", "/user/name", nil)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		p.ServeHTTP(nil, req)
	}
}
