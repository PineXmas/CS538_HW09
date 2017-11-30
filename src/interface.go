package main

// LambdaFunc the function provided by the clients
type LambdaFunc func(a GenericTypeA, s string, value GenericTypeV) GenericTypeA

// LookupResult to store the look-up result
type LookupResult struct {
	is_found bool
	value    GenericTypeV
}

// GenericTypeA this is to trick all types to be generic
type GenericTypeA interface {
}

// GenericTypeV this is to trick all types to be generic
type GenericTypeV interface {
}

// DictInterface our ADT dictionary
type DictInterface interface {
	Empty() DictInterface
	Lookup(s string) LookupResult
	Insert(s string, value GenericTypeV)
	Fold(a GenericTypeA, client_func LambdaFunc) GenericTypeA
}
