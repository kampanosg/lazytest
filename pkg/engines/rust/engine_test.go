package rust

import (
	"errors"
	"testing"

	"github.com/kampanosg/lazytest/pkg/models"
	"github.com/stretchr/testify/assert"
)

type mockRunner struct {
	runHandler func(cmd string) (string, error)
}

func (m *mockRunner) RunCmd(cmd string) (string, error) {
	return m.runHandler(cmd)
}

func TestRustEngine_GetIcon(t *testing.T) {
	g := NewRustEngine(nil)
	icon := g.GetIcon()
	if icon != "" {
		t.Errorf("Expected icon to be '', but got %s", icon)
	}
}

func TestRustEngine_Load(t *testing.T) {
	type fields struct {
		runner *mockRunner
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
			name: "runner returns error",
			fields: func() fields {
				return fields{
					runner: &mockRunner{
						runHandler: func(cmd string) (string, error) {
							return "", errors.New("an error")
						},
					},
				}
			},
			args: args{
				dir: ".",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
		{
			name: "no tests in the project",
			fields: func() fields {
				return fields{
					runner: &mockRunner{
						runHandler: func(cmd string) (string, error) {
							return "no tests to parse", nil
						},
					},
				}
			},
			args: args{
				dir: ".",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
		{
			name: "parse tests",
			fields: func() fields {
				return fields{
					runner: &mockRunner{
						runHandler: func(cmd string) (string, error) {
							return `
							
							`, nil
						},
					},
				}
			},
			args: args{
				dir: ".",
			},
			wantErr: false,
			wantNil: true,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			g := NewRustEngine(fields.runner)

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

			assert.Equal(t, len(tt.want.Root.Children), len(got.Root.Children))
		})
	}

}
