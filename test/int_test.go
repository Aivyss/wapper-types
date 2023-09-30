package test

import (
	"fmt"
	"testing"
	wrapperTypes "wrapper-types"
)

func TestTemp(t *testing.T) {
	a := wrapperTypes.Int(222)
	b := wrapperTypes.Int(-1)
	c := wrapperTypes.Int(222)

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a >= c)
	fmt.Println(a <= c)
	fmt.Println(a - c)
}
