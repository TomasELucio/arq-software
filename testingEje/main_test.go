package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {

	result := suma(2, 3)
	assert.Equal(t, result, 5)
}

func TestMayor(t *testing.T) {
	result := mayor(2, 3)
	assert.Equal(t, result, 3)

	result = mayor(3, 2)
	assert.Equal(t, result, 3)
}
