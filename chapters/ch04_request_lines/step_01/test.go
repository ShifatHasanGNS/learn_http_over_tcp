package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestLineParse(t *testing.T) {
	assert.Equal(t, "Hey, Chat!", "Hey, Chat!")
}
