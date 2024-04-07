package golang

import (
	"errors"
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

type GolangEngine struct {
}

type GolangEngine2 struct {
	FS FileSystem
}

func NewGolangEngine() *GolangEngine {
	return &GolangEngine{}
}

func (g *GolangEngine2) Vroom(dir string) (*models.LazyTree, error) {
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

	return nil, errors.New("not implemented")
}

func (g *GolangEngine2) loadFiles(path string) ([]fs.FileInfo, error) {
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

func (g *GolangEngine) ParseTestSuite(fp string) (*models.LazyTestSuite, error) {
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
