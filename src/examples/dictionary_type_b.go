package main

// Dictionary blah
type Dictionary struct {
	value GenericType
}

// ThongType dsjfjdskfds
type ThongType struct {
	num  int
	num1 int
	val  bool
}

// Speak blah
func (d Dictionary) Speak() string {
	return "I am Type B!"
}

// Init blah
func (d Dictionary) Init() DictInterface {
	var d1 Dictionary
	var val1 ThongType
	val1.num = 100
	val1.val = true
	val1.num1 = 101
	d1.value = val1
	return d1
}
