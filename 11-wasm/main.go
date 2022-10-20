package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `<h4>Hello, I'm a WebAssembly program!</h4>`

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	/*
		Notes:
		You can run Go programs in the browser using WebAssembly.
		To compile a Go program to WebAssembly you must use the GOOS and GOARCH flags.
	*/

	ch := make(chan struct{}, 0)

	fmt.Printf("Hello, WebAssembly!\n")

	js.Global().Set("getHtml", GetHtml())
	<-ch
}
