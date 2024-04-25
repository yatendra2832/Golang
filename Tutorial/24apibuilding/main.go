package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	// Routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/:id", getCourseById).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("courses/:id", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/:id", deleteOneCourse).Methods("DELETE")
	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 399, Author: &Author{Fullname: "Yatendra singh", Website: "codewithyatendra.com"}})

	// listen to port
	log.Fatal(http.ListenAndServe(":3000", r))

}

// controllers
//
//	server home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To APIs by Yatendra Singh Learning Golang</h1>"))
}

// This function handles the request to retrieve all courses. It sets the response content type to JSON, encodes the courses slice into JSON format, and sends the response.
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

/*
Notes
This function handles the request to retrieve a single course by its ID. It extracts the ID from the request parameters, then iterates through the courses slice to find the course with the matching ID. If found, it encodes the course into JSON and sends the response. If not found, it responds with a message indicating that no course was found with the given ID.
*/
func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from the request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return
}

/*
Notes
This function handles the request to create a new course. It first checks if the request body is empty. Then, it decodes the request body into a Course struct. If the decoded course is empty, it responds with a message indicating that there is no data inside the JSON. Otherwise, it generates a unique ID for the course, appends it to the courses slice, and responds with the created course in JSON format.
*/
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")
	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	// Decode the request body into a Course struct
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	// Check if the decoded course is empty
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No Data inside the json")
		return
	}

	// Generate a unique ID and append the course to the courses slice
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

/*
	Notes

This function handles the request to update an existing course. It first extracts the ID from the request parameters. Then, it iterates through the courses slice to find the course with the matching ID. If found, it removes the old course, decodes the request body into a new Course struct, updates the ID of the updated course, appends the updated course to the courses slice, and responds with a success message. If no course with the given ID is found, it responds with a "Course not found" error message.
*/
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the Course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)
	// Loop through the courses, find the matching ID, and update the course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove the old course
			courses = append(courses[:index], courses[index+1:]...)

			// Decode the request body into a new Course struct
			var course Course
			_ = json.NewDecoder(r.Body).Decode(course)

			// Set the ID of the updated course
			course.CourseId = params["id"]
			// Append the updated course to the courses slice
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course Updated Successfully")
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//loop,id,remove(index,index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			break

		}
	}

}
