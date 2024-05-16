package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/kampanosg/lazytest/internal/clipboard"
	"github.com/kampanosg/lazytest/internal/runner"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/internal/tui/elements"
	"github.com/kampanosg/lazytest/internal/tui/handlers"
	"github.com/kampanosg/lazytest/internal/tui/state"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/bashunit"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/engines/pytest"
	"github.com/kampanosg/lazytest/pkg/engines/rust"
	"github.com/rivo/tview"
	"github.com/spf13/afero"
)

const (
	Version = "v0.4.2"
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

	a := tview.NewApplication()
	h := handlers.NewHandlers()
	r := runner.NewRunner()
	e := elements.NewElements()
	c := clipboard.NewClipboardManager()
	s := state.NewState()

	excludedEngines := strings.Split(*exc, ",")
	var engines []engines.LazyEngine

	if !slices.Contains(excludedEngines, "golang") {
		engines = append(engines, golang.NewGoEngine(afero.NewOsFs()))
	}

	if !slices.Contains(excludedEngines, "bashunit") {
		engines = append(engines, bashunit.NewBashunitEngine(afero.NewOsFs()))
	}

	if !slices.Contains(excludedEngines, "rust") {
		engines = append(engines, rust.NewRustEngine(r))
	}

	if !slices.Contains(excludedEngines, "pytest") {
		engines = append(engines, pytest.NewPytestEngine(r))
	}

	t := tui.NewTUI(a, h, r, c, e, s, *dir, engines)

	if err := t.Run(); err != nil {
		panic(err)
	}
}
