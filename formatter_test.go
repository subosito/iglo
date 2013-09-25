package iglo

import (
	"bufio"
	"os"
	"testing"
	// "bytes"
)

func TestFormatter(t *testing.T) {
	// data := `{
	// 	"_version": "1.0",
	// 	"name": "Hello API",
	// 	"description": "A simple API demo",
	// 	"metadata": {
	// 		"FORMAT": {
	// 			"value": "1A"
	// 		},
	// 		"HOST": {
	// 			"value": "https://api.example.com/v1"
	// 		}
	// 	}
	// }`

	// err := HTML(os.Stdout, bytes.NewBufferString(data))
	f, _ := os.Open("api.json")
	defer f.Close()

	o, _ := os.Create("index.html")
	defer f.Close()

	w := bufio.NewWriter(o)

	err := HTML(w, f)
	if err != nil {
		t.Errorf("HTML returned an error %s", err)
	}

	w.Flush()
}
