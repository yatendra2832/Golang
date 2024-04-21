package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` // Specify JSON field names using struct tags.
	Price    int      // No struct tag specified, so field name will be used as JSON key.
	Platform string   `json:"website"`        // Specify JSON field names using struct tags.
	Password string   `json:"-"`              // Skip this field when marshaling to JSON.
	Tags     []string `json:"tags,omitempty"` // Omit this field if it's empty when marshaling to JSON.
}

func main() {
	fmt.Println("Welcome to the Journey for the json data in the Golang")
	EncodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"React JS", 299, "codewithyatendra", "abc123", []string{"web dev", "js"}},
		{"React JS", 299, "codewithyatendra", "abc123", []string{"web dev", "js"}},
		{"Angular", 249, "angulardev", "xyz789", []string{"web dev", "typescript"}},
		{"Vue.js", 199, "vuedev", "def456", []string{"web dev", "javascript"}},
		{"Ember.js", 179, "emberdev", "ghi789", []string{"web dev", "javascript"}},
		{"Ember.js", 179, "emberdev", "ghi789", nil},

		// now we will package this data as json

	}
	finalJson, err := json.MarshalIndent(Courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

/*Notes
The course struct is defined with fields representing course data. Struct tags are used to specify JSON field names and customize JSON encoding behavior.
The main function is the entry point of the program. It simply prints a welcome message and then calls the EncodeJson function.
The EncodeJson function encodes sample course data into JSON format. It demonstrates how to use the json.MarshalIndent function to marshal data into a JSON string with indentation for readability.
Error handling is included to handle any errors that may occur during JSON encoding. If an error occurs, the program panics with the error message.
The JSON data is printed to the console using fmt.Printf.
*/
