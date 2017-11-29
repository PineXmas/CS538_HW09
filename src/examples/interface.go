package main

// Dictionary I want this type to be exported, so I must add a comment
type Dictionary interface {
	Speak() string
}
