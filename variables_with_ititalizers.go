package main

import "fmt"

var i, j int = 1, 2

func main() {
	//short variation declarations can be made inside of functions using := to replace var at the beginning
	var c, python, java = true, false, "no!"
	// c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// export
// 1 2 true false "no!"
