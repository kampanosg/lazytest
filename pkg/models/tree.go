package models

type LazyTree struct {
	Root *LazyNode
}

type LazyNode struct {
	Name     string
	Ref      any
	Children []*LazyNode
}

func NewLazyTree(root *LazyNode) *LazyTree {
	return &LazyTree{
		Root: root,
	}
}

func NewLazyNode(name string, ref any) *LazyNode {
	return &LazyNode{
		Name: name,
		Ref:  ref,
	}
}

func (n *LazyNode) AddChild(node *LazyNode) {
	n.Children = append(n.Children, node)
}

func (n *LazyNode) FindChild(name string) *LazyNode {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}

	return nil
}

func (n *LazyNode) SetReference(ref any) {
	n.Ref = ref
}

func (n *LazyNode) IsTest() bool {
	if n.Ref == nil {
		return false
	}

	_, ok := n.Ref.(*LazyTest)
	return ok
}

func (n *LazyNode) IsTestSuite() bool {
	if n.Ref == nil {
		return false
	}

	_, ok := n.Ref.(*LazyTestSuite)
	return ok
}

func (n *LazyNode) IsDir() bool {
	return !n.IsTest() && !n.IsTestSuite()
}
