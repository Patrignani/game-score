package helpers

import "reflect"

func GetFields[S any](slice []S, nameField string) {
	val := reflect.Indirect(reflect.ValueOf(slice))
	structValue, exist := val.Type().FieldByName(nameField)

	if exist {
		println(structValue.Name)
	}
}
