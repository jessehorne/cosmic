package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jessehorne/cosmic/internal/pkg"
	"github.com/jessehorne/cosmic/internal/pkg/save"
)

const (
	screenWidth  = 900
	screenHeight = 600
)

func main() {
	// create save file if not exists
	err := save.CreateSaveIfNotExists()
	if err != nil {
		log.Fatal(err)
		return
	}

	// init raylib stuff
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	rl.SetConfigFlags(rl.FlagWindowTransparent)

	rl.InitWindow(screenWidth, screenHeight, "cosmic v0.0.1")
	defer rl.CloseWindow()

	daw := pkg.NewDAW()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 0, G: 0, B: 0, A: 150})

		daw.Update()

		rl.EndDrawing()
	}

	daw.Close()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
