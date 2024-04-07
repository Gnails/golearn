package my

import (
	"fmt"
	"reflect"
)

func M() {
	a := uint(0)
	b := uint(1)
	c := a - b
	fmt.Println(size(c))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
}
