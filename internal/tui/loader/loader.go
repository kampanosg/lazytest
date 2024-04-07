package loader

import (
	"fmt"

	"github.com/kampanosg/lazytest/pkg/engines"
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

func (l *LazyTestLoader) LoadLazyTests(dir string, root *tview.TreeNode) error {
	for _, engine := range l.Engines {
		t, err := engine.Load(dir)
		if err != nil {
			return fmt.Errorf("error parsing test suite: %w", err)
		}

		if t == nil || t.Root == nil || len(t.Root.Children) == 0 {
			continue
		}

		tuiNodes := toTUINodes(engine, t.Root)
		root.AddChild(tuiNodes)
	}
	return nil
}

// func toTUINodes(engine engines.LazyEngine, lazyNode *models.LazyNode) *tview.TreeNode {
// 	return doToTUINodes(engine, lazyNode)
// }
//
// func doToTUINodes(engine engines.LazyEngine, lazyNode *models.LazyNode) *tview.TreeNode {
// 	var node *tview.TreeNode
//
// 	if lazyNode.IsDir() {
// 		hasTests := false
// 		node = tview.NewTreeNode(fmt.Sprintf("[white]%s", lazyNode.Name))
//
// 		for _, child := range lazyNode.Children {
// 			childNode := doToTUINodes(engine, child)
//
// 			if childNode == nil {
// 				continue
// 			}
//
// 			node.AddChild(childNode)
// 			hasTests = true
// 		}
//
// 		if hasTests {
// 			return node
// 		}
// 	} else {
// 		node = tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", engine.GetIcon(), lazyNode.Name))
// 		node.SetReference(lazyNode.Ref)
// 		node.SetSelectable(true)
//
// 		for _, t := range lazyNode.Children {
// 			test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
// 			test.SetSelectable(true)
// 			test.SetReference(t)
// 			node.AddChild(test)
// 		}
// 		return node
// 	}
//
// 	return node
// }

// func (l *LazyTestLoader) doLoad(dir string, f fs.FileInfo) (*tview.TreeNode, error) {
// 	var node *tview.TreeNode
// 	if f.IsDir() {
// 		children, err := loadFiles(dir)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		hasTests := false
// 		node = tview.NewTreeNode(fmt.Sprintf("[white]%s", f.Name()))
//
// 		for _, child := range children {
// 			childNode, err := l.doLoad(filepath.Join(dir, child.Name()), child)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			if childNode == nil {
// 				continue
// 			}
//
// 			node.AddChild(childNode)
// 			hasTests = true
// 		}
//
// 		if hasTests {
// 			return node, nil
// 		}
// 	} else {
// 		suite, err := l.findLazyTestSuite(dir)
// 		if err != nil {
// 			return nil, fmt.Errorf("error finding lazy test suite: %w", err)
// 		}
//
// 		if suite != nil {
// 			node = tview.NewTreeNode(fmt.Sprintf("[bisque]%s %s", suite.Icon, f.Name()))
// 			node.SetReference(suite)
// 			node.SetSelectable(true)
//
// 			for _, t := range suite.Tests {
// 				test := tview.NewTreeNode(fmt.Sprintf("[darkturquoise] %s", t.Name))
// 				test.SetSelectable(true)
// 				test.SetReference(t)
// 				node.AddChild(test)
// 			}
// 			return node, nil
// 		}
// 	}
//
// 	return nil, nil
// }
//
// // findLazyTestSuite will use the engines to find a test suite in the given path and file
// func (l *LazyTestLoader) findLazyTestSuite(path string) (*models.LazyTestSuite, error) {
// 	for _, engine := range l.Engines {
// 		suite, err := engine.ParseTestSuite(path)
// 		if err != nil {
// 			return nil, fmt.Errorf("error parsing test suite: %w", err)
// 		}
// 		if suite != nil {
// 			return suite, nil
// 		}
// 	}
// 	return nil, nil
// }
//
// func loadFiles(dir string) ([]fs.FileInfo, error) {
// 	file, err := os.Open(filepath.Clean(dir))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
//
// 	fileInfos, err := file.Readdir(-1)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading directory: %w", err)
// 	}
// 	return fileInfos, err
// }
