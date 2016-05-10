package iglo

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestDrafterVersion(t *testing.T) {
	path, err := drafter()
	if err != nil {
		t.Fatalf("Drafter is not available on the system. err=%q", err)
	}

	cmd := exec.Command(path, "--version")

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("The drafter command failed. err=%q", err)
	}

	if want := []byte("v2.3."); !bytes.HasPrefix(output, want) {
		t.Fatalf("Got output %q, want prefix %q", output, want)
	}
}
