package runner

import (
	"os/exec"
	"time"

	"github.com/kampanosg/lazytest/pkg/models"
)

type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(cmd string) *models.LazyTestResult {
	now := time.Now()
	c := exec.Command("sh", "-c", cmd)
	out, err := c.Output()

	return &models.LazyTestResult{
		Passed:   err == nil,
		Output:   string(out),
		Duration: time.Since(now),
	}
}
