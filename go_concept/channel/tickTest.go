// Golang program to illustrate the usage of
// Tick() function

// Including main package
package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Defining UTCtime
func UTCtime() string {
	return ""
}

// Main function
func main() {

	// Calling Tick method
	// using range keyword
	//for tick := range time.Tick(3 * time.Second) {
	for tick := range time.Tick(time.Second) {
		// Prints UTC time and date
		fmt.Println(tick, UTCtime())
	}
}
