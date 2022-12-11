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
	// output:
	// 	 value Memory usage: 31 Bytes

	// decompose:
	// 	 string occupies 16 bytes
	// 	 int occupies 8 bytes
	// 	 bool occupies 1 bytes
	// 	 "jwping" occupies 6 bytes
	// 	 31 bytes in total
}
```