package rust

import (
	"fmt"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
)

const (
	suffix = ".rs"
	icon   = ""
)

type Runner interface {
	RunCmd(cmd string) (string, error)
}

type RustEngine struct {
	Runner Runner
}

type rustNode struct {
	Name     string
	Ref      any
	Children map[string]*rustNode
}

func NewRustEngine(r Runner) *RustEngine {
	return &RustEngine{
		Runner: r,
	}
}

func (r *RustEngine) GetIcon() string {
	return icon
}

func (r *RustEngine) Load(dir string) (*models.LazyTree, error) {
	o, err := r.Runner.RunCmd("cargo test -- --list --format=terse")
	if err != nil {
		return nil, nil
	}

	root := &rustNode{
		Name:     dir,
		Ref:      nil,
		Children: make(map[string]*rustNode),
	}

	lines := strings.Split(o, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": test")
		if len(parts) != 2 {
			continue
		}

		testLine := parts[0]
		testParts := strings.Split(testLine, "::")

		currentNode := root

		var testSuite *models.LazyTestSuite

		for i, part := range testParts {
			part = strings.TrimSpace(part)
			childNode, exists := currentNode.Children[part]
			if !exists {
				childNode = &rustNode{
					Name:     part,
					Children: make(map[string]*rustNode),
				}

				if i == len(testParts)-2 {
					suite := &models.LazyTestSuite{
						Path:  strings.Join(testParts[:i+1], "::"),
						Tests: make([]*models.LazyTest, 0),
					}
					childNode.Ref = suite
				}
				currentNode.Children[part] = childNode
			}
			currentNode = childNode

			if i == len(testParts)-2 {
				testSuite = currentNode.Ref.(*models.LazyTestSuite)
			}

			if i == len(testParts)-1 {
				test := &models.LazyTest{
					Name: part,
					RunCmd: fmt.Sprintf("cargo t %s -- --exact", testLine),
				}
				childNode.Ref = test
				testSuite.Tests = append(testSuite.Tests, test)
			}
		}
	}

	lazyRoot := toLazyTree(root)
	return models.NewLazyTree(lazyRoot), nil
}

func toLazyTree(r *rustNode) *models.LazyNode {
	children := make([]*models.LazyNode, 0)
	for _, child := range r.Children {
		children = append(children, toLazyTree(child))
	}

	return &models.LazyNode{
		Name:     r.Name,
		Ref:      r.Ref,
		Children: children,
	}
}
