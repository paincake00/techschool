package main

import (
	"fmt"
	"reflect"
)

func main() {
	defineType(1)
	defineType(true)
	defineType("some string")
	defineType(make(chan int))
	defineType(map[int]int{}) // но можно также проверить рефлексией
}

func defineType(v interface{}) {
	// получаем тип
	t := reflect.TypeOf(v)
	// Kind() находит категорию типа и проверяет принадлежность через switch-case
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("%v -> int\n", v)
		return
	case reflect.String:
		fmt.Printf("%v -> string\n", v)
		return
	case reflect.Bool:
		fmt.Printf("%v -> bool\n", v)
		return
	case reflect.Chan:
		fmt.Printf("%v -> channel of %s\n", v, t.Elem())
		return
	default:
		fmt.Printf("%v -> Undetected type\n", v)
	}
}
