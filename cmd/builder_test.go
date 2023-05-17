package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStructBuilder(t *testing.T) {
	got := NewStructBuilder([]byte{}, "", "")
	assert.NotNil(t, got)
}

func Test_structBuilder_Build(t *testing.T) {
	got := NewStructBuilder([]byte{}, "", "")
	res, err := got.Build()
	assert.Empty(t, res)
	assert.NotNil(t, err)
}
