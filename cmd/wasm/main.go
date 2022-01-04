// Program to be copiled to webAssembly and executed in the browser

package main

import (
	"fmt"
	"syscall/js"
)

var text string = "ceci est un texte que je veux mettre en majuscule qui est passé par webassembly"

func add(a, b int) int {
	return a + b
}

func greet(greet string) string {
	return greet
}

func main () {
	// var res = add(4, 5)

	// inp := js.Global().Get("document").Call("findElementById", "monText")
	// inp.value = text
	
	fmt.Println(text)
	fmt.Println("Hello, World!")

	js.Global().Set("greet", greet)
	js.Global().Set("varia", text)
}

