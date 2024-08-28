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
	springShape
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
	edges     uint32
	path      []vec2
	verts     []float32
	group     []Shape
}

func NewShape(name string, shape ShapeType, width, height, depth float32, position, rotation vec3, edges, col uint32, textureFile string) Shape {

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
		edges:     edges,
		colour:    col,
		texture:   tex,
		verts:     nil,
		path:      nil,
		group:     nil,
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
		DrawQuads(s.CreateCuboid())
	case planeShape:
		DrawQuads(s.CreatePlane())
	case sphereShape:
		DrawSharedQuads(s.CreateSphere(), int(s.edges))
	case cylinderShape:
		DrawSharedQuads(s.CreateCylinder(), int(s.edges))
	case coneShape:
		DrawSharedQuads(s.CreateCone(), int(s.edges))
	case tconeShape:
		DrawSharedQuads(s.CreateTCone(), int(s.edges))
	case tubeShape:
		DrawSharedQuads(s.CreateTube(), int(s.edges))
	case torusShape:
		DrawSharedQuads(s.CreateTorus(), int(s.edges))
	case springShape:
		DrawSharedQuads(s.CreateSpring(), int(s.edges))
	case latheShape:
	case extrudeShape:
	}
}

func DrawQuads(verts []float32) {
	vstep := 9
	nextQuad := 4 * vstep
	quadCount := len(verts) / nextQuad
	i := 0

	gl.Begin(gl.QUADS)
	for q := 0; q < quadCount; q++ {
		drawQuad(verts, i)
		drawQuad(verts, i+vstep)
		drawQuad(verts, i+vstep*2)
		drawQuad(verts, i+vstep*3)
		i += nextQuad
	}
	gl.End()
}

