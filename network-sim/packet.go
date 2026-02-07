package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

const (
	SIDEL_LEN_PACKET = 20
	SPEED            = 200

	INCREMENT   = 0
	DECEREMENT  = 1
	TIMER_RESET = 2
)

var (
	INCREMENT_COLOR   = rl.Green
	DECEREMENT_COLOR  = rl.Red
	TIMER_RESET_COLOR = rl.Brown
)

type Packet struct {
	ID      string
	Pos     rl.Vector2
	SideLen float32
	Color   rl.Color
	Dest    rl.Vector2
	Rect    rl.Rectangle
	Type    int
	SrcID   string // source node id
	DestID  string // dest node id
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

func (p *Packet) onCollisionNode(mgr *Manager, n *Node) {
	// todo delete itself...
	if n.ID == p.DestID {
		mgr.DeletePacket(p.ID)
	}
}

func NewPacket(Pos, Dest rl.Vector2, pt int, SrdID, DestID string) *Packet {
	var color rl.Color

	switch pt {
	case INCREMENT:
		color = INCREMENT_COLOR
	case DECEREMENT:
		color = DECEREMENT_COLOR
	case TIMER_RESET:
		color = TIMER_RESET_COLOR
	default:
		color = rl.Black
	}

	return &Packet{
		ID:      uuid.NewString(),
		Type:    pt,
		Pos:     Pos,
		SrcID:   SrdID,
		DestID:  DestID,
		SideLen: SIDEL_LEN_PACKET,
		Color:   color,
		Dest:    Dest,
	}
}
