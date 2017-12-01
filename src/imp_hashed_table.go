// Minh is responsible for this file

package main

// Dictionary a dictionary implemented using hashtable structure
type Dictionary struct {
	hashtable map[string]GenericTypeV
}

// Empty return an empty dictionary
func (d Dictionary) Empty() DictInterface {
	d.hashtable = make(map[string]GenericTypeV)
	return d
}

// Lookup lookup & return the given string in this dictionary
func (d Dictionary) Lookup(s string) LookupResult {
	value, ok := d.hashtable[s]
	if ok {
		var result LookupResult
		result.isFound = true
		result.value = value
		return result
	}
	var error LookupResult
	error.isFound = false
	error.value = nil
	return error
}

// Insert insert the given pair of string and value into this dictionary
func (d Dictionary) Insert(s string, value GenericTypeV) {
	d.hashtable[s] = value
}

// Fold perform the given function on all the elements in this dictionary
func (d Dictionary) Fold(a GenericTypeA, client_func LambdaFunc) GenericTypeA {
	if len(d.hashtable) == 0 {
		return a
	}
	var result = a
	for key, value := range d.hashtable {
		result = client_func(result, key, value)
	}
	return result
}
