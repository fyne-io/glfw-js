//go:build wasm

package glfw

import (
	"syscall/js"
)

// GetClipboardString returns the contents of the system clipboard, if it contains or is convertible to a UTF-8 encoded string.
//
// This function may only be called from the main thread.
func GetClipboardString() string {
	clipboard := js.Global().Get("navigator").Get("clipboard")
	clipboardChan := make(chan js.Value)

	clipboard.Call("readText").Call("then", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		clipboardContent := p[0]
		clipboardChan <- clipboardContent
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clipboardChan <- js.ValueOf(nil)
		return nil
	}))

	result := <-clipboardChan
	if !result.Truthy() {
		return ""
	}

	return result.String()
}

// SetClipboardString sets the system clipboard to the specified UTF-8 encoded string.
//
// This function may only be called from the main thread.
func SetClipboardString(str string) {
	js.Global().Get("navigator").Get("clipboard").Call("writeText", str)
}
