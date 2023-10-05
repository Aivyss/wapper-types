package test

import (
	types "github.com/aivyss/wrapper-types"
	"github.com/aivyss/wrapper-types/safe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var value any = types.Bool(true)

		convertType, err := safe.ConvertType[types.Bool](value)
		assert.Nil(t, err)
		assert.True(t, convertType.Raw())
	})

	t.Run("fail", func(t *testing.T) {
		var value any = types.Bool(true)

		convertType, err := safe.ConvertType[bool](value)
		assert.NotNil(t, err)
		assert.Nil(t, convertType)
	})

	t.Run("nil test", func(t *testing.T) {
		convertType, err := safe.ConvertType[bool](nil)
		assert.NotNil(t, err)
		assert.Nil(t, convertType)
	})

	t.Run("success - Must", func(t *testing.T) {
		var value any = types.Bool(true)

		convertType := safe.MustConvertType[types.Bool](value)
		assert.True(t, convertType.Raw())
	})

	t.Run("fail - Must", func(t *testing.T) {
		defer func() {
			rec := recover()
			err, ok := rec.(error)

			assert.True(t, ok)
			assert.NotNil(t, err)
		}()

		var value any = types.Bool(true)

		_ = safe.MustConvertType[bool](value)
	})
}
