package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {

	//os.Getwd()
	const str, want = "43a", 42
	got, err := strconv.Atoi(str)
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, nil, err)
	assert.True(t, err == nil)

	assert.Equal(t, want, got)

}