func DrawSharedQuads(verts []float32, edges int) {
	vstep := 9
	nextLevel := int(edges+1) * vstep //There's an extra edge for the quads to seamlessly join
	pathLength := len(verts) / nextLevel

	gl.Begin(gl.QUADS)
	for p := 0; p < pathLength-1; p++ {
		i := p * nextLevel
		for e := 0; e < edges; e++ {
			drawQuad(verts, i+vstep)
			drawQuad(verts, i)
			drawQuad(verts, i+nextLevel)
			drawQuad(verts, i+vstep+nextLevel)
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
	if c.verts != nil {
		return c.verts
	}
	c.verts = []float32{}
	c.verts = append(c.verts, storeVNTC2(0, vec3{-c.w, -c.h, 0}, vec3{0, 0, 1}, vec2{0, 0})...)
	c.verts = append(c.verts, storeVNTC2(1, vec3{c.w, -c.h, 0}, vec3{0, 0, 1}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(2, vec3{c.w, c.h, 0}, vec3{0, 0, 1}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(3, vec3{-c.w, c.h, 0}, vec3{0, 0, 1}, vec2{0, 1})...)
	return c.verts
}

func (c *Shape) CreateCuboid() []float32 {
	if c.verts != nil {
		return c.verts
	}
	c.verts = []float32{}
	c.verts = append(c.verts, storeVNTC2(0, vec3{-c.w, -c.h, c.d}, vec3{0, 0, 1}, vec2{0, 0})...)
	c.verts = append(c.verts, storeVNTC2(1, vec3{c.w, -c.h, c.d}, vec3{0, 0, 1}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(2, vec3{c.w, c.h, c.d}, vec3{0, 0, 1}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(3, vec3{-c.w, c.h, c.d}, vec3{0, 0, 1}, vec2{0, 1})...)

	c.verts = append(c.verts, storeVNTC2(4, vec3{-c.w, -c.h, -c.d}, vec3{0, 0, -1}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(5, vec3{-c.w, c.h, -c.d}, vec3{0, 0, -1}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(6, vec3{c.w, c.h, -c.d}, vec3{0, 0, -1}, vec2{0, 1})...)
	c.verts = append(c.verts, storeVNTC2(7, vec3{c.w, -c.h, -c.d}, vec3{0, 0, -1}, vec2{0, 0})...)

	c.verts = append(c.verts, storeVNTC2(8, vec3{-c.w, c.h, -c.d}, vec3{0, 1, 0}, vec2{0, 1})...)
	c.verts = append(c.verts, storeVNTC2(9, vec3{-c.w, c.h, c.d}, vec3{0, 1, 0}, vec2{0, 0})...)
	c.verts = append(c.verts, storeVNTC2(10, vec3{c.w, c.h, c.d}, vec3{0, 1, 0}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(11, vec3{c.w, c.h, -c.d}, vec3{0, 1, 0}, vec2{1, 1})...)

	c.verts = append(c.verts, storeVNTC2(8, vec3{-c.w, -c.h, -c.d}, vec3{0, -1, 0}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(9, vec3{c.w, -c.h, -c.d}, vec3{0, -1, 0}, vec2{0, 1})...)
	c.verts = append(c.verts, storeVNTC2(10, vec3{c.w, -c.h, c.d}, vec3{0, -1, 0}, vec2{0, 0})...)
	c.verts = append(c.verts, storeVNTC2(11, vec3{-c.w, -c.h, c.d}, vec3{0, -1, 0}, vec2{1, 0})...)

	c.verts = append(c.verts, storeVNTC2(8, vec3{c.w, -c.h, -c.d}, vec3{1, 0, 0}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(9, vec3{c.w, c.h, -c.d}, vec3{1, 0, 0}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(10, vec3{c.w, c.h, c.d}, vec3{1, 0, 0}, vec2{0, 1})...)
	c.verts = append(c.verts, storeVNTC2(11, vec3{c.w, -c.h, c.d}, vec3{1, 0, 0}, vec2{0, 0})...)

	c.verts = append(c.verts, storeVNTC2(8, vec3{-c.w, -c.h, -c.d}, vec3{-1, 0, 0}, vec2{0, 0})...)
	c.verts = append(c.verts, storeVNTC2(9, vec3{-c.w, -c.h, c.d}, vec3{-1, 0, 0}, vec2{1, 0})...)
	c.verts = append(c.verts, storeVNTC2(10, vec3{-c.w, c.h, c.d}, vec3{-1, 0, 0}, vec2{1, 1})...)
	c.verts = append(c.verts, storeVNTC2(11, vec3{-c.w, c.h, -c.d}, vec3{-1, 0, 0}, vec2{0, 1})...)

	return c.verts
}

func (c *Shape) CreateTube() []float32 {
	if c.verts != nil {
		return c.verts
	}
	c.path = make([]vec2, 5)
	c.path[0] = vec2{c.w, c.d / 2}
	c.path[1] = vec2{c.h, c.d / 2}
	c.path[2] = vec2{c.h, -c.d / 2}
	c.path[3] = vec2{c.w, -c.d / 2}
	c.path[4] = vec2{c.w, c.d / 2}
	c.verts = CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(c.edges), 0, vec3{0, 0, 0})
	return c.verts
}

func (c *Shape) CreateCylinder() []float32 {
	if c.verts != nil {
		return c.verts
	}
	return CreateVCone(c.w, c.w, c.h, int(c.edges))
}

func (c *Shape) CreateCone() []float32 {
	if c.verts != nil {
		return c.verts
	}
	return CreateVCone(0.01, c.w, c.h, int(c.edges))
}

func (c *Shape) CreateTCone() []float32 {
	if c.verts != nil {
		return c.verts
	}
	return CreateVCone(c.w, c.d, c.h, int(c.edges))
}

func CreateVCone(radius, radius2, height float32, sides int) []float32 {
	path := make([]vec2, 4)
	path[0] = vec2{0.01, height / 2}
	path[1] = vec2{radius, height / 2}
	path[2] = vec2{radius2, -height / 2}
	path[3] = vec2{0.01, -height / 2}
	return CreateLathe(path, 1, 0, 2*math32.Pi, 0, uint32(sides), 0, vec3{0, 0, 0})
}

func (c *Shape) CreateTorus() []float32 {
	if c.verts != nil {
		return c.verts
	}
	radius := c.w
	ringradius := c.h
	ringdivs := int(c.d)
	c.path = make([]vec2, ringdivs+1)
	st := (math32.Pi * 2) / float32(ringdivs)
	for r := 0; r <= ringdivs; r++ {
		c.path[r] = vec2{radius + ringradius*math32.Sin(float32(r)*st), ringradius * math32.Cos(float32(r)*st)}
	}
	c.verts = CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(c.edges), 0, vec3{0, 0, 0})
	return c.verts
}

func (c *Shape) CreateSphere() []float32 {
	if c.verts != nil {
		return c.verts
	}
	radius := c.w
	hemi := c.h
	c.path = make([]vec2, c.edges+1)
	st := (math32.Pi * (1 - hemi)) / float32(c.edges)
	ho := math32.Pi - (math32.Pi * (1 - hemi))

	for r := 0; r <= int(c.edges); r++ {
		c.path[r] = vec2{radius * math32.Sin(float32(r)*st+ho) * sign(c.d), radius * math32.Cos(float32(r)*st+ho)}
	}
	c.verts = CreateLathe(c.path, 1, 0, 2*math32.Pi, 0, uint32(c.edges), 0, vec3{0, 0, 0})
	return c.verts
}

func (c *Shape) CreateSpring() []float32 {
	if c.verts != nil {
		return c.verts
	}
	radius := c.w
	ringradius := c.h
	ringdivs := 12 //int(c.d)
	c.path = make([]vec2, ringdivs+1)
	st := (math32.Pi * 2) / float32(ringdivs)
	for r := 0; r <= ringdivs; r++ {
		c.path[r] = vec2{radius + ringradius*math32.Sin(float32(r)*st), ringradius * math32.Cos(float32(r)*st)}
	}
	c.verts = CreateLathe(c.path, 1, 0, 2*math32.Pi*10, c.d, uint32(c.edges), 0, vec3{0, 0, 0})
	return c.verts
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
	ab := sign(prod) * math32.Acos((v1.x*v2.x+v1.y*v2.y)/(v2.Length()*v2.Length()))
	return math32.Abs(ab)
}

func sign(v float32) float32 {
	if v < 0 {
		return -1
	}
	return 1
}
