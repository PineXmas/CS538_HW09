// Thong is responsible for this file

package main

import (
	"container/list"
)

// Dictionary a dictionary implemented using linked-list structure
type Dictionary struct {
	l *list.List
}

// ListElement to store info of one element in the linked list
type ListElement struct {
	s   string
	val GenericTypeV
}

// Empty return an empty dictionary
func (d Dictionary) Empty() DictInterface {
	d.l = list.New()
	return d
}

// Lookup lookup & return the given string in this dictionary
func (d Dictionary) Lookup(s string) LookupResult {
	var result LookupResult

	var found *list.Element
	found = d.Search(s)
	if found == nil {
		result.isFound = false
	} else {
		result.isFound = true
		result.value = found.Value.(ListElement).val
	}

	return result
}

// Insert insert the given pair of string and value into this dictionary
func (d Dictionary) Insert(s string, value GenericTypeV) {

	e := ListElement{}
	e.s = s
	e.val = value

	var found *list.Element
	found = d.Search(s)
	if found != nil {
		d.l.InsertBefore(e, found)
		d.l.Remove(found)
	} else {
		d.l.PushBack(e)
	}
}

// Fold perform the given function on all the elements in this dictionary
func (d Dictionary) Fold(a GenericTypeA, clientFunc LambdaFunc) GenericTypeA {
	var currElement *list.Element

	for currElement = d.l.Front(); currElement != nil; currElement = currElement.Next() {
		a = clientFunc(a, currElement.Value.(ListElement).s, currElement.Value.(ListElement).val)
	}

	return a
}

// Search search for an element in the list & return the element if found, otherwise return nil
func (d Dictionary) Search(s string) *list.Element {
	var foundElement *list.Element
	var currElement *list.Element

	foundElement = nil
	for currElement = d.l.Front(); currElement != nil; currElement = currElement.Next() {
		if currElement.Value.(ListElement).s == s {
			foundElement = currElement
			break
		}
	}

	return foundElement
}
