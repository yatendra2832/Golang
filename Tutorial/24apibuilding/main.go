package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB : slice having type course
var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {
	fmt.Println("Building APIs ...")
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))

}

// controllers
//
//	server home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To APIs by Yatendra Singh Learning Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}