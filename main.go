package main

import (
	"flag"
	"net/http"
	"log"
	"io"
	"os"
)

type loggingMiddleware struct {
	writer io.Writer
	next http.Handler
	}

func (h *loggingMiddleware) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	h.writer.Write([]byte(r.RequestURI + "\n"))
	h.next.ServeHTTP(w, r)
}

func main() {
	port := flag.String("p", "8080","port to serve")
	dir := flag.String("d", ".", "directory to serve")
	flag.Parse()

	http.Handle("/", &loggingMiddleware{
		os.Stdout,
		http.FileServer(http.Dir(*dir)),
	})

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
