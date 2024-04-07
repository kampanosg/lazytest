package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/spf13/afero"
)

const (
	suffix    = "_test.go"
	suiteType = "golang"
	icon      = "󰟓"
)

type FileSystem interface {
	Open(name string) (afero.File, error)
}

type GoEngine struct {
	FS FileSystem
}

func NewGoEngine(fs FileSystem) *GoEngine {
	return &GoEngine{
		FS: fs,
	}
}

func (g *GoEngine) GetIcon() string {
	return "󰟓"
}

func (g *GoEngine) Load(dir string) (*models.LazyTree, error) {
	fileInfos, err := g.loadFiles(dir)
	if err != nil {
		return nil, err
	}

	isGo := slices.ContainsFunc(fileInfos, func(fi fs.FileInfo) bool {
		return fi.Name() == "go.mod"
	})

	if !isGo {
		return nil, nil
	}

	root := models.NewLazyNode(dir, nil)

	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		node, err := g.doLoad(filepath.Join(dir, fileInfo.Name()), fileInfo)
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

func (g *GoEngine) doLoad(dir string, f fs.FileInfo) (*models.LazyNode, error) {
	var node *models.LazyNode
	if f.IsDir() {
		children, err := g.loadFiles(dir)
		if err != nil {
			return nil, fmt.Errorf("error loading files: %w", err)
		}

		hasTests := false
		node = models.NewLazyNode(f.Name(), nil)

		for _, child := range children {
			childNode, err := g.doLoad(filepath.Join(dir, child.Name()), child)
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
		suite, err := g.parseTestSuite(dir)
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

func (g *GoEngine) loadFiles(path string) ([]fs.FileInfo, error) {
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

func (g *GoEngine) parseTestSuite(fp string) (*models.LazyTestSuite, error) {
	if !strings.HasSuffix(fp, suffix) {
		return nil, nil
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fp, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file, %w", err)
	}

	suite := &models.LazyTestSuite{
		Path: fp,
		Type: suiteType,
		Icon: icon,
	}

	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if ok && (strings.HasPrefix(fn.Name.Name, "Test") || strings.HasSuffix(fn.Name.Name, "Test")) {
			suite.Tests = append(suite.Tests, &models.LazyTest{
				Name:   fn.Name.Name,
				RunCmd: fmt.Sprintf("go test -v -run %s ./%s", fn.Name.Name, removeFileFromFilepath(fp)),
			})
		}
	}

	return suite, nil
}

func removeFileFromFilepath(path string) string {
	if !strings.HasSuffix(path, ".go") {
		return path
	}

	parts := strings.Split(path, "/")
	return strings.Join(parts[:len(parts)-1], "/")
}
