package test

import (
	"github.com/aivyss/wrapper-types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("Append + Get", func(t *testing.T) {
		list := types.List[types.String]{"a", "b", "c"}
		list = list.Append("d")

		assert.Equal(t, list, types.List[types.String]{"a", "b", "c", "d"})
		assert.Equal(t, list.Get(3).Raw(), "d")
	})

	t.Run("ForEach", func(t *testing.T) {
		list := types.List[types.String]{"a", "b", "c"}
		replica := types.List[types.String]{}

		list.ForEach(func(elem types.String) {
			replica = replica.Append(elem)
		})

		assert.Equal(t, list, replica)
	})

	t.Run("Contains", func(t *testing.T) {
		list := types.List[types.String]{"a", "b", "c"}

		assert.True(t, list.Contains("b").Raw())
		assert.False(t, list.Contains("d").Raw())
	})

	t.Run("Map", func(t *testing.T) {
		list := types.List[types.String]{"1", "2", "3"}
		convertedList := types.Mappable[types.String, types.Int](list).Do(
			func(elem types.String) types.Int {
				result, _ := elem.ToInt()
				return result
			},
		)

		assert.True(t, convertedList.Contains(1).Raw())
		assert.True(t, convertedList.Contains(2).Raw())
		assert.True(t, convertedList.Contains(3).Raw())
	})
}
