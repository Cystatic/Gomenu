package main

import (
	"fmt"
)

//求数组[1, 3, 5, 7, 8]所有元素的和
func Sum() {
	a := [5]int{1, 3, 5, 7, 8}
	sum := 0
	for _, i := range a {
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	who := "world!"
	fmt.Println("hello", who)
	Sum()
}
