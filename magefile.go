//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "goresumake", "./cmd/goresumake.go")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s", b)
	return nil
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./resumake", "/usr/bin/resumake")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "mod", "download")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s", b)
	return nil
}

func Test() error {
	cmd := exec.Command("go", "test", "-cover", "./...")

	b, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s", b)
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("goresumake")
}
