package goengine

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
)

type Mesh struct {
	Verts       []float32
	VC          uint32
	MaterialRef int
	Stride      int
	BufRef      int
	VertOffset  int
	VertSize    int
	Mode        int
}

func (m *Mesh) Init() {
	m.VC = 0
	m.MaterialRef = 0
	m.Stride = VERTSIZE
	m.BufRef = 0
	m.VertOffset = 0
	m.VertSize = 0
	m.Mode = gl.TRIANGLES
}

func (m *Mesh) AddPackedVert(pos Vec3, normal Vec3, uv Vec2, col uint32) {
	m.Verts = append(m.Verts, pos.X)
	m.Verts = append(m.Verts, pos.Y)
	m.Verts = append(m.Verts, pos.Z)
	m.Verts = append(m.Verts, normal.X)
	m.Verts = append(m.Verts, normal.Y)
	m.Verts = append(m.Verts, normal.Z)
	m.Verts = append(m.Verts, uv.X)
	m.Verts = append(m.Verts, uv.Y)
	m.Verts = append(m.Verts, convertColToFloat(col))
}

func convertColToFloat(col uint32) float32 {
	return float32(col&255)/256 + float32((col>>8)&255) + float32((col>>16)&255)*256
}

func (m *Mesh) RenderMesh() {
	SetRenderBuffer(uint32(m.BufRef), int32(m.Stride))
	gl.DrawArrays(uint32(m.Mode), int32(m.VertOffset), int32(m.VertSize))
}

func (m *Mesh) Render() {
	gl.DrawArrays(uint32(m.Mode), int32(m.VertOffset), int32(m.VertSize))
}

func (m *Mesh) RenderIndexed(indexCo int32, indexes []*uint32) {
	gl.DrawElements(uint32(m.Mode), indexCo, gl.UNSIGNED_INT, unsafe.Pointer(&indexes))
}

func (m *Mesh) TransformVerts(matrix *Mat4s) {
	for i := m.VertOffset; i < len(m.Verts) && i < (m.VertOffset+int(m.VC)); i += m.Stride {
		v := Vec3{m.Verts[i], m.Verts[i+1], m.Verts[i+2]}
		v.MulMat4(matrix)
		m.Verts[i] = v.X
		m.Verts[i+1] = v.Y
		m.Verts[i+2] = v.Z
	}
}
