package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r* http.Request){
	method := r.Method;
	urlPath := r.URL.Path

	fmt.Printf("Method : %s\t Path : %s",method,urlPath)
	next.ServeHTTP(w,r)
	})
}