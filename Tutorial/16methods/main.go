package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	yatendra := User{Name: "Yatendra Singh", Email: "yatendr@gmail.com", Status: true, Age: 21}
	fmt.Println("User Data", yatendra)
	yatendra.GetStatus()
	yatendra.NewMail()
	fmt.Println("Yatendra old Email:",yatendra.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the User is :", u.Email)
}

/*Methods in the Golang
In Go, a method is a function associated with a particular type. Methods are similar to functions but are defined within the context of a specific type. They allow you to attach behavior to data structures by defining functions that operate on those data structures.

Creating a Method:
To create a method in Go, you define a function with a special receiver parameter. The receiver parameter binds the method to a specific type. The syntax for defining a method is:

func (receiver ReceiverType) methodName(parameters) returnType {
    // Method body
}
receiver: The receiver parameter specifies the type to which the method belongs.
ReceiverType: The type to which the method is attached.
methodName: The name of the method.
parameters: Optional input parameters.
returnType: Optional return type.
Example:
Let's create a simple example to illustrate how to define and use methods in Go:

package main

import "fmt"

// Define a struct type named Person
type Person struct {
    Name string
    Age  int
}

// Define a method named greet for the Person type
func (p Person) greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create a new Person instance
    john := Person{Name: "John", Age: 30}

    // Call the greet method on the Person instance
    john.greet() // Output: Hello, my name is John and I am 30 years old.
}
In this example:

We define a Person struct type with two fields: Name and Age.
We define a method named greet that belongs to the Person type. This method prints a greeting message using the Name and Age fields of the Person.
In the main function, we create a new Person instance named john.
We call the greet method on the john instance, which prints a greeting message.
Uses of Methods:
Encapsulation: Methods encapsulate behavior within the context of a type, allowing you to define actions that can be performed on instances of that type.
Code Organization: Methods help organize code by associating related behaviors with the types they operate on.
Promotes Reusability: Methods enable you to define reusable behaviors that can be used across multiple instances of the same type.
Methods are a powerful feature in Go that facilitate object-oriented programming principles like encapsulation and code organization while maintaining the simplicity and efficiency of the language.
*/