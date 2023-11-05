package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
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

        content = append(content, slurp[:]...)
    }

    return content
}

func index(w http.ResponseWriter, r *http.Request) {
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

    str := bytes.NewBuffer(data).String()
    fmt.Printf("%s\n", str);
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", index)

    err := http.ListenAndServe(":80", mux)
    
    if (err != nil) {
        log.Fatal(err)
    }
}
