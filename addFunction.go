package main

import "fmt"

//you must state the type of variable
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
