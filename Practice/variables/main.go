package main

import "fmt"

func main() {

	//short hand variable declaration

	variable1 := 10      //This is a variable that is an int
	variable2 := "Hello" //this is a variable that is a string
	variable3 := false   //this is a variale that is a boolean

	var variable4 string  // declaring a variable that is a string
	variable4 = "Goodbye" //assigning to a varaible

	fmt.Println(variable1, variable2, variable3, variable4)

	//finding the tye of a variable

	fmt.Printf("%T %T %T %T", variable1, variable2, variable3, variable4)

}
