/*

Scope -

Universal
Package
File
Block


{} - Braces, curly braces, curlies, mustashes
[] - brackets
() - parenheses, parens



*/
package main

import "fmt"

var yes = "Hello"

func main() {
	y := 12
	fmt.Println(y)
	printit() //exist in otherfile.go, but is in the same package thus we can use here
	yes := func() int {
		return 13
	}

	fmt.Println(yes()) //ANON function assigned to expression should print 13

}
