package utils

import (
	"encoding/json"
)

type Option[T any] struct {
	value *T
}

func (o *Option[T]) UnmarshalJSON(bytes []byte) error {
	o.value = new(T)
	return json.Unmarshal(bytes, o.value)
}

func (o *Option[T]) Value() (out T, ok bool) {
	if o.value == nil {
		return
	}
	return *o.value, true
}
