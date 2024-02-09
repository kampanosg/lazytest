package main

import (
	"github.com/kampanosg/lazytest/internal/loader"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"
)

func main() {
	currentDir := "."
	root := tree.NewFolder(currentDir)

	loader := loader.NewLazyTestLoader([]engines.LazyTestEngine{
		golang.NewGolangEngine(),
	})

	if err := loader.LoadLazyTests(currentDir, root); err != nil {
		panic(err)
	}

	root.TraverseDFS("")
}
