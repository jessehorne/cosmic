package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg"
	"github.com/jessehorne/cosmic/internal/pkg/ui"
)

const (
	screenWidth  = 900
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "cosmic v0.0.1")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	daw := pkg.NewDAW()

	playButton := ui.NewPlay(daw)
	pauseButton := ui.NewPause(daw)
	stopButton := ui.NewStop(daw)
	bpmCounter := ui.NewBpm(daw)
	which := ui.NewWhich(daw)
	metronome := ui.NewMetronome(daw)

	var widgets []ui.Widget
	widgets = append(widgets, playButton)
	widgets = append(widgets, pauseButton)
	widgets = append(widgets, stopButton)
	widgets = append(widgets, bpmCounter)
	widgets = append(widgets, which)
	widgets = append(widgets, metronome)

	daw.SetMetronome(metronome)

	for !rl.WindowShouldClose() {
		// --- 2. Drawing Phase ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		daw.Update()

		for _, w := range widgets {
			w.Update()
			w.Draw()
		}

		rl.EndDrawing()
	}

	for _, w := range widgets {
		w.Close()
	}

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
