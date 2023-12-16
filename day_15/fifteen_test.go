package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	input := "HASH"
	output := Hash(input)
	assert.Equal(t, output, 52)
}
