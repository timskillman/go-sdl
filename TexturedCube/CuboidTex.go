package main

import (
	"github.com/go-gl/gl/v2.1/gl"
)

type CuboidTex struct {
	w       float32
	h       float32
	d       float32
	pos     vec3
	rot     vec3
	colour  uint32
	texture Texture
}

func NewCuboid(width, depth, height float32, position, rotation vec3, col uint32, textureFile string) CuboidTex {
	tex := Texture{}
	if textureFile != "" {
		tex.LoadTexture(textureFile)
	}

	cuboid := CuboidTex{
		w:       width,
		h:       depth,
		d:       height,
		pos:     position,
		rot:     rotation,
		colour:  col,
		texture: tex,
	}
	return cuboid
}

func (c *CuboidTex) Init() {
	c.w = 1
	c.h = 1
	c.d = 1
	c.pos = vec3{0, 0, 0}
	c.rot = vec3{0, 0, 0}
	c.colour = 0xffffffff
	c.texture = Texture{}
}

func (c *CuboidTex) Draw() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Translatef(c.pos.x, c.pos.y, c.pos.z)
	gl.Rotatef(c.rot.x, 1, 0, 0)
	gl.Rotatef(c.rot.y, 0, 1, 0)
	gl.Rotatef(c.rot.y, 0, 0, 1)

	gl.BindTexture(gl.TEXTURE_2D, uint32(c.texture.id))

	gl.Begin(gl.QUADS)

	gl.Color4f(float32(c.colour&255)/255, float32((c.colour>>8)&255)/255, float32((c.colour>>16)&255)/255, float32((c.colour>>24)&255)/255)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-c.w, -c.h, c.d)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(c.w, -c.h, c.d)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(c.w, c.h, c.d)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-c.w, c.h, c.d)

	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-c.w, -c.h, -c.d)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-c.w, c.h, -c.d)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(c.w, c.h, -c.d)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(c.w, -c.h, -c.d)

	gl.Normal3f(0, 1, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-c.w, c.h, -c.d)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-c.w, c.h, c.d)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(c.w, c.h, c.d)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(c.w, c.h, -c.d)

	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-c.w, -c.h, -c.d)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(c.w, -c.h, -c.d)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(c.w, -c.h, c.d)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-c.w, -c.h, c.d)

	gl.Normal3f(1, 0, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(c.w, -c.h, -c.d)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(c.w, c.h, -c.d)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(c.w, c.h, c.d)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(c.w, -c.h, c.d)

	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-c.w, -c.h, -c.d)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-c.w, -c.h, c.d)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-c.w, c.h, c.d)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-c.w, c.h, -c.d)

	gl.End()
}
