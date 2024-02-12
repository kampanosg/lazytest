package loader

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/kampanosg/lazytest/pkg/tree"
)

type LazyTestLoader struct {
	Engines []engines.LazyTestEngine
}

func NewLazyTestLoader(e []engines.LazyTestEngine) *LazyTestLoader {
	return &LazyTestLoader{
		Engines: e,
	}
}

// LoadLazyTests will load the lazy tests from the given directory and add them to the parent node
func (l *LazyTestLoader) LoadLazyTests(dir string, parent *tree.LazyNode) error {
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
		var node *tree.LazyNode
		
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		if fileInfo.IsDir() {
			node = tree.NewFolder(fileInfo.Name())
			if err := l.LoadLazyTests(filepath.Join(dir, fileInfo.Name()), node); err != nil {
				return fmt.Errorf("error loading lazy tests: %w", err)
			}
		} else {
			suite, err := l.findLazyTestSuite(dir, fileInfo)
			if err != nil {
				return fmt.Errorf("error finding lazy test suite: %w", err)
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
