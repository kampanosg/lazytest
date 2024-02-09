package loader

import (
	"io/fs"

	"github.com/kampanosg/lazytest/pkg/engines"
	"github.com/kampanosg/lazytest/pkg/models"
)

type LazyTestLoader struct {
	Engines []engines.LazyTestEngine
}

func NewLazyTestLoader(e []engines.LazyTestEngine) *LazyTestLoader {
	return &LazyTestLoader{
		Engines: e,
	}
}

func (l *LazyTestLoader) LoadTestSuite(path string, f fs.FileInfo) (*models.LazyTestSuite, error) {
	for _, engine := range l.Engines {
		suite, err := engine.LoadTestSuite(path, f)
		if err != nil {
			return nil, err
		}
		if suite != nil {
			return suite, nil
		}
	}
	return nil, nil
}
