package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for the Course : ")
	// Comma ok  || err error 
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks For rating: ", input)
	fmt.Printf("Thanks For rating: %T ", input)



}
/* Package Imports:

bufio: This package provides buffered I/O. We'll use it to create a buffered reader for reading input from the user.
fmt: This package provides functions for formatted I/O. We'll use it to print messages to the console.
os: This package provides functions for operating system functionalities. We'll use it to access standard input (os.Stdin).
Main Function:

func main() { ... }: This is the entry point of the program.
Print Welcome Message:

welcome := "Welcome to user Input": Declares and initializes a string variable welcome with a welcome message.
fmt.Println(welcome): Prints the welcome message to the console.
Reading User Input:

reader := bufio.NewReader(os.Stdin): Creates a new buffered reader reader to read input from standard input (os.Stdin).
fmt.Println("Enter the Rating for the Course : "): Prints a prompt message asking the user to enter the rating for the course.
input, _ := reader.ReadString('\n'): Reads input from the user until a newline character ('\n') is encountered. The input is stored in the variable input. Any error returned by ReadString is ignored for simplicity.
Printing User Input:

fmt.Println("Thanks For rating: ", input): Prints a thank you message along with the input received from the user.
fmt.Printf("Thanks For rating: %T ", input): Prints the type of the input variable using %T format specifier in Printf function.
*/

/*Comma ok error syntax working
In Go, when calling a function that returns multiple values, such as the ReadString function from the bufio package, you can use the "comma, ok" idiom to check if an error occurred during the function call. This is particularly useful when dealing with functions that can return both a value and an error.

Here's how the "comma, ok" idiom works:

input, err := reader.ReadString('\n')
In this line of code, ReadString('\n') is a function call that returns two values: input (the string read from the input) and err (an error, if any, that occurred during the reading process).

The input variable will store the string read from the input.
The err variable will store any error that occurred during the reading process.
Now, let's say you want to check if an error occurred during the function call. You can use the "comma, ok" idiom:


input, err := reader.ReadString('\n')
if err != nil {
    // Handle the error
    fmt.Println("An error occurred:", err)
    // Additional error handling logic can go here
} else {
    // No error occurred, continue with processing the input
    fmt.Println("Input:", input)
}
In this code:

The if err != nil condition checks if the err variable contains a non-nil value, indicating that an error occurred during the function call.
If an error occurred (err is not nil), you can handle the error appropriately. In this example, we're simply printing the error message using fmt.Println.
If no error occurred (err is nil), you can continue with processing the input as desired.
The "comma, ok" idiom is commonly used in Go to handle errors returned by functions along with their respective values. It allows for concise and readable error handling in Go programs.*/