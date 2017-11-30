// Thong is responsible for this file

package main

import (
	"container/list"
)

// DictionaryLinkedList a dictionary implemented using linked-list structure
type DictionaryLinkedList struct {
	l interface{}
}

// Empty return an empty dictionary
func (d DictionaryLinkedList) Empty() DictInterface {
	d.l = list.New()
	return d
}

// Lookup lookup & return the given string in this dictionary
func (d DictionaryLinkedList) Lookup(s string) LookupResult {
	var result LookupResult
	return result
}

// Insert insert the given pair of string and value into this dictionary
func (d DictionaryLinkedList) Insert(s string, value GenericTypeV) {

}

// Fold perform the given function on all the elements in this dictionary
func (d DictionaryLinkedList) Fold(a GenericTypeA, clientFunc LambdaFunc) GenericTypeA {
	return a
}
