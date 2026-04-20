package main

import "net/http"
type Course struct {
	Name string
	Instructor string
	Price int
}

func main() {
	course:= Course{
		Name: "Go Programming",
		Instructor: "John Doe",
		Price: 100,
	}
	println("Course Name:", course.Name)
	println("Instructor:", course.Instructor)
	println("Price:", course.Price)
	
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
} 