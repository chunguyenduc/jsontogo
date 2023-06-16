package builder

import (
	"github.com/chunguyenduc/jsontogo/internal/config"
	"github.com/chunguyenduc/jsontogo/internal/handler"
)

type Builder interface {
	Build([]byte) (string, error)
}

type structBuilder struct {
	StructName string
}

func NewBuilder(conf *config.Config) Builder {
	return &structBuilder{
		StructName: conf.StructName,
	}
}

func (s *structBuilder) Build(input []byte) (string, error) {
	return handler.BuildJSONToGo(input, s.StructName)
}
