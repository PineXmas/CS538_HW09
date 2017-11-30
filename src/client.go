// Arash is responsible for this file

package main

import "fmt"

// type Test struct {
// 	custom_func LambdaFunc
// }

// func CustomFunc (a GenericTypeA, s string, value GenericTypeV) GenericTypeA {
// 	if (value == 100) {
// 		return true
// 	} else {
// 		return 1
// 	}
// }

func main() {
	var d Dictionary
	var Dict = d.Empty()
	fmt.Print(Dict)
	Dict.Insert("a", 1)
	Dict.Insert("b", 2)
	fmt.Print(Dict)
	fmt.Print(Dict.Lookup("a"))
	Dict.Insert("a", 5)
	fmt.Print(Dict.Lookup("a"))

}
