package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	Duration float64
	Elapsed  float64
	Repeat   bool
	Active   bool
	Callback func()
}

func NewTimer(duration float64, repeat bool, cb func()) *Timer {
	return &Timer{
		Duration: duration,
		Repeat:   repeat,
		Callback: cb,
		Active:   true,
	}
}

func (t *Timer) Update() {
	if !t.Active {
		return
	}

	t.Elapsed += float64(rl.GetFrameTime())

	if t.Elapsed >= t.Duration {
		t.Callback()

		if t.Repeat {
			t.Elapsed = 0
		} else {
			t.Active = false
		}
	}
}

func (t *Timer) Reset() {
	t.Elapsed = 0
	t.Active = true
}
