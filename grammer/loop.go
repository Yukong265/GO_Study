package main

import "fmt"

// golang don't have while
func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// for n<100{
	// role of while
	// }

	for {
		// infinite loop
	}
	names := []string("안녕", "티비", "어쩔티비")
	for index, name := range names { // can use range like python
		fmt.Println(index, name)
	}

	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	fmt.Println("OK")
}
