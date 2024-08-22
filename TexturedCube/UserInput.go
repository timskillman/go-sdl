package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
	X            int32
	Y            int32
	Z            int32
	LeftButton   bool
	RightButton  bool
	MiddleButton bool
}

func (m *Mouse) SetCoords(x, y int32) {
	m.X, m.Y = x, y
}

func (m Mouse) GetDelta(x, y int32) (int32, int32) {
	return m.X - x, m.Y - y
}

type UserInput struct {
	mouse     Mouse
	quit      bool
	playerPos vec3
}

func (ui *UserInput) GetUserInput() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			ui.quit = true
		case *sdl.MouseMotionEvent:
			ui.mouse.LeftButton = (ev.State == sdl.BUTTON_LEFT)
			ui.mouse.RightButton = (ev.State == sdl.BUTTON_RIGHT)
			ui.mouse.MiddleButton = (ev.State == sdl.BUTTON_MIDDLE)

			if ui.mouse.LeftButton {
				x, y := ui.mouse.GetDelta(ev.X, ev.Y)
				ui.playerPos.x -= float32(x)
				ui.playerPos.y += float32(y)
			}
			ui.mouse.SetCoords(ev.X, ev.Y)

		case *sdl.MouseWheelEvent:
			ui.playerPos.z += float32(ev.Y)
		case *sdl.KeyboardEvent:
			if ev.State == sdl.PRESSED {
				switch ev.Keysym.Sym {
				case sdl.K_LEFT:
					ui.playerPos.x -= 4
				case sdl.K_RIGHT:
					ui.playerPos.x += 4
				case sdl.K_UP:
					ui.playerPos.y -= 4
				case sdl.K_DOWN:
					ui.playerPos.y += 4
				case sdl.K_ESCAPE:
					ui.quit = true
				}
			}
			if ev.State == sdl.RELEASED {

			}
		}
	}
}
