package main

import "testing"

// cd cmd/web; go test
// or go test -v
func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
