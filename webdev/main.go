package main

import (
	"net/http"
	"strconv"
	"time"
)

var courses = map[int64]string{
	1: "Introduction to programming",
	2: "Introduction to algorithms",
	3: "Data structures",
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/courses/description", CourseDescHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go to /courses/description"))
}

func CourseDescHandler(w http.ResponseWriter, r *http.Request) {
	// BEGIN (write your solution here)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	courseIDStr := r.URL.Query().Get("course_id")
	if courseIDStr == "" {
		http.Error(w, "course_id parameter is required", http.StatusBadRequest)
		return
	}
	
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid course_id format", http.StatusBadRequest)
		return
	}
	
	courseName, exists := courses[courseID]
	if !exists {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(courseName))
	// END
}
