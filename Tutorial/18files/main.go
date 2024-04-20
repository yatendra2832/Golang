package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This need to go in a file - Golangtutorial"

	file, err := os.Create("golangtuts.txt")
	checkNilError(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is :", length)

	readFile("golangtuts.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
