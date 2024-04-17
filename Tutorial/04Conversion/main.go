package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Our Pizza App \nPlease Rate our Pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks For rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to Your rating: ", numRating+1)
	}

}

/* String Conversion
strings.TrimSpace(input):

strings.TrimSpace is a function provided by the strings package in Go.
It removes leading and trailing white spaces from the string.
input is the string variable that contains the user's input obtained from ReadString function.
So, strings.TrimSpace(input) ensures that any leading or trailing white spaces in the user's input are removed.
strconv.ParseFloat(..., 64):

strconv.ParseFloat is a function provided by the strconv package in Go.
It converts a string representing a floating-point number to its float64 equivalent.
The first argument (strings.TrimSpace(input)) is the string that represents the floating-point number. We use strings.TrimSpace(input) to ensure there are no leading or trailing white spaces in the input string.
The second argument (64) specifies the bit size (64 bits) of the floating-point number. It determines the precision and range of the floating-point number being parsed.
numRating, err := ...:

numRating is the variable where the parsed floating-point number will be stored.
err is the variable where any error encountered during the parsing process will be stored. If parsing is successful, err will be nil.
The := operator is the short variable declaration in Go, which declares and initializes the variables numRating and err.
So, putting it all together:

strconv.ParseFloat(strings.TrimSpace(input), 64) takes the user's input, trims leading and trailing white spaces, and attempts to parse it into a float64 value.
The parsed floating-point number is stored in the numRating variable, and any error encountered during parsing is stored in the err variable.
*/