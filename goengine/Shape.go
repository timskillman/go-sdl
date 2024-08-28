package goengine

import (
	"github.com/chewxy/math32"
	"github.com/go-gl/gl/v2.1/gl"
)

type ShapeType int

const (
	ShapeCuboid ShapeType = iota
	ShapePlane
	ShapeSphere
	ShapeCylinder
	ShapeCone
	ShapeTCone
	ShapeTube
	ShapeTorus
	ShapeLathe
	ShapeExtrude
	ShapeSpring
)

type Shape struct {
	Name      string
	ShapeType ShapeType
	Position  Vec3
	Rotation  Vec3
	Scale     Vec3
	Center    Vec3
	Colour    uint32
	Texture   Texture
	W         float32
	H         float32
	D         float32
	Edges     uint32
	Path      []Vec2
	Verts     []float32
	Group     []Shape
}

func NewShape(name string, shape ShapeType, width, height, depth float32, position, rotation Vec3, edges, col uint32, textureFile string) Shape {

	tex := Texture{}
	if textureFile != "" {
		tex.LoadTexture(textureFile)
	}

	return Shape{
		Name:      name,
		ShapeType: shape,
		W:         width,
		H:         height,
		D:         depth,
		Position:  position,
		Rotation:  rotation,
		Edges:     edges,
		Colour:    col,
		Texture:   tex,
		Verts:     nil,
		Path:      nil,
		Group:     nil,
	}
}

