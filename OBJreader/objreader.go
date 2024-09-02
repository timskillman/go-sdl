/*
Package main shows how to use the 'gwob' package to parse geometry data from OBJ files.

See also: https://github.com/udhos/gwob
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/udhos/gwob"
)

func main() {

	fileObj := os.Getenv("INPUT")
	if fileObj == "" {
		fileObj = "airport.obj"
	}
	log.Printf("env var INPUT=[%s] using input=%s", os.Getenv("INPUT"), fileObj)

	// Set options
	options := &gwob.ObjParserOptions{
		LogStats: true,
		Logger:   func(msg string) { fmt.Fprintln(os.Stderr, msg) },
	}

	// Load OBJ
	o, errObj := gwob.NewObjFromFile(fileObj, options)
	if errObj != nil {
		log.Printf("obj: parse error input=%s: %v", fileObj, errObj)
		return
	}

	fileMtl := o.Mtllib

	// Load material lib
	lib, errMtl := gwob.ReadMaterialLibFromFile(fileMtl, options)
	if errMtl != nil {
		log.Printf("mtl: parse error input=%s: %v", fileMtl, errMtl)
	} else {

		normOffset := o.StrideOffsetNormal / 4
		texOffset := o.StrideOffsetTexture / 4
		vertSize := 10

		// Scan OBJ groups
		verts := make([]float32, len(o.Indices)*vertSize) //x,y,z,nx,ny,nz,tx,ty,kd,ka
		vi := 0
		for _, g := range o.Groups {

			col := float32(0)
			alpha := float32(0)

			mtl, found := lib.Lib[g.Usemtl]
			if found {
				//log.Printf("obj=%s lib=%s group=%s material=%s MapKd=%s Kd=%v", fileObj, fileMtl, g.Name, g.Usemtl, mtl.MapKd, mtl.Kd)
				//continue
				col = (mtl.Kd[0] * 255) + (mtl.Kd[1]*255)*255 + (mtl.Kd[2]*255)*65536
				alpha = (mtl.Ka[0] * 255) + (mtl.Ka[1]*255)*255 + (mtl.Ka[2]*255)*65536
			}

			indexStart := g.IndexBegin
			for i := 0; i < g.IndexCount; i++ {
				ci := o.Indices[indexStart+i]
				verts[vi], verts[vi+1], verts[vi+2] = o.Coord[ci], o.Coord[ci+1], o.Coord[ci+2]
				if o.NormCoordFound {
					verts[vi+3], verts[vi+4], verts[vi+5] = o.Coord[ci+normOffset], o.Coord[ci+normOffset+1], o.Coord[ci+normOffset+2]
				}
				if o.TextCoordFound {
					verts[vi+6], verts[vi+7] = o.Coord[ci+texOffset], o.Coord[ci+texOffset+1]
				}
				verts[vi+8] = col   //diffuse taken from lib
				verts[vi+9] = alpha //alpha taken from lib
				vi += vertSize
			}

			log.Printf("obj=%s lib=%s group=%s material=%s NOT FOUND", fileObj, fileMtl, g.Name, g.Usemtl)
		}

	}

	if len(os.Args) < 2 {
		log.Printf("no cmd line args - dump to stdout suppressed")
		return
	}

	log.Printf("cmd line arg found - dumping to stdout")

	// Dump to stdout
	o.ToWriter(os.Stdout)
}
