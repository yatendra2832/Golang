package main

import "fmt"

func main() {
	// Print a message to indicate the start of the program
	fmt.Println("Hello From the Pointers")

	// Declare a variable 'myNumber' and initialize it with the value 23
	myNumber := 23

	// Declare a pointer variable 'ptr' to store the memory address of 'myNumber'
	// The '&' operator is used to get the memory address of 'myNumber' and assign it to 'ptr'
	var ptr = &myNumber

	// Print the memory address stored in 'ptr'
	fmt.Println("Memory address stored in 'ptr':", ptr)

	// Dereference 'ptr' using the '*' operator to access the value stored at the memory address it points to (which is 'myNumber')
	// Print the value stored at that memory address
	fmt.Println("Value stored at the memory address pointed by 'ptr':", *ptr)

	// Manipulate the value stored at the memory address pointed by 'ptr'
	// Multiply the value by 2 and store the result back at the same memory address
	*ptr = *ptr * 2

	// Print the updated value of 'myNumber'
	fmt.Println("Updated value of 'myNumber' after manipulation:", myNumber)
}
