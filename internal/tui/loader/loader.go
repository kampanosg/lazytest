package loader

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/rivo/tview"
)

type LazyTestLoader struct {
	Engines []engines.LazyEngine
}

func NewLazyTestLoader(e []engines.LazyEngine) *LazyTestLoader {
	return &LazyTestLoader{
		Engines: e,
	}
}

func (l *LazyTestLoader) LoadLazyTests(dir string, root *tview.TreeNode) error {
	fileInfos, err := loadFiles(dir)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		var node *tview.TreeNode

		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		node, err = l.doLoad(filepath.Join(dir, fileInfo.Name()), fileInfo)
		if err != nil {
			return err
		}

		if node == nil {
			continue
		}

		root.AddChild(node)

	}
	return nil
}

func (l *LazyTestLoader) doLoad(dir string, f fs.FileInfo) (*tview.TreeNode, error) {
	var node *tview.TreeNode
	if f.IsDir() {
		children, err := loadFiles(dir)
		if err != nil {
			return nil, err
		}

		hasTests := false
		node = tview.NewTreeNode(fmt.Sprintf("[white]%s", f.Name()))

		for _, child := range children {
			childNode, err := l.doLoad(filepath.Join(dir, child.Name()), child)
			if err != nil {
				return nil, err
			}

			if childNode == nil {
				continue
			}

			node.AddChild(childNode)
			hasTests = true
		}

		if hasTests {
			return node, nil
		}
	} else {
		suite, err := l.findLazyTestSuite(dir)
		if err != nil {
			return nil, fmt.Errorf("error finding lazy test suite: %w", err)
		}

		if suite != nil {
			node = tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", suite.Icon, f.Name()))
			node.SetReference(suite)
			node.SetSelectable(true)

			for _, t := range suite.Tests {
				test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
				test.SetSelectable(true)
				test.SetReference(t)
				node.AddChild(test)
			}
			return node, nil
		}
	}

	return nil, nil
}

// findLazyTestSuite will use the engines to find a test suite in the given path and file
func (l *LazyTestLoader) findLazyTestSuite(path string) (*models.LazyTestSuite, error) {
	for _, engine := range l.Engines {
		suite, err := engine.ParseTestSuite(path)
		if err != nil {
			return nil, fmt.Errorf("error parsing test suite: %w", err)
		}
		if suite != nil {
			return suite, nil
		}
	}
	return nil, nil
}

func loadFiles(dir string) ([]fs.FileInfo, error) {
	file, err := os.Open(filepath.Clean(dir))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}
	return fileInfos, err
}
