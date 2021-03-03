package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string `json:"first"`
}

func main() {

	http.HandleFunc("/encode", foo)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalln("Something went wrong starting web server:", err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Senny",
	}

	p2 := person{
		First: "James",
	}

	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("Bad data to encode", err)
	}
}
