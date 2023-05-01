package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Button Clicked")

	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func add(_ js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() + i[1].Int())
	js.Global().Set("output", result)
	fmt.Println(result)
	return result
}

func subtract(_ js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() - i[1].Int())
	js.Global().Set("output", result)
	fmt.Println(result)
	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
