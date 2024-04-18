package main

import "fmt"

func main() {
	fmt.Println("Hello From the Go Array")

	var fruitList [5]string 
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Peach"
	fruitList[4] = "Grapes"

	fmt.Println("Fruit List is :",fruitList)
	fmt.Println("Fruit List is :",len(fruitList))

	var vegList = [5]string{"potato","tomato","carrot","Ginger"}
	fmt.Println("Veggies List",vegList)
	fmt.Println("Length of the Veggies list Array",len(vegList))

	// the len method gives you the length of the array no matter what the number of element is present in that
	
}