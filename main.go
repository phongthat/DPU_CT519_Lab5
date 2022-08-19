package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func index(w http.ResponseWriter, req *http.Request){
	p := "html/index.html"
	http.ServeFile(w, req, p)
}

func about(w http.ResponseWriter, req *http.Request){
    p := "html/about.html"
    http.ServeFile(w, req, p)
}

func myresearch(w http.ResponseWriter, req *http.Request){
    p := "html/myresearch.html"
    http.ServeFile(w, req, p)
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/index", index)
    http.HandleFunc("/about", about)
    http.HandleFunc("/myresearch", myresearch)
    http.ListenAndServe(":8090", nil)
}
