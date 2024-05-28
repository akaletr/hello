package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8888",
	}

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	log.Fatal(server.ListenAndServe())
}
