//Struct based version of Goki Authors Matrix4.go
//Should be faster than array calculations

package goengine

import (
	"errors"

	"github.com/chewxy/math32"
)

// Mat4 is 4x4 matrix organized internally as column matrix.
type Mat4s struct {
	m0  float32
	m4  float32
	m8  float32
	m12 float32
	m1  float32
	m5  float32
	m9  float32
	m13 float32
	m2  float32
	m6  float32
	m10 float32
	m14 float32
	m3  float32
	m7  float32
	m11 float32
	m15 float32
}

// Set sets all the elements of this matrix row by row starting at row1, column1,
// row1, column2, row1, column3 and so forth.
func (m *Mat4s) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float32) {
	m.m0 = n11
	m.m4 = n12
	m.m8 = n13
	m.m12 = n14
	m.m1 = n21
	m.m5 = n22
	m.m9 = n23
	m.m13 = n24
	m.m2 = n31
	m.m6 = n32
	m.m10 = n33
	m.m14 = n34
	m.m3 = n41
	m.m7 = n42
	m.m11 = n43
	m.m15 = n44
}

// Identity4 returns a new identity [Mat4] matrix
func Identity4() *Mat4s {
	m := &Mat4s{}
	m.SetIdentity()
	return m
}

// FromArray set this matrix elements from the array starting at offset.
func (m *Mat4s) FromArray(array []float32, offset int) {
	m.m0 = array[0+offset]
	m.m4 = array[1+offset]
	m.m8 = array[2+offset]
	m.m12 = array[3+offset]
	m.m1 = array[4+offset]
	m.m5 = array[5+offset]
	m.m9 = array[6+offset]
	m.m13 = array[7+offset]
	m.m2 = array[8+offset]
	m.m6 = array[9+offset]
	m.m10 = array[10+offset]
	m.m14 = array[11+offset]
	m.m3 = array[12+offset]
	m.m7 = array[13+offset]
	m.m11 = array[14+offset]
	m.m15 = array[15+offset]
}

// Different from Goki ToArray - and with no offset
func (m *Mat4s) ToArray() []float32 {
	array := make([]float32, 16)
	array[0] = m.m0
	array[1] = m.m4
	array[2] = m.m8
	array[3] = m.m12
	array[4] = m.m1
	array[5] = m.m5
	array[6] = m.m9
	array[7] = m.m13
	array[8] = m.m2
	array[9] = m.m6
	array[10] = m.m10
	array[11] = m.m14
	array[12] = m.m3
	array[13] = m.m7
	array[14] = m.m11
	array[15] = m.m15
	return array
}

