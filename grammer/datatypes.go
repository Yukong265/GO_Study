package main

import "fmt"

// use `` and "" in string
const rawLiteral = `raw literal`
const interLiteral = "inter literal"

func main() {
	fmt.Println(rawLiteral, interLiteral)
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(i, u, f)

	str := "ABCD"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(str, bytes, str2)
} // have to specify data type, if you don't specify type, program rasie runtime error, not compile error
