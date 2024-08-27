package main

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/gl/v2.1/gl"
)

type ShapeType int

const (
	cuboidShape ShapeType = iota
	planeShape
	sphereShape
	cylinderShape
	coneShape
	tconeShape
	tubeShape
	torusShape
	latheShape
	extrudeShape
)

type Shape struct {
	name      string
	shapeType ShapeType
	position  vec3
	rotation  vec3
	scale     vec3
	center    vec3
	colour    uint32
	texture   Texture
	w         float32
	h         float32
	d         float32
	path      []vec2
}

func NewShape(name string, shape ShapeType, width, depth, height float32, position, rotation vec3, col uint32, textureFile string) Shape {

	tex := Texture{}
	if textureFile != "" {
		tex.LoadTexture(textureFile)
	}

	return Shape{
		name:      name,
		shapeType: shape,
		w:         width,
		h:         height,
		d:         depth,
		position:  position,
		rotation:  rotation,
		colour:    col,
		texture:   tex,
	}
}

func (s *Shape) Draw() {

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Translatef(s.position.x, s.position.y, s.position.z)

	if s.rotation.x != 0 {
		gl.Rotatef(s.rotation.x, 1, 0, 0)
	}
	if s.rotation.y != 0 {
		gl.Rotatef(s.rotation.y, 0, 1, 0)
	}
	if s.rotation.z != 0 {
		gl.Rotatef(s.rotation.z, 0, 0, 1)
	}

	gl.Color4f(float32(s.colour&255)/255, float32((s.colour>>8)&255)/255, float32((s.colour>>16)&255)/255, float32((s.colour>>24)&255)/255)

	gl.BindTexture(gl.TEXTURE_2D, uint32(s.texture.id))

	verts := []float32{}
	sides := int(s.h)

	switch s.shapeType {
	case cuboidShape:
		DrawVerts2(s.CreateCuboid())
		return
	case planeShape:
		DrawVerts2(s.CreatePlane())
		return
	case sphereShape:
		verts = s.CreateSphere(s.w, 0, sides, false)
	case cylinderShape:
		verts = s.CreateTCone(s.w, s.w, s.d, sides)
	case coneShape:
		verts = s.CreateTCone(0.001, s.w, s.d, sides)
	case tconeShape:
		sides = 32
		verts = s.CreateTCone(s.w, s.d, s.h, sides)
	case tubeShape:
		sides = 32
		verts = s.CreateTube(s.w, s.d, s.h, sides)
	case torusShape:
		verts = s.CreateTorus(s.w, s.d, sides, sides)
	case latheShape:
	case extrudeShape:
	}
	DrawVerts(verts, sides)
}

func DrawVerts2(verts []float32) {
	sz := len(verts) / (9 * 4)
	gl.Begin(gl.QUADS)
	i := 0
	vstep := 9
	nextlevel := 4 * vstep
	for p := 0; p < sz; p++ {
		drawQuad(verts, i)
		drawQuad(verts, i+vstep)
		drawQuad(verts, i+vstep*2)
		drawQuad(verts, i+vstep*3)
		i += nextlevel
	}
	gl.End()
}

func DrawVerts(verts []float32, edges int) {
	sz := (len(verts) / 9) / (edges + 1)
	gl.Begin(gl.QUADS)
	i := 0
	vstep := 9
	nextLevel := int(edges) * vstep

	for p := 0; p < sz-1; p++ {
		for r := 0; r <= edges; r++ {
			drawQuad(verts, i)
			drawQuad(verts, i+vstep)
			drawQuad(verts, i+vstep+nextLevel)
			drawQuad(verts, i+nextLevel)
			i += vstep
		}
	}
	gl.End()
}

func drawQuad(verts []float32, i int) {
	gl.Normal3f(verts[i+4], verts[i+5], verts[i+6])
	gl.TexCoord2f(verts[i+7], verts[i+8])
	gl.Vertex3f(verts[i+1], verts[i+2], verts[i+3])
}

