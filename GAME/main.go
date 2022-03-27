package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenHeight = 800
	screenWidth  = 600
)

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		fmt.Println("initializing sdl:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"GAMING WITH SDL",
		800, 600,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}

	defer renderer.Destroy()
	for {
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Present()

	}
}
