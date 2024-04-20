package pytest

import "testing"

type mockRunner struct {
	runHandler func(cmd string) (string, error)
}

func (m *mockRunner) RunCmd(cmd string) (string, error) {
	return m.runHandler(cmd)
}

func TestPytestEngine_GetIcon(t *testing.T) {
	p := NewPytestEngine(nil)
	icon := p.GetIcon()
	if icon != "" {
		t.Errorf("expected icont to be '', but got %s", icon)
	}
}
