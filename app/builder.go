package app

type StructBuilder interface {
	Build([]byte) (string, error)
}

type structBuilder struct {
	StructName string
}

func NewStructBuilder(conf *Config) StructBuilder {
	return &structBuilder{
		StructName: conf.StructName,
	}
}

func (s *structBuilder) Build(input []byte) (string, error) {
	return BuildJSONToGo(input, s.StructName)
}