func (s *Shape) Draw() {

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Translatef(s.Position.X, s.Position.Y, s.Position.Z)

	if s.Rotation.X != 0 {
		gl.Rotatef(s.Rotation.X, 1, 0, 0)
	}
	if s.Rotation.Y != 0 {
		gl.Rotatef(s.Rotation.Y, 0, 1, 0)
	}
	if s.Rotation.Z != 0 {
		gl.Rotatef(s.Rotation.Z, 0, 0, 1)
	}

	gl.Color4f(float32(s.Colour&255)/255, float32((s.Colour>>8)&255)/255, float32((s.Colour>>16)&255)/255, float32((s.Colour>>24)&255)/255)

	gl.BindTexture(gl.TEXTURE_2D, uint32(s.Texture.id))

	switch s.ShapeType {
	case ShapeCuboid:
		DrawQuads(s.CreateCuboid())
	case ShapePlane:
		DrawQuads(s.CreatePlane())
	case ShapeSphere:
		DrawSharedQuads(s.CreateSphere(), int(s.Edges))
	case ShapeCylinder:
		DrawSharedQuads(s.CreateCylinder(), int(s.Edges))
	case ShapeCone:
		DrawSharedQuads(s.CreateCone(), int(s.Edges))
	case ShapeTCone:
		DrawSharedQuads(s.CreateTCone(), int(s.Edges))
	case ShapeTube:
		DrawSharedQuads(s.CreateTube(), int(s.Edges))
	case ShapeTorus:
		DrawSharedQuads(s.CreateTorus(), int(s.Edges))
	case ShapeSpring:
		DrawSharedQuads(s.CreateSpring(), int(s.Edges))
	case ShapeLathe:
	case ShapeExtrude:
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
	if c.Verts != nil {
		return c.Verts
	}
	c.Verts = []float32{}
	c.Verts = append(c.Verts, storeVNTC2(0, Vec3{-c.W, -c.H, 0}, Vec3{0, 0, 1}, Vec2{0, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(1, Vec3{c.W, -c.H, 0}, Vec3{0, 0, 1}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(2, Vec3{c.W, c.H, 0}, Vec3{0, 0, 1}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(3, Vec3{-c.W, c.H, 0}, Vec3{0, 0, 1}, Vec2{0, 1})...)
	return c.Verts
}

func (c *Shape) CreateCuboid() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	c.Verts = []float32{}
	c.Verts = append(c.Verts, storeVNTC2(0, Vec3{-c.W, -c.H, c.D}, Vec3{0, 0, 1}, Vec2{0, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(1, Vec3{c.W, -c.H, c.D}, Vec3{0, 0, 1}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(2, Vec3{c.W, c.H, c.D}, Vec3{0, 0, 1}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(3, Vec3{-c.W, c.H, c.D}, Vec3{0, 0, 1}, Vec2{0, 1})...)

	c.Verts = append(c.Verts, storeVNTC2(4, Vec3{-c.W, -c.H, -c.D}, Vec3{0, 0, -1}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(5, Vec3{-c.W, c.H, -c.D}, Vec3{0, 0, -1}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(6, Vec3{c.W, c.H, -c.D}, Vec3{0, 0, -1}, Vec2{0, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(7, Vec3{c.W, -c.H, -c.D}, Vec3{0, 0, -1}, Vec2{0, 0})...)

	c.Verts = append(c.Verts, storeVNTC2(8, Vec3{-c.W, c.H, -c.D}, Vec3{0, 1, 0}, Vec2{0, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(9, Vec3{-c.W, c.H, c.D}, Vec3{0, 1, 0}, Vec2{0, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(10, Vec3{c.W, c.H, c.D}, Vec3{0, 1, 0}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(11, Vec3{c.W, c.H, -c.D}, Vec3{0, 1, 0}, Vec2{1, 1})...)

	c.Verts = append(c.Verts, storeVNTC2(8, Vec3{-c.W, -c.H, -c.D}, Vec3{0, -1, 0}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(9, Vec3{c.W, -c.H, -c.D}, Vec3{0, -1, 0}, Vec2{0, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(10, Vec3{c.W, -c.H, c.D}, Vec3{0, -1, 0}, Vec2{0, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(11, Vec3{-c.W, -c.H, c.D}, Vec3{0, -1, 0}, Vec2{1, 0})...)

	c.Verts = append(c.Verts, storeVNTC2(8, Vec3{c.W, -c.H, -c.D}, Vec3{1, 0, 0}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(9, Vec3{c.W, c.H, -c.D}, Vec3{1, 0, 0}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(10, Vec3{c.W, c.H, c.D}, Vec3{1, 0, 0}, Vec2{0, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(11, Vec3{c.W, -c.H, c.D}, Vec3{1, 0, 0}, Vec2{0, 0})...)

	c.Verts = append(c.Verts, storeVNTC2(8, Vec3{-c.W, -c.H, -c.D}, Vec3{-1, 0, 0}, Vec2{0, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(9, Vec3{-c.W, -c.H, c.D}, Vec3{-1, 0, 0}, Vec2{1, 0})...)
	c.Verts = append(c.Verts, storeVNTC2(10, Vec3{-c.W, c.H, c.D}, Vec3{-1, 0, 0}, Vec2{1, 1})...)
	c.Verts = append(c.Verts, storeVNTC2(11, Vec3{-c.W, c.H, -c.D}, Vec3{-1, 0, 0}, Vec2{0, 1})...)

	return c.Verts
}

func (c *Shape) CreateTube() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	c.Path = make([]Vec2, 5)
	c.Path[0] = Vec2{c.W, c.D / 2}
	c.Path[1] = Vec2{c.H, c.D / 2}
	c.Path[2] = Vec2{c.H, -c.D / 2}
	c.Path[3] = Vec2{c.W, -c.D / 2}
	c.Path[4] = Vec2{c.W, c.D / 2}
	c.Verts = CreateLathe(c.Path, 1, 0, 2*math32.Pi, 0, uint32(c.Edges), 0, Vec3{0, 0, 0})
	return c.Verts
}

func (c *Shape) CreateCylinder() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	return CreateVCone(c.W, c.W, c.H, int(c.Edges))
}

func (c *Shape) CreateCone() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	return CreateVCone(0.01, c.W, c.H, int(c.Edges))
}

func (c *Shape) CreateTCone() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	return CreateVCone(c.W, c.D, c.H, int(c.Edges))
}

func CreateVCone(radius, radius2, height float32, sides int) []float32 {
	path := make([]Vec2, 4)
	path[0] = Vec2{0.01, height / 2}
	path[1] = Vec2{radius, height / 2}
	path[2] = Vec2{radius2, -height / 2}
	path[3] = Vec2{0.01, -height / 2}
	return CreateLathe(path, 1, 0, 2*math32.Pi, 0, uint32(sides), 0, Vec3{0, 0, 0})
}

func (c *Shape) CreateTorus() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	radius := c.W
	ringradius := c.H
	ringdivs := int(c.D)
	c.Path = make([]Vec2, ringdivs+1)
	st := (math32.Pi * 2) / float32(ringdivs)
	for r := 0; r <= ringdivs; r++ {
		c.Path[r] = Vec2{radius + ringradius*math32.Sin(float32(r)*st), ringradius * math32.Cos(float32(r)*st)}
	}
	c.Verts = CreateLathe(c.Path, 1, 0, 2*math32.Pi, 0, uint32(c.Edges), 0, Vec3{0, 0, 0})
	return c.Verts
}

func (c *Shape) CreateSphere() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	radius := c.W
	hemi := c.H
	c.Path = make([]Vec2, c.Edges+1)
	st := (math32.Pi * (1 - hemi)) / float32(c.Edges)
	ho := math32.Pi - (math32.Pi * (1 - hemi))

	for r := 0; r <= int(c.Edges); r++ {
		c.Path[r] = Vec2{radius * math32.Sin(float32(r)*st+ho) * sign(c.D), radius * math32.Cos(float32(r)*st+ho)}
	}
	c.Verts = CreateLathe(c.Path, 1, 0, 2*math32.Pi, 0, uint32(c.Edges), 0, Vec3{0, 0, 0})
	return c.Verts
}

func (c *Shape) CreateSpring() []float32 {
	if c.Verts != nil {
		return c.Verts
	}
	radius := c.W
	ringradius := c.H
	ringdivs := 12 //int(c.d)
	c.Path = make([]Vec2, ringdivs+1)
	st := (math32.Pi * 2) / float32(ringdivs)
	for r := 0; r <= ringdivs; r++ {
		c.Path[r] = Vec2{radius + ringradius*math32.Sin(float32(r)*st), ringradius * math32.Cos(float32(r)*st)}
	}
	c.Verts = CreateLathe(c.Path, 1, 0, 2*math32.Pi*10, c.D, uint32(c.Edges), 0, Vec3{0, 0, 0})
	return c.Verts
}

func CreateLathe(lpath []Vec2, inverted, startAngle, endAngle, rise float32, edges, uvtype uint32, pos Vec3) []float32 {

	normals, path := calcPathNormals(lpath, 0.5, true, inverted)

	sz := len(path)

	angDiff := endAngle - startAngle
	angStep := angDiff / float32(edges)

	tcx, tcy, rdiv := 1/float32(edges), float32(0), rise/float32(edges)

	miny, maxy := path[0].Y, path[0].Y
	for p := 0; p < sz; p++ {
		if path[p].Y < miny {
			miny = path[p].Y
		}
		if path[p].Y > maxy {
			maxy = path[p].Y
		}
	}
	cy := (maxy + miny) / 2
	tc := 0

	verts := []float32{}
	for p := 0; p < sz; p++ {
		switch uvtype {
		case 0: //cylinder map
			tcy = 1 - ((path[p].Y - miny) / (maxy - miny))
		case 1: //sphere map
			tcy = math32.Atan2(path[p].X, path[p].Y-cy) / math32.Pi
			if tcy < 0 {
				tcy += 1
			}
		}

		risey := path[p].Y

		for r := 0; r < int(edges); r++ {
			verts = append(verts, storeVNTC(p, tc, startAngle+float32(r)*angStep, risey, Vec2{tcx * float32(r), tcy}, pos, path, normals)...)
			risey += rdiv
		}
		verts = append(verts, storeVNTC(p, tc, startAngle, risey, Vec2{0.9999, tcy}, pos, path, normals)...)
	}

	return verts
}

func storeVNTC(p, tc int, ang, risey float32, uv Vec2, pos Vec3, path, normals []Vec2) []float32 {
	sinr, cosr := math32.Sin(ang), math32.Cos(ang)
	v := Vec3{pos.X + path[p].X*sinr, pos.Y + risey, pos.Z + path[p].X*cosr}
	n := Vec3{normals[p].X * sinr, normals[p].Y, normals[p].X * cosr}
	return []float32{float32(tc), v.X, v.Y, v.Z, n.X, n.Y, n.Z, uv.X, uv.Y}
}

func storeVNTC2(tc int, pos, normal Vec3, uv Vec2) []float32 {
	return []float32{float32(tc), pos.X, pos.Y, pos.Z, normal.X, normal.Y, normal.Z, uv.X, uv.Y}
}

func calcPathNormals(path []Vec2, creaseAngle float32, joined bool, inverted float32) ([]Vec2, []Vec2) {
	sz := len(path)
	if sz < 2 {
		return nil, nil
	}

	p1 := Vec2{}
	p2 := path[0]
	p3 := path[1]

	newPath := []Vec2{}
	normals := []Vec2{}

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

func dotInvert(v1 Vec2, v2 Vec2, inverted float32) Vec2 {
	dp := v1.Dot(v2)
	return Vec2{dp.X * inverted, dp.Y * inverted}
}

func angleBetween(v1, v2 Vec2) float32 {
	prod := v1.X*v2.Y - v1.Y*v2.X
	ab := sign(prod) * math32.Acos((v1.X*v2.X+v1.Y*v2.Y)/(v2.Length()*v2.Length()))
	return math32.Abs(ab)
}

func sign(v float32) float32 {
	if v < 0 {
		return -1
	}
	return 1
}
