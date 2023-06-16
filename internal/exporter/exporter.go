package exporter

import (
	"bufio"
	"io"
	"os"

	"github.com/chunguyenduc/jsontogo/internal/config"
)

type Exporter interface {
	Export(string) (int, error)
}

type structExport struct {
	Filename      string
	ReadWriter    io.ReadWriteCloser
	ReadWriteFunc func(string) (*os.File, error)
}

func NewExporter(conf *config.Config, readWriteFunc func(string) (*os.File, error)) Exporter {
	return &structExport{
		Filename:      conf.FileOutput,
		ReadWriteFunc: readWriteFunc,
	}
}

func (s *structExport) Export(content string) (int, error) {
	var (
		writer io.ReadWriteCloser
		err    error
	)
	if len(s.Filename) > 0 {
		writer, err = s.ReadWriteFunc(s.Filename)
		if err != nil {
			return 0, err
		}

	} else {
		writer = os.Stdout
		content = "\n" + content
	}

	defer writer.Close()
	f := bufio.NewWriter(writer)
	defer f.Flush()
	return f.Write([]byte(content))
}
