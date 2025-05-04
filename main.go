package main

import (
	"fmt"
	"log"
	"net/http"
)

func parseHost(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Fprintf(w, "host: %s\n", host)
}

func parsePath(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	fmt.Fprintf(w, "path: %s\n", path)
}

func main() {
	http.HandleFunc("/parse-host", parseHost)
	http.HandleFunc("/parse-path", parsePath)

	log.Println("starting http server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
