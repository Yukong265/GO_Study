package main

import "fmt"

const k int = 2
const max = 4

func main() {
	if k == 1 {
		fmt.Println("one")
	} else if k == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("nothing")
	}

	if val := k * 2; val >= max {
		fmt.Println(val)
	}
}
