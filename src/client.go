// Arash is responsible for this file

package main

import (
	"fmt"
)

 type GenericTypeV struct {
 		integer int
		boolean bool
}


// func CustomFunc (a GenericTypeA, s string, value GenericTypeV) GenericTypeA {
// 	if (value == 100) {
// 		return true
// 	} else {
// 		return 1
// 	}
// }

func main() {
	 var d DictInterface
	 var pair_value GenericTypeV
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

}
