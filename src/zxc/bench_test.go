package zxc

import (
    "testing"
    "net/http"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
    return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
    return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
    return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func handler(writer http.ResponseWriter, request *Request) {}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
    w := new(mockResponseWriter)
    u := r.URL
    rq := u.RawQuery
    r.RequestURI = u.RequestURI()

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        u.RawQuery = rq
        router.ServeHTTP(w, r)
    }
}

func BenchmarkAddRoutes(b *testing.B) {
    z := GO()
    router := z.Router
    router.Add("GET", "/user/:id", handler)
    r, _ := http.NewRequest("GET", "/user/gordon", nil)
    benchRequest(b, router, r)
}
