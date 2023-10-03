package test

import (
	types "github.com/aivyss/wrapper-types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Raw", func(t *testing.T) {
		m := types.Map[types.String, types.String]{
			"key1": "value1",
			"key2": "value2",
		}

		expectedRaw := map[types.String]types.String{
			"key1": "value1",
			"key2": "value2",
		}

		raw := m.Raw()

		assert.Equal(t, expectedRaw, raw)
	})

	t.Run("Set + Get", func(t *testing.T) {
		m := types.Map[types.String, types.String]{
			"key1": "value1",
			"key2": "value2",
		}

		m["key3"] = "value3"
		m.Set("key4", "value4")

		assert.Equal(t, "value4", m.Get("key4").Raw())
		assert.Equal(t, "value4", m["key4"].Raw())
		assert.Equal(t, "value3", m.Get("key3").Raw())
		assert.Equal(t, "value3", m["key3"].Raw())
	})

	t.Run("Keys + Values => Entries", func(t *testing.T) {
		m := types.Map[types.String, types.String]{
			"key1": "value1",
			"key2": "value2",
		}

		assert.Equal(t, types.List[types.String]{"key1", "key2"}, m.Keys())
		assert.Equal(t, types.List[types.String]{"value1", "value2"}, m.Values())
	})
}
