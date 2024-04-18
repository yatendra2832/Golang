package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Jouney on the maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["py"] = "Python"
	languages["css"] = "Cascading Style Sheet"

	fmt.Println("List of all Languages : ", languages)
	fmt.Println(" JS of  Language : ", languages["JS"])

	delete(languages, "py")
	fmt.Println("After Deleting ", languages)

	// Loops
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v \n", key, value)
	}

	
}
