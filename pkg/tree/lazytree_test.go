package tree

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
)

func TestHasTestSuites(t *testing.T) {
	tests := []struct {
		name string
		l    *LazyNode
		want bool
	}{
		{
			name: "no test suite in flat folder",
			l: &LazyNode{
				Name:     "root",
				IsFolder: true,
				Children: []*LazyNode{},
			},
			want: false,
		},
		{
			name: "no test suite in nested folder",
			l: &LazyNode{
				Name:     "root",
				IsFolder: true,
				Children: []*LazyNode{
					{
						Name:     "two",
						IsFolder: true,
					},
				},
			},
			want: false,
		},
		{
			name: "test suite in flat folder",
			l: &LazyNode{
				Name:     "root",
				IsFolder: true,
				Children: []*LazyNode{
					{
						Name:     "test_suite",
						IsFolder: false,
						Suite: &models.LazyTestSuite{
							Tests: []models.LazyTest{},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "test suite in nested folder",
			l: &LazyNode{
				Name:     "root",
				IsFolder: true,
				Children: []*LazyNode{
					{
						Name:     "two",
						IsFolder: true,
						Children: []*LazyNode{
							{
								Name:     "test_suite",
								IsFolder: false,
								Suite: &models.LazyTestSuite{
									Tests: []models.LazyTest{},
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := tt

			t.Parallel()

			if got := tc.l.HasTestSuite(); got != tc.want {
				t.Errorf("HasTestSuites() = %v, want %v", got, tc.want)
			}
		})
	}

}
