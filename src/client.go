// Arash is responsible for this file

package main

import (
	"errors"
	"log"
)

// PairType : the values to be used in the dictionary
type PairType struct {
	integer int
	boolean bool
}

type aType struct {
	a int
}

// CustomFunc : use to support the Fold function
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

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// |PART 1|
	// Create a new empty dictionary
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	real_d = d.Empty()

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// |PART 2|
	// For each i, 1 <= i < 500, insert a (key,value) pair constructed as follows.
	// The key should be a string of 2i copies of the character ’a’.
	// The corresponding value should be the pair (i, b) where b is true if i is even and false
	// otherwise.
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	var key = ""
	var err error
	for i := 1; i < 500; i++ {
		key = key + "a" + "a"

		pair_value.integer = i
		if i%2 == 0 {
			pair_value.boolean = true
		} else {
			pair_value.boolean = false
		}

		// Test: to see if the key is in the dictionary yet, we expect it is not before insert
		if real_d.Lookup(key).isFound {
			err = errors.New("Error: Somehow this key:  " + key + "  already existed in the dictionary")
			log.Fatal(err)
		}
		real_d.Insert(key, pair_value)

		// Test: to see if the key is in the dictionary yet, we expect it is after insert
		if !real_d.Lookup(key).isFound {
			err = errors.New("Error: Cannot insert " + key + " into the dictionary properly")
			log.Fatal(err)
		}
	}

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// |PART 3|
	// For each i, 1 <= i < 333, lookup the key consisting of 3i copies of the character ’a’
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	key = ""
	for i := 1; i < 333; i++ {
		key = key + "a" + "a" + "a"

		// Test: if the key is really 3i copies of a
		if len(key)%3 != 0 {
			err = errors.New("Error: This key:  " + key + " is not 3i copies of a")
			log.Fatal(err)
		}

		// Test: if the key exists, we expect the value of that key to be a pair
		if real_d.Lookup(key).isFound {
			switch real_d.Lookup(key).value.(type) {
			case PairType:
				continue
			default:
				err = errors.New("Error: Expected this lookup value to be a pair")
				log.Fatal(err)
			}
		}

		// Test: if the key is 2k copies of a, the key should be found
		if len(key)%2 == 0 {
			if !real_d.Lookup(key).isFound {
				err = errors.New("Error: Expected key:  " + key + " is to be found")
				log.Fatal(err)
			}
		}
	}

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// |PART 4|
	// Use a fold to compute the total number of entries in the dictionary containing a value
	// whose boolean component is true.
	// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	a_fold.a = 0
	var total = real_d.Fold(a_fold, CustomFunc)

	// Test: if the total number of entries is 249
	if total.(aType).a != 249 {
		err = errors.New(" Error: Unexpected result for Fold function ")
		log.Fatal(err)
	}
}
