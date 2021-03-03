package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"first"`
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

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

func bar(w http.ResponseWriter, r *http.Request)  {
	people := []person{}
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("bad data to decode:", err)
	}
	fmt.Println(people)
}
