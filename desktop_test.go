package desktop

import (
	"runtime"
	"testing"
)

func TestCommand(t *testing.T) {
	if Command == nil {
		t.Fatalf("no Command on %q", runtime.GOOS)
	}
}
