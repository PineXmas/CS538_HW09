package main

import (
	"fmt"
)

func main() {
	dict1 := DictTypeA{}
	dict2 := DictTypeB{}
	fmt.Println(dict1.Speak())
	fmt.Println(dict2.Speak())
}
