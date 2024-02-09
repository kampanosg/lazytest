package engines

import (
	"io/fs"

	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyTestEngine interface {
	Load(dir string, f fs.FileInfo) (*models.LazyTestSuite, error)
}
