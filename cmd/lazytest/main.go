package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kampanosg/lazytest/internal/loader"
	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/engines/golang"
	"github.com/kampanosg/lazytest/pkg/tree"
)

func main() {
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("error getting current directory:", err)
	// 	return
	// }
	//
	currentDir := "."
	root := tree.NewFolder(currentDir)

	loader := loader.NewLazyTestLoader([]engines.LazyTestEngine{
		golang.NewGolangEngine(),
	})

	err := traverseDir(currentDir, root, loader)
	if err != nil {
		fmt.Println("Error traversing directory:", err)
		return
	}

	root.TraverseDFS("")
}

func traverseDir(dir string, parent *tree.LazyNode, loader *loader.LazyTestLoader) error {
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
			if err := traverseDir(filepath.Join(dir, fileInfo.Name()), node, loader); err != nil {
				return err
			}
		} else {
			suite, err := loader.LoadTestSuite(dir, fileInfo)
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
