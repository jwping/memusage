# memusage

Example code
```golang
package main

import (
	"fmt"
	"log"
	"github.com/jwping/memusage"
)

func main() {
	type Value struct {
		Name    string
		Age     int
		Married bool
	}

	value := Value{
		Name:    "jwping",
		Age:     18,
		Married: false,
	}

	size, err := memusage.GetSize(&value)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("value Memory usage: %d Bytes\n", size)
}
```