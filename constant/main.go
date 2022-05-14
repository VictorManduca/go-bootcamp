package main

import "fmt"

func main() {
	const daysInWeek = 7
	const days = 7

	const (
		a = iota
		b
		c
	)

	fmt.Println(daysInWeek)
	fmt.Println(a, b, c)
}
