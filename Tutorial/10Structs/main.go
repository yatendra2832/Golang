package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Journey on structs in the golang")

	yatendra := User{"Yatendra Singh", "yatendrayadav@gmail.com", true, 21}
	fmt.Println("User Data: ", yatendra)
	fmt.Printf("Yatendra details are %+v\n", yatendra) //it will give the structure with the value

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
