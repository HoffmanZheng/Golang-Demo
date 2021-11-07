package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Go web program"
	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)
	fmt.Printf("type: %T\n", t)
	fmt.Printf("value: %v\n", v)
	fmt.Println("settable: ", v.CanSet())

	// to enable the modifiability of reflected object
	pointer := reflect.ValueOf(&name)
	value := pointer.Elem()
	fmt.Println("settable of value: ", value.CanSet())
	value.SetString("new String")
	fmt.Println("value after SetString: ", value)
}
