package engines

import (
	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyEngine interface {
	ParseTestSuite(fp string) (*models.LazyTestSuite, error)
}
