package app

import (
	"bufio"
	"io"
	"os"
)

type StructExporter interface {
	Export(string) (int, error)
}

type structExport struct {
	Filename      string
	ReadWriter    io.ReadWriteCloser
	ReadWriteFunc func(string) (*os.File, error)
}

func NewStructExporter(conf *Config, readWriteFunc func(string) (*os.File, error)) StructExporter {
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
