package main

// Dictionary blah
type Dictionary struct {
	value GenericType
}

// Pair dsjfjdskfds
type Pair struct {
	num int
	val bool
}

// Speak blah
func (d Dictionary) Speak() string {
	return "I am Type A!"
}

// Init blah
func (d Dictionary) Init() DictInterface {
	var d1 Dictionary
	var val1 Pair
	val1.num = 100
	val1.val = true
	d1.value = val1
	return d1
}
