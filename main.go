package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    if (r.Method != http.MethodPost) {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", index)

    err := http.ListenAndServe(":80", mux)
    
    if (err != nil) {
        log.Fatal(err)
    }
}
