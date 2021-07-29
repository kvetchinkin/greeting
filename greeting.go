package greeting

import "fmt"

const TriangleSides int = 3

func Hello() {
	fmt.Println("Hello!")
}
func Hi() {
	fmt.Println("Hi!")
}

func All() {
	Hi()
	Hello()
}
