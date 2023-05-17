package cmd

import (
	"bufio"
	"os"
)

type StructExporter interface {
	Export(string) (int, error)
}

type structExport struct {
	Filename string
}

func NewStructExporter(filename string) StructExporter {
	return &structExport{
		Filename: filename,
	}
}

func (s *structExport) Export(content string) (int, error) {
	var (
		writer *os.File
		err    error
	)
	if len(s.Filename) > 0 {
		writer, err = os.Create(s.Filename)
		if err != nil {
			return 0, err
		}
	} else {
		writer = os.Stdout
	}
	defer writer.Close()

	f := bufio.NewWriter(writer)
	defer f.Flush()
	return f.Write([]byte(content))
}
