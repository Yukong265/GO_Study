package main

import "fmt"

var p *int // this can change * to ^
var k int = 10

var a = &k

func main() {
	fmt.Println(p, k, a)
}
