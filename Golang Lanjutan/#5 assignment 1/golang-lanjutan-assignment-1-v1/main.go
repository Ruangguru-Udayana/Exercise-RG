package main

import (
	"fmt"
	"net/http"
)




func GetStudyProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	}
}


func main() {
	http.HandleFunc("/study-program", GetStudyProgram())
	http.HandleFunc("/user/add", AddUser())
	http.HandleFunc("/user/delete", DeleteUser())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
