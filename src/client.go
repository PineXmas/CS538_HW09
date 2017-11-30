// Arash is responsible for this file

package main

import (
	"fmt"
)

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

// TestType hahaha
type TestType struct {
	a int
}

func main() {
	test()
}

func test() {
	temp := TestType{}
	temp.a = 100
	var alias *TestType
	alias = &(temp)
	alias.a = 200
	fmt.Println(temp.a)
}
