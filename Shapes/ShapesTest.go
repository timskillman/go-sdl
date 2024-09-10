package main

import (
	g "github.com/timskillman/go-sdl/goengine"

	"github.com/chewxy/math32"
)

const width, height = 800, 600
const imageDir = "../Resources/images/"

type vec3 = g.Vec3

func main() {

	scene := g.Scene{}
	scene.Setup("SDL 3D Shapes", width, height)

	scene.AddShape("cube1", g.ShapeCuboid, 3, 3, 3, vec3{-7, -10, -20}, vec3{0, 0, 0}, 6, 0xff00ffff, imageDir+"redsky.png")
	scene.AddShape("plane1", g.ShapePlane, 5, 5, 0, vec3{7, -10, -20}, vec3{0, 0, 0}, 1, 0xff00ffff, imageDir+"alps.jpg")
	scene.AddShape("sphere1", g.ShapeSphere, 5, 0, 1, vec3{-7, 0, -20}, vec3{0, 0.5, 0.5}, 30, 0xff00ffff, imageDir+"LavaRock.jpg")
	scene.AddShape("torus1", g.ShapeTorus, 5, 2, 30, vec3{7, 0, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, imageDir+"spanel.png")
	scene.AddShape("tube1", g.ShapeTube, 2, 4, 6, vec3{10, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, imageDir+"spanel.png")
	scene.AddShape("cylinder1", g.ShapeCylinder, 3, 5, 20, vec3{0, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, imageDir+"clouds.jpg")
	scene.AddShape("cone1", g.ShapeCone, 3, 5, 0, vec3{-10, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, imageDir+"spanel.png")
	scene.AddShape("tcone1", g.ShapeTCone, 2, 5, 3, vec3{-20, 0, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, imageDir+"barktile.jpg")
	scene.AddShape("spring1", g.ShapeSpring, 1.5, 0.2, 15, vec3{20, 0, -20}, vec3{0, 0.5, 0.5}, 180, 0xff00ffff, imageDir+"spanel.png")

	zang := float32(0.0)

	userInput := g.UserInput{}
	for !userInput.Quit {
		userInput.GetUserInput()

		scene.Draw()

		var shape = scene.Shape("cube1")
		shape.Rotation = vec3{shape.Rotation.X + 3, shape.Rotation.Y, shape.Rotation.Z + 1}

		scene.Shape("plane1").Rotation.Z += 1
		scene.Shape("plane1").Position.Z = math32.Cos(zang)*5 - 20
		zang += 0.1

		scene.Shape("sphere1").Rotation.X += 1
		scene.Shape("sphere1").Rotation.Y += 0.5

		scene.Shape("torus1").Rotation.Y += 2
		scene.Shape("torus1").Rotation.X += 0.7

		scene.Shape("tube1").Rotation.Y += 3
		scene.Shape("tube1").Rotation.X += 0.3

		scene.Shape("cone1").Rotation.Y += 3
		scene.Shape("cone1").Rotation.X += 0.3

		scene.Shape("tcone1").Rotation.Y += 3
		scene.Shape("tcone1").Rotation.X += 2

		scene.Shape("spring1").Rotation.Y += 3
		scene.Shape("spring1").Rotation.X += 2

		shape = scene.Shape("cylinder1")
		shape.Rotation = vec3{shape.Rotation.X + 0.5, shape.Rotation.Y + 0.3, shape.Rotation.Z + 0.1}

		scene.Window.GLSwap()
	}

	scene.Quit()
}
