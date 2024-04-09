package loader

import (
	"fmt"

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

func (l *LazyTestLoader) LoadLazyTests(dir string) (*tview.TreeNode, error) {
	root := tview.NewTreeNode(dir)
	for _, engine := range l.Engines {
		t, err := engine.Load(dir)
		if err != nil {
			return nil, fmt.Errorf("error parsing test suite: %w", err)
		}

		if t == nil || t.Root == nil || len(t.Root.Children) == 0 {
			continue
		}

		testNode := toTUINodes(engine, t.Root)
		root.AddChild(testNode)
	}
	return root, nil
}

func toTUINodes(engine engines.LazyEngine, lazyNode *models.LazyNode) *tview.TreeNode {
	var node *tview.TreeNode

	if lazyNode.IsDir() {
		hasTests := false
		node = tview.NewTreeNode(fmt.Sprintf("[white]%s", lazyNode.Name))

		for _, child := range lazyNode.Children {
			childNode := toTUINodes(engine, child)

			if childNode == nil {
				continue
			}

			node.AddChild(childNode)
			hasTests = true
		}

		if hasTests {
			return node
		}
	} else {
		node = tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", engine.GetIcon(), lazyNode.Name))
		node.SetReference(lazyNode.Ref)
		node.SetSelectable(true)

		for _, t := range lazyNode.Children {
			test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
			test.SetSelectable(true)
			test.SetReference(t.Ref)
			node.AddChild(test)
		}
		return node
	}

	return node
}