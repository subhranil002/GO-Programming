package router

import (
	"net/http"

	"github.com/subhranil002/GO-CRUD/internal/employee"
	"github.com/subhranil002/GO-CRUD/pkg/response"
)

func Setup(employeeHandler *course.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, "employee API is running", nil)
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, "ok", nil)
	})

	mux.HandleFunc("GET /employees", employeeHandler.List)
	mux.HandleFunc("POST /employees", employeeHandler.Create)
	mux.HandleFunc("GET /employees/{id}", employeeHandler.Get)
	mux.HandleFunc("PATCH /employees/{id}", employeeHandler.Update)
	mux.HandleFunc("DELETE /employees/{id}", employeeHandler.Delete)

	return mux
}
