package builder

import (
	"testing"

	"github.com/chunguyenduc/jsontogo/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	got := NewBuilder(&config.Config{})
	assert.NotNil(t, got)
}

func Test_Builder_Build(t *testing.T) {
	got := NewBuilder(&config.Config{})
	res, err := got.Build([]byte{})
	assert.Empty(t, res)
	assert.NotNil(t, err)
}
