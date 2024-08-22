package main

const width, height = 800, 600

func main() {

	scene := Scene{}
	scene.Setup("SDL 3D Cude", width, height)

	scene.AddShape("cube1", cuboidShape, 20, 1, 20, vec3{0, 0, -10}, vec3{0, 0, 0}, 0xff00ffff, "redsky.png")
	scene.AddShape("plane1", planeShape, 5, 0, 5, vec3{0, 0, -6}, vec3{0, 0, 0}, 0xff00ffff, "me.png")

	userInput := UserInput{}
	for !userInput.quit {
		userInput.GetUserInput()

		scene.Draw()

		scene.Shape("cube1").rotation.z += 3
		scene.Shape("plane1").rotation.z += 1
		//var shape = scene.Shape("cube1")
		//shape.rotation.x += 0.5
		//shape.rotation.y += 0.3

		scene.window.GLSwap()
	}

	scene.Quit()
}
