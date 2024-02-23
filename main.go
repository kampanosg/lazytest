package main

import (
	"github.com/kampanosg/lazytest/internal/runner"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
)

func main() {
	dir := "."

	r := runner.NewRunner()

	engines := []engines.LazyEngine{
		golang.NewGolangEngine(),
	}

	t := tui.NewTUI(dir, r, engines)

	if err := t.Run(); err != nil {
		panic(err)
	}
}
