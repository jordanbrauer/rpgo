package main

import (
	"math/rand"
	"time"

	ray "github.com/gen2brain/raylib-go/raylib"
	"github.com/jordanbrauer/rpgo/game"
)

var config *game.Configuration

func init() {
	config = game.Config()
	// storage := game.Storage()

	rand.Seed(time.Now().UnixNano())
	ray.InitWindow(config.Window.Width, config.Window.Height, config.Window.Title)
	ray.SetTargetFPS(config.FPS)

	if config.Window.Fullscreen {
		width := ray.GetMonitorWidth(ray.GetCurrentMonitor())
		height := ray.GetMonitorHeight(ray.GetCurrentMonitor())

		ray.SetWindowSize(width, height)
		ray.ToggleFullscreen()
	}
}

func main() {
	defer ray.CloseWindow()

	world := game.CreateWorld()
	rendering := game.EnableRendering(world)
	scripting, script := game.EnableScripting(world)

	defer script.Close()

	for !ray.WindowShouldClose() {
		delta := ray.GetFrameTime()

		world.Update(scripting.Name(), delta)
		world.Update(rendering.Name(), delta)
	}
}
