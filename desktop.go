/* The desktop package exposes Open function which acts like the user clicked on a uri. */
package desktop

import (
	"fmt"
	"os/exec"
	"runtime"
)

var commands = map[string][]string{
	"windows": []string{"cmd", "/c", "start"},
	"darwin":  []string{"open"},
	"linux":   []string{"xdg-open"},
}

var Version = "0.1.1"

// Open calls the OS default program for uri
// e.g. Open("http://www.google.com") will open the default browser on www.google.com
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	run = append(run, uri)
	cmd := exec.Command(run[0], run[1:]...)

	return cmd.Start()
}
