package main

import (
	"fmt"
)

func main() {
	// dict := Dictionary{}
	// fmt.Println(dict.Speak())

	var a Dictionary
	var b = a.Init()
	fmt.Println(b.Speak())
}
