package iglo

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestDrafterVersion(t *testing.T) {
	path, err := drafter()
	if err != nil {
		t.Error("Drafter is not available on the system.")
		panic(err)
	}

	cmd := exec.Command(path, "--version")

	output, err := cmd.Output()
	if err != nil {
		t.Error("The drafter command failed.")
		panic(err)
	}

	expected := []byte("v2.3.0-pre.2\n")
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected:\n%+v\nActual:\n%+v", string(expected), string(output))
	}
}
