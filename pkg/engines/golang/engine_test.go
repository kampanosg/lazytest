package golang

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_GetIcon(t *testing.T) {
	g := NewGoEngine(nil)
	icon := g.GetIcon()
	if icon != "󰟓" {
		t.Errorf("Expected icon to be '󰟓', but got %s", icon)
	}
}

func Test_Load(t *testing.T) {
	type fields struct {
		fs afero.Fs
	}

	type args struct {
		dir string
	}

	tests := []struct {
		name    string
		fields  func() fields
		args    args
		wantErr bool
		wantNil bool
		want    *models.LazyTree
	}{
		{
			name: "no go.mod file",
			fields: func() fields {
				appFS := afero.NewMemMapFs()
				appFS.MkdirAll("src", 0755)
				afero.WriteFile(appFS, "src/main.go", []byte("package main"), 0644)
				return fields{
					fs: appFS,
				}
			},
			args: args{
				dir: "/",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
		{
			name: "no tests",
			fields: func() fields {
				appFS := afero.NewMemMapFs()
				appFS.MkdirAll("src", 0755)
				afero.WriteFile(appFS, "src/go.mod", []byte("module test"), 0644)
				afero.WriteFile(appFS, "src/main.go", []byte("package main"), 0644)
				return fields{
					fs: appFS,
				}
			},
			args: args{
				dir: "/",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
		{
			name: "load tests",
			fields: func() fields {
				appFS := afero.NewMemMapFs()
				appFS.MkdirAll("src", 0755)
				afero.WriteFile(appFS, "src/go.mod", []byte("module test"), 0644)
				afero.WriteFile(appFS, "src/main.go", []byte("package main"), 0644)
				afero.WriteFile(appFS, "src/main_test.go", []byte("package main \n\n func TestMain(t *testing.T){t.skip()}"), 0644)
				return fields{
					fs: appFS,
				}
			},
			args: args{
				dir: "src",
			},
			wantErr: false,
			wantNil: false,
			want: &models.LazyTree{
				Root: &models.LazyNode{
					Name: "src",
					Children: []*models.LazyNode{
						{
							Name:     "main_test.go",
							Children: []*models.LazyNode{
								{
									Name: "TestMain",
									Children: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			g := NewGoEngine(fields.fs)

			got, err := g.Load(tt.args.dir)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			if tt.wantNil {
				assert.Nil(t, got)
				return
			}

			assert.NotNil(t, got)
			assert.Equal(t, len(tt.want.Root.Children), len(got.Root.Children))
		})
	}

}
