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

		api, err := ParseJSON(f)
		if err != nil {
			t.Error("ParseJSON returned an error %s", err)
		}

		o, _ := os.Create("index.html")
		defer f.Close()

		w := bufio.NewWriter(o)

		err = HTML(w, api)
		if err != nil {
			t.Errorf("HTML returned an error %s", err)
		}

		w.Flush()
	}
}
