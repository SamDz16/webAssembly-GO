package main

import (
	"strings"
	"syscall/js"
)

var text string = "ceci est un texte que je veux mettre en majuscule!"

func mettreEnMajuscule() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "IL faut entrer un et un seul argument"
		}
		return strings.ToUpper(args[0].String())
	})
}

func main() {
	println("Go WebAssembly Initialisé")

	js.Global().Set("mettreEnMajuscule", mettreEnMajuscule())

	// TO PERVENT PROGRAM TO EXIT
	<-make(chan struct{}, 0)
}
