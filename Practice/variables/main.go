package main

import "fmt"

func main() {

	//short hand variable declaration

	variable1 := 10      //This is a variable that is an int
	variable2 := "Hello" //this is a variable that is a string
	variable3 := false   //this is a variale that is a boolean

	var variable4 string  // declaring a variable that is a string
	variable4 = "Goodbye" //assigning to a varaible

	variable5 := 'M' //this is a variable that is a character

	var a, b, c int = 1, 2, 3 //declare and initialize multiple variables

	fmt.Println(variable1, variable2, variable3, variable4)

	//finding the tye of a variable

	fmt.Printf("%v %v %v %v %v \n", variable1, variable2, variable3, variable4, variable5) //print them all out
	fmt.Printf("%T %T %T %T %T \n", variable1, variable2, variable3, variable4, variable5) //prints their types out
	fmt.Printf("%d %d %d \n", a, b, c)                                                     //prints ints for a,b,c
}

/*
Notes:

Declare, you just say there is a varaible and it has a type
assign you have a variable and you give it a value
initialize is when you declare and assign at the same time

to initialize to zero, type var name type

types in golang
bool, int, float64, string

you can declare multiple variables of the same type
var a,b,c,d string will declare 4 varaibles of type string

you can also declare and initialize many at once
var a,b,c int = 1,2,3 this assigns 1 to a, 2 to b, 3 to c

in the above example it's possible to remove the type and let the compiler infer
var a,b,c = 1,2,3 same effect

its even possible to infer mixed types
var a,b,c = 1 true "hello world"

a will equal 1, b will equal boolean true, c will equal string "hello world"

you can even ommit the var and use shorthand

a,b,c := 1, 2, 3

or  mix type
a,b,c := "hello", 1, true



*/
