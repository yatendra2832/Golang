package main

import "fmt"

// Constant variable declaration
const LoginToken string = "qwertyuiop"
// - Constants in Go are declared using the 'const' keyword.
// - Constants are immutable and cannot be changed during the program's execution.
// - Conventionally, constant names are written in uppercase to indicate their visibility and immutability.

func main() {
    // Variable declarations and assignments
    var username string = "Yatendra Singh Yadav"
    // - Variables are declared using the 'var' keyword followed by the variable name and type.
    // - Explicit type declaration is done using ':', followed by the type.
    // - The 'string' type indicates that 'username' will hold text data.
    // - Variables can be assigned values using the '=' operator.

    // Printing the value of 'username'
    fmt.Println("Username :", username)
    // - 'fmt.Println' is used to print data to the console.
    // - The value of 'username' is printed to the console.

    // Printing the type of 'username'
    fmt.Printf("Variable is of type : %T \n", username)
    // - 'fmt.Printf' is used for formatted printing.
    // - '%T' is a format specifier to print the type of the variable.

    // Boolean variable declaration and assignment
    var isLoggedIn bool = true
    // - 'bool' is a data type in Go representing Boolean values, which can be either 'true' or 'false'.
    // - The variable 'isLoggedIn' is assigned a value of 'true', indicating the user is logged in.

    // Printing the value of 'isLoggedIn'
    fmt.Println("isLoggedIn: ", isLoggedIn)

    // Printing the type of 'isLoggedIn'
    fmt.Printf("Variable is of Type :%T \n", isLoggedIn)

    // Integer variable declaration and assignment
    var smallVal int = 255
    // - 'int' is a data type in Go representing integer numbers.
    // - The variable 'smallVal' is assigned a value of '255'.

    // Printing the value of 'smallVal'
    fmt.Println("smallValue: ", smallVal)

    // Printing the type of 'smallVal'
    fmt.Printf("Variable is of Type :%T \n", smallVal)

    // Implicit type variable declaration and assignment
    var website = "http://localhost:3001"
    // - Go supports implicit type declaration using the 'var' keyword without specifying the type explicitly.
    // - The variable 'website' is assigned a string value.

    // Printing the value of 'website'
    fmt.Println("Website Link:", website)

    // Short variable declaration (no 'var' style)
    numberOfUser := 3000
    // - Short variable declaration ':=' is used for declaring and initializing variables without specifying the type explicitly.
    // - The variable 'numberOfUser' is assigned an integer value of '3000'.

    // Printing the value of 'numberOfUser'
    fmt.Println("Number of User on the Website:", numberOfUser)

    // Printing the value of 'LoginToken'
    fmt.Println("LoginToken:",LoginToken);
    // - The value of the constant 'LoginToken' is printed to the console.

    // Printing the type of 'LoginToken'
    fmt.Printf("LoginToken Type :%T \n",LoginToken)
    // - The type of the constant 'LoginToken' is printed to the console.
}
