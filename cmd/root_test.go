/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	root := RootCmd()
	err := Execute(root)
	assert.Error(t, err)

	root1 := RootCmd()
	root1.SetArgs([]string{"1"})
	err = Execute(root1)
	assert.NoError(t, err)
}
