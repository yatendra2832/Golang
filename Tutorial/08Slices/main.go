package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome To Journey on Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Typeof FruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("fruitList after adding :", fruitList)

	// fruitList = append(fruitList[1:])  that will start from the index 1 and till the end
	// fruitList = append(fruitList[1:3]) that will strt with the index 1 and go upto range-1 which is 3-1 =2
	// fruitList = append(fruitList[:3]) that will start with the index 0 which is default index and go upto the index 2 as 3 will not included
	fruitList = append(fruitList[1:3])
	fmt.Println("Print the FruitList by adding the range", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 321
	highScores[2] = 1799
	highScores[3] = 21
	fmt.Println(highScores)
	// highScores[4] = 777 as this is not possible as we have only the 4 as limit for that but when we use append we can add more element as the append will reallocate the memory]
	highScores = append(highScores, 234, 876, 9876)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("After Sorting the Scores", highScores)

	fmt.Println(sort.IntsAreSorted((highScores))) // that will tell is the value is sorted or not

	// How to remove a value from slice based on index
	var courses = []string{"react", "js", "nodejs", "python"}
	fmt.Println("Courses", courses)
	var deletingIndex int = 2
	courses = append(courses[:deletingIndex], courses[deletingIndex+1:]...) // that will delete the nodejs as their index is 2
	fmt.Println("courses after deleting ", courses)
	/*
			The courses slice is initialized with four strings: "react", "js", "nodejs", and "python".
		The deletingIndex variable is set to 2, indicating that the element at index 2 ("nodejs") should be removed.
		The append function is used to delete the element at deletingIndex from the courses slice. Here's how it works:
		courses[:deletingIndex] extracts the elements from the beginning of the slice up to (but not including) deletingIndex. This part retains all elements before the one to be deleted.
		courses[deletingIndex+1:] extracts the elements after deletingIndex. This part retains all elements after the one to be deleted.
		The ellipsis (...) after courses[deletingIndex+1:] unpacks the elements from the slice.
		By appending these two slices together, the element at deletingIndex is effectively removed.
		The modified courses slice, now without the element at index 2, is printed to the console.*/
}

/*
The program starts with the main package declaration.
The fmt package is imported for printing messages.
The main function is defined, which is the entry point of the program.
A welcome message is printed to the console.
A slice named fruitList is declared and initialized with three string elements: "Apple", "Tomato", and "Peach".
The type of fruitList is printed using the Printf function.
Two new elements ("Mango" and "Banana") are appended to the fruitList slice using the append function.
The updated fruitList is printed to the console.
A slice of elements from index 1 to 2 (excluding index 3) of fruitList is appended to fruitList itself using the append function.
The updated fruitList after appending the slice is printed to the console
*/
