package safe

import (
	"errors"
)

func ConvertType[T any](original interface{}) (*T, error) {
	if original == nil {
		return nil, errors.New("fail to assert type to generic type(nil value)")
	}

	t, ok := original.(T)

	if !ok {
		return nil, errors.New("fail to assert type to generic type(not same type)")
	}

	var tmp *T
	if &t == tmp {
		return nil, errors.New("fail to assert type to generic type(the converted result is null)")
	}

	return &t, nil
}

func MustConvertType[T any](original interface{}) *T {
	converted, err := ConvertType[T](original)
	if err != nil {
		panic(err)
	}

	return converted
}
