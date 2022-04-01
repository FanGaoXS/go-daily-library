package main

import "reflect"

func Vo2Do(vo, do interface{}) {
	// vo的types和values
	voTypes := reflect.TypeOf(vo)
	voValues := reflect.ValueOf(vo)

	// do的types
	doTypes := reflect.TypeOf(do)
	if voTypes.NumField() > doTypes.NumField() {
		return
	}

}
