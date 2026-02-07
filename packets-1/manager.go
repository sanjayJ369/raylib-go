package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Manager struct {
	packets map[string]*Packet
	server  *Server
}

func NewManager() *Manager {
	return &Manager{
		packets: map[string]*Packet{},
	}
}

func (m *Manager) Update() {
	// update packets
	for _, pkt := range m.packets {
		pkt.Update()
	}
	// update server
	m.server.Update()

	// check collusions

	for _, pkt := range m.packets {
		// check colussion
		if rl.CheckCollisionCircleRec(m.server.Pos, m.server.Radius, pkt.Rect) {
			m.server.onCollisionPacket(pkt)
			pkt.onCollisionServer(m, m.server)
		}
	}

	// on click create new packets
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		m.AddPacket(NewPacket(rl.GetMousePosition(), m.server.Pos))
	}
}

func (m *Manager) Render() {
	for _, pkt := range m.packets {
		pkt.Render()
	}

	m.server.Render()
}

func (m *Manager) AddServer(srv *Server) {
	m.server = srv
}

func (m *Manager) AddPacket(p *Packet) {
	m.packets[p.ID] = p
}

func (m *Manager) DeletePacket(id string) {
	delete(m.packets, id)
}

func (m *Manager) DeleteServer(id string) {
	m.server = nil
}
