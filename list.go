package types

import (
	"reflect"
)

type List[T any] []T
type MappableList[T, V any] List[T]

func (l *List[T]) Append(t T) List[T] {
	return append(*l, t)
}

func (l *List[T]) Get(i Int) T {
	return []T(*l)[i]
}

func (l *List[T]) ForEach(consumer func(t T)) {
	size := len(*l)

	for i := 0; i < size; i++ {
		consumer(l.Get(Int(i)))
	}
}

func (l *List[T]) Contains(elem T) Bool {
	flag := Bool(false)

	for i := 0; i < len(*l); i++ {
		value := l.Get(Int(i))
		valueType := reflect.TypeOf(value)
		elemType := reflect.TypeOf(elem)

		if Bool(valueType == elemType).And(Bool(reflect.DeepEqual(value, elem))) {
			flag = true
			break
		}
	}

	return flag
}

func Mappable[T, V any](l List[T]) MappableList[T, V] {
	return MappableList[T, V](l)
}

func (m MappableList[T, V]) Do(converter func(t T) V) List[V] {
	size := len(m)
	var result List[V]

	for i := 0; i < size; i++ {
		l := List[T](m)

		result = result.Append(converter(l.Get(Int(i))))
	}

	return result
}
