package main

import (
	"flag"
	"log"
	"os"

	"github.com/subosito/iglo"
)

var fname = flag.String("out", "output.html", "Filename of the HTML output")

func main() {
	flag.Parse()

	f, err := os.Open("../API.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w, err := os.Create(*fname)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	err = iglo.MarkdownToHTML(w, f)
	if err != nil {
		log.Fatal(err)
	}
}
