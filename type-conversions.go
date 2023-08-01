package main

import (
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var f int = int(x*x + y*y)
	var z int = int(f)
	fmt.Println(x, y, z)
}
