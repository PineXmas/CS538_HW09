// Minh is responsible for this file

package main

type Dictionary struct {
	hashtable map[string]GenericTypeV
}

func (d Dictionary) Empty() DictInterface {
	d.hashtable = make(map[string]GenericTypeV)
	return d
}

func (d Dictionary) Lookup(s string) LookupResult {
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

func (d Dictionary) Insert(s string, value GenericTypeV) {
	d.hashtable[s] = value
}

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
