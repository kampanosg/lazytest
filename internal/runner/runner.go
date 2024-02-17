package runner

import (
	"fmt"
	"os/exec"

	"github.com/kampanosg/lazytest/pkg/models"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(cmd string) (*models.LazyTestResult, error) {
	c := exec.Command("sh", "-c", cmd)
	out, err := c.Output()

	if err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}

	res := &models.LazyTestResult{
		IsSuccess: true,
		Output:    string(out),
	}

	return res, nil
}
