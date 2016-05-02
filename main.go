package main

import (
	"net/http"
	"log"
	"flag"
	"fmt"
)

var (
	addr = flag.String("addr", "0.0.0.0:3000", "bind adderess and port.")
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	flag.Parse()
	log.Println("Starting server @", *addr)

	http.HandleFunc("/", Handler) 

	err := http.ListenAndServe(*addr, nil) 

	if err != nil {
		log.Println("ERROR:", err)
	}
}
