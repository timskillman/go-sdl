# go-sdl
Experiments with Golang and SDL2

## Why?

Very simply I wanted to get SDL2 up and running using Go and found it wasn't straight forward (Aug '24) using the following setup;

- Windows
- VSCode
- C++ support (for compiling C programs for Go)
  
Windows and VSCode probably isn't the best environment for writing Go-SDL programs but it's probably the most accessible.
On linux this appear to be easier as gcc is part of the Linux OS.

Before I begin I want to thank [Ve & Co](https://github.com/veandco) for their amazing work on their [go-sdl2](https://github.com/veandco/go-sdl2) Go bindings and examples

One more note, you may be asking why add C support? There are 1000's of C functions that can be used for Go, one of which I aim to implement; the [Glutess library](https://github.com/mlabbe/glutess/tree/master).

Here are some basic, step-by-step instructions on how I got SDL2 running in Go with VSCode with C support ...

## Setup on Windows 64 bit

- Install VSCode from https://code.visualstudio.com/download
- Download GO language from https://go.dev/dl/
- Install MinGW-W64 installer (easiest method) from https://github.com/nixman/mingw-builds-binaries
    - Install to 'C:/' root path with the 64bit version
    - Check that 'C:/mingw64' exists before continuing
- Set the PATH variable
    - Go to Windows search bar and find 'environment variables' (then select Environment Variables button from the window)
    - Select 'Path' and 'Edit'
    - Add 'C:\mingw64\bin' as a new path variable and save
    - Open a new command prompt from the search bar (type 'cmd' and enter)
    - Check that gcc, g++ and gdb exist by adding '--version' after them e.g. 'gcc --version'
      This should return the version number of each to confirm that the C compilers are working
- Download SDL2-devel-2.30.6.mingw.zip (or latest version) and unzip 
    - Drag the x86_64-w64-mingw32 and i686-w64-mingw32 folders into the 'C:/mingw64' folder
- Start VSCode
  - Add extensions (4x cube icon on left bar)
    - Install Go extension
    - Install C/C++ 

## Setup on Linux

- Install VSCode from https://code.visualstudio.com/download
- Install Go from a terminal window;
  ```
  $ sudo snap install go --classic
  ```
- Install SDL2
  ```
  $ sudo apt-get install libsdl2-dev
  ```

## Write a basic SDL program in Go 
(based on [Ve & Co's examples](https://github.com/veandco/go-sdl2-examples/tree/master/examples))

- Create a 'go' folder in your 'user/\<name\>' folder
- Create an 'SDLtest' folder in the 'go' folder
- Open the 'SDLtest' folder in VSCode (File \> Open Folder...)
- Create an 'SDLtest.go' file in the 'SDLtest' folder with the following contents ...
  
```Go
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

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// Draw initial frame
	draw(window, surface)

	// Event loop

	for event := sdl.WaitEvent(); event != nil; event = sdl.WaitEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
			println("Quitting..")
			return
		case *sdl.KeyboardEvent:
			if t.State == sdl.PRESSED {
				if t.Keysym.Sym == sdl.K_LEFT {
					playerX -= 4
				} else if t.Keysym.Sym == sdl.K_RIGHT {
					playerX += 4
				}
				if t.Keysym.Sym == sdl.K_UP {
					playerY -= 4
				} else if t.Keysym.Sym == sdl.K_DOWN {
					playerY += 4
				}
				if playerX < 0 {
					playerX = WIDTH
				} else if playerX > WIDTH {
					playerX = 0
				}
				if playerY < 0 {
					playerY = HEIGHT
				} else if playerY > HEIGHT {
					playerY = 0
				}
				draw(window, surface)
			}
		}
	}

}

func draw(window *sdl.Window, surface *sdl.Surface) {
	// Clear surface
	surface.FillRect(nil, 0)

	// Draw on the surface
	rect := sdl.Rect{int32(playerX), int32(playerY), 40, 40}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

	window.UpdateSurface()
}
```

## Running the programs

- **Windows Only** Before you can the program, you will need the **SDL2.dll** found in the 'SDL-devel-2.30.6-mingw.zip' file you downloaded earlier.
  Drop the **SDL2.dll** into your *SDLtest* folder otherwise your program won't work.
  Alternatively drop SDL2.dll into your *Windows/System32* folder so it's always accessible

  Note that installing SDL2 on Linux will run the program with no issues
  
- Now 'Run & Debug' the program - there may be a few other things VSCode wants and it may take a while on first run.

  You should see a purple square in an SDL2 window.  Press the cursor keys to move it around
