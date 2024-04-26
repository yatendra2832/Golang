package router

import (
	"github.com/gorilla/mux"
	controller "github.com/yatendrayadav2832/mongoapi/controllers"
)

func Router() *mux.Router {
	router:=mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies",controller.DeleteAllMovies).Methods("DELETE")
	return router
}