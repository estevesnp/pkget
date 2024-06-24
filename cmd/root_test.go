package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyVersion(t *testing.T) {
	_, ok := verifyVersion("")
	assert.False(t, ok)

	pkg, ok := verifyVersion("latest")
	assert.Equal(t, "@latest", pkg)
	assert.True(t, ok)

	pkg, ok = verifyVersion("@latest")
	assert.Equal(t, "@latest", pkg)
	assert.True(t, ok)
}
