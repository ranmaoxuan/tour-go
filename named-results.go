package main

import "fmt"

func split(sum int) (x int, y int, d int, n int) {
	x = sum * 4 / 9
	y = sum - x
	d = 999
	n = 123
	return
}
func main() {
	fmt.Println(split(17))
}
