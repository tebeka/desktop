package desktop

import (
	"runtime"
	"testing"
)

func TestOpenBadPlatform(t *testing.T) {
	saved := commands[runtime.GOOS]
	delete(commands, runtime.GOOS)
	err := Open("")
	commands[runtime.GOOS] = saved

	if err == nil {
		t.Fatalf("open succeeded on unkown platform")
	}
}
