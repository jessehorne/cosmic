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

	daw := pkg.NewDAW()

	playButton := ui.NewPlay(daw)
	pauseButton := ui.NewPause(daw)
	stopButton := ui.NewStop(daw)
	bpmCounter := ui.NewBpm(daw)
	which := ui.NewWhich(daw)

	var widgets []ui.Widget
	widgets = append(widgets, playButton)
	widgets = append(widgets, pauseButton)
	widgets = append(widgets, stopButton)
	widgets = append(widgets, bpmCounter)
	widgets = append(widgets, which)

	for !rl.WindowShouldClose() {
		// --- 2. Drawing Phase ---
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for _, w := range widgets {
			w.Update()
			w.Draw()

			// reset styles
			//gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, gui.GetStyle(gui.DEFAULT, gui.TEXT_SIZE))
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
