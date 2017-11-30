// Minh is responsible for this file

package main

type DictionaryHashTable struct {
	hashtable map[string]GenericTypeV
}

func (d DictionaryHashTable) Empty() DictInterface {
	d.hashtable = make(map[string]GenericTypeV)
	return d
}

func (d DictionaryHashTable) Lookup(s string) LookupResult {
	value, ok := d.hashtable[s]
	if ok {
		var result LookupResult
		result.is_found = true
		result.value = value
		return result
	}
	var error LookupResult
	error.is_found = false
	error.value = nil
	return error
}

func (d DictionaryHashTable) Insert(s string, value GenericTypeV) {
	d.hashtable[s] = value
}

func (d DictionaryHashTable) Fold(a GenericTypeA, client_func LambdaFunc) GenericTypeA {
	var result GenericTypeA
	return result
}
