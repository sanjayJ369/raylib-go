package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type Node struct {
	ID          string
	Pos         rl.Vector2
	timer       *Timer
	Radius      float32
	timerLoader *CircularLoader
	Counter     int32
	Manager     *Manager
	Color       rl.Color
}

func (n *Node) Update() {
	n.timer.Update()
	angle := (n.timer.Elapsed / n.timer.Duration) * 360
	n.timerLoader.Update(float32(angle))
}

func (n *Node) Render() {
	n.timerLoader.Render()
	rl.DrawCircleV(n.Pos, n.Radius-1, n.Color)
	rl.DrawText(fmt.Sprintf("%d", n.Counter), int32(n.Pos.X), int32(n.Pos.Y), 20, rl.White)
}

func (n *Node) OnCollisionPacket(p *Packet) {
	if p.DestID == n.ID {
		switch p.Type {
		case INCREMENT:
			atomic.AddInt32(&n.Counter, 1)
		case DECEREMENT:
			atomic.AddInt32(&n.Counter, -1)
		case TIMER_RESET:
			n.timer.Reset()
		}
	}
}

func (n *Node) SendPacketToNodes() {
	// send the copy of packet to all the nodes in the
	// service mesh
	for _, node := range n.Manager.nodes {
		if node.ID != n.ID {
			newp := NewPacket(n.Pos, node.Pos, rand.Intn(3), n.ID, node.ID)
			n.Manager.AddPacket(newp)
			fmt.Println("sending: ", newp.Type)
		}
	}
}

func NewNode(pos rl.Vector2, duration float64) *Node {
	// todo create callback
	// on timer timeout
	// create a new packet of random type and send it to all other nodes

	node := &Node{
		ID:          uuid.NewString(),
		Pos:         pos,
		Counter:     0,
		timerLoader: NewCircularLoader(pos.X, pos.Y),
		Radius:      50,
		Color:       rl.Green,
	}

	callback := func() {
		node.SendPacketToNodes()
	}
	node.timer = NewTimer(duration, true, callback)

	return node
}
