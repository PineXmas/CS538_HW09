// Arash is responsible for this file

package main

import "fmt"

// PairType : the values to be used in the dictionary
type PairType struct {
	integer int
	boolean bool
}

type aType struct {
	a int
}

//type MyLambdaFunc func() aType
func CustomFunc(a1 GenericTypeA, s string, value GenericTypeV) GenericTypeA {
	var count aType
	count.a = a1.(aType).a
	if value.(PairType).boolean == true {
		count.a = count.a + 1
	}

	return count
}

func main() {
	var d Dictionary
	var pair_value PairType
	var a_fold aType
	var real_d DictInterface

	// create a new empty dictionary
	real_d = d.Empty()

	// For each i, 1 <= i < 500, insert a (key,value) pair constructed as follows. The key should be a string of 2i copies
	// of the character ’a’. The corresponding value should be the pair (i, b) where b is true if i is even and false
	// otherwise.
	var key = ""
	for i := 1; i < 500; i++ {
		key = key + "a" + "a"

		pair_value.integer = i
		if i%2 == 0 {
			pair_value.boolean = true
		} else {
			pair_value.boolean = false
		}
		real_d.Insert(key, pair_value)

	}

	// For each i, 1 <= i < 333, lookup the key consisting of 3i copies of the character ’a’
	key = ""
	for i := 1; i < 333; i++ {
		key = key + "a" + "a" + "a"
		fmt.Print(len(key))
		fmt.Print(" = ")
		fmt.Print(real_d.Lookup(key).isFound)

		if real_d.Lookup(key).isFound == true {
			fmt.Print("   ")
			fmt.Print(real_d.Lookup(key).value)
		}

		fmt.Println(" ")
	}

	//  Use a fold to compute the total number of entries in the dictionary containing a value whose boolean component
	// is true.
	a_fold.a = 0
	fmt.Print("Total # of entries: ")
	fmt.Println(real_d.Fold(a_fold, CustomFunc))
}
