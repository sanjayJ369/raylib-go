package main

import rl "github.com/gen2brain/raylib-go/raylib"

// on click create a new packet
// and there will only be one server
// packest should hit the server
// and disappear
// on hit the server should change it's color

const (
	WINDOW_HEIGHT = 450
	WINDOW_WIDTH  = 800
)

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "packet-1")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	mgr := NewManager()
	mgr.AddServer(NewServer(rl.NewVector2(400, 200)))

	for !rl.WindowShouldClose() {

		mgr.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		mgr.Render()
		rl.EndDrawing()
	}
}
