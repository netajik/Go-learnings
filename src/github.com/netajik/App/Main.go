// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

var i = 14

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
	var foo bool = false
	fmt.Println(foo)
	var a float32 = 4;
	fmt.Println(a)

	//functions
	//test()
}

func test() {
	fmt.Println(" World")
}
