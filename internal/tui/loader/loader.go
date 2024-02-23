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

// LoadLazyTests will load the lazy tests from the given directory and add them to the parent node
func (l *LazyTestLoader) LoadLazyTests(dir string, root *tview.TreeNode) error {
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, fileInfo := range fileInfos {
		var node *tview.TreeNode

		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		if fileInfo.IsDir() {
			node = tview.NewTreeNode(fmt.Sprintf("[white]%s", fileInfo.Name()))
			node.SetSelectable(true)
			if err := l.LoadLazyTests(filepath.Join(dir, fileInfo.Name()), node); err != nil {
				return fmt.Errorf("error loading lazy tests: %w", err)
			}
		} else {
			suite, err := l.findLazyTestSuite(dir, fileInfo)
			if err != nil {
				return fmt.Errorf("error finding lazy test suite: %w", err)
			}

			if suite != nil {
				node = tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", suite.Icon, fileInfo.Name()))
				node.SetReference(suite)
				node.SetSelectable(true)

				for _, t := range suite.Tests {
					test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
					test.SetSelectable(true)
					test.SetReference(t)
					node.AddChild(test)
				}

			}
		}

		if node != nil {
			root.AddChild(node)
		}
	}
	return nil
}

// findLazyTestSuite will use the engines to find a test suite in the given path and file
func (l *LazyTestLoader) findLazyTestSuite(path string, f fs.FileInfo) (*models.LazyTestSuite, error) {
	for _, engine := range l.Engines {
		suite, err := engine.ParseTestSuite(path, f)
		if err != nil {
			return nil, fmt.Errorf("error parsing test suite: %w", err)
		}
		if suite != nil {
			return suite, nil
		}
	}
	return nil, nil
}
