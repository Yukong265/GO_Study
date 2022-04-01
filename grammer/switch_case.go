package main

import (
	"fmt"
)

const i int = 3

func main() {
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("nothing")
	}

	// switch reflect.TypeOf(i) {
	// case int:
	// 	println("int")
	// case bool:
	// 	println("bool")
	// case string:
	// 	println("string")
	// default:
	// 	println("unknown")
	// }
	// can check parameter's datatype lik this
}
