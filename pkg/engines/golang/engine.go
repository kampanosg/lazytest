package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
)

const (
	suffix = "_test.go"
)

type GolangEngine struct {
}

func NewGolangEngine() *GolangEngine {
	return &GolangEngine{}
}

func (g *GolangEngine) Load(dir string, f fs.FileInfo) (*models.LazyTestSuite, error) {
	if !strings.HasSuffix(f.Name(), suffix) {
		return nil, nil
	}

	fp := filepath.Join(dir, f.Name())
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fp, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file, %w", err)
	}

	suite := &models.LazyTestSuite{
		Path: fp,
	}

	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if ok && (strings.HasPrefix(fn.Name.Name, "Test") || strings.HasSuffix(fn.Name.Name, "Test")) {
			suite.Tests = append(suite.Tests, models.LazyTest{
				Name: fn.Name.Name,
			})
		}
	}

	return suite, nil
}
