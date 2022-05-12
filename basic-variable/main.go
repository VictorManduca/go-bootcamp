package main

import "fmt"

func main() {
	var age int = 23
	fmt.Println("My age is", age, "years")

	var name = "Victor"
	fmt.Println("My name is", name)

	var varWithNoUse = "I'm not using this variable"
	_ = varWithNoUse

	inferType := "I'm inferring the type"
	fmt.Println(inferType)

	car, cost := "Audi", 50_000
	fmt.Println(car, cost)

	car, year := "BMW", 2010
	fmt.Println(car, year)

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)
}
