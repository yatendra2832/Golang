package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")
	greeter()

	r:= mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000",r))

}

func greeter() {
	fmt.Println("Welcome to Journey on mod")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang series on YT </h1>"))


}


// go list -m all --> that will print the list of the modules that are used in the code
// go list -m version module_name --> that will print the all the version which are available for the use of that modules

