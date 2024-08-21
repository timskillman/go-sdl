package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	TITLE  = "Square"
	WIDTH  = 800
	HEIGHT = 600
)

var (
	playerX, playerY = int32(WIDTH / 2), int32(HEIGHT / 2)
)

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(TITLE, int32(sdl.WINDOWPOS_UNDEFINED), int32(sdl.WINDOWPOS_UNDEFINED), WIDTH, HEIGHT, uint32(sdl.WINDOW_SHOWN))
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, uint32(sdl.RENDERER_ACCELERATED))
	if err != nil {
		panic(err)
	}
	renderer.Clear()
	defer renderer.Destroy()

	//gl.Enable(gl.DEPTH_TEST)

	// Draw initial frame
	draw(renderer, playerX, playerY)

	// Event loop

	for event := sdl.WaitEvent(); event != nil; event = sdl.WaitEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
			println("Quitting..")
			return
		case *sdl.KeyboardEvent:
			if t.State == sdl.PRESSED {
				switch t.Keysym.Sym {
				case sdl.K_LEFT:
					playerX -= 4
					if playerX < 0 {
						playerX = WIDTH
					}
				case sdl.K_RIGHT:
					playerX += 4
					if playerX > WIDTH {
						playerX = 0
					}
				case sdl.K_UP:
					playerY -= 4
					if playerY < 0 {
						playerY = HEIGHT
					}
				case sdl.K_DOWN:
					playerY += 4
					if playerY > HEIGHT {
						playerY = 0
					}
				}

				draw(renderer, playerX, playerY)
			}
		}
	}

}

func draw(renderer *sdl.Renderer, X int32, Y int32) {
	// Clear surface
	//surface.FillRect(nil, 0)
	//
	//// Draw on the surface
	//rect := sdl.Rect{X: int32(playerX - 20), Y: int32(playerY - 20), W: 40, H: 40}
	//surface.FillRect(&rect, 0xff00ff)
	//
	//window.UpdateSurface()

	//gfx.StringColor(renderer, 16, 16, "GFX Demo", sdl.Color{R: 0, G: 255, B: 0, A: 255})
	//gfx.CharacterColor(renderer, X, Y, 'X', sdl.Color{R: 255, G: 0, B: 0, A: 255})

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(X, Y, X+10, Y+10)
	renderer.Present()
	sdl.PollEvent()
}
