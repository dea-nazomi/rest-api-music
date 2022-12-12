package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/music", controller.GetAllMusic).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/music/{id}", controller.GetMusic).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/music", controller.AddMusic).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/music/{id}", controller.UpdateMusic).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/music/{id}", controller.DeleteMusic).Methods("DELETE", "OPTIONS")

	return router
}
