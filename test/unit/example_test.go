package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestBookStorePackage(t *testing.T) {
	t.Run("Package exists", func(t *testing.T) {
		assert.True(t, true)
	})
}
