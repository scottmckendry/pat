package main

import (
	"testing"
)

func TestPathExists(t *testing.T) {
	path := "main.go"
	if !pathExists(path) {
		t.Errorf("Expected %s to exist", path)
	}
}

func TestPathDoesNotExist(t *testing.T) {
	path := "nonexistentfile"
	if pathExists(path) {
		t.Errorf("Expected %s to not exist", path)
	}
}
