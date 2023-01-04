package main

import (
	"fmt"
	"reflect"
)

func main() {
	test_reflect()
}

type T1 struct {
	Name    string
	Content string
	Val     int
}

func test_reflect() {
	t1 := T1{}
	setObjValue(&t1, "Name", "ibas")
	fmt.Println(t1)
}

func setObjValue(obj interface{}, name string, value string) {
	ps := reflect.ValueOf(obj)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName(name)
		if f.IsValid() && f.CanSet() {
			f.SetString(value)
		}
	}
}
