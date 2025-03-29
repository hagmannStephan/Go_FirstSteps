# Go: First Steps
The tutorial I used is from [w3-schools](https://www.w3schools.com/go/index.php).
## Initialize Go-Project
```sh
go mod init <project_name>
```
This creates a go.mod file.
## Run a Go-File
```sh
go run .
```
You also can replace the dot with a specific file.
## Hello World
```go
/* Tell the compiler that this is an executable file
(Not a library) */
package main

// Import the format library
import (
	"fmt"
)

/* Define the main function
Go always starts with the main function */
func main() {
    // Print "Hello World!" with a line break at the end
	fmt.Println("Hello World!")
}
```
If you would want to compact this code, you could reolace each line break with a `;`, because they do the same thing.
## Variables
The **naming rules** are as follows: 
- Must start with a letter or an underscore
- Can only contain alpha-numeric characters and `_`
- camelCase is advised for multi-word variables
### Standard `var`
```go
  	var student1 string = "John" //type is string
  	var student2 = "Jane" //type is inferred
	var isActive bool // Without initial value
	var a, b, c, d int = 1, 3, 5, 7	// Multiple variable declaration

	student1 = "Max"
```
### Shortened `:=`
This version can only be used inside of functions and there has to be an initial value.
```go
x := 2 //type is inferred
c, d := 7, "World!" // Multiple variable declaration
```

### Constants
Constants can be created inside and outside of functions. The need to have an initial value and their name should be written in upper case.