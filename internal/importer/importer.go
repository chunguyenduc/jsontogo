package importer

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/chunguyenduc/jsontogo/internal/config"
)

type Importer interface {
	Import() ([]byte, error)
}

type structImport struct {
	Filename   string
	Input      []byte
	Reader     io.ReadCloser
	ReaderFunc func(string) (*os.File, error)
}

func NewImporter(conf *config.Config, readerFunc func(string) (*os.File, error)) Importer {
	return &structImport{
		Input:      conf.Input,
		Filename:   conf.FileInput,
		ReaderFunc: readerFunc,
	}
}

func (s *structImport) Import() ([]byte, error) {
	if len(s.Filename) > 0 {
		if ext := filepath.Ext(s.Filename); ext != ".json" {
			return nil, fmt.Errorf("open %v: not JSON file", s.Filename)
		}
		jsonFile, err := s.ReaderFunc(s.Filename)
		if err != nil {
			return nil, err
		}
		defer jsonFile.Close()

		return ioutil.ReadAll(jsonFile)
	}
	return s.Input, nil
}
