package main

import "github.com/chewxy/math32"

const width, height = 800, 600

func main() {

	scene := Scene{}
	scene.Setup("SDL 3D Shapes", width, height)

	scene.AddShape("cube1", cuboidShape, 3, 3, 3, vec3{-7, -10, -20}, vec3{0, 0, 0}, 6, 0xff00ffff, "redsky.png")
	scene.AddShape("plane1", planeShape, 5, 5, 0, vec3{7, -10, -20}, vec3{0, 0, 0}, 1, 0xff00ffff, "alps.jpg")
	scene.AddShape("sphere1", sphereShape, 5, 0, 1, vec3{-7, 0, -20}, vec3{0, 0.5, 0.5}, 30, 0xff00ffff, "LavaRock.jpg")
	scene.AddShape("torus1", torusShape, 5, 2, 30, vec3{7, 0, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, "spanel.png")
	scene.AddShape("tube1", tubeShape, 2, 4, 6, vec3{10, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, "spanel.png")
	scene.AddShape("cylinder1", cylinderShape, 3, 5, 20, vec3{0, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, "clouds.jpg")
	scene.AddShape("cone1", coneShape, 3, 5, 0, vec3{-10, 10, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, "spanel.png")
	scene.AddShape("tcone1", tconeShape, 2, 5, 3, vec3{-20, 0, -20}, vec3{0, 0.5, 0.5}, 20, 0xff00ffff, "barktile.jpg")
	scene.AddShape("spring1", springShape, 1.5, 0.2, 15, vec3{20, 0, -20}, vec3{0, 0.5, 0.5}, 180, 0xff00ffff, "spanel.png")

	zanim := float32(0.0)
	zang := float32(0.0)

	userInput := UserInput{}
	for !userInput.quit {
		userInput.GetUserInput()

		scene.Draw()

		scene.Shape("cube1").rotation.x += 3
		scene.Shape("cube1").rotation.z += 1

		scene.Shape("plane1").rotation.z += 1
		scene.Shape("plane1").position.z = zanim - 20

		zanim = math32.Cos(zang) * 5
		zang += 0.1

		scene.Shape("sphere1").rotation.x += 1
		scene.Shape("sphere1").rotation.y += 0.5

		scene.Shape("torus1").rotation.y += 2
		scene.Shape("torus1").rotation.x += 0.7

		scene.Shape("tube1").rotation.y += 3
		scene.Shape("tube1").rotation.x += 0.3

		scene.Shape("cone1").rotation.y += 3
		scene.Shape("cone1").rotation.x += 0.3

		scene.Shape("tcone1").rotation.y += 3
		scene.Shape("tcone1").rotation.x += 2

		scene.Shape("spring1").rotation.y += 3
		scene.Shape("spring1").rotation.x += 2

		scene.Shape("cylinder1").rotation.x += 0.7
		scene.Shape("cylinder1").rotation.y += 3
		scene.Shape("cylinder1").rotation.z += 0.1
		//var shape = scene.Shape("cube1")
		//shape.rotation.x += 0.5
		//shape.rotation.y += 0.3

		scene.window.GLSwap()
	}

	scene.Quit()
}
