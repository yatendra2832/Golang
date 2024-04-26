package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yatendrayadav2832/mongoapi/router"
)

func main() {
	fmt.Println("Creating MongoDb API")
	r :=router.Router()
	fmt.Println("Server is Getting started")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000...")
}