package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// https://golangbot.com/webassembly-using-go/

// main 1. need to use 'GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm'
// to compile this golang file into web-assembly
// 2. cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
func main() {

	// need go keep waiting on a channel
	c := make(chan struct{}, 0)

	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper()) // expose operation
	<-c
}

// prettyJson is for returning the string in beautified json format
// after adding the function, need to expose it to JavaScript
func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// jsonWrapper is to expose the function to JavaScript
// it return a function -> as a result
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(
		func(this js.Value, args []js.Value) any {
			if len(args) != 1 {
				return "Invalid no of arguments passed"
			}
			inputJSON := args[0].String()
			fmt.Printf("input %s\n", inputJSON)
			pretty, err := prettyJson(inputJSON)
			if err != nil {
				fmt.Printf("unable to convert to json %s\n", err)
				return err.Error()
			}
			return pretty
		})
	return jsonFunc
}
