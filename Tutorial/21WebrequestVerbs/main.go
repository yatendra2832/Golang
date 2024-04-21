package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web request Verbs in Golang")
    PerformGetRequest()

}

func PerformGetRequest() {
	const myURL = "http://localhost:3000/get"

	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content Length is :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	defer response.Body.Close()

}

/*

func PerformGetRequest() {
    // Define the URL to send the GET request to
    const myURL = "http://localhost:3000/get"

    // Send a GET request to the specified URL
    response, err := http.Get(myURL)
    if err != nil {
        // If there's an error while making the request, panic
        panic(err)
    }

    // Print the status code of the response
    fmt.Println("Status Code:", response.StatusCode)

    // Print the content length of the response
    fmt.Println("Content Length:", response.ContentLength)

    // Read the response body into a byte slice
    content, _ := ioutil.ReadAll(response.Body)

    // Print the response body as a string
    fmt.Println("Response Body:", string(content))

    // Close the response body to free up resources
    defer response.Body.Close()
}
*/