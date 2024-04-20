package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	/*for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}


	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}
	*/
	/*
		roughValue := 1
		for roughValue < 10{

			if roughValue ==5 {
				roughValue ++
				continue
			}
			fmt.Println("Value is :",roughValue)
			roughValue++
		} // that is just similar to the while loop

	*/

}

/*
The for loop in Go is a versatile construct used for iterating over collections, repeating a block of code a specific number of times, and executing code until a certain condition is met. Here are some common use cases for the for loop in Go:
***************************************

Iterating Over Arrays and Slices:

numbers := []int{1, 2, 3, 4, 5}
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
***************************************

Iterating Over Maps:

ages := map[string]int{
    "Alice": 30,
    "Bob":   35,
    "Eve":   25,
}
for key, value := range ages {
    fmt.Printf("%s is %d years old\n", key, value)
}
***************************************

Processing Each Character of a String:

str := "Hello, World!"
for i := 0; i < len(str); i++ {
    fmt.Println(string(str[i]))
}
***************************************

Repeating Code a Specific Number of Times:

for i := 0; i < 3; i++ {
    fmt.Println("Hello, Golang!")
}
***************************************

Iterating Over a Range of Numbers:

for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
***************************************

Looping Until a Certain Condition is Met:

sum := 0
for sum < 100 {
    sum += rand.Intn(10) // Adding random numbers until sum >= 100
}
Looping Until Break Condition is Met:
go
Copy code
for {
    // code block
    if condition {
        break
    }
}
***************************************
Pattern Printing:

for i := 0; i < 5; i++ {
    for j := 0; j <= i; j++ {
        fmt.Print("* ")
    }
    fmt.Println()
}
These are just a few examples of how the for loop can be used in Go. It's a fundamental control structure that provides great flexibility for various looping scenarios in Go programs.






*/
