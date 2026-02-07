package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "servers.....")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	mgr := NewManager()
	mgr.RegisterNode(NewNode(rl.NewVector2(100, 100), 1))
	mgr.RegisterNode(NewNode(rl.NewVector2(300, 300), 2))
	mgr.RegisterNode(NewNode(rl.NewVector2(600, 300), 1.5))

	for !rl.WindowShouldClose() {

		mgr.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		mgr.Render()
		rl.EndDrawing()
	}
}
