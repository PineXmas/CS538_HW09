package main

// DictInterface I want this type to be exported, so I must add a comment
type DictInterface interface {
	Speak() string
	Init() DictInterface
}

// GenericType fhhffh
type GenericType interface {
}
