package main

import (
	"log"
	"net/http"
	"os"

	"github.com/subosito/iglo"
)

func httpError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("../API.md")
		if err != nil {
			httpError(w, err)
			return
		}
		defer f.Close()

		err = iglo.MarkdownToHTML(w, f)
		if err != nil {
			httpError(w, err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