// SetIdentity sets this matrix as the identity matrix.
func (m *Mat4s) SetIdentity() {
	m.Set(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

// SetZero sets this matrix as the zero matrix.
func (m *Mat4s) SetZero() {
	m.Set(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)
}

// CopyFrom copies from source matrix into this matrix
// (a regular = assign does not copy data, just the pointer!)
func (m *Mat4s) CopyFrom(src *Mat4s) {
	m.Set(src.m0, src.m4, src.m8, src.m12, src.m1, src.m5, src.m9, src.m13, src.m2, src.m6, src.m10, src.m14, src.m3, src.m7, src.m11, src.m15)
}

// CopyPos copies the position elements of the src matrix into this one.
func (m *Mat4s) CopyPos(src *Mat4s) {
	m.m12 = src.m12
	m.m13 = src.m13
	m.m14 = src.m14
}

// MulMatrices sets this matrix as matrix multiplication a by b (i.e. b*a).
func (m *Mat4s) MulMatrices(a, b *Mat4s) {
	m.m0 = a.m0*b.m0 + a.m4*b.m1 + a.m8*b.m2 + a.m12*b.m3
	m.m4 = a.m0*b.m4 + a.m4*b.m5 + a.m8*b.m6 + a.m12*b.m7
	m.m8 = a.m0*b.m8 + a.m4*b.m9 + a.m8*b.m10 + a.m12*b.m11
	m.m12 = a.m0*b.m12 + a.m4*b.m13 + a.m8*b.m14 + a.m12*b.m15
	m.m1 = a.m1*b.m0 + a.m5*b.m1 + a.m9*b.m2 + a.m13*b.m3
	m.m5 = a.m1*b.m4 + a.m5*b.m5 + a.m9*b.m6 + a.m13*b.m7
	m.m9 = a.m1*b.m8 + a.m5*b.m9 + a.m9*b.m10 + a.m13*b.m11
	m.m13 = a.m1*b.m12 + a.m5*b.m13 + a.m9*b.m14 + a.m13*b.m15
	m.m2 = a.m2*b.m0 + a.m6*b.m1 + a.m10*b.m2 + a.m14*b.m3
	m.m6 = a.m2*b.m4 + a.m6*b.m5 + a.m10*b.m6 + a.m14*b.m7
	m.m10 = a.m2*b.m8 + a.m6*b.m9 + a.m10*b.m10 + a.m14*b.m11
	m.m14 = a.m2*b.m12 + a.m6*b.m13 + a.m10*b.m14 + a.m14*b.m15
	m.m3 = a.m3*b.m0 + a.m7*b.m1 + a.m11*b.m2 + a.m15*b.m3
	m.m7 = a.m3*b.m4 + a.m7*b.m5 + a.m11*b.m6 + a.m15*b.m7
	m.m11 = a.m3*b.m8 + a.m7*b.m9 + a.m11*b.m10 + a.m15*b.m11
	m.m15 = a.m3*b.m12 + a.m7*b.m13 + a.m11*b.m14 + a.m15*b.m15
}

// Mul returns this matrix times other matrix (this matrix is unchanged)
func (m *Mat4s) Mul(other *Mat4s) *Mat4s {
	nm := &Mat4s{}
	nm.MulMatrices(m, other)
	return nm
}

// SetMul sets this matrix to this matrix times other
func (m *Mat4s) SetMul(other *Mat4s) {
	m.MulMatrices(m, other)
}

// SetMulScalar multiplies each element of this matrix by the specified scalar.
func (m *Mat4s) MulScalar(s float32) {
	m.m0 *= s
	m.m4 *= s
	m.m8 *= s
	m.m12 *= s
	m.m1 *= s
	m.m5 *= s
	m.m9 *= s
	m.m13 *= s
	m.m2 *= s
	m.m6 *= s
	m.m10 *= s
	m.m14 *= s
	m.m3 *= s
	m.m7 *= s
	m.m11 *= s
	m.m15 *= s
}

// MulVec3Array multiplies count vectors (i.e., 3 sequential array values per each increment in count)
// in the array starting at start index by this matrix.
// func (m *Mat4s) MulVec3Array(array []float32, start, count int) {
// 	var v1 Vec3
// 	j := start
// 	for i := 0; i < count; i++ {
// 		v1.FromArray(array, j)
// 		mv := v1.MulMat4(m)
// 		mv.ToArray(array, j)
// 		j += 3
// 	}
// }

// Determinant calculates and returns the determinat of this matrix.
func (m *Mat4s) Determinant() float32 {
	return m.m3*(+m.m12*m.m9*m.m6-m.m8*m.m13*m.m6-m.m12*m.m5*m.m10+m.m4*m.m13*m.m10+m.m8*m.m5*m.m14-m.m4*m.m9*m.m14) +
		m.m7*(+m.m0*m.m9*m.m14-m.m0*m.m13*m.m10+m.m12*m.m1*m.m10-m.m8*m.m1*m.m14+m.m8*m.m13*m.m2-m.m12*m.m9*m.m2) +
		m.m11*(+m.m0*m.m13*m.m6-m.m0*m.m5*m.m14-m.m12*m.m1*m.m6+m.m4*m.m1*m.m14+m.m12*m.m5*m.m2-m.m4*m.m13*m.m2) +
		m.m15*(-m.m8*m.m5*m.m2-m.m0*m.m9*m.m6+m.m0*m.m5*m.m10+m.m8*m.m1*m.m6-m.m4*m.m1*m.m10+m.m4*m.m9*m.m2)
}

// SetInverse sets this matrix to the inverse of the src matrix.
// If the src matrix cannot be inverted returns error and
// sets this matrix to the identity matrix.
func (m *Mat4s) SetInverse(src *Mat4s) error {
	t11 := src.m9*src.m14*src.m7 - src.m13*src.m10*src.m7 + src.m13*src.m6*src.m11 - src.m5*src.m14*src.m11 - src.m9*src.m6*src.m15 + src.m5*src.m10*src.m15
	t12 := src.m12*src.m10*src.m7 - src.m8*src.m14*src.m7 - src.m12*src.m6*src.m11 + src.m4*src.m14*src.m11 + src.m8*src.m6*src.m15 - src.m4*src.m10*src.m15
	t13 := src.m8*src.m13*src.m7 - src.m12*src.m9*src.m7 + src.m12*src.m5*src.m11 - src.m4*src.m13*src.m11 - src.m8*src.m5*src.m15 + src.m4*src.m9*src.m15
	t14 := src.m12*src.m9*src.m6 - src.m8*src.m13*src.m6 - src.m12*src.m5*src.m10 + src.m4*src.m13*src.m10 + src.m8*src.m5*src.m14 - src.m4*src.m9*src.m14

	det := src.m0*t11 + src.m1*t12 + src.m2*t13 + src.m3*t14

	if det == 0 {
		m.SetIdentity()
		return errors.New("cannot invert matrix, determinant is 0")
	}

	detInv := 1 / det

	m.m0 = t11 * detInv
	m.m1 = (src.m13*src.m10*src.m3 - src.m9*src.m14*src.m3 - src.m13*src.m2*src.m11 + src.m1*src.m14*src.m11 + src.m9*src.m2*src.m15 - src.m1*src.m10*src.m15) * detInv
	m.m2 = (src.m5*src.m14*src.m3 - src.m13*src.m6*src.m3 + src.m13*src.m2*src.m7 - src.m1*src.m14*src.m7 - src.m5*src.m2*src.m15 + src.m1*src.m6*src.m15) * detInv
	m.m3 = (src.m9*src.m6*src.m3 - src.m5*src.m10*src.m3 - src.m9*src.m2*src.m7 + src.m1*src.m10*src.m7 + src.m5*src.m2*src.m11 - src.m1*src.m6*src.m11) * detInv
	m.m4 = t12 * detInv
	m.m5 = (src.m8*src.m14*src.m3 - src.m12*src.m10*src.m3 + src.m12*src.m2*src.m11 - src.m0*src.m14*src.m11 - src.m8*src.m2*src.m15 + src.m0*src.m10*src.m15) * detInv
	m.m6 = (src.m12*src.m6*src.m3 - src.m4*src.m14*src.m3 - src.m12*src.m2*src.m7 + src.m0*src.m14*src.m7 + src.m4*src.m2*src.m15 - src.m0*src.m6*src.m15) * detInv
	m.m7 = (src.m4*src.m10*src.m3 - src.m8*src.m6*src.m3 + src.m8*src.m2*src.m7 - src.m0*src.m10*src.m7 - src.m4*src.m2*src.m11 + src.m0*src.m6*src.m11) * detInv
	m.m8 = t13 * detInv
	m.m9 = (src.m12*src.m9*src.m3 - src.m8*src.m13*src.m3 - src.m12*src.m1*src.m11 + src.m0*src.m13*src.m11 + src.m8*src.m1*src.m15 - src.m0*src.m9*src.m15) * detInv
	m.m10 = (src.m4*src.m13*src.m3 - src.m12*src.m5*src.m3 + src.m12*src.m1*src.m7 - src.m0*src.m13*src.m7 - src.m4*src.m1*src.m15 + src.m0*src.m5*src.m15) * detInv
	m.m11 = (src.m8*src.m5*src.m3 - src.m4*src.m9*src.m3 - src.m8*src.m1*src.m7 + src.m0*src.m9*src.m7 + src.m4*src.m1*src.m11 - src.m0*src.m5*src.m11) * detInv
	m.m12 = t14 * detInv
	m.m13 = (src.m8*src.m13*src.m2 - src.m12*src.m9*src.m2 + src.m12*src.m1*src.m10 - src.m0*src.m13*src.m10 - src.m8*src.m1*src.m14 + src.m0*src.m9*src.m14) * detInv
	m.m14 = (src.m12*src.m5*src.m2 - src.m4*src.m13*src.m2 - src.m12*src.m1*src.m6 + src.m0*src.m13*src.m6 + src.m4*src.m1*src.m14 - src.m0*src.m5*src.m14) * detInv
	m.m15 = (src.m4*src.m9*src.m2 - src.m8*src.m5*src.m2 + src.m8*src.m1*src.m6 - src.m0*src.m9*src.m6 - src.m4*src.m1*src.m10 + src.m0*src.m5*src.m10) * detInv

	return nil
}

// Inverse returns the inverse of this matrix.
// If the matrix cannot be inverted returns error and
// sets this matrix to the identity matrix.
func (m *Mat4s) Inverse() (*Mat4s, error) {
	nm := &Mat4s{}
	err := nm.SetInverse(m)
	return nm, err
}

// SetTranspose transposes this matrix.
func (m *Mat4s) SetTranspose() {
	m.m1, m.m4 = m.m4, m.m1
	m.m2, m.m8 = m.m8, m.m2
	m.m6, m.m9 = m.m9, m.m6
	m.m3, m.m12 = m.m12, m.m3
	m.m7, m.m13 = m.m13, m.m7
	m.m11, m.m14 = m.m14, m.m11
}

// Transpose returns the transpose of this matrix.
func (m *Mat4s) Transpose() *Mat4s {
	nm := *m
	nm.SetTranspose()
	return &nm
}

/////////////////////////////////////////////////////////////////////////////
//   Translation, Rotation, Scaling transform

// ScaleCols returns matrix with first column of this matrix multiplied by the vector X component,
// the second column by the vector Y component and the third column by
// the vector Z component. The matrix fourth column is unchanged.
func (m *Mat4s) ScaleCols(v Vec3) *Mat4s {
	nm := &Mat4s{}
	nm.SetScaleCols(v)
	return nm
}

// SetScaleCols multiplies the first column of this matrix by the vector X component,
// the second column by the vector Y component and the third column by
// the vector Z component. The matrix fourth column is unchanged.
func (m *Mat4s) SetScaleCols(v Vec3) {
	m.m0 *= v.X
	m.m4 *= v.Y
	m.m8 *= v.Z
	m.m1 *= v.X
	m.m5 *= v.Y
	m.m9 *= v.Z
	m.m2 *= v.X
	m.m6 *= v.Y
	m.m10 *= v.Z
	m.m3 *= v.X
	m.m7 *= v.Y
	m.m11 *= v.Z
}

// GetMaxScaleOnAxis returns the maximum scale value of the 3 axes.
func (m *Mat4s) GetMaxScaleOnAxis() float32 {
	scaleXSq := m.m0*m.m0 + m.m1*m.m1 + m.m2*m.m2
	scaleYSq := m.m4*m.m4 + m.m5*m.m5 + m.m6*m.m6
	scaleZSq := m.m8*m.m8 + m.m9*m.m9 + m.m10*m.m10
	return math32.Sqrt(math32.Max(scaleXSq, math32.Max(scaleYSq, scaleZSq)))
}

// SetTranslation sets this matrix to a translation matrix from the specified x, y and z values.
func (m *Mat4s) SetTranslation(x, y, z float32) {
	m.Set(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

// SetRotationX sets this matrix to a rotation matrix of angle theta around the X axis.
func (m *Mat4s) SetRotationX(theta float32) {
	c := math32.Cos(theta)
	s := math32.Sin(theta)

	m.Set(
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	)
}

// SetRotationY sets this matrix to a rotation matrix of angle theta around the Y axis.
func (m *Mat4s) SetRotationY(theta float32) {
	c := math32.Cos(theta)
	s := math32.Sin(theta)
	m.Set(
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	)
}

// SetRotationZ sets this matrix to a rotation matrix of angle theta around the Z axis.
func (m *Mat4s) SetRotationZ(theta float32) {
	c := math32.Cos(theta)
	s := math32.Sin(theta)
	m.Set(
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

// SetRotationAxis sets this matrix to a rotation matrix of the specified angle around the specified axis.
func (m *Mat4s) SetRotationAxis(axis *Vec3, angle float32) {
	c := math32.Cos(angle)
	s := math32.Sin(angle)
	t := 1 - c
	x := axis.X
	y := axis.Y
	z := axis.Z
	tx := t * x
	ty := t * y
	m.Set(
		tx*x+c, tx*y-s*z, tx*z+s*y, 0,
		tx*y+s*z, ty*y+c, ty*z-s*x, 0,
		tx*z-s*y, ty*z+s*x, t*z*z+c, 0,
		0, 0, 0, 1,
	)
}

// SetScale sets this matrix to a scale transformation matrix using the specified x, y and z values.
func (m *Mat4s) SetScale(x, y, z float32) {
	m.Set(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

// SetPos sets this transformation matrix position fields from the specified vector v.
func (m *Mat4s) SetPos(v Vec3) {
	m.m12 = v.X
	m.m13 = v.Y
	m.m14 = v.Z
}

// Pos returns the position component of the matrix
func (m *Mat4s) Pos() Vec3 {
	pos := Vec3{}
	pos.X = m.m12
	pos.Y = m.m13
	pos.Z = m.m14
	return pos
}

// SetTransform sets this matrix to a transformation matrix for the specified position,
// rotation specified by the quaternion and scale.
func (m *Mat4s) SetTransform(pos Vec3, quat Quat, scale Vec3) {
	m.SetRotationFromQuat(quat)
	m.SetScaleCols(scale)
	m.SetPos(pos)
}

// ExtractRotation sets this matrix as rotation matrix from the src transformation matrix.
func (m *Mat4s) ExtractRotation(src *Mat4s) {
	scaleX := 1 / V3(src.m0, src.m1, src.m2).Length()
	scaleY := 1 / V3(src.m4, src.m5, src.m6).Length()
	scaleZ := 1 / V3(src.m8, src.m9, src.m10).Length()

	m.m0 = src.m0 * scaleX
	m.m1 = src.m1 * scaleX
	m.m2 = src.m2 * scaleX

	m.m4 = src.m4 * scaleY
	m.m5 = src.m5 * scaleY
	m.m6 = src.m6 * scaleY

	m.m8 = src.m8 * scaleZ
	m.m9 = src.m9 * scaleZ
	m.m10 = src.m10 * scaleZ
}

// SetRotationFromEuler set this a matrix as a rotation matrix from the specified euler angles.
func (m *Mat4s) SetRotationFromEuler(euler Vec3) {
	x := euler.X
	y := euler.Y
	z := euler.Z
	a := math32.Cos(x)
	b := math32.Sin(x)
	c := math32.Cos(y)
	d := math32.Sin(y)
	e := math32.Cos(z)
	f := math32.Sin(z)

	ae := a * e
	af := a * f
	be := b * e
	bf := b * f
	m.m0 = c * e
	m.m4 = -c * f
	m.m8 = d
	m.m1 = af + be*d
	m.m5 = ae - bf*d
	m.m9 = -b * c
	m.m2 = bf - ae*d
	m.m6 = be + af*d
	m.m10 = a * c
	m.m3 = 0
	m.m7 = 0
	m.m11 = 0
	m.m12 = 0
	m.m13 = 0
	m.m14 = 0
	m.m15 = 1
}

// SetRotationFromQuat sets this matrix as a rotation matrix from the specified quaternion.
func (m *Mat4s) SetRotationFromQuat(q Quat) {
	x := q.X
	y := q.Y
	z := q.Z
	w := q.W
	x2 := x + x
	y2 := y + y
	z2 := z + z
	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	m.m0 = 1 - (yy + zz)
	m.m4 = xy - wz
	m.m8 = xz + wy
	m.m1 = xy + wz
	m.m5 = 1 - (xx + zz)
	m.m9 = yz - wx
	m.m2 = xz - wy
	m.m6 = yz + wx
	m.m10 = 1 - (xx + yy)
	m.m3 = 0
	m.m7 = 0
	m.m11 = 0
	m.m12 = 0
	m.m13 = 0
	m.m14 = 0
	m.m15 = 1
}

// LookAt sets this matrix as view transform matrix with origin at eye,
// looking at target and using the up vector.
func (m *Mat4s) LookAt(eye, target, up Vec3) {
	z := eye.Sub(target)
	if z.LengthSq() == 0 {
		// Eye and target are in the same position
		z.Z = 1
	}
	z.SetNormal()

	x := up.Cross(z)
	if x.LengthSq() == 0 { // Up and Z are parallel
		if math32.Abs(up.Z) == 1 {
			z.X += 0.0001
		} else {
			z.Z += 0.0001
		}
		z.SetNormal()
		x = up.Cross(z)
	}
	x.SetNormal()

	y := z.Cross(x)

	m.m0 = x.X
	m.m1 = x.Y
	m.m2 = x.Z
	m.m4 = y.X
	m.m5 = y.Y
	m.m6 = y.Z
	m.m8 = z.X
	m.m9 = z.Y
	m.m10 = z.Z
}

// NewLookAt returns Mat4 matrix as view transform matrix with origin at eye,
// looking at target and using the up vector.
func NewLookAt(eye, target, up Vec3) *Mat4s {
	rotMat := &Mat4s{}
	rotMat.LookAt(eye, target, up)
	return rotMat
}

// SetFrustum sets this matrix to a projection frustum matrix bounded by the specified planes.
func (m *Mat4s) SetFrustum(left, right, bottom, top, near, far float32) {
	fmn := far - near
	m.m0 = 2 * near / (right - left)
	m.m1 = 0
	m.m2 = 0
	m.m3 = 0
	m.m4 = 0
	m.m5 = 2 * near / (top - bottom)
	m.m6 = 0
	m.m7 = 0
	m.m8 = (right + left) / (right - left)
	m.m9 = (top + bottom) / (top - bottom)
	m.m10 = -(far + near) / fmn
	m.m11 = -1
	m.m12 = 0
	m.m13 = 0
	m.m14 = -(2 * far * near) / fmn
	m.m15 = 0
}

// SetPerspective sets this matrix to a perspective projection matrix
// with the specified field of view in degrees,
// aspect ratio (width/height) and near and far planes.
func (m *Mat4s) SetPerspective(fov, aspect, near, far float32) {
	ymax := near * math32.Tan(DegToRad(fov*0.5))
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect
	m.SetFrustum(xmin, xmax, ymin, ymax, near, far)
}

// SetOrthographic sets this matrix to an orthographic projection matrix.
func (m *Mat4s) SetOrthographic(width, height, near, far float32) {
	p := far - near
	z := (far + near) / p

	m.m0 = 2 / width
	m.m4 = 0
	m.m8 = 0
	m.m12 = 0
	m.m1 = 0
	m.m5 = 2 / height
	m.m9 = 0
	m.m13 = 0
	m.m2 = 0
	m.m6 = 0
	m.m10 = -2 / p
	m.m14 = -z
	m.m3 = 0
	m.m7 = 0
	m.m11 = 0
	m.m15 = 1
}

// SetVkFrustum sets this matrix to a projection frustum matrix bounded by the specified planes.
// This version is for use with Vulkan, and does the equivalent of GLM_DEPTH_ZERO_ONE in glm
// and also multiplies the Y axis by -1, preserving the original OpenGL Y-up system.
// OpenGL provides a "natural" coordinate system for the physical world
// so it is useful to retain that for the world system and just convert
// on the way out to the render using this projection matrix.
func (m *Mat4s) SetVkFrustum(left, right, bottom, top, near, far float32) {
	fmn := far - near
	m.m0 = 2 * near / (right - left)
	m.m1 = 0
	m.m2 = 0
	m.m3 = 0
	m.m4 = 0
	m.m5 = -2 * near / (top - bottom)
	m.m6 = 0
	m.m7 = 0
	m.m8 = (right + left) / (right - left)
	m.m9 = (top + bottom) / (top - bottom)
	m.m10 = -far / fmn
	m.m11 = -1
	m.m12 = 0
	m.m13 = 0
	m.m14 = -(far * near) / fmn
	m.m15 = 0
}

// SetVkPerspective sets this matrix to a vulkan appropriate perspective
// projection matrix, assuming the use of the OpenGL Y-up
// coordinate system for the geometry points.
// OpenGL provides a "natural" coordinate system for the physical world
// so it is useful to retain that for the world system and just convert
// on the way out to the render using this projection matrix.
// The specified field of view is in degrees,
// aspect ratio (width/height) and near and far planes.
func (m *Mat4s) SetVkPerspective(fov, aspect, near, far float32) {
	ymax := near * math32.Tan(DegToRad(fov*0.5))
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect
	m.SetVkFrustum(xmin, xmax, ymin, ymax, near, far)
}

const (
	// DegToRadFactor is the number of radians per degree.
	DegToRadFactor = math32.Pi / 180

	// RadToDegFactor is the number of degrees per radian.
	RadToDegFactor = 180 / math32.Pi
)

// Infinity is positive infinity.
var Infinity = float32(math32.Inf(1))

// DegToRad converts a number from degrees to radians
func DegToRad(degrees float32) float32 {
	return degrees * DegToRadFactor
}

// RadToDeg converts a number from radians to degrees
func RadToDeg(radians float32) float32 {
	return radians * RadToDegFactor
}
