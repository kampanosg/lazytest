package pytest

import (
	"fmt"
	"strings"

	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
)

const icon = ""

type pyNode struct {
	Name     string
	Ref      any
	Children map[string]*pyNode
}

type PytestEngine struct {
	Runner engines.Runner
}

func NewPytestEngine(r engines.Runner) *PytestEngine {
	return &PytestEngine{
		Runner: r,
	}
}

func (p *PytestEngine) GetIcon() string {
	return icon
}

func (p *PytestEngine) Load(dir string) (*models.LazyTree, error) {
	o, err := p.Runner.RunCmd("python3 -m pytest --collect-only -q | head -n -2")
	if err != nil {
		return nil, nil
	}

	root := &pyNode{
		Name:     dir,
		Ref:      nil,
		Children: make(map[string]*pyNode),
	}

	lines := strings.Split(o, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		testLine := strings.ReplaceAll(line, "/", " ")
		testLine = strings.ReplaceAll(testLine, "::", " ")

		parts := strings.Split(testLine, " ")
		if len(parts) == 0 {
			continue
		}

		currentNode := root
		var testSuite *models.LazyTestSuite

		for i, part := range parts {
			childNode, exists := currentNode.Children[part]
			if !exists {
				childNode = &pyNode{
					Name:     part,
					Children: make(map[string]*pyNode),
				}

				if i == len(parts)-2 {
					childNode.Ref = &models.LazyTestSuite{
						Path:  strings.Join(parts[:i+1], "::"),
						Tests: make([]*models.LazyTest, 0),
					}
				}
				currentNode.Children[part] = childNode
			}
			currentNode = childNode

			if i == len(parts)-2 {
				testSuite = currentNode.Ref.(*models.LazyTestSuite)
			}

			if i == len(parts)-1 {
				test := &models.LazyTest{
					Name:   part,
					RunCmd: fmt.Sprintf("python3 -m pytest -x --verbose %s", line),
				}
				childNode.Ref = test
				testSuite.Tests = append(testSuite.Tests, test)
			}
		}
	}

	if len(root.Children) == 0 {
		return nil, nil
	}

	lazyRoot := toLazyTree(root)
	return models.NewLazyTree(lazyRoot), nil
}

func toLazyTree(r *pyNode) *models.LazyNode {
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
