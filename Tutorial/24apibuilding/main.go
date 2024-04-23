package main

import "fmt"

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
}
