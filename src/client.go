// Arash is responsible for this file

package main
//import "fmt"

 type PairType struct {
 		integer int
		boolean bool
}
type aType struct {
	a int
}

type MyLambdaFunc func() aType

func (m MyLambdaFunc) CustomFunc (a1 aType, s string, value PairType) aType{
	var count aType
	count.a = a1.a
 if value.boolean == true {
	count.a = count.a + 1
}
 return count
}

func main() {
	 var d Dictionary
	 var pair_value PairType
	 var a_fold aType
	 var real_d DictInterface
	 var key = ""
	 var key1 = ""
	 var m MyLambdaFunc

	 // create a new empty dictionary
	 real_d = d.Empty()

	 for i := 1; i < 500; i++ {
		key := key + "a" + "a"
		pair_value.integer = i
		if i%2 == 0 {
			pair_value.boolean = true
		} else {
			pair_value.boolean = false
		}
		real_d.Insert(key, pair_value)
		//fmt.Println(real_d.Lookup(key))
	 }
	 //key1 = "bb"
	 //fmt.Println(real_d.Lookup(key1))
	 for i:=1; i<333; i++ {
		 key1 := key1 + "a" + "a" + "a"
		 real_d.Lookup(key1)
	 }
	 a_fold.a = 0
	 real_d.Fold(a_fold, m.CustomFunc)
}
