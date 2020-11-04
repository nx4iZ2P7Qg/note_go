package main

import (
	"fmt"
	"net/http"
)

func WithRequestInfo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("1")
		handler.ServeHTTP(w, req)
	})
}

func WithRequestInfo2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("2")
		handler.ServeHTTP(w, req)
	})
}

// 怪异的 handler 链
// 看作入栈操作，输出2 1 3
func demo01() {
	server := &http.Server{
		Addr: ":9999",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Println("3")
		}),
	}
	h := server.Handler
	h = WithRequestInfo(h)
	h = WithRequestInfo2(h)
	server.Handler = h

	server.ListenAndServe()
}
