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

	//defer
	fmt.Println("defer")
	fmt.Println(1)
	//defer fmt.Println(2)
	fmt.Println(3) 

	//pointers used to sotore address of memory location
	fmt.Println("\npointers")
	var foo *Foo
	fmt.Println(foo)
	foo= new(Foo) // assigning memory for foo
	fmt.Println(foo) // reference
	fmt.Println((*foo)) //dereference
	fmt.Println((*foo).bar)

	//or compiler dereference during compile time
	fmt.Println(foo.bar)

	//functions
	fmt.Println("\n------functions------------")
	// pass by value
	total := sum(1,2,3)
	fmt.Println("sum:",total)
	// pass by reference
	msg := "Hello World"
	passbyreference(&msg)
	fmt.Println(msg)

	value,err := divide(2,1)
	if err != nil {
		fmt.Println(err)
		return
	} 
	fmt.Println(value)

	//anonymaous function
	fun:= func() {
		fmt.Println("This is anonymous function")
	}
	fun()
}

func sum(values ...int) int {
	total:= 0
	for _,value:= range values {
		total +=value
	}
	return total
}

func passbyreference(msg *string) {
	fmt.Println(*msg)
	*msg = "Hello"
}

func divide(a, b float32) (float32, error){
	if(b==0.0 || b==0){
		return 0.0, fmt.Errorf("divide by zero")
	}

	return a/b, nil
}