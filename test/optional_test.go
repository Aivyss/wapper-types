package test

import (
	"github.com/aivyss/wrapper-types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptional(t *testing.T) {
	t.Run("NewOption + IsPresent", func(t *testing.T) {
		value1 := types.String("test value")
		opt1 := types.NewOptional(&value1)
		opt2 := types.NewOptional[types.String](nil)

		assert.True(t, opt1.IsPresent())
		assert.False(t, opt2.IsPresent())
	})

	t.Run("SetDefault", func(t *testing.T) {
		defaultStr := types.String("test")
		opt := types.NewOptional[types.String](nil).SetDefault(defaultStr)

		assert.True(t, opt.IsPresent())
		value, ok := opt.Get()
		assert.True(t, ok)
		assert.Equal(t, defaultStr, *value)
	})

	t.Run("IfPresent", func(t *testing.T) {
		value := types.String("test value")
		opt1 := types.NewOptional(&value)
		opt2 := types.NewOptional[types.String](nil)

		opt1.IfPresent(func(s *types.String) {
			if s == nil || !opt1.IsPresent() {
				panic("no case")
			}

			assert.Equal(t, value, *s)
		})

		opt2.IfPresent(func(s *types.String) {
			panic("no case")
		})
	})
}
