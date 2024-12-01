//go:build js && wasm

package main

import (
	"syscall/js"
)

// Add 函数实现简单的加法，并暴露到 JavaScript
func add(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	return a + b
}

func main() {
	// 暴露函数给 JavaScript
	js.Global().Set("add", js.FuncOf(add))

	// 阻塞主线程，让 WASM 一直运行
	select {}
}
