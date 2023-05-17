package cmd

type StructBuilder interface {
	Build() (string, error)
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

func (s *structBuilder) Build() (string, error) {
	return BuildJSONToGo(s.Input, s.StructName)
}
