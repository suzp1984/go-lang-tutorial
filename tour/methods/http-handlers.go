package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	mux := http.NewServeMux()
	
	mux.Handle("/string", String("I'm a frayed knot."))
	mux.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	
	log.Println("Listening...")
	err := http.ListenAndServe("localhost:4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
