package main

import (
	"log"

	"github.com/jessehorne/cosmic/internal/pkg/save"
	"github.com/jessehorne/cosmic/internal/pkg/ui"
)

func main() {
	// create save file if not exists
	err := save.CreateSaveIfNotExists()
	if err != nil {
		log.Fatal(err)
		return
	}

	app := ui.NewApp()
	app.Run()
}
