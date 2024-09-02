package goengine

import (
	"fmt"
	"testing"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var context sdl.GLContext

func TestGetAttrib(t *testing.T) {
	vs := OpenTextFile("Resources/vs.txt")
	attributes, err := GetAttributes(vs)
	fmt.Printf("Attributes: %v\n", attributes)
	if err != "" {
		t.Error(err)
	}
}

func TestLoadShader(t *testing.T) {
	initSDL()

	//Setup vertex shader
	vertexSrc := OpenTextFile("Resources/vs.txt")
	vertexShader, err := LoadShaderStr(gl.VERTEX_SHADER, vertexSrc)
	fmt.Printf("Vertex shader ref: %v\n", vertexShader)
	if err != "" {
		t.Error(err)
	}

	//Setup fragment shader
	fragShader, err := LoadShaderStr(gl.FRAGMENT_SHADER, OpenTextFile("Resources/fs.txt"))
	fmt.Printf("Fragment shader ref: %v\n", fragShader)
	if err != "" {
		t.Error(err)
	}

	//Read vertex attributes vars
	attributes, err := GetAttributes(vertexSrc)
	fmt.Printf("Attributes: %v\n", attributes)
	if err != "" {
		t.Error(err)
	}

	//Create and link shader program
	program, err := CreateShaderProgram(vertexShader, fragShader, attributes)
	if err != "" {
		t.Error(err)
	}

	//Fetch all the uniform references from the shader
	settings := ShaderSettings{}
	refs := settings.SetupShaderSettings(program)

	for i, ref := range refs {
		fmt.Printf("Shader uniform ref: %v = %v\n", i, ref)
	}

	sdlquit()
}

func TestGLVersion(t *testing.T) {
	initSDL()
	fmt.Printf("OpenGL version: %v\n", GLVersion())
	sdlquit()
}

func initSDL() {
	var err error
	//runtime.LockOSThread()
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	//defer sdl.Quit()
	window, err = sdl.CreateWindow("TestSDL", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_OPENGL) //|sdl.WINDOW_HIDDEN
	if err != nil {
		panic(err)
	}
	//defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	//defer sdl.GLDeleteContext(context)
	if err = gl.Init(); err != nil {
		panic(err)
	}
}

func sdlquit() {
	sdl.GLDeleteContext(context)
	window.Destroy()
	sdl.Quit()
}
