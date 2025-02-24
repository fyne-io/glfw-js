//go:build wasm

package glfw

import (
	"syscall/js"
)

// GetClipboardString returns the contents of the system clipboard, if it contains or is convertible to a UTF-8 encoded string.
//
// This function may only be called from the main thread.
func GetClipboardString() string {
	text := make(chan string)

	clipboard := js.Global().Get("navigator").Get("clipboard")
	clipboard.Call("readText").Call("then", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		content := p[0]
		if !content.Truthy() {
			text <- ""
			return nil
		}

		text <- content.String()
		return nil
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text <- ""
		return nil
	}))

	return <-text
}

// SetClipboardString sets the system clipboard to the specified UTF-8 encoded string.
//
// This function may only be called from the main thread.
func SetClipboardString(str string) {
	js.Global().Get("navigator").Get("clipboard").Call("writeText", str)
}
