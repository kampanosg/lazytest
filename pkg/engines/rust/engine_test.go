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
Finished test [unoptimized + debuginfo] target(s) in 0.66s
     Running unittests src/lib.rs (target/debug/deps/wallet_api-93c222fb9d897c74)
api::account_handlers::tests::test_post_accounts_existing_account: test
api::account_handlers::tests::test_post_accounts_no_account: test
api::account_handlers::tests::test_post_accounts_no_account_with_generate_ephemeral_attr: test
api::account_handlers::tests::test_post_accounts_no_account_without_generate_ephemeral_attr: test
api::error::tests::test_error_handler_bad_request: test
api::error::tests::test_error_handler_bad_request_validation: test
api::error::tests::test_error_handler_unauthorised: test
api::item_handlers::tests::test_delete_item: test
api::item_handlers::tests::test_get_items: test
api::item_handlers::tests::test_identity_failure: test
api::item_handlers::tests::test_patch_item: test
api::item_handlers::tests::test_post_initial_item: test
models::wallet::tests::test_card_brand_from_str: test
     Running unittests src/main.rs (target/debug/deps/wallet_api-67ae81328fc5acea)
   Doc-tests wallet-api`, nil
						},
					},
				}
			},
			args: args{
				dir: ".",
			},
			wantErr: false,
			wantNil: false,
			want: &models.LazyTree{
				Root: &models.LazyNode{
					Name: "wallet_api",
					Children: []*models.LazyNode{
						{
							Name: "api",
						},
						{
							Name: "models",
						},
					},
				},
			},
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
