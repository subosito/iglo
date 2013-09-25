package iglo

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestParseJSON(t *testing.T) {
	a, err := ParseJSON(strings.NewReader(dummyJSON))
	if err != nil {
		t.Errorf("ParseJSON() returned an error %s", err)
	}

	if !reflect.DeepEqual(a, dummyAPI) {
		t.Errorf("ParseJSON() returned %+v, want %+v", a, dummyAPI)
	}
}

func TestParseMarkdown(t *testing.T) {
	data, err := ParseMarkdown(strings.NewReader(dummyMarkdown))
	if err != nil {
		t.Errorf("ParseMarkdown() returned an error %s", err)
	}

	a, err := ParseJSON(bytes.NewBuffer(data))
	if err == nil {
		if !reflect.DeepEqual(a, dummyAPI) {
			t.Errorf("ParseJSON() returned %+v, want %+v", a, dummyAPI)
		}
	}
}
