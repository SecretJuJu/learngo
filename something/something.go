package something

import "fmt"

// this is not exported
func sayBye() {
	fmt.Println("Bye")
}

// this is exported
func SayHello() {
	fmt.Println("Hello")
}
