package test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	wrapperTypes "wrapper-types"
)

func TestString(t *testing.T) {
	t.Run("Length", func(t *testing.T) {
		s := wrapperTypes.String("test")
		assert.Equal(t, s.Length(), wrapperTypes.Int(len(s.Raw())))
	})

	t.Run("Bool", func(t *testing.T) {
		ss := []wrapperTypes.String{
			wrapperTypes.String("true"),
			wrapperTypes.String("True"),
			wrapperTypes.String("1"),
			wrapperTypes.String("T"),
			wrapperTypes.String("t"),
			wrapperTypes.String("aa"),
			wrapperTypes.String("false"),
		}

		expectedResults := []wrapperTypes.Bool{
			true, true, true, true, true, false, false,
		}

		for i, s := range ss {
			assert.Equal(t, expectedResults[i], s.Bool())
		}
	})

	t.Run("Raw", func(t *testing.T) {
		uid := uuid.New()
		w := wrapperTypes.String(uid.String())

		assert.Equal(t, uid.String(), w.Raw())
	})

	t.Run("Split", func(t *testing.T) {
		s := "ab-fef-asdf-gawef-asdfasdf"
		ss := strings.Split(s, "-")
		w := wrapperTypes.String(s)
		ws := w.Split("-")

		assert.Equal(t, len(ss), len(ws))
		for i := 0; i < len(ss); i++ {
			assert.Equal(t, ss[i], ws[i].Raw())
		}
	})

	t.Run("Concat", func(t *testing.T) {
		s := wrapperTypes.String("")
		texts := []wrapperTypes.String{
			"a", "b", "c", "d", "e",
		}

		result := s.Concat("-", texts...)
		assert.Equal(t, "-a-b-c-d-e", result.Raw())
	})

	t.Run("ReplaceAll", func(t *testing.T) {
		s := wrapperTypes.String("ab\tcd\nef\rgh")
		replaceAll := s.ReplaceAll("\t", " ").ReplaceAll("\n", " ").ReplaceAll("\r", " ")
		assert.Equal(t, "ab cd ef gh", replaceAll.Raw())

		s = "ab{{devide}}cd{{devide}}ef{{devide}}gh"
		assert.Equal(t, s, s.Replace("{{devide}}", " ", 0))
		assert.Equal(t, "ab cd{{devide}}ef{{devide}}gh", s.Replace("{{devide}}", " ", 1).Raw())
		assert.Equal(t, "ab cd ef{{devide}}gh", s.Replace("{{devide}}", " ", 2).Raw())
		assert.Equal(t, "ab cd ef gh", s.Replace("{{devide}}", " ", 3).Raw())
		assert.Equal(t, "ab cd ef gh", s.ReplaceAll("{{devide}}", " ").Raw())
	})

	t.Run("TrimSpace", func(t *testing.T) {
		s := wrapperTypes.String("abcdef")
		w := " " + s + "   "

		assert.Equal(t, s, w.TrimSpace())
	})

	t.Run("ToInt", func(t *testing.T) {
		t.Run("case: success", func(t *testing.T) {
			num := 123
			numStr := wrapperTypes.String(fmt.Sprintf("%d", num))

			converted, err := numStr.ToInt()
			doPanicIfNotNil(err)

			assert.Equal(t, num, converted.Raw())
		})

		t.Run("case: Fail", func(t *testing.T) {
			numStr := wrapperTypes.String("not number")

			_, err := numStr.ToInt()

			assert.NotNil(t, err)
		})
	})

	t.Run("IsBlank", func(t *testing.T) {
		s1 := wrapperTypes.String(" \r abcdef \r\n fjhgjhk \n fasef ")
		s2 := wrapperTypes.String(" \r  \r\n  \n  ")

		assert.False(t, s1.IsBlank().Raw())
		assert.True(t, s2.IsBlank().Raw())
	})

}

func doPanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
