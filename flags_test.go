package main_test

import (
	"testing"

	. "github.com/goldeneggg/gat"
)

func TestGlobalFlags(t *testing.T) {
	if len(GlobalFlags) != 2 {
		t.Errorf("Invalid globalFlags length: %d\n", len(GlobalFlags))
	}
}
