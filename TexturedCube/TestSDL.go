package main

const width, height = 800, 600

func main() {

	scene := Scene{}
	scene.Setup("SDL 3D Cude", width, height)

	scene.AddShape("cube1", cuboidShape, 3, 3, 3, vec3{-7, -10, -20}, vec3{0, 0, 0}, 0xff00ffff, "redsky.png")
	scene.AddShape("plane1", planeShape, 5, 0, 5, vec3{7, -10, -20}, vec3{0, 0, 0}, 0xff00ffff, "me.png")
	scene.AddShape("sphere1", sphereShape, 5, 3, 20, vec3{-7, 0, -20}, vec3{0, 0.5, 0.5}, 0xff00ffff, "spanel.png")
	scene.AddShape("torus1", torusShape, 5, 2, 30, vec3{7, 0, -20}, vec3{0, 0.5, 0.5}, 0xff00ffff, "spanel.png")
	scene.AddShape("tube1", tubeShape, 2, 4, 6, vec3{10, 10, -20}, vec3{0, 0.5, 0.5}, 0xff00ffff, "spanel.png")
	scene.AddShape("cylinder1", cylinderShape, 3, 5, 20, vec3{0, 10, -20}, vec3{0, 0.5, 0.5}, 0xff00ffff, "me.png")

	userInput := UserInput{}
	for !userInput.quit {
		userInput.GetUserInput()

		scene.Draw()

		scene.Shape("cube1").rotation.x += 3
		scene.Shape("cube1").rotation.z += 1
		scene.Shape("plane1").rotation.z += 1
		scene.Shape("sphere1").rotation.x += 1
		scene.Shape("sphere1").rotation.y += 0.5

		scene.Shape("torus1").rotation.y += 2
		scene.Shape("torus1").rotation.x += 0.7

		scene.Shape("tube1").rotation.y += 3
		scene.Shape("tube1").rotation.x += 0.3

		scene.Shape("cylinder1").rotation.x += 0.7
		//var shape = scene.Shape("cube1")
		//shape.rotation.x += 0.5
		//shape.rotation.y += 0.3

		scene.window.GLSwap()
	}

	scene.Quit()
}
