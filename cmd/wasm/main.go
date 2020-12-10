package main

import (
	"fmt"
	"io/ioutil"
	"syscall/js"

	"github.com/tamada/samplewasm"
)

func showError(message string) {
	document := js.Global().Get("document")
	app := document.Call("getElementById", "message")
	app.Set("innerText", message)
}

func callbackFunc(this js.Value, args []js.Value) interface{} {
	urlString := args[0].String()
	fmt.Println(`callbackFunc("%s")`, urlString)

	reader, err := samplewasm.Get(urlString)
	if err != nil {
		fmt.Println(err.Error())
		showError(err.Error())
		return nil
	}
	defer reader.Close()
	fullyText, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		showError(err.Error())
		return nil
	}

	document := js.Global().Get("document")
	app := document.Call("getElementById", "app")
	app.Set("innerText", string(fullyText))

	return nil
}

func registerCallbacks() {
	js.Global().Set("wasmCallback", js.FuncOf(callbackFunc))
}

func initWasm() {
	fmt.Println("Hello Wasm")
}

func main() {
	channel := make(chan interface{}, 0)
	initWasm()
	registerCallbacks()
	<-channel
}
