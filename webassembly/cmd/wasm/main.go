package main

import "fmt"

// main 1. need to use 'GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm'
// to compile this golang file into web-assembly
// 2. cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
func main() {
	fmt.Println("Go Web Assembly")
}
