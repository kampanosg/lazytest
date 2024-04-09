package bashunit

import (
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestBashEngine_GetIcon(t *testing.T) {
	b := NewBashunitEngine(nil)
	icon := b.GetIcon()
	if icon != "󱆃" {
		t.Errorf("Expected icon to be '󱆃', but got %s", icon)
	}
}

func TestBashEngine_Load(t *testing.T) {
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
			name: "no files",
			fields: func() fields {
				appFS := afero.NewMemMapFs()
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
				afero.WriteFile(appFS, "src/test.sh", []byte("#!/bin/bash\necho 'hello'"), 0644)
				return fields{
					fs: appFS,
				}
			},
			args: args{
				dir: "src",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
		{
			name: "with tests",
			fields: func() fields {
				appFS := afero.NewMemMapFs()
				appFS.MkdirAll("src", 0755)
				afero.WriteFile(appFS, "src/test.sh", []byte("#!/bin/bash \n\n function test_echo() { \n echo 1 \n }"), 0644)
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
							Name: "test.sh",
							Children: []*models.LazyNode{
								{
									Name: "test",
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
			f := tt.fields()
			b := NewBashunitEngine(f.fs)
			got, err := b.Load(tt.args.dir)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			if tt.wantNil {
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, len(tt.want.Root.Children), len(got.Root.Children))
		})
	}
}
