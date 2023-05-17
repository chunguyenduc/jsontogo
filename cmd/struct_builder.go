package cmd

import (
	"bufio"
	"os"
)

type StructBuilder interface {
	Run() error
}

type structBuilder struct {
	Input      []byte
	Output     string
	FileOutput string
	StructName string
}

func NewStructBuilder(input []byte, fileOutput string, name string) StructBuilder {
	return &structBuilder{
		Input:      input,
		FileOutput: fileOutput,
		StructName: name,
	}
}

func (s *structBuilder) Run() error {
	var err error

	s.Output, err = jsonToGo(s.Input, s.StructName)
	if err != nil {
		return err
	}

	return s.Export()
}

func (s *structBuilder) Export() error {
	var (
		writer *os.File
		err    error
	)
	if len(s.FileOutput) > 0 {
		writer, err = os.Create(s.FileOutput)
		if err != nil {
			return err
		}
	} else {
		writer = os.Stdout
	}
	defer writer.Close()

	f := bufio.NewWriter(writer)
	defer f.Flush()
	_, err = f.Write([]byte(s.Output))
	return err
}
