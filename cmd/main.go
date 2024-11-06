package main

import (
	"context"
	"github.com/Killered672/game/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	// Exit приводит к завершению программы с заданным кодом.
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := applications.Config{
		Width:  100,
		Height: 100,
	}
	app := applications.New(cfg)

	return app.Run(ctx)
}
