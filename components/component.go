package components

import (
	"github.com/mpetavy/common"
	"reflect"
)

const (
	ID    = "id"
	NAME  = "name"
	CLASS = "class"
	STYLE = "style"
)

type Component struct {
	Id    StringProperty
	Name  StringProperty
	Class StringSliceProperty
	Style StringSliceProperty
}

func (component Component) updateElement(element *common.Element) {
	valComponent := reflect.ValueOf(component)

	for i := 0; i < valComponent.NumField(); i++ {
		valField := reflect.ValueOf(valComponent.Field(i))

		valHTML := valField.MethodByName("HTML")
		results := valHTML.Call(nil)[0]

		common.Info("Field %d: %s", i, results.String())
	}
}

func (component Component) HTML() (string, error) {
	return "", nil
}
