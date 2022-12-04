/* The desktop package exposes Open function which acts like the user clicked on a uri. */
package desktop

import (
	"fmt"
	"os/exec"
	"runtime"
)

var Version = "0.2.0"

// Open calls the OS default program for uri
// e.g. Open("http://www.google.com") will open the default browser on www.google.com
func Open(uri string) error {
	if Command == nil {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	args := append(Command, uri)
	cmd := exec.Command(args[0], args[1:]...)

	return cmd.Start()
}
