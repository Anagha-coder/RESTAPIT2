package controller

import (
	"empsystem/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/employees", handlers.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", handlers.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/search", handlers.SearchEmployees).Methods("GET")
	r.HandleFunc("/employees", handlers.CreateEmployeeHandler).Methods("POST")
	r.HandleFunc("/employees/{id}", handlers.UpdateEmployeeHandler).Methods("PUT")
	r.HandleFunc("/employees/{id}", handlers.DeleteEmployeeHandler).Methods("DELETE")
}
