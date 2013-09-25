package iglo

import (
	"bufio"
	"os"
	"testing"
)

func TestFormatter(t *testing.T) {
	f, err := os.Open("api.json")

	if err == nil {
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
}