func (c *Shape) CreatePlane() []float32 {
	verts := []float32{}
	verts = append(verts, storeVNTC2(0, vec3{-c.w, -c.h, 0}, vec3{0, 0, 1}, vec2{0, 0})...)
	verts = append(verts, storeVNTC2(1, vec3{c.w, -c.h, 0}, vec3{0, 0, 1}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(2, vec3{c.w, c.h, 0}, vec3{0, 0, 1}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(3, vec3{-c.w, c.h, 0}, vec3{0, 0, 1}, vec2{0, 1})...)
	return verts
}

func (c *Shape) CreateCuboid() []float32 {
	verts := []float32{}
	verts = append(verts, storeVNTC2(0, vec3{-c.w, -c.h, c.d}, vec3{0, 0, 1}, vec2{0, 0})...)
	verts = append(verts, storeVNTC2(1, vec3{c.w, -c.h, c.d}, vec3{0, 0, 1}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(2, vec3{c.w, c.h, c.d}, vec3{0, 0, 1}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(3, vec3{-c.w, c.h, c.d}, vec3{0, 0, 1}, vec2{0, 1})...)

	verts = append(verts, storeVNTC2(4, vec3{-c.w, -c.h, -c.d}, vec3{0, 0, -1}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(5, vec3{-c.w, c.h, -c.d}, vec3{0, 0, -1}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(6, vec3{c.w, c.h, -c.d}, vec3{0, 0, -1}, vec2{0, 1})...)
	verts = append(verts, storeVNTC2(7, vec3{c.w, -c.h, -c.d}, vec3{0, 0, -1}, vec2{0, 0})...)

	verts = append(verts, storeVNTC2(8, vec3{-c.w, c.h, -c.d}, vec3{0, 1, 0}, vec2{0, 1})...)
	verts = append(verts, storeVNTC2(9, vec3{-c.w, c.h, c.d}, vec3{0, 1, 0}, vec2{0, 0})...)
	verts = append(verts, storeVNTC2(10, vec3{c.w, c.h, c.d}, vec3{0, 1, 0}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(11, vec3{c.w, c.h, -c.d}, vec3{0, 1, 0}, vec2{1, 1})...)

	verts = append(verts, storeVNTC2(8, vec3{-c.w, -c.h, -c.d}, vec3{0, -1, 0}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(9, vec3{c.w, -c.h, -c.d}, vec3{0, -1, 0}, vec2{0, 1})...)
	verts = append(verts, storeVNTC2(10, vec3{c.w, -c.h, c.d}, vec3{0, -1, 0}, vec2{0, 0})...)
	verts = append(verts, storeVNTC2(11, vec3{-c.w, -c.h, c.d}, vec3{0, -1, 0}, vec2{1, 0})...)

	verts = append(verts, storeVNTC2(8, vec3{c.w, -c.h, -c.d}, vec3{1, 0, 0}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(9, vec3{c.w, c.h, -c.d}, vec3{1, 0, 0}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(10, vec3{c.w, c.h, c.d}, vec3{1, 0, 0}, vec2{0, 1})...)
	verts = append(verts, storeVNTC2(11, vec3{c.w, -c.h, c.d}, vec3{1, 0, 0}, vec2{0, 0})...)

	verts = append(verts, storeVNTC2(8, vec3{-c.w, -c.h, -c.d}, vec3{-1, 0, 0}, vec2{0, 0})...)
	verts = append(verts, storeVNTC2(9, vec3{-c.w, -c.h, c.d}, vec3{-1, 0, 0}, vec2{1, 0})...)
	verts = append(verts, storeVNTC2(10, vec3{-c.w, c.h, c.d}, vec3{-1, 0, 0}, vec2{1, 1})...)
	verts = append(verts, storeVNTC2(11, vec3{-c.w, c.h, -c.d}, vec3{-1, 0, 0}, vec2{0, 1})...)

	return verts
}

// func (c *Shape) DrawCuboid() {

// 	gl.Begin(gl.QUADS)

// 	gl.Normal3f(0, 0, 1)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(-c.w, -c.h, c.d)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(c.w, -c.h, c.d)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(c.w, c.h, c.d)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(-c.w, c.h, c.d)

// 	gl.Normal3f(0, 0, -1)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(-c.w, -c.h, -c.d)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(-c.w, c.h, -c.d)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(c.w, c.h, -c.d)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(c.w, -c.h, -c.d)

// 	gl.Normal3f(0, 1, 0)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(-c.w, c.h, -c.d)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(-c.w, c.h, c.d)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(c.w, c.h, c.d)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(c.w, c.h, -c.d)

// 	gl.Normal3f(0, -1, 0)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(-c.w, -c.h, -c.d)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(c.w, -c.h, -c.d)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(c.w, -c.h, c.d)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(-c.w, -c.h, c.d)

// 	gl.Normal3f(1, 0, 0)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(c.w, -c.h, -c.d)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(c.w, c.h, -c.d)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(c.w, c.h, c.d)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(c.w, -c.h, c.d)

// 	gl.Normal3f(-1, 0, 0)
// 	gl.TexCoord2f(0, 0)
// 	gl.Vertex3f(-c.w, -c.h, -c.d)
// 	gl.TexCoord2f(1, 0)
// 	gl.Vertex3f(-c.w, -c.h, c.d)
// 	gl.TexCoord2f(1, 1)
// 	gl.Vertex3f(-c.w, c.h, c.d)
// 	gl.TexCoord2f(0, 1)
// 	gl.Vertex3f(-c.w, c.h, -c.d)

// 	gl.End()
// }

func (c *Shape) CreateTube(innerRadius, outerRadius, height float32, sides int) []float32 {
	c.path = make([]vec2, 5)
	c.path[0] = vec2{innerRadius, height / 2}
	c.path[1] = vec2{outerRadius, height / 2}
	c.path[2] = vec2{outerRadius, -height / 2}
	c.path[3] = vec2{innerRadius, -height / 2}
	c.path[4] = vec2{innerRadius, height / 2}
	return CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(sides), 0, vec3{0, 0, 0})
}

func (c *Shape) CreateTCone(radius, radius2, height float32, sides int) []float32 {
	c.path = make([]vec2, 4)
	c.path[0] = vec2{0.001, height / 2}
	c.path[1] = vec2{radius, height / 2}
	c.path[2] = vec2{radius2, -height / 2}
	c.path[3] = vec2{0.001, -height / 2}
	return CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(sides), 0, vec3{0, 0, 0})
}

func (c *Shape) CreateTorus(radius, ringradius float32, ringdivs, sides int) []float32 {
	c.path = make([]vec2, ringdivs+1)
	st := (math32.Pi * 2) / float32(ringdivs)
	for r := 0; r <= ringdivs; r++ {
		c.path[r] = vec2{radius + ringradius*math32.Sin(float32(r)*st), ringradius * math32.Cos(float32(r)*st)}
	}
	return CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(sides), 0, vec3{0, 0, 0})
}

func (c *Shape) CreateSphere(radius, hemi float32, edges int, invert bool) []float32 {
	c.path = make([]vec2, edges+1)
	st := (math32.Pi * (1 - hemi)) / float32(edges)
	ho := math32.Pi - (math32.Pi * (1 - hemi))
	inv := float32(1)
	if invert {
		inv = -1
	}
	for r := 0; r <= edges; r++ {
		c.path[r] = vec2{radius * math32.Sin(float32(r)*st+ho) * inv, radius * math32.Cos(float32(r)*st+ho)}
	}
	return CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(edges), 0, vec3{0, 0, 0})
}

func CreateLathe(lpath []vec2, inverted, startAngle, endAngle, rise float32, edges, uvtype uint32, pos vec3) []float32 {

	normals, path := calcPathNormals(lpath, 0.5, true, inverted)

	sz := len(path)

	angDiff := endAngle - startAngle
	angStep := angDiff / float32(edges)

	tcx, tcy, rdiv := 1/float32(edges), float32(0), rise/float32(edges)

	miny, maxy := path[0].y, path[0].y
	for p := 0; p < sz; p++ {
		if path[p].y < miny {
			miny = path[p].y
		}
		if path[p].y > maxy {
			maxy = path[p].y
		}
	}
	cy := (maxy + miny) / 2
	tc := 0

	verts := []float32{}
	for p := 0; p < sz; p++ {
		switch uvtype {
		case 0: //cylinder map
			tcy = 1 - ((path[p].y - miny) / (maxy - miny))
		case 1: //sphere map
			tcy = math32.Atan2(path[p].x, path[p].y-cy) / math32.Pi
			if tcy < 0 {
				tcy += 1
			}
		}

		risey := path[p].y

		for r := 0; r < int(edges); r++ {
			verts = append(verts, storeVNTC(p, tc, startAngle+float32(r)*angStep, risey, vec2{tcx * float32(r), tcy}, pos, path, normals)...)
			risey += rdiv
		}
		verts = append(verts, storeVNTC(p, tc, startAngle, risey, vec2{0.9999, tcy}, pos, path, normals)...)
	}

	return verts
}

func storeVNTC(p, tc int, ang, risey float32, uv vec2, pos vec3, path, normals []vec2) []float32 {
	sinr, cosr := math32.Sin(ang), math32.Cos(ang)
	v := vec3{pos.x + path[p].x*sinr, pos.y + risey, pos.z + path[p].x*cosr}
	n := vec3{normals[p].x * sinr, normals[p].y, normals[p].x * cosr}
	return []float32{float32(tc), v.x, v.y, v.z, n.x, n.y, n.z, uv.x, uv.y}
}

func storeVNTC2(tc int, pos, normal vec3, uv vec2) []float32 {
	return []float32{float32(tc), pos.x, pos.y, pos.z, normal.x, normal.y, normal.z, uv.x, uv.y}
}

func calcPathNormals(path []vec2, creaseAngle float32, joined bool, inverted float32) ([]vec2, []vec2) {
	sz := len(path)
	if sz < 2 {
		return nil, nil
	}

	p1 := vec2{}
	p2 := path[0]
	p3 := path[1]

	newPath := []vec2{}
	normals := []vec2{}

	newPath = append(newPath, p2)

	if joined && sz > 2 {
		p1 = path[sz-2]
		if angleBetween(p2.Minus(p1), p3.Minus(p2)) > creaseAngle {
			normals = append(normals, dotInvert(p2, p3, inverted))
		} else {
			normals = append(normals, dotInvert(p1.Minus(p2), p3.Minus(p2), inverted))
		}
		p1 = path[0]
	} else {
		p1 = path[0]
		normals = append(normals, dotInvert(p2, p3, inverted))
	}

	if sz > 2 {
		for i := 1; i < sz-1; i++ {
			p2, p3 := path[i], path[i+1]
			newPath = append(newPath, p2)
			if angleBetween(p2.Minus(p1), p3.Minus(p2)) > creaseAngle {
				normals = append(normals, dotInvert(p1, p2, inverted))
				newPath = append(newPath, p2)
				normals = append(normals, dotInvert(p2, p3, inverted))
			} else {
				normals = append(normals, dotInvert(p1.Minus(p2), p3.Minus(p2), inverted))
			}
			p1 = p2
		}

		p2 = path[sz-1]
		newPath = append(newPath, p2)

		if joined {
			p3 = path[1]
			if angleBetween(p2.Minus(p1), p3.Minus(p2)) > creaseAngle {
				normals = append(normals, dotInvert(p1, p2, inverted))
			} else {
				normals = append(normals, dotInvert(p1.Minus(p2), p3.Minus(p2), inverted))
			}
		} else {
			normals = append(normals, normals[len(normals)-1])
		}

	} else {
		normals = append(normals, normals[0])
		newPath = append(newPath, p3)
	}

	return normals, newPath
}

func dotInvert(v1 vec2, v2 vec2, inverted float32) vec2 {
	dp := v1.Dot(v2)
	return vec2{dp.x * inverted, dp.y * inverted}
}

func angleBetween(v1, v2 vec2) float32 {
	prod := v1.x*v2.y - v1.y*v2.x
	psign := float32(1)
	if prod < 0 {
		psign = -1
	}
	ab := psign * math32.Acos((v1.x*v2.x+v1.y*v2.y)/(v2.Length()*v2.Length()))
	return math32.Abs(ab)
}
