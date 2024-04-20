package main

import "fmt"

func main() {
	fmt.Println("Functions in the Golang")
	// main is act as entry point for the go program
	greeter()

	result := adder(3, 5)
	fmt.Println("The Sum of the number is :", result)

	proResult := proAdder(8, 9, 5, 6, 77)
	fmt.Println("Pro Result is: ", proResult)

}

func greeter() {
	fmt.Println("Hi, from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

/* Function in Golang
In Go, a function is a reusable block of code that performs a specific task. Functions allow you to break down your code into smaller, modular pieces, making it easier to read, understand, and maintain. Here are some key points about functions in Go:

Definition and Syntax:
To define a function in Go, you use the func keyword followed by the function name, a list of parameters (if any), an optional return type, and the function body enclosed in curly braces.
Here's the basic syntax:

func functionName(parameters) returnType {
    // Function body
}
Use Cases:
Modularity and Reusability: Functions help break down complex tasks into smaller, manageable chunks, making your code more modular and easier to reuse.
Abstraction: Functions allow you to abstract away the implementation details of a certain task, providing a higher-level interface for other parts of your code.
Encapsulation: Functions can encapsulate certain behaviors, hiding the internal details and exposing only the necessary functionality to the rest of the program.
Code Organization: Functions help in organizing your code logically by grouping related tasks together.
Types of Functions in Go:
**********************************************************
Functions with Parameters and Return Values:
These are the most common types of functions, where you can pass input parameters and receive output values.

func add(a, b int) int {
    return a + b
}
**********************************************************
Variadic Functions:
Variadic functions can accept a variable number of arguments of the same type. They are denoted by ... followed by the type of the argument.

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
Anonymous Functions:
Anonymous functions are functions without a name. They are useful when you need to define a function inline or pass a function as an argument to another function.
go
Copy code
func main() {
    add := func(a, b int) int {
        return a + b
    }
    result := add(3, 5)
    fmt.Println(result) // Output: 8
}
*********************************************************
Higher-Order Functions:
In Go, functions are first-class citizens, meaning you can pass them as arguments to other functions or return them from other functions. Functions that operate on other functions are called higher-order functions.

func applyOperation(num int, operation func(int) int) int {
    return operation(num)
}

func double(num int) int {
    return num * 2
}

func main() {
    result := applyOperation(5, double)
    fmt.Println(result) // Output: 10
}
*****************************************************
Closure Functions:
Closure functions capture and retain the environment in which they are defined. They are particularly useful for creating functions with shared state.

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment()) // Output: 1
    fmt.Println(increment()) // Output: 2
}
Functions are fundamental building blocks in Go programming, enabling you to create flexible, modular, and maintainable code.
*/

/* Pro Adder function
Here's a breakdown of the proAdder function:

Function Name: proAdder
This function is named proAdder, short for "progressive adder."
Parameters: values ...int
The function accepts a variadic number of integers, denoted by ...int. This means it can take zero or more integer arguments.
Return Type: int
The function returns an integer, which represents the total sum of all the input values.
Function Body:
The function initializes a variable named total to store the sum of the input values. It's initialized to 0.
It then iterates over each value passed to the function using a for loop and adds each value to the total variable.
The _ in for _, val := range values is a blank identifier, which is used when you want to ignore the index of the slice/array while iterating. Here, we're only interested in the value itself.
Inside the loop, each value is added to the total variable.
After iterating over all values, the function returns the total sum.
Usage:
This function is useful when you want to add up a variable number of integers without having to specify the number of arguments explicitly. For example, proAdder(1, 2, 3, 4) will return 10, and proAdder() will return 0.
*/
