package goengine

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

func (m *Mesh) AddPackedVert(pos Vec3, normal Vec3, uv Vec2, col uint32) {
	m.verts = append(m.verts, pos.X)
	m.verts = append(m.verts, pos.Y)
	m.verts = append(m.verts, pos.Z)
	m.verts = append(m.verts, normal.X)
	m.verts = append(m.verts, normal.Y)
	m.verts = append(m.verts, normal.Z)
	m.verts = append(m.verts, uv.X)
	m.verts = append(m.verts, uv.Y)
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
