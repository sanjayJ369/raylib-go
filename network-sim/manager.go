package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Manager struct {
	packets map[string]*Packet
	nodes   map[string]*Node
}

func NewManager() *Manager {
	return &Manager{
		packets: map[string]*Packet{},
		nodes:   map[string]*Node{},
	}
}

func (m *Manager) Update() {
	// update packets
	for _, pkt := range m.packets {
		pkt.Update()
	}
	// update server
	for _, n := range m.nodes {
		n.Update()
	}

	// check collusions
	for _, node := range m.nodes {
		for _, pkt := range m.packets {
			// check colussion
			if rl.CheckCollisionCircleRec(node.Pos, node.Radius, pkt.Rect) {
				node.OnCollisionPacket(pkt)
				pkt.onCollisionNode(m, node)
			}
		}
	}
}

func (m *Manager) Render() {
	for _, pkt := range m.packets {
		pkt.Render()
	}

	for _, node := range m.nodes {
		node.Render()
	}
}

func (m *Manager) RegisterNode(n *Node) {
	n.Manager = m
	m.nodes[n.ID] = n
}

func (m *Manager) AddPacket(p *Packet) {
	m.packets[p.ID] = p
}

func (m *Manager) DeletePacket(id string) {
	delete(m.packets, id)
}

func (m *Manager) DeleteServer(id string) {
	delete(m.nodes, id)
}
