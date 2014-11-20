package main

import "testing"

func TestGlobalFlags(t *testing.T) {
	if len(globalFlags) != 2 {
		t.Errorf("Invalid globalFlags length: %d\n", len(globalFlags))
	}
}
