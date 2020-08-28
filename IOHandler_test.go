package internalutils

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	err := SetEnvir("test_config.txt")
	if err != nil {
		t.Errorf("Setting environment error: %s", err)
	}

	v1 := os.Getenv("test")
	if v1 != "test" {
		t.Errorf("Error setting 1 variable: %s", err)
	}
	v2 := os.Getenv("test2")
	if v2 != "test2" {
		t.Errorf("Error setting 2 variable: %s", err)
	}
}
