package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/kampanosg/lazytest/pkg/tree"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory:", err)
		return
	}

	root := tree.NewFolder(currentDir)

	err = traverseDir(currentDir, root)
	if err != nil {
		fmt.Println("Error traversing directory:", err)
		return
	}

	root.TraverseDFS("")
}

func traverseDir(dir string, parent *tree.LazyNode) error {
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

			err := traverseDir(filepath.Join(dir, fileInfo.Name()), node)
			if err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(fileInfo.Name(), "_test.go") {
				node = &tree.LazyNode{
					Name:     fileInfo.Name(),
					IsFolder: false,
					Suite:    models.LazyTestSuite{Path: filepath.Join(dir, fileInfo.Name())},
				}
			}
		}

		if node != nil {
			parent.AddChild(node)
		}
	}
	return nil
}
