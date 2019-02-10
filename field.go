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

func (field *Field) SetParent(parent_resolver interface{}) {
	parent_field := FromResolver(parent_resolver)

	field.Depth = parent_field.Depth + 1

	field.ParentTree = append(
		parent_field.ParentTree,
		NameFromResolver(parent_resolver))
}

func FromResolver(resolver interface{}) Field {
	field_value := reflect.ValueOf(resolver).Elem().FieldByName("Field")

	return field_value.Interface().(Field)
}

func NameFromResolver(resolver interface{}) string {
	field := FromResolver(resolver)

	var field_name string

	if field.CustomName != "" {
		field_name = field.CustomName
	} else {
		field_type, _ := reflect.TypeOf(
			resolver).Elem().FieldByName("Field")

		field_name = fmt.Sprintf("%s", field_type.Tag)
	}

	return field_name
}
