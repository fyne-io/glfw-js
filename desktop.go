// +build !js

package goglfw

import (
	"runtime"

	glfw "github.com/shurcooL/glfw3"
)

func Init() error {
	runtime.LockOSThread()

	return glfw.Init()
}

func Terminate() error {
	return glfw.Terminate()
}

func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error) {
	var m *glfw.Monitor
	if monitor != nil {
		m = monitor.Monitor
	}
	var s *glfw.Window
	if share != nil {
		s = share.Window
	}

	w, err := glfw.CreateWindow(width, height, title, m, s)

	var window *Window
	if w != nil {
		window = &Window{w}
	}

	return window, err
}

type Window struct {
	*glfw.Window
}

type Monitor struct {
	*glfw.Monitor
}

func PollEvents() error {
	return glfw.PollEvents()
}

type CursorPositionCallback func(w *Window, xpos float64, ypos float64)

func (w *Window) SetCursorPositionCallback(cbfun CursorPositionCallback) (previous CursorPositionCallback, err error) {
	wrappedCbfun := func(_ *glfw.Window, xpos float64, ypos float64) {
		cbfun(w, xpos, ypos)
	}

	p, err := w.Window.SetCursorPositionCallback(wrappedCbfun)
	_ = p

	// TODO: Handle previous.
	return nil, err
}

type FramebufferSizeCallback func(w *Window, width int, height int)

func (w *Window) SetFramebufferSizeCallback(cbfun FramebufferSizeCallback) (previous FramebufferSizeCallback, err error) {
	wrappedCbfun := func(_ *glfw.Window, width int, height int) {
		cbfun(w, width, height)
	}

	p, err := w.Window.SetFramebufferSizeCallback(wrappedCbfun)
	_ = p

	// TODO: Handle previous.
	return nil, err
}