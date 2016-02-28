package vis

import "fmt"

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)   //capital letter so it's exported
	fmt.Println(yourName) //works because we're in the same package
}
