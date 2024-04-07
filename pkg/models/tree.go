package models

type LazyTree struct {
	Root *LazyNode
}

type LazyNode struct {
	Name     string
	Path     string
	Ref      any
	Children []*LazyNode
}

func NewLazyTree(root *LazyNode) *LazyTree {
	return &LazyTree{
		Root: root,
	}
}

func NewLazyNode(name, path string, ref any) *LazyNode {
	return &LazyNode{
		Name: name,
		Path: path,
		Ref:  ref,
	}
}

func (n *LazyNode) AddChild(node *LazyNode) {
	if n.Children == nil {
		n.Children = append(n.Children, node)
	}
}

func (n *LazyNode) SetReference(ref any) {
	if n.Ref == nil {
		n.Ref = ref
	}
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
