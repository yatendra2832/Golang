package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Form the time study of golang")
	presentTime := time.Now()
	fmt.Println("Present Time", presentTime)
	fmt.Println("Present Time", presentTime.Format("01-02-2006 15:04:05 Monday")) //this is a fix for formating we have to always use that value
	createdDate := time.Date(2020, time.August,10 ,23,23,0,0, time.UTC)
	fmt.Println(createdDate)

}
