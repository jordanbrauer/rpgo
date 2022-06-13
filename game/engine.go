package game

import (
	ray "github.com/gen2brain/raylib-go/raylib"
	"github.com/jordanbrauer/rpgo/game/component"
	"github.com/jordanbrauer/rpgo/game/system"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func EnableRendering(world World) *system.Rendering {
	textures := component.Texture{}.Name()
	models := component.Model{}.Name()
	positioning := component.Position{}.Name()

	world.RegisterComponent(textures)
	world.RegisterComponent(models)
	world.RegisterComponent(positioning)

	camera := &ray.Camera3D{
		Position:   ray.NewVector3(4.0, 2.0, 4.0),
		Target:     ray.NewVector3(0.0, 1.8, 0.0),
		Up:         ray.NewVector3(0.0, 1.0, 0.0),
		Fovy:       60.0,
		Projection: ray.CameraPerspective,
	}
	rendering := &system.Rendering{Camera: camera}

	// TODO: implement custom camera mode
	ray.SetCameraMode(*camera, ray.CameraFirstPerson)
	world.RegisterSystem(rendering, textures, positioning, models)

	return rendering
}

func EnableScripting(world World) (*system.Scripting, *lua.LState) {
	library := Library{World: world}
	script := lua.NewState()
	behaviour := component.Scriptable{}.Name()
	scripting := &system.Scripting{Global: script}

	script.SetGlobal("configuration", luar.New(script, config))
	script.SetGlobal("spawn", script.NewFunction(library.Spawn))

	world.RegisterComponent(behaviour)
	world.RegisterSystem(scripting, behaviour)

	return scripting, script
}
