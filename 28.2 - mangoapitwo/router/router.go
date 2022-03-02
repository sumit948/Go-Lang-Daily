package router

import (
	"github.com/gorilla/mux"
	"main.go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/employee", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/getemp", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/deleteOneEmp/{id}", controller.DeleteAEmployee).Methods("DELETE")
	return router
}
