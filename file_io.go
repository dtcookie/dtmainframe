package main

import (
	"os"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
