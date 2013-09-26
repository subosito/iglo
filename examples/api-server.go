package main

import (
	"github.com/subosito/iglo"
	"net/http"
	"log"
	"os"
	"bytes"
)

func httpError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("API.md")
		if err != nil {
			httpError(w, err); return
		}

		data, err := iglo.ParseMarkdown(f)
		if err != nil {
			httpError(w, err); return
		}

		api, err := iglo.ParseJSON(bytes.NewBuffer(data))
		if err != nil {
			httpError(w, err); return
		}

		err = iglo.HTML(w, api)
		if err != nil {
			httpError(w, err); return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
