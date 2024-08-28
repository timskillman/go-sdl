package goengine

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
	Mouse     Mouse
	Quit      bool
	PlayerPos Vec3
}

func (ui *UserInput) GetUserInput() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.QuitEvent:
			ui.Quit = true
		case *sdl.MouseMotionEvent:
			ui.Mouse.LeftButton = (ev.State == sdl.BUTTON_LEFT)
			ui.Mouse.RightButton = (ev.State == sdl.BUTTON_RIGHT)
			ui.Mouse.MiddleButton = (ev.State == sdl.BUTTON_MIDDLE)

			if ui.Mouse.LeftButton {
				x, y := ui.Mouse.GetDelta(ev.X, ev.Y)
				ui.PlayerPos.X -= float32(x)
				ui.PlayerPos.Y += float32(y)
			}
			ui.Mouse.SetCoords(ev.X, ev.Y)

		case *sdl.MouseWheelEvent:
			ui.PlayerPos.Z += float32(ev.Y)
		case *sdl.KeyboardEvent:
			if ev.State == sdl.PRESSED {
				switch ev.Keysym.Sym {
				case sdl.K_LEFT:
					ui.PlayerPos.X -= 4
				case sdl.K_RIGHT:
					ui.PlayerPos.X += 4
				case sdl.K_UP:
					ui.PlayerPos.Y -= 4
				case sdl.K_DOWN:
					ui.PlayerPos.Y += 4
				case sdl.K_ESCAPE:
					ui.Quit = true
				}
			}
			if ev.State == sdl.RELEASED {

			}
		}
	}
}
