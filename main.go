package main

import (
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func readFormFile(mr *multipart.Reader) []byte {
    var content []byte
    for {
        p, err := mr.NextPart()

        if (err != nil) {
            break
        }

        slurp, err := io.ReadAll(p)
        if (err != nil) {
            log.Fatal(err)
        }

        content = append(content, slurp...)
    }

    return content
}

func export(w http.ResponseWriter, r *http.Request) {
    if (r.Method != http.MethodPost) {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }

    contentType := r.Header.Get("Content-type")
    mediaType, params, err := mime.ParseMediaType(contentType)

    if (err != nil) {
        log.Fatal(err)
    }

    if (!strings.HasPrefix(mediaType, "multipart/")) {
        log.Fatal("Unsupported content-type")
    }

    mr := multipart.NewReader(r.Body, params["boundary"])
    data := readFormFile(mr)

    file, err := os.CreateTemp("uploads", "aseprite_")
    if (err != nil) {
        log.Fatal(err)
    }

    defer file.Close()

    file.Write(data)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/export", export)

    err := http.ListenAndServe(":80", mux)
    
    if (err != nil) {
        log.Fatal(err)
    }
}
