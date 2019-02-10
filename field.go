package fields

import (
	"fmt"
	"reflect"
)

type Field struct {
	CustomName string
	Depth      int
	ParentTree []string
}

func (field *Field) SetCustomName(name string) {
	field.CustomName = name
}

func (field *Field) SetParent(parentResolver interface{}) {
	parentField := FromResolver(parentResolver)

	field.Depth = parentField.Depth + 1

	field.ParentTree = append(
		parentField.ParentTree,
		NameFromResolver(parentResolver))
}

func FromResolver(resolver interface{}) Field {
	fieldValue := reflect.ValueOf(resolver).Elem().FieldByName("Field")

	return fieldValue.Interface().(Field)
}

func NameFromResolver(resolver interface{}) string {
	field := FromResolver(resolver)

	var fieldName string

	if field.CustomName != "" {
		fieldName = field.CustomName
	} else {
		fieldType, _ := reflect.TypeOf(
			resolver).Elem().FieldByName("Field")

		fieldName = fmt.Sprintf("%s", fieldType.Tag)
	}

	return fieldName
}
