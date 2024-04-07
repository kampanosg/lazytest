package main

import (
	"flag"
	"fmt"

	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/spf13/afero"
)

const (
	Version = "v.0.2.0"
)

func main() {
	dir := flag.String("dir", ".", "the directory to start searching for tests")
	// exc := flag.String("excl", "", "engines to exclude")
	vsn := flag.Bool("version", false, "the current version of LazyTest")
	flag.Parse()

	if *vsn {
		fmt.Printf("LazyTest %s\n", Version)
		return
	}

	g := golang.GolangEngine2{
		FS: afero.NewOsFs(),
	}

	t, err := g.Vroom(*dir)
	fmt.Println(t, err)

	// excludedEngines := strings.Split(*exc, ",")
	// var engines []engines.LazyEngine
	//
	// if !slices.Contains(excludedEngines, "golang") {
	// 	engines = append(engines, golang.NewGolangEngine())
	// }
	//
	// if !slices.Contains(excludedEngines, "bashunit") {
	// 	engines = append(engines, bashunit.NewBashunitEngine())
	// }
	//
	// a := tview.NewApplication()
	// h := handlers.NewHandlers()
	// r := runner.NewRunner()
	// e := elements.NewElements()
	// c := clipboard.NewClipboardManager()
	// s := state.NewState()
	//
	// t := tui.NewTUI(a, h, r, c, e, s, *dir, engines)
	//
	// if err := t.Run(); err != nil {
	// 	panic(err)
	// }
}
