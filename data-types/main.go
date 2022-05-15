package main

import "fmt"

func main() {
	var a int = 10
	var b int8 = -127
	var c uint16 = 65_535
	var d float64 = 3.141592654

	// Rune is an alias for int32
	var e rune = 'S'
	fmt.Printf("%T\n", e)
	fmt.Printf("%v\n", e)

	var f bool = true
	var st string = "string"

	// Array is fixed size
	var numbers = [4]int{-1, 1, 0, 1}
	fmt.Printf("%T \n", numbers)

	// Slice is dynamic size
	var cities = []string{"São Paulo", "Rio de Janeiro", "Belo Horizonte"}
	fmt.Printf("%T \n", cities)

	// MAP
	balances := map[string]float64{
		"USD":  5.06,
		"Euro": 5.26,
	}
	fmt.Println("\n MAP")
	fmt.Printf("%T \n", balances)

	// Struct
	type Person struct {
		name string
		age  int
	}

	var someOne Person
	fmt.Println("\n Struct")
	fmt.Printf("%T \n", someOne)

	// Pointer
	var x int = 2
	p := &x
	fmt.Println("\n Pointer")
	fmt.Printf("%v \n", p)  // Get x address
	fmt.Printf("%v \n", *p) // Get x value

	// Function
	fmt.Println("\n Function type")
	fmt.Printf("%T \n", fun)

	_ = a
	_ = b
	_ = c
	_ = d
	_ = e
	_ = f
	_ = st
}

func fun() {}
