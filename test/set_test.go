package test

import (
	types "github.com/aivyss/wrapper-types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("Set + Contains", func(t *testing.T) {
		set := types.Set[types.String]{
			"key2": false,
		}

		set.Set("key1")
		assert.True(t, set.Contains("key1").Raw())
		assert.False(t, set.Contains("key2").Raw())
	})
}
