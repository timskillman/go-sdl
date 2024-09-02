package goengine

import (
	s "strings"

	"github.com/go-gl/gl/v2.1/gl"
)

type Shader struct {
	handle uint32
}

func GetAttributes(vertShaderSource string) ([]string, string) {
	attributes := make([]string, 0)
	err := ""

	attkey := "attribute"
	if s.Contains(vertShaderSource, "#Version 3") {
		attkey = "in"
	}

	m := s.Index(vertShaderSource, "void")
	if m < 0 {
		return attributes, "Missing void in vertex shader"
	}
	header := vertShaderSource[:m]
	m = s.Index(vertShaderSource, attkey)
	if m < 0 {
		return attributes, "Cant find attributes in vertex shader header"
	}

	header = header[m:]

	i := 0
	for i >= 0 {
		i = s.Index(header, attkey)
		if i >= 0 {
			substr := header[i:]
			j := Find(header, ";", i)
			if j >= 0 {
				substr := substr[:j]
				k := s.Index(substr, " ")
				if k >= 0 {
					attributes = append(attributes, substr[k+1:])
					//attributes += substr[k+1:] + " "
				}
				header = header[j:]
			} else {
				i = -1
				err = "Missing ; at end of attribute"
			}
		}
	}
	return attributes, err
}

func LoadShaderStr(shaderType uint32, shaderSource string) (*Shader, string) {

	var shader = gl.CreateShader(shaderType)
	glSrcs, freeFn := gl.Strs(shaderSource + "\x00")
	defer freeFn()
	gl.ShaderSource(shader, 1, glSrcs, nil)
	gl.CompileShader(shader)

	compileOK := int32(0)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &compileOK)

	if compileOK == 0 {
		loglength := int32(0)
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &loglength)
		logErr := gl.Str(s.Repeat(" ", int(loglength))) //Create a string long enough to hold error message
		gl.GetShaderInfoLog(shader, loglength, nil, logErr)
		return nil, "Compile error: " + string(*logErr)
	}

	return &Shader{handle: shader}, ""
}

func CreateShaderProgram(vertexShader, fragShader *Shader, attributes []string) (uint32, string) {
	program := gl.CreateProgram()
	if program == 0 {
		return 0, "Program not created"
	}

	gl.AttachShader(program, vertexShader.handle)
	gl.AttachShader(program, fragShader.handle)

	for i, att := range attributes {
		gl.BindAttribLocation(program, uint32(i), gl.Str(att+"\x00"))
	}

	gl.LinkProgram(program)

	linkOK := int32(0)
	gl.GetProgramiv(program, gl.LINK_STATUS, &linkOK)
	if linkOK == 0 {
		loglength := int32(0)
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &loglength)
		logErr := gl.Str(s.Repeat(" ", int(loglength)))
		gl.GetShaderInfoLog(program, loglength, nil, logErr)
		return 0, "Shader link failure: " + string(*logErr)
	}

	gl.DetachShader(program, vertexShader.handle)
	gl.DetachShader(program, fragShader.handle)

	return program, ""
}

func GLVersion() string {
	var logErr = gl.GetString(gl.VERSION)
	return string(*logErr)
}
