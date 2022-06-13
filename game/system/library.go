package system

import (
	"fmt"

	ray "github.com/gen2brain/raylib-go/raylib"
	"github.com/jordanbrauer/rpgo/game/component"
	lua "github.com/yuin/gopher-lua"
)

type Rendering struct {
	SystemAccess
	Camera *ray.Camera3D
}

func (Rendering) Name() string {
	return "rendering"
}

func (system *Rendering) Update(delta float32) {
	ray.BeginDrawing()

	// World
	ray.UpdateCamera(system.Camera)
	ray.ClearBackground(ray.White)
	ray.BeginMode3D(*system.Camera)
	ray.DrawGrid(10, 5.0)

	for _, entity := range system.Entities() {
		position := system.Component(entity, "position").(*component.Position)
		model := system.Component(entity, "model").(*component.Model)
		// texture := Component(entity, "texture").(*component.Texture)

		// ray.DrawBillboard(*camera, texture.Loaded, position.Vector3, 2.0, ray.White)
		ray.DrawModel(model.Loaded, position.Vector3, position.Scale, ray.White)
	}

	ray.EndMode3D()

	// HUD
	ray.DrawText(
		fmt.Sprintf("FPS: %.0f", ray.GetFPS()),
		16, 16, 24, ray.DarkGreen,
	)
	ray.DrawText(
		fmt.Sprintf("Entities: %d", system.World().Entities()),
		16, 40, 24, ray.DarkGreen,
	)

	ray.EndDrawing()
}

type Scripting struct {
	SystemAccess
	Global *lua.LState
}

func (Scripting) Name() string {
	return "scripting"
}

func (system *Scripting) Update(delta float32) {
	// for _, entity := range system.Entities() {
	// }

	if ray.IsKeyReleased(ray.KeySpace) {
		if err := system.Global.DoFile("content/plants/script.lua"); err != nil {
			panic(err)
		}
	}
}
