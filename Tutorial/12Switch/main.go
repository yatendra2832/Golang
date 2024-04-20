package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case statement in Golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Random Number is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("confused")
	}

}

/*Use cases
Multiple Conditions: When you need to check multiple conditions based on the value of a single expression.
********************************************************
switch day {
case "Monday":
    fmt.Println("It's Monday!")
case "Tuesday":
    fmt.Println("It's Tuesday!")
// More cases...
default:
    fmt.Println("It's not a weekday!")
}
******************************************************
Enumerated Types: Switch is particularly useful when working with enumerated types.

type Status int

const (
    New Status = iota
    InProgress
    Completed
)

func getStatusText(status Status) string {
    switch status {
    case New:
        return "New"
    case InProgress:
        return "In Progress"
    case Completed:
        return "Completed"
    default:
        return "Unknown"
    }
}
*********************************************************
Fallthrough: Go's switch statement does not fall through by default, but you can explicitly specify fallthrough using the fallthrough keyword.

switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
// Output: 
// One
// Two
*****************************************************
Type Switch: switch can also be used to check the type of an interface value.

var x interface{} = 5

switch x.(type) {
case int:
    fmt.Println("x is an integer")
case string:
    fmt.Println("x is a string")
default:
    fmt.Println("x is of unknown type")
}
Expressionless Switch: In Go, switch can be used without an expression, and each case clause is evaluated for truthiness.

*********************************************************
hour := time.Now().Hour()

switch {
case hour < 12:
    fmt.Println("Good morning!")
case hour < 18:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}

*/