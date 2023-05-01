package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// main 1. need to use 'GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm'
// to compile this golang file into web-assembly
// 2. cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
func main() {
	fmt.Println("Go Web Assembly")
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
func jsonWrapper() js.Func {

}
