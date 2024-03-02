package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/kampanosg/lazytest/internal/runner"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
)

const (
	Version = "v.0.1.0"
)

func main() {
	dir := flag.String("dir", ".", "the directory to start searching for tests")
	exc := flag.String("excl", "", "engines to exclude")
	vsn := flag.Bool("version", false, "the current version of LazyTest")
	flag.Parse()

	if *vsn {
		fmt.Printf("LazyTest %s\n", Version)
		return
	}

	excludedEngines := strings.Split(*exc, ",")
	var engines []engines.LazyEngine

	if !slices.Contains("golang", excludedEngines) {
		engines = append(engines, golang.NewGolangEngine())
	}

	r := runner.NewRunner()
	t := tui.NewTUI(*dir, r, engines)

	if err := t.Run(); err != nil {
		panic(err)
	}
}
