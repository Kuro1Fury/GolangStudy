package main // package declaration

/*
import "fmt" // import statement
import "time"
*/

import (
	"fmt"
	"time"
)

// main function
// The curly braces must be on the same line as the function declaration
func main() {
	fmt.Println("Hello World!")
	time.Sleep(1 * time.Second)
}