package engines

import (
	"io/fs"

	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyTestEngine interface {
	ParseTestSuite(dir string, f fs.FileInfo) (*models.LazyTestSuite, error)
}
