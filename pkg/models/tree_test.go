package models

import "testing"

func TestAddChild(t *testing.T) {
	tests := []struct {
		name string
		node *LazyNode
		want int
	}{
		{
			name: "success",
			node: &LazyNode{
				Name: "root",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child := &LazyNode{
				Name: "child",
			}
			tt.node.AddChild(child)
			if len(tt.node.Children) != tt.want {
				t.Errorf("AddChild() got = %d, want %d", len(tt.node.Children), tt.want)
			}
		})
	}
}

func TestSetReference(t *testing.T) {
	tests := []struct {
		name string
		node *LazyNode
		want any
	}{
		{
			name: "success",
			node: &LazyNode{
				Name: "root",
			},
			want: "reference",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.node.SetReference(tt.want)
			if tt.node.Ref != tt.want {
				t.Errorf("SetReference() got = %v, want %v", tt.node.Ref, tt.want)
			}
		})
	}
}

func TestIsTest(t *testing.T) {
	tests := []struct {
		name string
		node *LazyNode
		want bool
	}{
		{
			name: "success",
			node: &LazyNode{
				Name: "root",
				Ref:  &LazyTest{},
			},
			want: true,
		},
		{
			name: "nil ref",
			node: &LazyNode{
				Name: "root",
			},
			want: false,
		},
		{
			name: "not a test",
			node: &LazyNode{
				Name: "root",
				Ref:  "not a test",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.IsTest(); got != tt.want {
				t.Errorf("IsTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTestSuite(t *testing.T) {
	tests := []struct {
		name string
		node *LazyNode
		want bool
	}{
		{
			name: "success",
			node: &LazyNode{
				Name: "root",
				Ref:  &LazyTestSuite{},
			},
			want: true,
		},
		{
			name: "nil ref",
			node: &LazyNode{
				Name: "root",
			},
			want: false,
		},
		{
			name: "not a test suite",
			node: &LazyNode{
				Name: "root",
				Ref:  "not a test suite",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.IsTestSuite(); got != tt.want {
				t.Errorf("IsTestSuite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	tests := []struct {
		name string
		node *LazyNode
		want bool
	}{
		{
			name: "success",
			node: &LazyNode{
				Name: "root",
			},
			want: true,
		},
		{
			name: "test",
			node: &LazyNode{
				Name: "root",
				Ref:  &LazyTest{},
			},
			want: false,
		},
		{
			name: "test suite",
			node: &LazyNode{
				Name: "root",
				Ref:  &LazyTestSuite{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.IsDir(); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
