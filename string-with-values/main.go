package main

import "fmt"

func main() {
	age := 23
	fmt.Println("This is string with a value", age)
	fmt.Printf("My age is %d years old \n", age)

	name := "John"
	fmt.Printf("Random name is %s \n", name)
	fmt.Printf("Name with qutoes %q \n", name)

	pi := 3.1515
	fmt.Printf("Pi is equal to %f \n", pi)
	fmt.Printf("Pi is equal to %.2f \n", pi)

	// %s for string
	// %q for string with quotes

	// %d for integer
	// %f for float

	// %t for boolean

	// %v for any type
	// %T for print the type of the value

	// %b for base 2
	// %x for base 10

}
