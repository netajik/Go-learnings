// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

var i = 14

//Foo struct
type Foo struct {
	bar int
}

func main() {
	//variables
	//var i = 12
	var j float32 = 13 
	i := 14
	var k string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf(" %v, %T\n", j, j)
	fmt.Printf(" %v, %T\n", k, k)

	//data types
	var fo bool = false
	fmt.Println(fo)
	var a float32 = 4;
	fmt.Println(a)

	//functions
	//test()

	//defer
	fmt.Println("defer")
	fmt.Println(1)
	//defer fmt.Println(2)
	fmt.Println(3)

	//pointers used to sotore address of a location
	fmt.Println("\npointers")
	var foo *Foo
	fmt.Println(foo)
	foo= new(Foo) // assigning memory for foo
	fmt.Println(foo) // reference
	fmt.Println((*foo)) //dereference
	fmt.Println((*foo).bar)

	//or compiler dereference during compile time
	fmt.Println(foo.bar)
}

func test() {
	fmt.Println(" World")
}
