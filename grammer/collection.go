package main

import "fmt"

// Array
var a [3]int
var a1 = [3]int{1, 2, 3}
var a2 = [...]int{4, 5, 6}

var a3 = [2][3]int{
	{1, 2, 3},
	{4, 5, 6},
}

// Slice
func slice() {
	var s []int
	var a = [3]int{1, 2, 3}

	if s == nil {
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(s), cap(s))

	i := make([]int, 5, 10)

	j := []int{0, 1, 2, 3, 4, 5}
	// j = s[2:5] 2,3,4
	// j = j[1:] 3,4

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	// {1,2,3,4,5,6}
	fmt.Println(a, i, j)
	fmt.Println(sliceA) // [1,2,3,4,5,6]
}

source := []int{0,1,2}
target := make([]int, len(source), cap(source)*2)
copy(target, source)

// Mapping
var idMap map[int]string
idMap1 = make(map[int]string)

func main() {
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	val, exist := tickers["MSFT"]
	fmt.Println(val, exist)
	if !exist {
		fmt.Println("No MSFT ticker")
	}
}
