package main

import "reflect"

// Fields get fieldsName of struct
func Fields(obj interface{}) []interface{} {
	var fieldsName []interface{}
	typeOf := reflect.TypeOf(obj)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fieldsName = append(fieldsName, field.Name)
	}
	return fieldsName
}

// Values get fieldsValue of struct
func Values(obj interface{}) []interface{} {
	var fieldsValue []interface{}
	valueOf := reflect.ValueOf(obj)
	for i := 0; i < valueOf.NumField(); i++ {
		value := valueOf.Field(i)
		fieldsValue = append(fieldsValue, value.Interface())
	}
	return fieldsValue
}

// ToMap get fields as key-value
func ToMap(obj interface{}) map[string]interface{} {
	fieldsMap := make(map[string]interface{})
	typeOf := reflect.TypeOf(obj)
	valueOf := reflect.ValueOf(obj)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		value := valueOf.Field(i)

		fieldsMap[field.Name] = value.Interface()
	}
	return fieldsMap
}
