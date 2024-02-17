package main

import (
	"github.com/kampanosg/lazytest/internal/loader"
	"github.com/kampanosg/lazytest/internal/runner"
	"github.com/kampanosg/lazytest/internal/tui"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"
)

func main() {
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	currentDir := "."
	root := tree.NewFolder(currentDir)

	loader := loader.NewLazyTestLoader([]engines.LazyTestEngine{
		golang.NewGolangEngine(),
	})

	if err := loader.LoadLazyTests(currentDir, root); err != nil {
		panic(err)
	}

	r := runner.NewRunner()

	t := tui.NewTUI(root, r)
	if err := t.Run(); err != nil {
		panic(err)
	}
}
