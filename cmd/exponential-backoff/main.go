package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cenkalti/backoff"
	"github.com/jonstacks/backoff/internal/cli"
)

var (
	retries int
)

func init() {
	flag.IntVar(&retries, "retries", 3, "Number of times to retry the command")
	flag.Parse()
}

func main() {
	var cmd string
	var args []string

	bck := backoff.NewExponentialBackOff()
	switch len(flag.Args()) {
	case 0:
		fmt.Fprintf(os.Stderr, "Not enough arguments supplied.")
	case 1:
		cmd = flag.Args()[0]
	default:
		cmd = flag.Args()[0]
		args = flag.Args()[1:]
	}
	err := cli.ExecStreamingWithBackoff(bck, retries, cmd, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
