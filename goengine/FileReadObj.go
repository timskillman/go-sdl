package goengine

import (
	"fmt"
	"log"
	"os"

	"github.com/udhos/gwob"
)

func ReadOBJ(file string, scene Scene) []Shape {

	// Set options
	options := &gwob.ObjParserOptions{
		LogStats: true,
		Logger:   func(msg string) { fmt.Fprintln(os.Stderr, msg) },
	}

	// Load OBJ
	o, errObj := gwob.NewObjFromFile(file, options)
	if errObj != nil {
		log.Printf("obj: parse error input=%s: %v", file, errObj)
		return nil
	}

	fileMtl := o.Mtllib
	sc := 0
	s := make([]Shape, len(o.Groups))

	// Load material lib
	lib, errMtl := gwob.ReadMaterialLibFromFile(fileMtl, options)
	if errMtl != nil {
		log.Printf("mtl: parse error input=%s: %v", fileMtl, errMtl)
	} else {

		normOffset := o.StrideOffsetNormal / 4
		texOffset := o.StrideOffsetTexture / 4
		vertSize := 9
		s[sc] = Shape{}

		// Scan OBJ groups
		s[sc].Verts = make([]float32, len(o.Indices)*vertSize) //x,y,z,nx,ny,nz,tx,ty,kd,ka
		vi := 0
		for _, g := range o.Groups {

			col := float32(0)
			//alpha := float32(0)

			mtl, found := lib.Lib[g.Usemtl]
			if found {
				log.Printf("obj=%s lib=%s group=%s material=%s MapKd=%s Kd=%v", file, fileMtl, g.Name, g.Usemtl, mtl.MapKd, mtl.Kd)
				//specular := (mtl.Ks[0] * 255) + (mtl.Ks[1]*255)*255 + (mtl.Ks[2]*255)*65536
				//alpha = (mtl.Ka[0] * 255) + (mtl.Ka[1]*255)*255 + (mtl.Ka[2]*255)*65536
				col = (mtl.Kd[0] * 255) + (mtl.Kd[1]*255)*255 + (mtl.Kd[2]*255)*65536 //need to combine this with alpha
			} else {
				log.Printf("obj=%s lib=%s group=%s material=%s NOT FOUND", file, fileMtl, g.Name, g.Usemtl)
			}

			indexStart := g.IndexBegin
			for i := 0; i < g.IndexCount; i++ {
				ci := o.Indices[indexStart+i]
				s[sc].Verts[vi], s[sc].Verts[vi+1], s[sc].Verts[vi+2] = o.Coord[ci], o.Coord[ci+1], o.Coord[ci+2]
				if o.NormCoordFound {
					s[sc].Verts[vi+3], s[sc].Verts[vi+4], s[sc].Verts[vi+5] = o.Coord[ci+normOffset], o.Coord[ci+normOffset+1], o.Coord[ci+normOffset+2]
				}
				if o.TextCoordFound {
					s[sc].Verts[vi+6], s[sc].Verts[vi+7] = o.Coord[ci+texOffset], o.Coord[ci+texOffset+1]
				}
				s[sc].Verts[vi+8] = col //diffuse taken from lib
				//s.Verts[vi+9] = alpha //alpha taken from lib
				vi += vertSize
			}
		}

	}
	return s
}
