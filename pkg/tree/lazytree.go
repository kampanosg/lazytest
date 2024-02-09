package tree

import (
	"fmt"

	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyTree struct {
	Nodes []LazyNode
}

type LazyNode struct {
	Name     string
	IsFolder bool
	Children []*LazyNode
	Suite    *models.LazyTestSuite
}

func NewFolder(name string) *LazyNode {
	return &LazyNode{
		Name:     name,
		IsFolder: true,
	}
}

func (n *LazyNode) AddChild(child *LazyNode) {
	n.Children = append(n.Children, child)
}

func (n *LazyNode) TraverseDFS(padding string) {
	if n.IsFolder {
		fmt.Printf("%s%s/\n", padding, n.Name)
	} else {
		fmt.Printf("%s%s\n", padding, n.Name)
		for _, t := range n.Suite.Tests {
			fmt.Printf("  %s- %s\n", padding, t.Name)
		}
	}

	for _, child := range n.Children {
		child.TraverseDFS(padding + "  ")
	}
}
