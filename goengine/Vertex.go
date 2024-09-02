package goengine

import (
	"github.com/chewxy/math32"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

// V3 returns a new [Vec3] with the given x, y and z components.
func V3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

// V3Scalar returns a new [Vec3] with all components set to the given scalar value.
func V3Scalar(s float32) Vec3 {
	return Vec3{s, s, s}
}

// Set sets this vector X, Y and Z components.
func (v *Vec3) Set(x, y, z float32) {
	v.X = x
	v.Y = y
	v.Z = z
}

// SetScalar sets all vector X, Y and Z components to same scalar value.
func (v *Vec3) SetScalar(s float32) {
	v.X = s
	v.Y = s
	v.Z = s
}

// SetZero sets this vector X, Y and Z components to be zero.
func (v *Vec3) SetZero() {
	v.SetScalar(0)
}

///////////////////////////////////////////////////////////////////////
//  Basic math operations

// Add adds other vector to this one and returns result in a new vector.
func (v Vec3) Add(other Vec3) Vec3 {
	return V3(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
}

// AddScalar adds scalar s to each component of this vector and returns new vector.
func (v Vec3) AddScalar(s float32) Vec3 {
	return V3(v.X+s, v.Y+s, v.Z+s)
}

// SetAdd sets this to addition with other vector (i.e., += or plus-equals).
func (v *Vec3) SetAdd(other Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// SetAddScalar sets this to addition with scalar.
func (v *Vec3) SetAddScalar(s float32) {
	v.X += s
	v.Y += s
	v.Z += s
}

// Sub subtracts other vector from this one and returns result in new vector.
func (v Vec3) Sub(other Vec3) Vec3 {
	return V3(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
}

// SubScalar subtracts scalar s from each component of this vector and returns new vector.
func (v Vec3) SubScalar(s float32) Vec3 {
	return V3(v.X-s, v.Y-s, v.Z-s)
}

// SetSub sets this to subtraction with other vector (i.e., -= or minus-equals).
func (v *Vec3) SetSub(other Vec3) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

// SetSubScalar sets this to subtraction of scalar.
func (v *Vec3) SetSubScalar(s float32) {
	v.X -= s
	v.Y -= s
	v.Z -= s
}

// Mul multiplies each component of this vector by the corresponding one from other
// and returns resulting vector.
func (v Vec3) Mul(other Vec3) Vec3 {
	return V3(v.X*other.X, v.Y*other.Y, v.Z*other.Z)
}

// MulScalar multiplies each component of this vector by the scalar s and returns resulting vector.
func (v Vec3) MulScalar(s float32) Vec3 {
	return V3(v.X*s, v.Y*s, v.Z*s)
}

// SetMul sets this to multiplication with other vector (i.e., *= or times-equals).
func (v *Vec3) SetMul(other Vec3) {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z
}

// SetMulScalar sets this to multiplication by scalar.
func (v *Vec3) SetMulScalar(s float32) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Div divides each component of this vector by the corresponding one from other vector
// and returns resulting vector.
func (v Vec3) Div(other Vec3) Vec3 {
	return V3(v.X/other.X, v.Y/other.Y, v.Z/other.Z)
}

// DivScalar divides each component of this vector by the scalar s and returns resulting vector.
// If scalar is zero, returns zero.
func (v Vec3) DivScalar(scalar float32) Vec3 {
	if scalar != 0 {
		return v.MulScalar(1 / scalar)
	} else {
		return Vec3{}
	}
}

// SetDiv sets this to division by other vector (i.e., /= or divide-equals).
func (v *Vec3) SetDiv(other Vec3) {
	v.X /= other.X
	v.Y /= other.Y
	v.Z /= other.Z
}

// SetDivScalar sets this to division by scalar.
func (v *Vec3) SetDivScalar(s float32) {
	if s != 0 {
		v.SetMulScalar(1 / s)
	} else {
		v.SetZero()
	}
}

// Min returns min of this vector components vs. other vector.
func (v Vec3) Min(other Vec3) Vec3 {
	return V3(math32.Min(v.X, other.X), math32.Min(v.Y, other.Y), math32.Min(v.Z, other.Z))
}

// SetMin sets this vector components to the minimum values of itself and other vector.
func (v *Vec3) SetMin(other Vec3) {
	v.X = math32.Min(v.X, other.X)
	v.Y = math32.Min(v.Y, other.Y)
	v.Z = math32.Min(v.Z, other.Z)
}

// Max returns max of this vector components vs. other vector.
func (v Vec3) Max(other Vec3) Vec3 {
	return V3(math32.Max(v.X, other.X), math32.Max(v.Y, other.Y), math32.Max(v.Z, other.Z))
}

// SetMax sets this vector components to the maximum value of itself and other vector.
func (v *Vec3) SetMax(other Vec3) {
	v.X = math32.Max(v.X, other.X)
	v.Y = math32.Max(v.Y, other.Y)
	v.Z = math32.Max(v.Z, other.Z)
}

// Clamp sets this vector components to be no less than the corresponding components of min
// and not greater than the corresponding component of max.
// Assumes min < max, if this assumption isn't true it will not operate correctly.
func (v *Vec3) Clamp(min, max Vec3) {
	if v.X < min.X {
		v.X = min.X
	} else if v.X > max.X {
		v.X = max.X
	}
	if v.Y < min.Y {
		v.Y = min.Y
	} else if v.Y > max.Y {
		v.Y = max.Y
	}
	if v.Z < min.Z {
		v.Z = min.Z
	} else if v.Z > max.Z {
		v.Z = max.Z
	}
}

// Clamp clamps x to the provided closed interval [a, b]
func Clamp(x, a, b float32) float32 {
	if x < a {
		return a
	}
	if x > b {
		return b
	}
	return x
}

// ClampScalar sets this vector components to be no less than minVal and not greater than maxVal.
func (v *Vec3) ClampScalar(minVal, maxVal float32) {
	v.Clamp(V3Scalar(minVal), V3Scalar(maxVal))
}

// Floor returns vector with mat32.Floor() applied to each of this vector's components.
func (v Vec3) Floor() Vec3 {
	return V3(math32.Floor(v.X), math32.Floor(v.Y), math32.Floor(v.Z))
}

// SetFloor applies mat32.Floor() to each of this vector's components.
func (v *Vec3) SetFloor() {
	v.X = math32.Floor(v.X)
	v.Y = math32.Floor(v.Y)
	v.Z = math32.Floor(v.Z)
}

// Ceil returns vector with mat32.Ceil() applied to each of this vector's components.
func (v Vec3) Ceil() Vec3 {
	return V3(math32.Ceil(v.X), math32.Ceil(v.Y), math32.Ceil(v.Z))
}

// SetCeil applies mat32.Ceil() to each of this vector's components.
func (v *Vec3) SetCeil() {
	v.X = math32.Ceil(v.X)
	v.Y = math32.Ceil(v.Y)
	v.Z = math32.Ceil(v.Z)
}

// Round returns vector with mat32.Round() applied to each of this vector's components.
func (v Vec3) Round() Vec3 {
	return V3(math32.Round(v.X), math32.Round(v.Y), math32.Round(v.Z))
}

// SetRound rounds each of this vector's components.
func (v *Vec3) SetRound() {
	v.X = math32.Round(v.X)
	v.Y = math32.Round(v.Y)
	v.Z = math32.Round(v.Z)
}

// Negate returns vector with each component negated.
func (v Vec3) Negate() Vec3 {
	return V3(-v.X, -v.Y, -v.Z)
}

// SetNegate negates each of this vector's components.
func (v *Vec3) SetNegate() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Abs returns vector with Abs of each component.
func (v Vec3) Abs() Vec3 {
	return V3(math32.Abs(v.X), math32.Abs(v.Y), math32.Abs(v.Z))
}

//////////////////////////////////////////////////////////////////////////////////
//  Distance, Norm

// IsEqual returns if this vector is equal to other.
func (v Vec3) IsEqual(other Vec3) bool {
	return (other.X == v.X) && (other.Y == v.Y) && (other.Z == v.Z)
}

// AlmostEqual returns whether the vector is almost equal to another vector within the specified tolerance.
func (v *Vec3) AlmostEqual(other Vec3, tol float32) bool {
	if (math32.Abs(v.X-other.X) < tol) &&
		(math32.Abs(v.Y-other.Y) < tol) &&
		(math32.Abs(v.Z-other.Z) < tol) {
		return true
	}
	return false
}

// Dot returns the dot product of this vector with other.
func (v Vec3) Dot(other Vec3) float32 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// LengthSq returns the length squared of this vector.
// LengthSq can be used to compare vectors' lengths without the need to perform a square root.
func (v Vec3) LengthSq() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Length returns the length of this vector.
func (v Vec3) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normal returns this vector divided by its length
func (v Vec3) Normal() Vec3 {
	return v.DivScalar(v.Length())
}

// SetNormal normalizes this vector so its length will be 1.
func (v *Vec3) SetNormal() {
	v.SetDivScalar(v.Length())
}

// Normalize normalizes this vector so its length will be 1.
func (v *Vec3) Normalize() {
	v.SetDivScalar(v.Length())
}

// DistTo returns the distance of this point to other.
func (v Vec3) DistTo(other Vec3) float32 {
	return math32.Sqrt(v.DistToSquared(other))
}

// DistToSquared returns the distance squared of this point to other.
func (v Vec3) DistToSquared(other Vec3) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return dx*dx + dy*dy + dz*dz
}

// SetLength sets this vector to have the specified length.
// If the current length is zero, does nothing.
func (v *Vec3) SetLength(l float32) {
	oldLength := v.Length()
	if oldLength != 0 && l != oldLength {
		v.SetMulScalar(l / oldLength)
	}
}

// Lerp returns vector with each components as the linear interpolated value of
// alpha between itself and the corresponding other component.
func (v Vec3) Lerp(other Vec3, alpha float32) Vec3 {
	return V3(v.X+(other.X-v.X)*alpha, v.Y+(other.Y-v.Y)*alpha, v.Z+(other.Z-v.Z)*alpha)
}

// SetLerp sets each of this vector's components to the linear interpolated value of
// alpha between itself and the corresponding other component.
func (v *Vec3) SetLerp(other Vec3, alpha float32) {
	v.X += (other.X - v.X) * alpha
	v.Y += (other.Y - v.Y) * alpha
	v.Z += (other.Z - v.Z) * alpha
}

/////////////////////////////////////////////////////////////////////////////
//  Matrix operations

// RotateAxisAngle returns vector rotated around axis by angle.
func (v Vec3) RotateAxisAngle(axis Vec3, angle float32) Vec3 {
	return v.MulQuat(NewQuatAxisAngle(axis, angle))
}

// SetRotateAxisAngle sets vector rotated around axis by angle.
func (v *Vec3) SetRotateAxisAngle(axis Vec3, angle float32) {
	*v = v.RotateAxisAngle(axis, angle)
}

// // MulMat3 returns vector multiplied by specified 3x3 matrix.
// func (v Vec3) MulMat3(m *Mat3) Vec3 {
// 	return Vec3{m[0]*v.X + m[3]*v.Y + m[6]*v.Z,
// 		m[1]*v.X + m[4]*v.Y + m[7]*v.Z,
// 		m[2]*v.X + m[5]*v.Y + m[8]*v.Z}
// }

// // SetMulMat3 sets vector multiplied by specified 3x3 matrix.
// func (v *Vec3) SetMulMat3(m *Mat3) {
// 	*v = v.MulMat3(m)
// }

// MulMat4 returns vector multiplied by specified 4x4 matrix.
func (v Vec3) MulMat4(m *Mat4s) Vec3 {
	return Vec3{m.m0*v.X + m.m4*v.Y + m.m8*v.Z + m.m12,
		m.m1*v.X + m.m5*v.Y + m.m9*v.Z + m.m13,
		m.m2*v.X + m.m6*v.Y + m.m10*v.Z + m.m14}
}

// MulMat4AsVec4 returns 3-dim vector multiplied by specified 4x4 matrix
// using a 4-dim vector with given 4th dimensional value, then reduced back to
// a 3-dimensional vector.  This is somehow different from just straight
// MulMat4 on the 3-dim vector.  Use 0 for normals and 1 for positions
// as the 4th dim to set.
// func (v Vec3) MulMat4AsVec4(m *Mat4s, w float32) Vec3 {
// 	return V3FromV4(V4FromV3(v, w).MulMat4(m))
// }

// SetMulMat4 sets vector multiplied by specified 4x4 matrix.
func (v *Vec3) SetMulMat4(m *Mat4s) {
	*v = v.MulMat4(m)
}

// MVProjToNDC project given vector through given MVP model-view-projection Mat4
// and do perspective divide to return normalized display coordinates (NDC).
// w is value for 4th coordinate -- use 1 for positions, 0 for normals.
// func (v Vec3) MVProjToNDC(m *Mat4s, w float32) Vec3 {
// 	clip := V4FromV3(v, w).MulMat4(m)
// 	return clip.PerspDiv()
// }

// NDCToWindow converts normalized display coordinates (NDC) to window
// (pixel) coordinates, using given window size parameters.
// near, far are 0, 1 by default (glDepthRange defaults).
// flipY if true means flip the Y axis (top = 0 for windows vs. bottom = 0 for 3D coords)
func (v Vec3) NDCToWindow(size, off Vec2, near, far float32, flipY bool) Vec3 {
	w := Vec3{}
	half := size.MulScalar(0.5)
	w.X = half.X*v.X + half.X
	w.Y = half.Y*v.Y + half.Y
	w.Z = 0.5*(far-near)*v.Z + 0.5*(far+near)
	if flipY {
		w.Y = size.Y - w.Y
	}
	w.X += off.X
	w.Y += off.Y
	return w
}

// WindowToNDC converts window (pixel) coordinates to
// normalized display coordinates (NDC), using given window size parameters.
// The Z depth coordinate (0-1) must be set manually or by reading from framebuffer
// flipY if true means flip the Y axis (top = 0 for windows vs. bottom = 0 for 3D coords)
func (v Vec2) WindowToNDC(size, off Vec2, flipY bool) Vec3 {
	n := Vec3{}
	half := size.MulScalar(0.5)
	n.X = v.X - off.X
	n.Y = v.Y - off.Y
	if flipY {
		n.Y = size.Y - n.Y
	}
	n.X = n.X/half.X - 1
	n.Y = n.Y/half.Y - 1
	return n
}

// MulProjection returns vector multiplied by the projection matrix m
func (v Vec3) MulProjection(m *Mat4s) Vec3 {
	d := 1 / (m.m3*v.X + m.m7*v.Y + m.m11*v.Z + m.m15) // perspective divide
	return Vec3{(m.m0*v.X + m.m4*v.Y + m.m8*v.Z + m.m12) * d,
		(m.m1*v.X + m.m5*v.Y + m.m9*v.Z + m.m13) * d,
		(m.m2*v.X + m.m6*v.Y + m.m10*v.Z + m.m14) * d}
}

// MulQuat returns vector multiplied by specified quaternion and
// then by the quaternion inverse.
// It basically applies the rotation encoded in the quaternion to this vector.
func (v Vec3) MulQuat(q Quat) Vec3 {
	qx := q.X
	qy := q.Y
	qz := q.Z
	qw := q.W
	// calculate quat * vector
	ix := qw*v.X + qy*v.Z - qz*v.Y
	iy := qw*v.Y + qz*v.X - qx*v.Z
	iz := qw*v.Z + qx*v.Y - qy*v.X
	iw := -qx*v.X - qy*v.Y - qz*v.Z
	// calculate result * inverse quat
	return Vec3{ix*qw + iw*-qx + iy*-qz - iz*-qy,
		iy*qw + iw*-qy + iz*-qx - ix*-qz,
		iz*qw + iw*-qz + ix*-qy - iy*-qx}
}

// SetMulQuat multiplies vector by specified quaternion and
// then by the quaternion inverse.
// It basically applies the rotation encoded in the quaternion to this vector.
func (v *Vec3) SetMulQuat(q Quat) {
	*v = v.MulQuat(q)
}

// Cross returns the cross product of this vector with other.
func (v Vec3) Cross(other Vec3) Vec3 {
	return V3(v.Y*other.Z-v.Z*other.Y, v.Z*other.X-v.X*other.Z, v.X*other.Y-v.Y*other.X)
}

// ProjectOnVector returns vector projected on other vector.
func (v *Vec3) ProjectOnVector(other Vec3) Vec3 {
	on := other.Normal()
	return on.MulScalar(v.Dot(on))
}

// ProjectOnPlane returns vector projected on the plane specified by normal vector.
func (v *Vec3) ProjectOnPlane(planeNormal Vec3) Vec3 {
	return v.Sub(v.ProjectOnVector(planeNormal))
}

// Reflect returns vector reflected relative to the normal vector (assumed to be
// already normalized).
func (v *Vec3) Reflect(normal Vec3) Vec3 {
	return v.Sub(normal.MulScalar(2 * v.Dot(normal)))
}

// CosTo returns the cosine (normalized dot product) between this vector and other.
func (v Vec3) CosTo(other Vec3) float32 {
	return v.Dot(other) / (v.Length() * other.Length())
}

// AngleTo returns the angle between this vector and other.
// Returns angles in range of -PI to PI (not 0 to 2 PI).
func (v Vec3) AngleTo(other Vec3) float32 {
	ang := math32.Acos(Clamp(v.CosTo(other), -1, 1))
	cross := v.Cross(other)
	switch {
	case math32.Abs(cross.Z) >= math32.Abs(cross.Y) && math32.Abs(cross.Z) >= math32.Abs(cross.X):
		if cross.Z > 0 {
			ang = -ang
		}
	case math32.Abs(cross.Y) >= math32.Abs(cross.Z) && math32.Abs(cross.Y) >= math32.Abs(cross.X):
		if cross.Y > 0 {
			ang = -ang
		}
	case math32.Abs(cross.X) >= math32.Abs(cross.Z) && math32.Abs(cross.X) >= math32.Abs(cross.Y):
		if cross.X > 0 {
			ang = -ang
		}
	}
	return ang
}

// SetFromMatrixPos set this vector from the translation coordinates
// in the specified transformation matrix.
func (v *Vec3) SetFromMatrixPos(m *Mat4s) {
	v.X = m.m12
	v.Y = m.m13
	v.Z = m.m14
}

// SetFromMatrixCol set this vector with the column at index of the m matrix.
// func (v *Vec3) SetFromMatrixCol(index int, m *Mat4s) {
// 	offset := index * 4
// 	v.X = moffset
// 	v.Y = moffset+1
// 	v.Z = moffset+2
// }

// SetEulerAnglesFromMatrix sets this vector components to the Euler angles
// from the specified pure rotation matrix.
func (v *Vec3) SetEulerAnglesFromMatrix(m *Mat4s) {
	m11 := m.m0
	m12 := m.m4
	m13 := m.m8
	m22 := m.m5
	m23 := m.m9
	m32 := m.m6
	m33 := m.m10

	v.Y = math32.Asin(Clamp(m13, -1, 1))
	if math32.Abs(m13) < 0.99999 {
		v.X = math32.Atan2(-m23, m33)
		v.Z = math32.Atan2(-m12, m11)
	} else {
		v.X = math32.Atan2(m32, m22)
		v.Z = 0
	}
}

// NewEulerAnglesFromMatrix returns a Vec3 with components as the Euler angles
// from the specified pure rotation matrix.
func NewEulerAnglesFromMatrix(m *Mat4s) Vec3 {
	rot := Vec3{}
	rot.SetEulerAnglesFromMatrix(m)
	return rot
}

// SetEulerAnglesFromQuat sets this vector components to the Euler angles
// from the specified quaternion.
func (v *Vec3) SetEulerAnglesFromQuat(q Quat) {
	mat := Identity4()
	mat.SetRotationFromQuat(q)
	v.SetEulerAnglesFromMatrix(mat)
}

// RandomTangents computes and returns two arbitrary tangents to the vector.
func (v *Vec3) RandomTangents() (Vec3, Vec3) {
	t1 := Vec3{}
	t2 := Vec3{}
	length := v.Length()
	if length > 0 {
		n := v.Normal()
		randVec := Vec3{}
		if math32.Abs(n.X) < 0.9 {
			randVec.X = 1
			t1 = n.Cross(randVec)
		} else if math32.Abs(n.Y) < 0.9 {
			randVec.Y = 1
			t1 = n.Cross(randVec)
		} else {
			randVec.Z = 1
			t1 = n.Cross(randVec)
		}
		t2 = n.Cross(t1)
	} else {
		t1.X = 1
		t2.Y = 1
	}
	return t1, t2
}

///////////////////////////////////////////////////

func Vec3toCol(v *Vec3) uint32 {
	r := v.X * 255
	g := v.Y * 255
	b := v.Z * 255

	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}

	return uint32(r + g*255 + b*65536)
}

func Vec3toFloats(v *Vec3) []float32 {
	v3 := make([]float32, 3)
	v3[0] = v.X
	v3[1] = v.Y
	v3[2] = v.Z
	return v3
}
