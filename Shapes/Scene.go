package main

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	width  int32
	height int32

	textures map[string]uint32
	shapes   map[string]*Shape

	window  *sdl.Window
	context sdl.GLContext
}

func (s *Scene) Setup(title string, w, h int32) {

	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		panic(err)
	}
	//defer sdl.Quit()

	window, err := sdl.CreateWindow(title, int32(sdl.WINDOWPOS_UNDEFINED), int32(sdl.WINDOWPOS_UNDEFINED), width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	s.window = window

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_BUFFER_SIZE, 32)
	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 16)
	sdl.GLSetAttribute(sdl.GL_ACCELERATED_VISUAL, 1)
	sdl.GLSetSwapInterval(-1)

	context, err := s.window.GLCreateContext()
	s.context = context
	if err != nil {
		panic(err)
	}
	s.window.GLMakeCurrent(context)

	gl.Init()
	gl.Viewport(0, 0, w, h)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.CULL_FACE)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := float64(width)/height - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	s.width = w
	s.height = h

}

func (s *Scene) Quit() {
	sdl.GLDeleteContext(s.context)
	s.window.Destroy()
	sdl.Quit()
}

func (s *Scene) AddShape(name string, shape ShapeType, width, depth, height float32, position, rotation vec3, edges, col uint32, textureFile string) {
	newshape := NewShape(name, shape, width, depth, height, position, rotation, edges, col, textureFile)
	if s.shapes == nil {
		s.shapes = make(map[string]*Shape)
	}
	s.shapes[name] = &newshape
}

func (s *Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, shape := range s.shapes {
		shape.Draw()
	}
}

func (s *Scene) Shape(name string) *Shape {
	shape, ok := s.shapes[name]
	if ok {
		return shape
	}
	log.Println("Shape '" + name + "' does not exist. Ignoring")
	return &Shape{}
}
