package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

var courses []Course

func main() {
	courses = append(courses, Course{
		CourseId: "1",
		CourseName: "GO_1",
		CoursePrice: 500,
		Author: &Author{
			Fullname: "Subhranil",
			Website: "subhranil.dev",
		},
	})
	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "GO_2",
		CoursePrice: 700,
		Author: &Author{
			Fullname: "Subhranil",
			Website: "subhranil.dev",
		},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", serveHome)
	mux.HandleFunc("GET /courses", getAllCourses)
	mux.HandleFunc("POST /courses", createCourse)
	mux.HandleFunc("GET /courses/{id}", getCourse)
	mux.HandleFunc("PATCH /courses/{id}", updateCourse)
	mux.HandleFunc("DELETE /courses/{id}", deleteCourse)

	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
	}
	fmt.Println("Server is UP and RUNNING on localhost:3000")
	log.Fatal(server.ListenAndServe())
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server is UP and RUNNING</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by id")
	
	id := r.PathValue("id")

	for _, course := range courses {
		if course.CourseId == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Not found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)

	fmt.Println("Create a course : ", course)

	course.CourseId = strconv.Itoa(rand.Int())
	courses = append(courses, course)

	json.NewEncoder(w).Encode(courses)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	id := r.PathValue("id")

	fmt.Println("Update a course by id : ", id)

	for i, course := range courses {
		if course.CourseId == id {
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses[i] = course
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Println("Delete a course with id : ", id)

	for i, course := range courses {
		if course.CourseId == id {
			courses = append(courses[:i], courses[i+1:]... )
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Not found")
}