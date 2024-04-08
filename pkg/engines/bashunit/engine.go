package bashunit

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/spf13/afero"
)

const (
	suiteType = "bashunit"
	suffix    = ".sh"
	icon      = "󱆃"
)

type FileSystem interface {
	Open(name string) (afero.File, error)
}

type BashEngine struct {
	FS FileSystem
}

func NewBashunitEngine(fs FileSystem) *BashEngine {
	return &BashEngine{
		FS: fs,
	}
}

func (b *BashEngine) GetIcon() string {
	return icon
}

func (b *BashEngine) Load(dir string) (*models.LazyTree, error) {
	fileInfos, err := b.loadFiles(dir)
	if err != nil {
		return nil, err
	}

	root := models.NewLazyNode(dir, nil)

	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		node, err := b.doLoad(filepath.Join(dir, fileInfo.Name()), fileInfo)
		if err != nil {
			return nil, fmt.Errorf("error loading tests: %w", err)
		}

		if node == nil {
			continue
		}

		root.AddChild(node)
	}

	return models.NewLazyTree(root), nil
}

func (b *BashEngine) loadFiles(path string) ([]fs.FileInfo, error) {
	dir, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	return fileInfos, err
}

func (b *BashEngine) doLoad(dir string, f fs.FileInfo) (*models.LazyNode, error) {
	var node *models.LazyNode
	if f.IsDir() {
		children, err := b.loadFiles(dir)
		if err != nil {
			return nil, fmt.Errorf("error loading files: %w", err)
		}

		hasTests := false
		node = models.NewLazyNode(f.Name(), nil)

		for _, child := range children {
			childNode, err := b.doLoad(filepath.Join(dir, child.Name()), child)
			if err != nil {
				return nil, fmt.Errorf("error loading child: %w", err)
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
		suite, err := b.parseTestSuite(dir)
		if err != nil {
			return nil, fmt.Errorf("error finding lazy test suite: %w", err)
		}

		if suite != nil {
			node = models.NewLazyNode(f.Name(), suite)

			for _, t := range suite.Tests {
				test := models.NewLazyNode(t.Name, t)
				node.AddChild(test)
			}
			return node, nil
		}
	}

	return nil, nil
}

func (b *BashEngine) parseTestSuite(fp string) (*models.LazyTestSuite, error) {
	if !strings.HasSuffix(fp, suffix) {
		return nil, nil
	}

	file, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	suite := &models.LazyTestSuite{
		Path: fp,
		Icon: icon,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "function test_") {
			name := strings.Fields(line)[1]
			name = strings.TrimSuffix(name, "()")
			test := &models.LazyTest{
				Name:   name,
				RunCmd: fmt.Sprintf("bashunit -v -S -f \"%s\" %s", name, fp),
			}
			suite.Tests = append(suite.Tests, test)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(suite.Tests) == 0 {
		return nil, nil
	}

	return suite, nil
}
