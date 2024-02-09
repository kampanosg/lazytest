package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory:", err)
		return
	}

	root := tree.NewFolder(currentDir)

	ge := golang.NewGolangEngine()

	err = traverseDir(currentDir, root, ge)
	if err != nil {
		fmt.Println("Error traversing directory:", err)
		return
	}

	root.TraverseDFS("")
}

func traverseDir(dir string, parent *tree.LazyNode, engine engines.LazyTestEngine) error {
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		var node *tree.LazyNode
		if fileInfo.IsDir() {
			node = tree.NewFolder(fileInfo.Name())

			err := traverseDir(filepath.Join(dir, fileInfo.Name()), node, engine)
			if err != nil {
				return err
			}
		} else {
			suite, err := engine.Load(dir, fileInfo)
			if err != nil {
				return err
			}

			if suite != nil {
				node = &tree.LazyNode{
					Name:  fileInfo.Name(),
					Suite: *suite,
				}
			}
		}

		if node != nil {
			parent.AddChild(node)
		}
	}
	return nil
}
