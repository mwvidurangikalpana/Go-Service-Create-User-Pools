package router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/userPool", controller.createUserPool).Methods("POST")
	router.HandleFunc("/userPool/{userPoolID}", controller.readUserPool).Methods("GET")
	router.HandleFunc("/userPool/{userPoolID}", controller.updateUserPool).Methods("PUT")
	router.HandleFunc("/userPool/{userPoolID}", controller.deleteUserPool).Methods("DELETE")

	return router
}
