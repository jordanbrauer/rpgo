package game

import (
	"log"
	"math/rand"

	ray "github.com/gen2brain/raylib-go/raylib"
	"github.com/jordanbrauer/rpgo/game/component"
	"github.com/jordanbrauer/rpgo/game/entity"
)

func Spawn(world World, obj, tex string, x, y, z, scale float32) entity.Entity {
	object := world.CreateEntity()

	// TODO: share and manage unique component data (textures, scripts, etc.)
	model := ray.LoadModel(obj)
	texture := ray.LoadTexture(tex)

	ray.SetMaterialTexture(model.Materials, ray.MapDiffuse, texture)
	world.AttachComponent(object, &component.Model{Loaded: model})
	world.AttachComponent(object, &component.Texture{Loaded: texture})
	world.AttachComponent(object, &component.Position{
		Vector3: ray.NewVector3(x, y, z),
		Scale:   1.0,
	})

	return object
}

// Abort handles an error by checking for a nil value and panicing otherwise.
func Abort(caught error) {
	if caught != nil {
		log.Panicln(caught)
	}
}

// Lerp is a linear interpolation implementation from many shader languages.
// Used to find a given distance between two known locations (coordinates).
//
// The formula used here is taken from the Wikipedia article on the
// subject: https://en.wikipedia.org/wiki/Linear_interpolation#Programming_language_support
func Lerp(a, b, distance float32) float32 {
	return a + ((b - a) * distance)
}

func RandomInt(min, max int) int {
	return rand.Intn((max - min + 1)) + min
}

func RandomInt32(min, max int) int32 {
	return int32(RandomInt(min, max))
}

func RandomFloat32(min, max int) float32 {
	return float32(RandomInt(min, max))
}
