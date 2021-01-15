package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	//NOTE: It is very important that the name of the variables of your structures begins with a capital letter
	Title            string `json:"title"`
	Number_of_videos int    `json:"number_videos"`
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		course := Course{"Golang Course", 27}
		json.NewEncoder(w).Encode(course)
	})

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		courses := Courses{
			Course{"Golang Course", 27},
			Course{"Python Course", 25},
			Course{"NodeJS Course", 50},
			Course{"VueJS Course", 40},
		}
		json.NewEncoder(w).Encode(courses)
	})

	http.ListenAndServe(":8000", nil)
}
