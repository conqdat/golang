package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTriangle(t *testing.T) {
	assertions := assert.New(t)
	assertions.True(isTriangle(3, 4, 5))
}
func TestIsNotTriangle(t *testing.T) {
	assertions := assert.New(t)
	assertions.False(isTriangle(3, 4, 15))
}
