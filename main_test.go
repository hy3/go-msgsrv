package main

import (
	"flag"
	"os"
	"testing"
)

func TestFetchArgs(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	os.Args = append(os.Args, "-p", "1234")
	defer func() {
		os.Args = os.Args[:len(os.Args)-2]
	}()

	args := fetchArgs()
	if args.port != "1234" {
		t.Errorf("args.port => %s, want %s", args.port, "1234")
	}
}

func TestFetchArgs_NoOptions(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	args := fetchArgs()
	if args.port != defaultPort {
		t.Errorf("args.port => %s, want %s", args.port, defaultPort)
	}
}
