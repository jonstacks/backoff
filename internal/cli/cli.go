package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cenkalti/backoff"
)

// ExecStreamingWithBackoff executes the command in a streaming fashion with
// the given backoff policy
func ExecStreamingWithBackoff(bck backoff.BackOff, maxRetries int, name string, args []string) error {
	var err error
	ticker := backoff.NewTicker(bck)

	for _ = range ticker.C {
		cmd := exec.Command(name, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Start()
		if err == nil {
			err = cmd.Wait()
			if err == nil {
				return nil
			}
		}

		maxRetries--
		fmt.Fprintf(os.Stderr, "Failed to run '%s': %s. %d attempts remaining...\n", name, err, maxRetries)

		if maxRetries <= 0 {
			return fmt.Errorf("max number of retries exceded: %s", err)
		}
	}

	return err
}
