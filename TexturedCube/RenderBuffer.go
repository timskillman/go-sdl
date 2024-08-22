package main

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type RenderBuffer struct {
	BufferID []uint32
}

func (b *RenderBuffer) SetRenderBuffer(buf, stride int32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.BufferID[buf])

	attribCount := uint32(0)
	pos := uint32(0)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Vertices
	gl.EnableVertexArrayAttrib(b.BufferID[buf], attribCount)
	attribCount++
	pos += 3 * 4

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Normals
	gl.EnableVertexArrayAttrib(b.BufferID[buf], attribCount)
	attribCount++
	pos += 3 * 4

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //UVs
	gl.EnableVertexArrayAttrib(b.BufferID[buf], attribCount)
	attribCount++
	pos += 2 * 4

	gl.VertexAttribPointer(0, 1, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Colour
	gl.EnableVertexArrayAttrib(b.BufferID[buf], attribCount)
	attribCount++
	pos += 1 * 4

}
