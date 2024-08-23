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

	switch s.shapeType {
	case cuboidShape:
		s.DrawCuboid()
	case planeShape:
		s.DrawPlane()
	case sphereShape:
	case cylinderShape:
	case coneShape:
	case tubeShape:
	case torusShape:
	case latheShape:
	case extrudeShape:
	}
}

func (c *Shape) DrawPlane() {
	gl.Begin(gl.QUADS)
	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-c.w, -c.h, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(c.w, -c.h, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(c.w, c.h, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-c.w, c.h, 0)
	gl.End()
}

func (c *Shape) DrawCuboid() {

	gl.Begin(gl.QUADS)

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

type lathe struct {
	vertTemp []float32
}

func (c *Shape) DrawLathe(inverted, endAngle, startAngle, rise float32, edges, uvtype uint32, pos vec3) {

	gl.Begin(gl.QUADS)

	normals, path := calcPathNormals(c.path, 0.5, true, inverted)
	c.path = path

	sz := len(c.path)

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

	vertTemp := []float32{}
	tc := 0

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
		for r := uint32(0); r < edges; r++ {
			storeVNTC(p, tc, vertTemp, startAngle+float32(r)*angStep, risey, vec2{tcx * float32(r), tcy}, pos, path, normals)
			// ang := startAngle + float32(r) * angStep
			// sinr, cosr := math32.Sin(ang), math32.Cos(ang)
			// v := vec3{pos.x + path[p].x * sinr, pos.y + risey, pos.z + path[p].x * cosr}
			// n := vec3{normals[p].x * sinr, normals[p].y, normals[p].x * cosr}
			// uv := vec2{tcx * float32(r), tcy}
			// vertTemp = append(vertTemp, float32(tc))
			// vertTemp = append(vertTemp, v.x)
			// vertTemp = append(vertTemp, v.y)
			// vertTemp = append(vertTemp, v.z)
			// vertTemp = append(vertTemp, n.x)
			// vertTemp = append(vertTemp, n.y)
			// vertTemp = append(vertTemp, n.z)
			// vertTemp = append(vertTemp, uv.x)
			// vertTemp = append(vertTemp, uv.y)
			risey += rdiv
		}
		storeVNTC(p, tc, vertTemp, startAngle+float32(0)*angStep, risey, vec2{tcx * float32(0), tcy}, pos, path, normals)
	}
}

func storeVNTC(p, tc int, vertTemp []float32, ang, risey float32, uv vec2, pos vec3, path, normals []vec2) {
	sinr, cosr := math32.Sin(ang), math32.Cos(ang)
	v := vec3{pos.x + path[p].x*sinr, pos.y + risey, pos.z + path[p].x*cosr}
	n := vec3{normals[p].x * sinr, normals[p].y, normals[p].x * cosr}
	//uv := vec2{tcx , tcy}
	vertTemp = append(vertTemp, float32(tc))
	vertTemp = append(vertTemp, v.x)
	vertTemp = append(vertTemp, v.y)
	vertTemp = append(vertTemp, v.z)
	vertTemp = append(vertTemp, n.x)
	vertTemp = append(vertTemp, n.y)
	vertTemp = append(vertTemp, n.z)
	vertTemp = append(vertTemp, uv.x)
	vertTemp = append(vertTemp, uv.y)
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
