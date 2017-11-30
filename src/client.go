// Arash is responsible for this file

package main

import (
	"fmt"
)

 type PairType struct {
 		integer int
		boolean bool
}
type aType struct {
	a int
}

func CustomFunc (a aType, s string, value PairType) aType {
	var count int
	count = a
 if boolean.value == true{
	 count = count + 1
 }
}

func main() {
	 var d Dictionary
	 var pair_value PairType
	 var a_fold aType
	 var real_d DictInterface
	 var key = ""
	 var key1 = ""
	 // create a new empty dictionary
	 real_d = d.Empty()

	 for i := 1; i < 500; i++ {
		key := key + "a" + "a"
		integer.pair_value = i
		if i%2 == 0 {
			boolean.pair_value = true
		} else {
			boolean.pair_value = false
		}
		real_d.Insert(key, pair_value)
	 }
	 for i:=1; i<333; i++ {
		 key1 := key1 + "a" + "a" + "a"
		 real_d.Lookup(key1)
	 }
	 a.a_fold = 0
	 real_d.Fold(a_fold, CustomFunc) aType
}
