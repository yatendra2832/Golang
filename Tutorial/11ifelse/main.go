package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login Count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {
		
	// }


}

/*USE CASES

In Go, the if-else statement is a fundamental control structure used for decision-making. Here are some common use cases for if-else statements in Go:

Conditional Branching: This is the most common use case. Depending on the evaluation of a condition, different blocks of code are executed.

if condition {
    // code block if condition is true
} else {
    // code block if condition is false
}
**********************************************************
Error Handling: Checking for errors returned by functions and handling them accordingly.

if err != nil {
    // handle error
} else {
    // proceed with normal execution
}
**********************************************************
Setting Variables: Assigning different values to a variable based on certain conditions.
var status string
if age >= 18 {
    status = "adult"
} else {
    status = "minor"
}
*********************************************************
Validation: Validating user input or data before processing.

if len(username) < 6 {
    fmt.Println("Username must be at least 6 characters long")
} else {
    // proceed with username validation
}
*********************************************************
Switching Between Options: Using if-else in conjunction with logical operators to select between multiple options.

if choice == "A" || choice == "a" {
    // execute option A
} else if choice == "B" || choice == "b" {
    // execute option B
} else {
    // handle invalid choice
}
*******************************************************
Iterating Over Conditions: Checking multiple conditions in sequence until one is satisfied.

if condition1 {
    // handle condition1
} else if condition2 {
    // handle condition2
} else {
    // handle default case
}

*****************************************************
Resource Cleanup: Using if-else to handle resource cleanup or closing resources conditionally.

file, err := os.Open("filename.txt")
if err != nil {
    // handle error
} else {
    defer file.Close() // Close the file when done
    // proceed with file processing
}*/
