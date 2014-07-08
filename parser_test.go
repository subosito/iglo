package iglo

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	data, err := ParseMarkdown(strings.NewReader(dummyMarkdown))
	if err != nil {
		t.Errorf("ParseMarkdown() returned an error %s", err)
	}

	ioutil.WriteFile("x.json", data, os.ModePerm)

	a, err := ParseJSON(bytes.NewBuffer(data))
	if err == nil {
		if !reflect.DeepEqual(a, dummyAPI) {
			t.Errorf("ParseJSON() returned %+v, want %+v", a, dummyAPI)
		}
	}
}

func TestCheckVersion(t *testing.T) {
	versions := map[string]bool{
		"0.12.0": true,
		"0.11.0": true,
		"0.10.1": true,
		"0.10.0": true,
		"0.9.0":  true,
		"0.8.1":  false,
		"0.8.0":  false,
		"0.7.5":  false,
		"0.7.4":  false,
		"0.7.3":  false,
		"0.7.2":  false,
		"0.7.1":  false,
		"0.7.0":  false,
		"0.1.1":  false,
	}

	for v, b := range versions {
		err := CheckVersion(v)

		if (err != nil) == b {
			t.Errorf("CheckVersion() is returning %s with version %s", err, v)
		}
	}
}
