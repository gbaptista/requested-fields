package fields

import (
	"fmt"
	"reflect"
)

// Field type to be used in Resolvers
type Field struct {
	CustomName string
	Depth      int
	ParentTree []string
}

// SetCustomName define a custom nome for Field.
func (field *Field) SetCustomName(name string) {
	field.CustomName = name
}

// SetParent define the parent resolver of a Feild.
func (field *Field) SetParent(parentResolver interface{}) {
	parentField := fromResolver(parentResolver)

	field.Depth = parentField.Depth + 1

	field.ParentTree = append(
		parentField.ParentTree,
		nameFromResolver(parentResolver))
}

func fromResolver(resolver interface{}) Field {
	fieldValue := reflect.ValueOf(resolver).Elem().FieldByName("Field")

	return fieldValue.Interface().(Field)
}

func nameFromResolver(resolver interface{}) string {
	field := fromResolver(resolver)

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
