package pytest

import "github.com/kampanosg/lazytest/pkg/models"

const icon = ""

type PytestEngine struct {
}

func (p *PytestEngine) NewPytestEngine() *PytestEngine {
	return &PytestEngine{}
}

func (p *PytestEngine) Icon() string {
	return icon
}

func (p *PytestEngine) Load(dir string) (*models.LazyTree, error) {
	return nil, nil
}

// pytest --co
// pytest -x --verbose test_divide.py::test_sum
