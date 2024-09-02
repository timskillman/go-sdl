package goengine

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type RenderBuffer struct {
	BufferID      []uint32
	Verts         [][]float32
	VertsPtr      []int
	MaxBufSize    uint32
	CurrentBuffer int
}

func (b *RenderBuffer) Init() {
	b.Verts = append(b.Verts, make([]float32, b.MaxBufSize))
	b.VertsPtr = append(b.VertsPtr, 0)
	b.CurrentBuffer = 0

}

func (b *RenderBuffer) RemainingBuffer(bufRef int) int {
	return len(b.Verts[bufRef]) - b.VertsPtr[bufRef]
}

func (b *RenderBuffer) AddMesh(mesh *Mesh) (int32, string) {
	meshSize := len(mesh.Verts)
	if meshSize > int(b.MaxBufSize) {
		return -1, "Cant fit mesh into render buffer - too big"
	}
	if meshSize == 0 {
		return -1, "Mesh has no vertices"
	}

	//Search for space to fit mesh in the main Verts buffer
	bufferWithSpace := -1
	if b.CurrentBuffer > 0 {
		for buf := 0; buf < len(b.Verts) && bufferWithSpace < 0; buf++ {
			if b.RemainingBuffer(buf) > meshSize {
				bufferWithSpace = buf
			}
		}
	}

	offset := 0
	if bufferWithSpace < 0 {
		//No buffer found - create a new one with enough space
		b.Verts = append(b.Verts, make([]float32, b.MaxBufSize))
		b.VertsPtr = append(b.VertsPtr, 0)
		b.CurrentBuffer++
		bufferWithSpace = b.CurrentBuffer
		copy(b.Verts[bufferWithSpace], mesh.Verts)
		gl.GenBuffers(1, &b.BufferID[b.CurrentBuffer])
		gl.BindBuffer(gl.ARRAY_BUFFER, b.BufferID[b.CurrentBuffer])
		gl.BufferData(gl.ARRAY_BUFFER, int(b.MaxBufSize*4), unsafe.Pointer(&b.Verts[bufferWithSpace][0]), gl.DYNAMIC_DRAW)
	} else {
		//Copy mesh verts into the main buffer
		offset = b.VertsPtr[bufferWithSpace]
		for v := 0; v < len(mesh.Verts); v++ {
			//b.Verts[bufferWithSpace] = append(b.Verts[bufferWithSpace], mesh.Verts) //get this to work?
			b.Verts[bufferWithSpace][offset+v] = mesh.Verts[v]
		}
		gl.BindBuffer(gl.ARRAY_BUFFER, b.BufferID[b.CurrentBuffer])
		gl.BufferSubData(gl.ARRAY_BUFFER, offset*4, len(mesh.Verts), unsafe.Pointer(&b.Verts[bufferWithSpace][0]))
	}
	b.VertsPtr[bufferWithSpace] += len(mesh.Verts)
	mesh.BufRef = bufferWithSpace
	mesh.VertOffset = offset

	return 0, ""
}

func SetRenderBuffer(bufID uint32, stride int32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, bufID)

	attribCount := uint32(0)
	pos := uint32(0)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Vertices
	gl.EnableVertexArrayAttrib(bufID, attribCount)
	attribCount++
	pos += 3 * 4

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Normals
	gl.EnableVertexArrayAttrib(bufID, attribCount)
	attribCount++
	pos += 3 * 4

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //UVs
	gl.EnableVertexArrayAttrib(bufID, attribCount)
	attribCount++
	pos += 2 * 4

	gl.VertexAttribPointer(0, 1, gl.FLOAT, false, stride*4, unsafe.Pointer(&pos)) //Colour
	gl.EnableVertexArrayAttrib(bufID, attribCount)
	attribCount++
	pos += 1 * 4

}
