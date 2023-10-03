package types

type Optional[T any] struct {
	value *T
	def   *T
}

func NewOptional[T any](t *T) *Optional[T] {
	return &Optional[T]{value: t}
}

func (o *Optional[T]) Get() (value *T, ok bool) {
	if !o.IsPresent() {
		return nil, false
	}

	o.IfPresent(func(t *T) {
		value = t
	})

	return value, true
}

func (o *Optional[T]) IfPresent(consumer func(t *T)) {
	if o.value != nil {
		consumer(o.value)

		return
	}

	if o.def != nil {
		consumer(o.def)
	}
}

func (o *Optional[T]) IsPresent() bool {
	return o.value != nil || o.def != nil
}

func (o *Optional[T]) SetDefault(value T) *Optional[T] {
	o.def = &value

	return o
}
