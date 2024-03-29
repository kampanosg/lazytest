package golang

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
)

const (
	suffix    = "_test.go"
	suiteType = "golang"
	icon      = "󰟓"
)

type GolangEngine struct {
}

func NewGolangEngine() *GolangEngine {
	return &GolangEngine{}
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
	return strings.Join(parts[:len(parts) - 1], "/")
}