// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	"os"
	"strings"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// tidy code
func Fmt() error {
	f, _ := os.Open("./")
	fs, _ := f.Readdir(-1)

	items := make([]string, 0, 32)
	for _, it := range fs {
		if it.IsDir() {
			if strings.HasPrefix(it.Name(), ".") {
				continue
			}
			items = append(items, it.Name())
			continue
		}

		if strings.HasSuffix(it.Name(), ".go") {
			items = append(items, it.Name())
		}
	}
	return sh.Run("gofmt", append([]string{"-s", "-l", "-w"}, items...)...)
}

// golangci-lint
func Lint() error {
	return sh.RunV("golangci-lint", "run", "--fix")
}
