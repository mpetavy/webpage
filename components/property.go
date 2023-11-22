package components

import (
	"fmt"
	"strings"
)

type Property[T any] struct {
	data T
}

func (property *Property[T]) Set(data T) {
	property.data = data
}

func (property Property[T]) Get() T {
	return property.data
}

func (property Property[T]) HTML() string {
	return fmt.Sprintf("%v", property.data)
}

type StringProperty struct {
	Property[string]
}

func (stringProperty StringProperty) HTML() string {
	return stringProperty.data
}

type StringSliceProperty struct {
	Property[[]string]
}

func (stringSliceProperty StringSliceProperty) HTML() string {
	return strings.Join(stringSliceProperty.data, " ")
}
