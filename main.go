package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func readFormFile(mr *multipart.Reader) []byte {
	var content []byte

	for {
		p, err := mr.NextPart()

		if err != nil {
			break
		}

		slurp, err := io.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}

		content = append(content, slurp...)
	}

	return content
}

func export(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	contentType := r.Header.Get("Content-type")
	mediaType, params, err := mime.ParseMediaType(contentType)

	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		log.Fatal("Unsupported content-type")
	}

	mr := multipart.NewReader(r.Body, params["boundary"])
	data := readFormFile(mr)

	file, err := os.CreateTemp("uploads", "aseprite_")
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Base(file.Name())

	defer file.Close()

	file.Write(data)

	destFile := fmt.Sprintf("out/%s.png", fileName)

	cmd := exec.Command(os.Getenv("ASEPRITE"), "-b", file.Name(), "--sheet", destFile)
	res, err := cmd.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	f, err := os.ReadFile(destFile)

	if err != nil {
		http.Error(w, bytes.NewBuffer(res).String(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/octet-stream")
	w.Write(f)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("ASEPRITE") == "" {
		log.Fatal("missing ASEPRITE environment variable")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/export", export)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}
