package rest

import (
	"net/http"

	rest "github.com/mir-asif01/students_api/rest/handlers"
	"github.com/mir-asif01/students_api/rest/middleware"
)


func Router(mux *http.ServeMux) {
	mux.Handle("GET /",middleware.Logger(http.HandlerFunc(rest.Health)))
}