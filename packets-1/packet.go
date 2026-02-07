package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

const (
	SIDEL_LEN_PACKET = 20
	SPEED            = 200
)

// they must be spawner on mouse click
// it is square
type Packet struct {
	ID      string // uuid
	Pos     rl.Vector2
	SideLen float32
	Color   rl.Color
	Dest    rl.Vector2
	Rect    rl.Rectangle
}

func (p *Packet) Update() {
	delta := rl.Vector2Subtract(p.Dest, p.Pos)
	if rl.Vector2Length(delta) > 0 {
		dir := rl.Vector2Normalize(delta)
		vel := rl.Vector2Scale(dir, SPEED*rl.GetFrameTime())
		p.Pos = rl.Vector2Add(p.Pos, vel)
	}

	p.Rect = rl.NewRectangle(p.Pos.X, p.Pos.Y, p.SideLen, p.SideLen)
}

func (p *Packet) Render() {
	rl.DrawRectangleRec(p.Rect, p.Color)
}

func (p *Packet) onCollisionServer(mgr *Manager, s *Server) {
	// todo delete itself...
	mgr.DeletePacket(p.ID)
}

func NewPacket(Pos, Dest rl.Vector2) *Packet {
	return &Packet{
		ID:      uuid.NewString(),
		Pos:     Pos,
		SideLen: SIDEL_LEN_PACKET,
		Color:   rl.Black,
		Dest:    Dest,
	}
}
