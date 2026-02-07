package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

const (
	MIN_SERVER_RADIUS = 10
	MAX_SERVER_RADIUS = 50
)

type Server struct {
	ID     string
	Pos    rl.Vector2
	Color  rl.Color
	Radius float32
}

func (s *Server) Update() {
	// ahhh.... nothing to do for now so //todo
}

func (s *Server) Render() {
	rl.DrawCircle(int32(s.Pos.X), int32(s.Pos.Y), s.Radius, s.Color)
}

func (s *Server) onCollisionPacket(p *Packet) {
	// for now just change the server color
	R := uint8(rand.Intn(256))
	G := uint8(rand.Intn(256))
	B := uint8(rand.Intn(256))
	s.Color = rl.NewColor(R, G, B, 255)
}

// it is just a circle
func NewServer(pos rl.Vector2) *Server {
	return &Server{
		ID:     uuid.NewString(),
		Pos:    pos,
		Color:  rl.Blue,
		Radius: MAX_SERVER_RADIUS,
	}
}
