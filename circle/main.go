package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	Pos    rl.Vector2
	Color  rl.Color
	Radius float32
}

const (
	MIN_RADIUS   = 1
	MAX_RADIUS   = 100
	DELTA_RADIUS = 1
)

func (b *Ball) render() {
	rl.DrawCircle(int32(b.Pos.X), int32(b.Pos.Y), b.Radius, b.Color)
}

func (b *Ball) update() {
	// if mouse hover -> change from blue to green
	if rl.CheckCollisionPointCircle(rl.GetMousePosition(), b.Pos, float32(b.Radius)) {
		b.Color = rl.Green
	} else {
		b.Color = rl.Blue
	}

	// if mouse down -> reduce Radius
	change := rl.GetFrameTime() * DELTA_RADIUS * 10

	if rl.CheckCollisionPointCircle(rl.GetMousePosition(), b.Pos, float32(b.Radius)) && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		b.Radius -= change
		if b.Radius < MIN_RADIUS {
			b.Radius = MIN_RADIUS
		}
	} else {
		if b.Radius < MAX_RADIUS {
			b.Radius += change
		}

		if b.Radius >= MAX_RADIUS {
			b.Radius = MAX_RADIUS
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ball := Ball{
		Pos:    rl.NewVector2(400, 200),
		Color:  rl.Blue,
		Radius: MAX_RADIUS,
	}

	for !rl.WindowShouldClose() {
		ball.update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		ball.render()

		rl.EndDrawing()
	}
}
