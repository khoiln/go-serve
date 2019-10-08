package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		os.Stdout.Write([]byte(r.RemoteAddr + " - " + r.RequestURI + "\n"))
		next.ServeHTTP(w, r)
	})
}

func noCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.String("p", "8080", "port to serve")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()
	fileServer := http.FileServer(http.Dir(*dir))

	http.Handle("/", noCacheMiddleware(loggingMiddleware(fileServer)))

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
