package main

import (
	"github.com/refs/j/cmd"
	_ "github.com/refs/j/pkg/vcs/github" // load github defaults.
)

func main() {
	cmd.Execute()
}
