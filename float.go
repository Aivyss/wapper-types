package types

import (
	"fmt"
	"math"
)

type Float64 float64

type Float32 float32

func (f *Float64) Raw() float64 {
	return float64(*f)
}

func (f *Float64) Floor(x Float64) Float64 {
	return Float64(math.Floor(x.Raw()))
}

func (f *Float64) Ceil(x Float64) Float64 {
	return Float64(math.Ceil(x.Raw()))
}

func (f *Float64) Round(x Float64) Float64 {
	return Float64(math.Round(x.Raw()))
}

func (f *Float64) Pow10(n Int) Float64 {
	return Float64(math.Pow10(n.Raw()))
}

func (f *Float64) Pow(n Float64) Float64 {
	return Float64(math.Pow(f.Raw(), n.Raw()))
}

func (f *Float64) Abs() Float64 {
	if *f < 0 {
		return -*f
	}

	return *f
}

func (f *Float64) ToStr() String {
	return String(fmt.Sprintf("%v", f.Raw()))
}

func (f *Float64) ToInt() Int {
	return Int(*f)
}

func (f *Float32) Raw() float64 {
	return float64(*f)
}

func (f *Float32) Floor(x Float32) Float32 {
	return Float32(math.Floor(x.Raw()))
}

func (f *Float32) Ceil(x Float32) Float32 {
	return Float32(math.Ceil(x.Raw()))
}

func (f *Float32) Round(x Float32) Float32 {
	return Float32(math.Round(x.Raw()))
}

func (f *Float32) Pow10(n Int) Float32 {
	return Float32(math.Pow10(n.Raw()))
}

func (f *Float32) Pow(n Float32) Float32 {
	return Float32(math.Pow(f.Raw(), n.Raw()))
}

func (f *Float32) Abs() Float32 {
	if *f < 0 {
		return -*f
	}

	return *f
}

func (f *Float32) ToStr() String {
	return String(fmt.Sprintf("%v", f.Raw()))
}

func (f *Float32) ToInt() Int {
	return Int(*f)
}
