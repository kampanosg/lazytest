package runner

import (
	"os/exec"

	"github.com/kampanosg/lazytest/pkg/models"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(cmd string) *models.LazyTestResult {
	c := exec.Command("sh", "-c", cmd)
	out, err := c.Output()

	return &models.LazyTestResult{
		IsSuccess: err == nil,
		Output:    string(out),
	}
}
