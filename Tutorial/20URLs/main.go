package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://jsonplaceholder.typicode.com?coursename=reactjs&payment_id=ys2021"

func main() {
	fmt.Println("Handling URLs in Golang")
	fmt.Println(myURL)

	// parsing
	result,_ := url.Parse(myURL)
	fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery) 

	qparams := result.Query()
	fmt.Printf("The type of the query params are : %T \n",qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["payment_id"])

	for _,values := range qparams{
		fmt.Println("Param is",values)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/posts/1",
		RawPath: "",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)




}