package main

import (
	"strconv"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add)) //
	js.Global().Set("min", js.FuncOf(min)) //
}

func min(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
	return nil
}

func add(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	return nil
}

func main() {
	c := make(chan struct{}, 0) //チャネルを生成
	//println("Go WebAssembly Initialized")
	registerCallbacks() //registerCallbacks()関数を実行
	<-c
	<-c
}
