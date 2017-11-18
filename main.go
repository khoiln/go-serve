package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	port := flag.String("p", "8080","port to serve")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
