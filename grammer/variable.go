package main

import "fmt"

// variable
var a int = 0
var b float32 = 1
var i, j, k int = 1, 2, 3
var d = 123

//constant
const hi = "hi"

const (
	Sky      = "blue"
	Rose     = "red"
	Jinseong = "idiot"
)

func main() {
	fmt.Println("var :", a, b, i, j, k, d)
	fmt.Println("const :", hi, Sky, Rose, Jinseong)
}
