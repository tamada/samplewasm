package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"syscall/js"

	"github.com/tamada/samplewasm"
)

func showError(message string) {
	document := js.Global().Get("document")
	app := document.Call("getElementById", "message")
	app.Set("innerText", message)
}

func gotItem(reader io.ReadCloser, err error) {
	if err != nil {
		fmt.Println(err.Error())
		showError(err.Error())
		return
	}
	defer reader.Close()
	fullyText, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		showError(err.Error())
		return
	}
	document := js.Global().Get("document")
	app := document.Call("getElementById", "app")
	app.Set("innerText", string(fullyText))
}

func callbackFunc(this js.Value, args []js.Value) interface{} {
	urlString := args[0].String()
	fmt.Printf(`callbackFunc("%s")`, urlString)

	// http.Get and etc causes deadlock error in wasm.
	// Therefore, use goroutine and not use channel.
	reader, err := samplewasm.Get(urlString)
	gotItem(reader, err)
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
