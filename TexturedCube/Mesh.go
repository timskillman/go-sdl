package main

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Mesh struct {
	verts       []float32
	vc          uint32
	materialRef int
	stride      int
	bufRef      int
	vertOffset  int
	vertSize    int
	mode        int
}

func (m *Mesh) Init() {
	m.vc = 0
	m.materialRef = 0
	m.stride = 9
	m.bufRef = 0
	m.vertOffset = 0
	m.vertSize = 0
	m.mode = gl.TRIANGLES
}

func (m *Mesh) AddPackedVert(pos vec3, normal vec3, uv vec2, col uint32) {
	m.verts = append(m.verts, pos.x)
	m.verts = append(m.verts, pos.y)
	m.verts = append(m.verts, pos.z)
	m.verts = append(m.verts, normal.x)
	m.verts = append(m.verts, normal.y)
	m.verts = append(m.verts, normal.z)
	m.verts = append(m.verts, uv.x)
	m.verts = append(m.verts, uv.y)
	m.verts = append(m.verts, convertColToFloat(col))
}

func convertColToFloat(col uint32) float32 {
	return float32(col&255)/256 + float32((col>>8)&255) + float32((col>>16)&255)*256
}

func (m *Mesh) Render() {
	gl.DrawArrays(uint32(m.mode), int32(m.vertOffset), int32(m.vertSize))
}

func (m *Mesh) RenderIndexed(indexCo int32, indexes []*uint32) {
	gl.DrawElements(uint32(m.mode), indexCo, gl.UNSIGNED_INT, unsafe.Pointer(&indexes))
}
