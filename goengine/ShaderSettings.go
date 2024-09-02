package goengine

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Note: these settings are particular to the provided vertex shader and are not generic

type ShaderSettings struct {
	fogColour  uint32
	fogMaxDist float32
	fogMinDist float32
	lightPos   Vec3
}

type shaderRef = int32

const (
	fogColourRef shaderRef = iota
	fogRangeRef
	fogMaxRef
	specularRef
	diffuseRef
	ambientRef
	emissiveRef
	texAnimRef
	lightPosRef
	lightColRef
	textureRef
	reflectRef
	perspectiveMatrixRef
	modelMatrixRef
	illuminationModelRef
	lastRef //dont remove this and always leave it last
)

func (settings *ShaderSettings) SetupShaderSettings(program uint32) []int32 {
	refs := make([]int32, lastRef)
	refs[illuminationModelRef] = gl.GetUniformLocation(program, gl.Str("u_illuminationModel\x00"))
	refs[fogColourRef] = GetSetVec3(program, gl.Str("u_fogColour\x00"), ColToFloats(settings.fogColour))
	refs[fogRangeRef] = GetSetFloat(program, gl.Str("u_fogRange\x00"), 1/(settings.fogMaxDist-settings.fogMinDist))
	refs[fogMaxRef] = GetSetFloat(program, gl.Str("u_fogMaxDist\x00"), settings.fogMaxDist)
	refs[specularRef] = gl.GetUniformLocation(program, gl.Str("u_specularColour\x00"))
	refs[diffuseRef] = gl.GetUniformLocation(program, gl.Str("u_diffuseColour\x00"))
	refs[ambientRef] = gl.GetUniformLocation(program, gl.Str("u_ambientColour\x00"))
	refs[emissiveRef] = gl.GetUniformLocation(program, gl.Str("u_emissiveColour\x00"))
	refs[texAnimRef] = gl.GetUniformLocation(program, gl.Str("u_animoffset\x00"))
	refs[lightPosRef] = GetSetVec3(program, gl.Str("u_LightPos\x00"), Vec3toFloats(&settings.lightPos))
	refs[lightColRef] = gl.GetUniformLocation(program, gl.Str("u_lightColour\x00"))
	refs[reflectRef] = gl.GetUniformLocation(program, gl.Str("u_reflective\x00"))
	refs[perspectiveMatrixRef] = gl.GetUniformLocation(program, gl.Str("u_ProjMatrix\x00"))
	refs[modelMatrixRef] = gl.GetUniformLocation(program, gl.Str("u_ModelMatrix\x00"))
	refs[textureRef] = GetSetInt(program, gl.Str("u_Texture\x00"), 0)
	return refs
}

func ActiveTexture(texID uint32, uniformId string, program uint32, texActive uint32) int32 {
	texh := gl.GetUniformLocation(program, gl.Str(uniformId+"\x00"))
	if texh >= 0 {
		gl.ActiveTexture(gl.TEXTURE0 + texActive)
		gl.BindTexture(uint32(texh), texActive)
	}
	return texh
}

func SetTexture(uniformId string, program uint32, texLoc int32) int32 {
	texh := gl.GetUniformLocation(program, gl.Str(uniformId+"\x00"))
	if texh >= 0 {
		gl.Uniform1i(texh, texLoc)
	}
	return texh
}

func SetFog(refs []int32, minDist, maxDist float32, colour []float32) {
	gl.Uniform3fv(refs[fogColourRef], 1, &colour[0])
	gl.Uniform1f(refs[fogRangeRef], 1/(maxDist-minDist))
	gl.Uniform1f(refs[fogMaxRef], maxDist)
}

func GetSetInt(program uint32, name *uint8, v int32) int32 {
	loc := gl.GetUniformLocation(program, name)
	if loc >= 0 {
		gl.Uniform1i(loc, v)
	}
	return loc
}

func GetSetFloat(program uint32, name *uint8, v float32) int32 {
	loc := gl.GetUniformLocation(program, name)
	if loc >= 0 {
		gl.Uniform1f(loc, v)
	}
	return loc
}

func GetSetVec3(program uint32, name *uint8, vec []float32) int32 {
	loc := gl.GetUniformLocation(program, name)
	if loc >= 0 {
		gl.Uniform3fv(loc, 1, &vec[0])
	}
	return loc
}

func ColToFloats(col uint32) []float32 {
	c := make([]float32, 3)
	c[0] = float32(col&255) / 255
	c[1] = float32((col>>8)&255) / 255
	c[2] = float32((col>>16)&255) / 255
	return c
}
