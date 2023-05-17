package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStructBuilder(t *testing.T) {
	got := NewStructBuilder(&Config{})
	assert.NotNil(t, got)
}

func Test_structBuilder_Build(t *testing.T) {
	got := NewStructBuilder(&Config{})
	res, err := got.Build([]byte{})
	assert.Empty(t, res)
	assert.NotNil(t, err)
}
