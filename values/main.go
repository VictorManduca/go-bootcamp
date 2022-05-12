package main

import "fmt"

func main() {
	var a int = 1
	var b float32 = 3.51

	a = int(b)
	fmt.Println(a, b)

	var x float64
	fmt.Println(x == 0.0)
}
