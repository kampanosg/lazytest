package bashunit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kampanosg/lazytest/pkg/models"
)

const (
	suiteType = "bashunit"
	icon      = "󱆃"
)

type BashunitEngine struct {
}

func NewBashunitEngine() *BashunitEngine {
	return &BashunitEngine{}
}

func (g *BashunitEngine) ParseTestSuite(fp string) (*models.LazyTestSuite, error) {
	if !strings.HasSuffix(fp, ".sh") {
		return nil, nil
	}

	suite := &models.LazyTestSuite{
		Path: fp,
		Type: suiteType,
		Icon: icon,
	}

	err := filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		tests, err := extractTests(path)
		if err != nil {
			return err
		}

		suite.Tests = append(suite.Tests, tests...)
		return nil
	})

	return suite, err
}

func extractTests(f string) ([]*models.LazyTest, error) {
	file, err := os.Open(filepath.Clean(f))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tests []*models.LazyTest
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "function test_") {
			name := strings.Fields(line)[1]
			name = strings.TrimSuffix(name, "()")
			test := &models.LazyTest{
				Name:   name,
				RunCmd: fmt.Sprintf("bashunit -v -S -f \"%s\" %s", name, f),
			}
			tests = append(tests, test)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}
