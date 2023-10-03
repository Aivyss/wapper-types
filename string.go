package types

import (
	"strconv"
	"strings"
)

type String string

var trueSlice = []String{"True", "true", "1", "T", "t"}

func (s String) Length() Int {
	return Int(len(string(s)))
}

func (s String) Bool() Bool {
	for _, value := range trueSlice {
		if s == value {
			return true
		}
	}

	return false
}

func (s String) Raw() string {
	return string(s)
}

func (s String) Split(sep String) List[String] {
	slice := strings.Split(s.Raw(), sep.Raw())
	var list List[String]

	for _, value := range slice {
		list = list.Append(String(value))
	}

	return list
}

func (s String) Concat(sep String, texts ...String) String {
	raw := s.Raw()
	for _, text := range texts {
		raw = raw + sep.Raw() + text.Raw()
	}

	return String(raw)
}

func (s String) ReplaceAll(old String, new String) String {
	return String(strings.ReplaceAll(s.Raw(), old.Raw(), new.Raw()))
}

func (s String) Replace(old String, new String, n Int) String {
	return String(strings.Replace(s.Raw(), old.Raw(), new.Raw(), n.Raw()))
}

func (s String) TrimSpace() String {
	return String(strings.TrimSpace(s.Raw()))
}

func (s String) ToInt() (Int, error) {
	num, err := strconv.Atoi(s.Raw())
	return Int(num), err
}

func (s String) ToInt64(base Int, bitSize Int) (Int64, error) {
	num, err := strconv.ParseInt(s.Raw(), base.Raw(), bitSize.Raw())
	return Int64(num), err
}

func (s String) ToFloat64(bitSize Int) (Float64, error) {
	num, err := strconv.ParseFloat(s.Raw(), bitSize.Raw())
	return Float64(num), err
}

func (s String) ToUint64(base Int, bitSize Int) (Uint64, error) {
	num, err := strconv.ParseUint(s.Raw(), base.Raw(), bitSize.Raw())
	return Uint64(num), err
}

func (s String) IsBlank() Bool {
	return s.TrimSpace().
		ReplaceAll(" ", "").
		ReplaceAll("\r", "").
		ReplaceAll("\n", "").
		Length() == 0
}

func (s String) AppendInt(num Int) String {
	return s + num.ToStr()
}

func (s String) AppendF64(num Float64) String {
	return s + num.ToStr()
}

func (s String) AppendF32(num Float32) String {
	return s + num.ToStr()
}
