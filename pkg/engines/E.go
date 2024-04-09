package engines

import (
	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyEngine interface {
	Load(dir string) (*models.LazyTree, error)
	GetIcon() string
}
